package services

import (
	"golang.org/x/crypto/bcrypt"
)

type Cipher struct {
}

func NewCipher() *Cipher {
	return &Cipher{}
}

func (c *Cipher) Encrypt(text string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), 8)
	return hashed, err
}

func (c *Cipher) Compare(text string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(text), hash)
	return err != nil, err
}
