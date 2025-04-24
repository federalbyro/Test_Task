package cryptserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
        Addr:    ":" + s.config.ServerPort,
        Handler: s.router,
    }
	
	go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()
    log.Println("Server started on", s.config.ServerPort)

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit

    log.Println("Shutdown...")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return srv.Shutdown(ctx)
}
