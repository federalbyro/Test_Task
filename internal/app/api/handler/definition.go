package handler

import (
	"github.com/federalbyro/encryptServer/internal/app/service"
	"github.com/gin-gonic/gin"
)

type RouterRegister interface {
	RegisterRoutes(router *gin.Engine)
}

type TokenHandler struct {
	service service.ServiceRegister
}

func NewTokenHandler(serv service.ServiceRegister)*TokenHandler{
	return &TokenHandler{
		service: serv,
	}
}


func (h *TokenHandler) RegisterRoutes(router *gin.Engine) {
	tokens := router.Group("/tokens")
	tokens.POST("/access", h.CreateTokens)
	tokens.PUT("/refresh", h.Update)

}
