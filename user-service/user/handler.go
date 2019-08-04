package user

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    "context"
    "github.com/google/uuid"
    "github.com/micro/go-micro/metadata"
    "errors"
    "fmt"
)

type service struct {
    repo Repository
}

func (srv *service) Register(ctx context.Context, req *pb.RegisterRequest, resp *pb.Empty) error {
    fmt.Printf("%+v\n", req)
    if err := srv.repo.CreateUser(req); err != nil {
        return err
    }
    return nil
}

func (srv *service) GetVerifier(ctx context.Context, req *pb.VerifierRequest, resp *pb.VerifierResponse) error {
    res, err := srv.repo.GetVerifier(req)
    if err != nil {
        return err
    }
    *resp = *res
    return nil
}

func (srv *service) AddDevice(ctx context.Context, req *pb.AddDeviceRequest, resp *pb.Empty) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    session_id := md["Session-Id"]
    if session_id == "" {
        return errors.New("unauthorized")
    }
    if err := srv.repo.AddDevice(req); err != nil {
        return err
    }
    return nil
}

func (srv *service) GetAccessKey(ctx context.Context, req *pb.Empty, resp *pb.AccessKeyResponse) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    session_id := md["Session-Id"]
    username := md["Username"]
    if session_id == "" || username == "" {
        return errors.New("unauthorized")
    }
    accessKey := uuid.New().String()
    if err := srv.repo.InsertAccessKey(accessKey, username); err != nil {
        return err
    }
    resp.AccessKey = accessKey
    return nil
}
