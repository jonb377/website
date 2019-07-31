package auth

import (
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/auth-service/proto/auth"
    "context"
    "errors"
    "github.com/google/uuid"
    "github.com/jinzhu/gorm"
    "github.com/micro/go-micro/metadata"
    "fmt"
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
    if device == "" {
        return errors.New("device header required")
    }
    verifierResponse, err := s.userClient.GetVerifier(context.Background(), &userProto.VerifierRequest{
        Username: req.Username,
        Device: device,
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
    }
    if err := s.db.Create(&session).Error; err != nil {
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
    device := md["Device"]
    if device == "" {
        return errors.New("device header required")
    }
    token := md["Token"]
    claims, err := s.tokenService.Decode(token)
    if err != nil {
        fmt.Println("Error response from token service ", err)
        return err
    }
    if claims.Device != device {
        fmt.Println("Invalid device ", claims.Device, " != ", device)
        return errors.New(fmt.Sprintf("invalid device %v != %v", claims.Device, device))
    }
    verifierResponse, err := s.userClient.GetVerifier(context.Background(), &userProto.VerifierRequest{
        Username: claims.Username,
        Device: claims.Device,
    })
    if err != nil {
        fmt.Println("Error response from user service ", err)
        return err
    }
    var session Session
    if err := s.db.Table("sessions").Where("session_id = ?", claims.SessionId).First(&session).Error; err != nil {
        return err
    }
    srp := NewSRPServerWithB(claims.Username, verifierResponse.Salt, verifierResponse.Verifier, session.SRPA, session.SRPb)
    HAMK := srp.VerifyM(req.M)
    if HAMK == nil {
        return errors.New("authorization failed")
    }

    // Store the session key
    key := srp.getKey()
    if err := s.db.Table("sessions").Where("session_id = ?", claims.SessionId).Update("key", key).Error; err != nil {
        return err
    }

    resp.HAMK = HAMK
    return nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest, resp *pb.ValidateTokenResponse) error {
    claims, err := s.tokenService.Decode(req.Token)
    if err != nil {
        return err
    }
    var session Session
    if err = s.db.Table("sessions").Where("session_id = ?", claims.SessionId).First(&session).Error; err != nil {
        return err
    }
    if session.Device != req.Device {
        return errors.New("device is not authorized for session")
    }
    resp.SessionKey = session.Key
    resp.SessionID = session.SessionID
    return nil
}

func (s *AuthService) CloseConnection(ctx context.Context, req *pb.ConnectionCloseRequest, resp *pb.Empty) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    device := md["Device"]
    if device == "" {
        return errors.New("device header required")
    }
    token := md["Token"]
    claims, err := s.tokenService.Decode(token)
    if err != nil {
        fmt.Println("Error response from token service ", err)
        return err
    }
    if claims.Device != device {
        fmt.Println("Invalid device ", claims.Device, " != ", device)
        return errors.New(fmt.Sprintf("invalid device %v != %v", claims.Device, device))
    }
    if err := s.db.Table("sessions").Where("session_id = ?", claims.SessionId).Delete(&Session{}).Error; err != nil {
        return err
    }
    return nil
}
