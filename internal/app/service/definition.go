package service

import (
	"context"

	"github.com/federalbyro/encryptServer/internal/app/repository"
	"github.com/federalbyro/encryptServer/internal/app/token"
	"github.com/federalbyro/encryptServer/internal/infra"
)

type ServiceRegister interface {
	CreateTokens(ctx context.Context, userID, ipAddress string) (string ,string, error)
	Refresh(ctx context.Context,nowClientIP, oldRefreshToken string)(string,string,error)
}

type ServiceWorker struct {
	tokenGen     token.TokenRegister
	emailService infra.NotificationRegister
	Repository   repository.DBRegister
}

func NewServiceWorker(tokenGen token.TokenRegister, emailService infra.NotificationRegister, repo repository.DBRegister) *ServiceWorker {
	return &ServiceWorker{
		tokenGen:     tokenGen,
		emailService: emailService,
		Repository:   repo,
	}
}
