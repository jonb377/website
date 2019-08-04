package user

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    "github.com/jinzhu/gorm"
    "time"
)

type Repository interface {
    GetVerifier(*pb.VerifierRequest) (*pb.VerifierResponse, error)
    CreateUser(*pb.RegisterRequest) error
    AddDevice(*pb.AddDeviceRequest) error
    InsertAccessKey(string, string) error
}

type UserRepository struct {
    db *gorm.DB
}

func (repo *UserRepository) GetVerifier(req *pb.VerifierRequest) (*pb.VerifierResponse, error) {
    var user User
    deviceQuery := repo.db.Table("devices").Select("*").Where(
        "username = ? and guid = ?",
        req.Username,
        req.Device,
    ).QueryExpr()
    accessKeyQuery := repo.db.Table("access_keys").Select("*").Where(
        "username = ? and key = ? and created_at > extract(epoch from now() - '10 minutes'::interval)",
        req.Username,
        req.AccessKey,
    ).QueryExpr()
    if err := repo.db.Table("users").Select("salt, verifier").Where(
            "username = ? and (exists (?) or exists (?))",
            req.Username, deviceQuery, accessKeyQuery,
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
    if err := repo.db.Create(&user).Error; err != nil {
        return err
    }
    device := Device{
        Guid: req.Device,
        Username: req.Username,
    }
    if err := repo.db.Create(&device).Error; err != nil {
        return err
    }
    return nil
}

func (repo *UserRepository) AddDevice(req *pb.AddDeviceRequest) error {
    device := Device {
        Guid: req.Device,
        Username: req.Username,
    }
    if err := repo.db.Create(&device).Error; err != nil {
        return err
    }
    return nil
}

func (repo *UserRepository) InsertAccessKey(key string, username string) error {
    accessKey := AccessKey{
        Username: username,
        Key: key,
        CreatedAt: time.Now().Unix(),
    }
    return repo.db.Create(&accessKey).Error
}
