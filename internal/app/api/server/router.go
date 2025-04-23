package cryptserver

import (
	"log"

	"github.com/federalbyro/encryptServer/internal/app/api/handler"
	"github.com/federalbyro/encryptServer/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

type CryptServer struct {
	config   *config.Config
	router   *gin.Engine
	register []handler.RouterRegister
}

func New(config *config.Config, registars ...handler.RouterRegister) *CryptServer {
	return &CryptServer{
		config:   config,
		router:   gin.Default(),
		register: registars,
	}
}

func (s *CryptServer) Run() error {
	log.Print("Start server...")
	for _, registrar := range s.register {
		registrar.RegisterRoutes(s.router)
	}
	return s.router.Run(s.config.ServerPort)
}
