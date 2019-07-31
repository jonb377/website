package user

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    "context"
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
    if err := srv.repo.AddDevice(req); err != nil {
        return err
    }
    return nil
}
