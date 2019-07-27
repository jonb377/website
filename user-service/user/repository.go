package user

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    "github.com/jinzhu/gorm"
)

type Repository interface {
    GetVerifier(*pb.VerifierRequest) (*pb.VerifierResponse, error)
    CreateUser(*pb.RegisterRequest) error
    AddDevice(*pb.AddDeviceRequest) error
}

type UserRepository struct {
    db *gorm.DB
}

func (repo *UserRepository) GetVerifier(req *pb.VerifierRequest) (*pb.VerifierResponse, error) {
    var user User
    subquery := repo.db.Table("devices").Select("*").Where("guid = ?", req.Device).QueryExpr()
    if err := repo.db.Table("users").Select("salt, verifier").Where(
            "username = ? and exists ?",
            req.Username,
            subquery,
        ).First(&user).Error; err != nil {
        return nil, err
    }
    res := pb.VerifierResponse {
        Verifier: user.Verifier,
        Salt: user.Salt,
    }
    return &res, nil
}

func (repo *UserRepository) CreateUser(req *pb.RegisterRequest) error {
    user := User {
        Username:  req.Username,
        FirstName: req.Firstname,
        LastName:  req.Lastname,
        Email:     req.Email,
        Salt:      req.Salt,
        Verifier:  req.Verifier,
    }
    if err := repo.db.Create(user).Error; err != nil {
        return err
    }
    return nil
}

func (repo *UserRepository) AddDevice(req *pb.AddDeviceRequest) error {
    device := Device {
        Guid: req.Device,
        Username: req.Username,
    }
    if err := repo.db.Create(device).Error; err != nil {
        return err
    }
    return nil
}
