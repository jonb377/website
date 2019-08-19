package passwordmanager

import (
    pb "github.com/jonb377/website/password-manager-service/proto/password-manager"
    util "github.com/jonb377/website/router-service/router"
    "context"
    "github.com/jinzhu/gorm"
)

type PasswordManagerService struct {
    db *gorm.DB
}

func (s *PasswordManagerService) UpdatePassword(ctx context.Context,  req *pb.UpdatePasswordRequest, resp *pb.Empty) error {
    session := util.GetSessionData(ctx)
    entry := PasswordEntry{
        Userkey: session.Username,
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
    session := util.GetSessionData(ctx)
    var dbpasswords []PasswordEntry
    if err := s.db.Table("password_entries").Where("userkey = ? and date > ?", session.Username, req.Date).Find(&dbpasswords).Error; err != nil {
        return err
    }
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
