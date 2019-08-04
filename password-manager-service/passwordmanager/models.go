package passwordmanager

// Tuple (Userkey, Domain, Username) defines a unique entry
type PasswordEntry struct {
    Userkey string `gorm:"primary_key"`
    Domain string `gorm:"primary_key"`
    Username string `gorm:"primary_key"`
    Password string
    Date int64 // Date is milliseconds
    Deleted bool
}
