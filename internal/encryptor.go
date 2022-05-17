package internal

import (
	"crypto/md5"
	"encoding/hex"
)

type Encryption interface {
	Encrypt(url string) string
}

type Encryptor struct {
	Salt string
}

func NewEncryptor(salt string) *Encryptor {
	return &Encryptor{
		Salt: salt,
	}
}

func (e *Encryptor) Encrypt(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[0:4])
}
