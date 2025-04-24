package main

import (
	"log"

	"github.com/federalbyro/encryptServer/internal/app/api/handler"
	cryptserver "github.com/federalbyro/encryptServer/internal/app/api/server"
	"github.com/federalbyro/encryptServer/internal/app/repository"
	"github.com/federalbyro/encryptServer/internal/app/service"
	"github.com/federalbyro/encryptServer/internal/app/token"
	"github.com/federalbyro/encryptServer/internal/infra"
	"github.com/federalbyro/encryptServer/internal/pkg/config"
	"github.com/federalbyro/encryptServer/internal/pkg/db"
)

func main() {

	InitConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbPool,err := db.NewPostgresPool(InitConfig)

	if err!=nil{
		log.Fatal(err)
	}

	tokenGen := token.NewTokenGenerator(InitConfig.JWTSecret)

	repo:= repository.NewPostrgeDB(dbPool)
	serviceEmail := infra.NewNotificationService()
	service:=service.NewServiceWorker(tokenGen,serviceEmail,repo)
	handler := handler.NewTokenHandler(service)
	server := cryptserver.New(InitConfig,handler)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
