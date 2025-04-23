package main

import (
	"log"

	cryptserver "github.com/federalbyro/encryptServer/internal/app/api/server"
	"github.com/federalbyro/encryptServer/internal/pkg/config"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := cryptserver.New(config)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
