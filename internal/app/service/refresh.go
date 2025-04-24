package service

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrTokenMalformed = errors.New("token malformed")
	ErrTokenMismatch  = errors.New("invalid refresh token")
	ErrTokenReused    = errors.New("was reuse")
	ErrGenerateToken  = errors.New("problem with generate in refresh")
)

func (s *ServiceWorker) Refresh(ctx context.Context, nowClientIP, oldRefreshToken string) (string, string, error) {

	parts := strings.SplitN(oldRefreshToken, ".", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", ErrTokenMalformed
	}
	oldRtID, secret := parts[0], parts[1]

	oldHash, oldClientIP, userID, flag, err := s.Repository.GetRefreshToken(ctx, oldRtID)

	if err != nil {
		return "", "", err
	}
	if flag != false {
		return "", "", ErrTokenReused
	}
	if err := bcrypt.CompareHashAndPassword([]byte(oldHash), []byte(secret)); err != nil {
		return "", "", ErrTokenMismatch
	}

	if oldClientIP != nowClientIP {
		go s.emailService.Notify(oldClientIP, nowClientIP, userID)
	}

	rtID, access, refresh, hash, err := s.GenerateToken(ctx, userID, nowClientIP)
	if err != nil {
		return "", "", ErrGenerateToken
	}

	err = s.Repository.UpdateRefreshToken(ctx, oldRtID, rtID, nowClientIP, hash, userID)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil

}
