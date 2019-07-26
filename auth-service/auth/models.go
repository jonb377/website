package auth

import (
    "time"
)

type Session struct {
    SessionID   string      `gorm:"PRIMARY_KEY"`
    Username    string
    Device      string
    Key         []byte
    SRPb        []byte
    SRPA        []byte
    LastUsed    time.Time   `sql:"type: timestamp without time zone default now()"`
}
