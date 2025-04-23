package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *TokenAlgoritm) CreateAccessToken(userID, tokenID, ipAddress string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"ip":      ipAddress,
		"rtid":    tokenID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(s.secret))
}
