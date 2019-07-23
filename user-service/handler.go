package main

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    "context"
)

type service struct {
    repo Repository
}

func (srv *service) RegisterUser(ctx context.Context, req *pb.RegisterRequest, resp *pb.Empty) error {
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
