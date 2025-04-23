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

func (h *TokenHandler) RegisterRoutes(router *gin.Engine) {
	tokens := router.Group("/tokens")

	tokens.GET("/access", h.CreateTokens)
	tokens.POST("/refresh", h.Update)

}
