package main

import (
    userProto "github.com/jonb377/website/user-service/proto/user"
    pb "github.com/jonb377/website/auth-service/proto/auth"
    "context"
    "errors"
    "github.com/google/uuid"
    "github.com/jinzhu/gorm"
    "github.com/micro/go-micro/metadata"
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
    if err := s.db.Create(Session{
            SessionID: sessionId,
            Username: req.Username,
            Device: device,
            Key: nil,
            SRPb: srp.b.Bytes(),
            SRPA: req.A,
        }).Error; err != nil {
        return err
    }
    resp.B = srp.B.Bytes()
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
    verifierResponse, err := s.userClient.GetVerifier(context.Background(), &userProto.VerifierRequest{
        Username: req.Username,
        Device: device,
    })
    if err != nil {
        return err
    }
    var session Session
    if err := s.db.Table("sessions").Where("session_id = ?", ).First(&session).Error; err != nil {
        return err
    }
    srp := NewSRPServerWithB(req.Username, verifierResponse.Salt, verifierResponse.Verifier, session.SRPA, session.SRPb)
    HAMK := srp.VerifyM(req.M)
    if HAMK == nil {
        return errors.New("authorization failed")
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

func (s *AuthService) CloseConnection(ctx context.Context, req *pb.CloseConnectionRequest, resp *pb.Empty) error {
    if err := s.db.Table("sessions").Where("session_id = ?", req.SessionID).Delete(&Session{}).Error; err != nil {
        return err
    }
    return nil
}
