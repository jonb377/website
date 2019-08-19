package user

type User struct {
    Username    string `gorm:"primary_key"`
    Firstname   string
    Lastname    string
    Email       string
    Admin       bool
}

type UserKey struct {
    Username      string `sql:"type: varchar REFERENCES users(username)"`
    SRPSalt       []byte
    SRPVerifier   []byte
    MUKSalt       []byte
    PublicKey     []byte
    EncPrivateKey []byte
    EncVaultKey   []byte
}

type Device struct {
    Guid        string `gorm:"primary_key"`
    Username    string `sql:"type:varchar REFERENCES users(username)"`
}


// Login using a shared secret
// Lifespan: 10 minutes
type AccessKey struct {
    Username string
    Key string
    CreatedAt int64
}
