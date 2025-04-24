package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TokenHandler) CreateTokens(c *gin.Context) {
	userId := c.Query("GUID")
	ctx := c.Request.Context()

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad GUID"})
		return
	}

	clientIP := c.ClientIP()

	accessToken, refreshToken, err := h.service.CreateTokens(ctx, userId, clientIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
