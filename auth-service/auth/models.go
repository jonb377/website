package auth

type Session struct {
    SessionID   string      `gorm:"PRIMARY_KEY"`
    Username    string
    Device      string
    Key         []byte
    SRPb        []byte
    SRPA        []byte
    LastUsed    int64
}

type Permission struct {
    Username    string  `gorm:"PRIMARY_KEY"`
    Permission  string  `gorm:"PRIMARY_KEY"`
}
