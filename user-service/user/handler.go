package user

import (
    pb "github.com/jonb377/website/user-service/proto/user"
    util "github.com/jonb377/website/router-service/router"
    "context"
    "github.com/google/uuid"
    "github.com/micro/go-micro/metadata"
    "github.com/jinzhu/gorm"
    "errors"
    "time"
    "log"
)

type service struct {
    db *gorm.DB
}

func (srv *service) Register(ctx context.Context, req *pb.RegisterRequest, resp *pb.Empty) error {
    user := User {
        Username:  req.Username,
        Firstname: req.Firstname,
        Lastname:  req.Lastname,
        Email:     req.Email,
    }
    if err := srv.db.Create(&user).Error; err != nil {
        return err
    }
    keyInfo := UserKey {
        Username:      req.Username,
        SRPSalt:       req.SRPSalt,
        SRPVerifier:   req.SRPVerifier,
        MUKSalt:       req.MUKSalt,
        PublicKey:     req.PublicKey,
        EncPrivateKey: req.EncPrivateKey,
        EncVaultKey:   req.EncVaultKey,
    }
    if err := srv.db.Create(&keyInfo).Error; err != nil {
        return err
    }
    device := Device {
        Guid: req.Device,
        Username: req.Username,
    }
    return srv.db.Create(&device).Error
}

func (srv *service) GetVerifier(ctx context.Context, req *pb.VerifierRequest, resp *pb.VerifierResponse) error {
    var keyInfo UserKey
    deviceQuery := srv.db.Table("devices").Select("*").Where(
        "username = ? and guid = ?",
        req.Username,
        req.Device,
    ).QueryExpr()
    if err := srv.db.Table("user_keys").Select("srp_salt, srp_verifier").Where(
            "username = ? and exists (?)",
            req.Username, deviceQuery,
        ).First(&keyInfo).Error; err != nil {
        return err
    }
    resp.Verifier = keyInfo.SRPVerifier
    resp.Salt = keyInfo.SRPSalt
    return nil
}

func (srv *service) RegisterDevice(ctx context.Context, req *pb.Empty, resp *pb.RegisterDeviceResponse) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    accessKeyString := md["Access-Key"]
    device := md["Device"]
    if accessKeyString == "" {
        return errors.New("unauthorized")
    }
    var accessKey AccessKey
    if err := srv.db.Where("key = ?", accessKeyString).First(&accessKey).Error; err != nil {
        return err
    }

    // Delete the access key; single-use
    if err := srv.db.Delete(&accessKey).Error; err != nil {
        return err
    }
    newDevice := Device {
        Guid: device,
        Username: accessKey.Username,
    }

    // Register the new device
    if err := srv.db.Create(&newDevice).Error; err != nil {
        return err
    }

    // Retrieve key info
    var keyInfo UserKey
    var user User
    if err := srv.db.Where("username = ?", accessKey.Username).First(&keyInfo).Error; err != nil {
        return err
    }
    if err := srv.db.Where("username = ?", accessKey.Username).First(&user).Error; err != nil {
        return err
    }
    resp.SRPSalt       = keyInfo.SRPSalt
    resp.MUKSalt       = keyInfo.MUKSalt
    resp.PublicKey     = keyInfo.PublicKey
    resp.EncPrivateKey = keyInfo.EncPrivateKey
    resp.EncVaultKey   = keyInfo.EncVaultKey
    resp.Username      = user.Username
    resp.Firstname     = user.Firstname
    resp.Lastname      = user.Lastname
    resp.Email         = user.Email
    return nil
}

func (srv *service) GetAccessKey(ctx context.Context, req *pb.Empty, resp *pb.AccessKeyResponse) error {
    username, _, _, err := util.RequireAuth(ctx)
    if err != nil {
        return err
    }
    accessKey := AccessKey{
        Username: username,
        Key: uuid.New().String(),
        CreatedAt: time.Now().Unix(),
    }
    resp.AccessKey = accessKey.Key
    return srv.db.Create(&accessKey).Error
}

func (srv *service) Logout(ctx context.Context, req *pb.Empty, resp *pb.LogoutResponse) error {
    username, device, _, err := util.RequireAuth(ctx)
    if err != nil {
        return err
    }
    var count int
    log.Println("username: ", username, ", device: ", device)
    if err := srv.db.Table("devices").Where("username = ?", username).Count(&count).Error; err != nil {
        return err
    }
    if count < 2 {
        resp.Approved = false
        return nil
    }

    // Delete the device
    if err := srv.db.Delete(&Device{
        Guid: device,
        Username: username,
    }).Error; err != nil {
        return err
    }

    // TODO: Delete session
    resp.Approved = true
    return nil
}
