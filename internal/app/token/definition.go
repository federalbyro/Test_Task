package token

type TokenRegister interface {
	CreateAccessToken(userID, tokenID, ipAddress string) (string, error)
	GenerateRefreshToken(rtid string) (string, string, error)
	HashRefreshToken(token string) (string, error)
}

type TokenManager struct {
	secret string
}

func NewTokenGenerator(secret string) *TokenManager {
	return &TokenManager{secret: secret}
}
