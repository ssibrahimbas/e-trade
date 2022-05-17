package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Jwt struct {
	secretKey []byte
}

func (j *Jwt) setSecret(s string) {
	j.secretKey = []byte(s)
}

func (j *Jwt) Sign(p interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["payload"] = p
	claims["expire"] = time.Now().Add(time.Minute * 30).Unix()
	return token.SignedString(j.secretKey)
}

func (j *Jwt) Parse(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})
}
