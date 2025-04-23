package service

import "github.com/google/uuid"

func (s *ServiceWorker) CreateTokens(userID, ipAddress string) (string, string, error) {

	refreshAccessID := uuid.New().String()

	accessToken, err := s.tokenGen.CreateAccessToken(userID, refreshAccessID, ipAddress)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenGen.CreateRefreshToken()
	if err != nil {
		return "", "", err
	}

	refreshHash, err := s.tokenGen.HashRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	err = s.Repository.SaveRefreshToken(refreshAccessID, userID, ipAddress, refreshHash)
	return accessToken, refreshToken, err
}
