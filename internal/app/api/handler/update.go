package handler

import (
	"errors"
	"net/http"

	"github.com/federalbyro/encryptServer/internal/app/service"
	"github.com/gin-gonic/gin"
)


type RefreshDTO struct{
	OldRefreshToken string `json:"refresh_token"`
}


func (h *TokenHandler) Update(c *gin.Context) {
	nowClientIP := c.ClientIP()
	ctx :=c.Request.Context()
	var input RefreshDTO

	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	access,refresh,err := h.service.Refresh(ctx,nowClientIP,input.OldRefreshToken)
	
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTokenMalformed):
			c.JSON(http.StatusBadRequest,  gin.H{"error":"token malformed"})
		case errors.Is(err, service.ErrTokenMismatch):
			c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid refresh token"})
		case errors.Is(err, service.ErrTokenReused):
			c.JSON(http.StatusUnauthorized, gin.H{"error":"was reuse"})
		case errors.Is(err, service.ErrGenerateToken):
			c.JSON(http.StatusUnauthorized, gin.H{"error":"problem with generate in refresh"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error":"internal_error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})

}
