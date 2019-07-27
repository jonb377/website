package user

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username    string `gorm:"PRIMARY_KEY"`
    FirstName   string
    LastName    string
    Email       string
    Salt        []byte
    Verifier    []byte
}

type Device struct {
    gorm.Model
    Guid        string `gorm:"PRIMARY_KEY"`
    Username    string `sql:"type:varchar REFERENCES users(username)"`
}
