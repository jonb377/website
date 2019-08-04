package user

type User struct {
    Username    string `gorm:"primary_key"`
    FirstName   string
    LastName    string
    Email       string
    Salt        []byte
    Verifier    []byte
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
