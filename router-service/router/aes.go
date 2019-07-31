package router

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

type AESCipher interface {
    Encrypt([]byte) ([]byte, error)
    Decrypt([]byte) ([]byte, error)
}

type AES struct {
    key []byte
}

func NewAESCipher(key []byte) AESCipher {
    return &AES{
        key: key,
    }
}

func (aesCipher *AES) Encrypt(message []byte) ([]byte, error) {
    block, err := aes.NewCipher(aesCipher.key)
    if err != nil {
        return nil, err
    }
    cipherText := make([]byte, 12)
    iv := cipherText[:]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    res := aesgcm.Seal(nil, iv, message, nil)
    if err != nil {
        return nil, err
    }
    cipherText = make([]byte, 12+len(res))
    copy(cipherText[:12], iv)
    copy(cipherText[12:], res)
    return cipherText, nil
}

func (aesCipher *AES) Decrypt(message []byte) ([]byte, error) {
    block, err := aes.NewCipher(aesCipher.key)
    if err != nil {
        return nil, err
    }
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    res, err := aesgcm.Open(nil, message[:12], message[12:], nil)
    if err != nil {
        return nil, err
    }
    return res, nil
}
