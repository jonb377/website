package main

import "time"

// Tuple (User, Domain, Username) defines a unique entry
type PasswordEntry struct {
    User string `gorm:"primary_key"`
    Domain string `gorm:"primary_key"`
    Username string `gorm:"primary_key"`
    Password string
    Date time.Time `sql:"type: timestamp without time zone"`
    Deleted bool
}
