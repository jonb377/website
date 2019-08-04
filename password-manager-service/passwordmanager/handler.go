package passwordmanager

import (
    pb "github.com/jonb377/website/password-manager-service/proto/password-manager"
    "context"
    "github.com/jinzhu/gorm"
    "github.com/micro/go-micro/metadata"
    "errors"
    "log"
)

type PasswordManagerService struct {
    db *gorm.DB
}

func (s *PasswordManagerService) UpdatePassword(ctx context.Context,  req *pb.UpdatePasswordRequest, resp *pb.Empty) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    user := md["Username"]
    if user == "" {
        return errors.New("unauthorized")
    }
    entry := PasswordEntry{
        Userkey: user,
        Domain: req.Domain,
        Username: req.Username,
        Password: req.Password,
        Date: req.Date,
        Deleted: req.Deleted,
    }
    if err := s.db.Table("password_entries").Set(
        "gorm:insert_option",
        "ON CONFLICT ON CONSTRAINT password_entries_pkey DO UPDATE SET username = excluded.username, password = excluded.password, date = excluded.date, deleted = excluded.deleted",
        ).Create(&entry).Error; err != nil {
        return err
    }
    return nil
}

func (s *PasswordManagerService) ListPasswords(ctx context.Context, req *pb.ListPasswordRequest, resp *pb.ListPasswordResponse) error {
    md, ok := metadata.FromContext(ctx)
    if !ok {
        md = metadata.Metadata{}
    }
    user := md["Username"]
    if user == "" {
        return errors.New("unauthorized")
    }
    log.Printf("Finding passwords for user %s after date %d\n", user, req.Date)
    var dbpasswords []PasswordEntry
    if err := s.db.Table("password_entries").Where("userkey = ? and date > ?", user, req.Date).Find(&dbpasswords).Error; err != nil {
        return err
    }
    log.Printf("Found %d passwords: %v\n", len(dbpasswords), dbpasswords)
    resp.Passwords = make([]*pb.PasswordEntry, len(dbpasswords))
    for i, p := range dbpasswords {
        resp.Passwords[i] = &pb.PasswordEntry{
            Domain: p.Domain,
            Username: p.Username,
            Password: p.Password,
            Date: p.Date,
            Deleted: p.Deleted,
        }
    }
    return nil
}
