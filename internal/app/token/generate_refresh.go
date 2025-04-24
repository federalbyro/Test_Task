package token

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func (t *TokenManager) GenerateRefreshToken(rtid string) (full string, bcryptHash string, err error) {
	raw := make([]byte, 32)
	if _, err = rand.Read(raw); err != nil {
		return
	}
	secret := base64.StdEncoding.EncodeToString(raw)
	full = rtid + "." + secret

	hash, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	bcryptHash = string(hash)
	return
}

func (t *TokenManager) HashRefreshToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(hash), err
}
