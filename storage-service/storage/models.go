package storage

type Blob struct {
    User string `gorm:"primary_key"`
    URI string `gorm:"primary_key"`
    Data []byte
    Modified int64
}
