package service

import (
	"github.com/federalbyro/encryptServer/internal/app/repository"
	"github.com/federalbyro/encryptServer/internal/app/token"
	"github.com/federalbyro/encryptServer/internal/infra"
)

type ServiceRegister interface {
	CreateTokens(userID, ipAddress string) (accessToken, refreshToken string, err error)
	Update()
}

type ServiceWorker struct {
	tokenGen     *token.TokenRegister
	emailService infra.NotificationRegister
	Repository   repository.DBRegister
}

func NewServiceWorker(tokenGen *token.TokenRegister, emailService infra.NotificationRegister, repo repository.DBRegister) *ServiceWorker {
	return &ServiceWorker{
		tokenGen:     tokenGen,
		emailService: emailService,
		Repository:   repo,
	}
}
