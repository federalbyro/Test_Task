package token

type TokenRegister interface {
	CreateAccessToken(userID, tokenID, ipAddress string) (string, error)
	GenerateRefreshToken() (string, error)
	HashRefreshToken(token string) (string, error)
}

type TokenAlgoritm struct {
	secret string
}

func NewTokenGenerator(secret string) *TokenAlgoritm {
	return &TokenAlgoritm{secret: secret}
}
