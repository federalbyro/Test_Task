package token

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func (t *TokenAlgoritm) GenerateRefreshToken() (token string, err error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	refreshToken := base64.StdEncoding.EncodeToString(bytes)
	return refreshToken, nil
}

func (t *TokenAlgoritm) HashRefreshToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(hash), err
}
