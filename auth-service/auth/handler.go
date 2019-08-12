package auth

import (
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/auth-service/proto/auth"
    "context"
    "errors"
    "github.com/google/uuid"
    "github.com/jinzhu/gorm"
    "github.com/micro/go-micro/metadata"
    "time"
    "fmt"
    "log"
)

type AuthService struct {
    db *gorm.DB
    userClient userProto.UserService
    tokenService Authable
}

func (s *AuthService) CreateConnection(ctx context.Context, req *pb.CreateConnectionRequest, resp *pb.CreateConnectionResponse) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    device := md["Device"]
    accessKey := md["Access-Key"]
    if device == "" && accessKey == "" {
        return errors.New("device header or access key required")
    }
    verifierResponse, err := s.userClient.GetVerifier(context.Background(), &userProto.VerifierRequest{
        Username: req.Username,
        Device: device,
        AccessKey: accessKey,
    })
    if err != nil {
        return err
    }
    srp := NewSRPServer(req.Username, verifierResponse.Salt, verifierResponse.Verifier, req.A)
    if srp == nil {
        return errors.New("failed to start SRP handshake")
    }
    // Create the Session object
    sessionId := uuid.New().String()
    session := Session{
        SessionID: sessionId,
        Username: req.Username,
        Device: device,
        Key: nil,
        SRPb: srp.b.Bytes(),
        SRPA: req.A,
        LastUsed: time.Now().Unix(),
    }
    if err := s.db.Table("sessions").Create(&session).Error; err != nil {
        return err
    }
    resp.B = srp.B.Bytes()
    token, err := s.tokenService.Encode(&session)
    if err != nil {
        return err
    }
    resp.Token = token
    return nil
}

func (s *AuthService) ConnectionChallenge(ctx context.Context, req *pb.ConnectionChallengeRequest, resp *pb.ConnectionChallengeResponse) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    session_id := md["Session-Id"]
    device := md["Device"]
    accessKey := md["Access-Key"]
    username := md["Username"]
    if session_id == "" || (device == "" && accessKey == "") || username == "" {
        return errors.New(fmt.Sprintf("unauthorized: %s %s %s %s", session_id, device, accessKey, username))
    }
    verifierResponse, err := s.userClient.GetVerifier(context.Background(), &userProto.VerifierRequest{
        Username: username,
        Device: device,
        AccessKey: accessKey,
    })
    if err != nil {
        fmt.Println("Error response from user service ", err)
        return err
    }
    var session Session
    if err := s.db.Table("sessions").Where("session_id = ? and last_used > extract(epoch from now() - '1 hour'::interval)", session_id).First(&session).Error; err != nil {
        return err
    }
    srp := NewSRPServerWithB(username, verifierResponse.Salt, verifierResponse.Verifier, session.SRPA, session.SRPb)
    HAMK := srp.VerifyM(req.M)
    if HAMK == nil {
        return errors.New("authorization failed")
    }

    // Store the session key
    key := srp.getKey()
    if err := s.db.Model(&session).Update("key", key).Error; err != nil {
        return err
    }

    resp.HAMK = HAMK
    return nil
}

func (s *AuthService) CloseConnection(ctx context.Context, req *pb.Empty, resp *pb.Empty) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    session_id := md["Session-Id"]
    if session_id == "" {
        return errors.New("No active session on CloseConnection call")
    }
    log.Println("Attempting to delete session_id ", session_id)
    if err := s.db.Delete(&Session{SessionID: session_id}).Error; err != nil {
        return err
    }
    log.Println("Successfully deleted session")
    return nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest, resp *pb.ValidateTokenResponse) error {
    claims, err := s.tokenService.Decode(req.Token)
    if err != nil {
        return err
    }
    log.Printf("Found claims from token: %v\n", claims)
    var session Session
    if err = s.db.Table("sessions").Where("session_id = ? and last_used > extract(epoch from now() - '1 hour'::interval)", claims.SessionId).First(&session).Error; err != nil {
        return err
    }
    log.Println("Found session: ", session.SessionID)

    // Verify the request is authorized to use the token
    if session.Device != req.Device {
        return errors.New("device is not authorized for session")
    } else if session.Username != claims.Username {
        return errors.New("user is not authorized for session")
    }

    session.LastUsed = time.Now().Unix()
    s.db.Save(&session)

    resp.SessionKey = session.Key
    resp.SessionID = session.SessionID
    resp.Username = session.Username
    return nil
}

