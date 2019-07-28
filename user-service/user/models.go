package user

//import (
//    "github.com/jinzhu/gorm"
//)

type User struct {
    //gorm.Model
    Username    string `gorm:"primary_key"`
    FirstName   string
    LastName    string
    Email       string
    Salt        []byte
    Verifier    []byte
}

type Device struct {
    //gorm.Model
    Guid        string `gorm:"primary_key"`
    Username    string `sql:"type:varchar REFERENCES users(username)"`
}
