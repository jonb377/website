package main

import (
    pb "github.com/jonb377/website/password-manager-service/proto/password-manager"
    "context"
    "github.com/jinzhu/gorm"
    "time"
)

type PasswordManagerService struct {
    db *gorm.DB
}

func (s *PasswordManagerService) UpdatePassword(ctx context.Context,  req *pb.UpdatePasswordRequest, resp *pb.Empty) error {
    if err := s.db.Table("logins").Set(
        "gorm:insert_option",
        "ON CONFLICT (user, username, domain) DO UPDATE SET username = excluded.username, password = excluded.password, date = excluded.date, deleted = excluded.deleted",
		).Create(&PasswordEntry{
			User: req.User,
			Domain: req.Entry.Domain,
			Username: req.Entry.Username,
			Password: req.Entry.Password,
			Date: time.Unix(req.Entry.Date, 0),
			Deleted: req.Entry.Deleted,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (s *PasswordManagerService) ListPasswords(ctx context.Context, req *pb.ListPasswordRequest, resp *pb.ListPasswordResponse) error {
	if err := s.db.Table("logins").Where("user = ? and date > ?", req.User, time.Unix(req.Date, 0)).Find(&resp.Passwords).Error; err != nil {
		return err
	}
	return nil
}
