package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token requerido"})
		return
	}

	username, exists := refreshTokens[request.RefreshToken]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	newAccessToken, err := h.service.GenerateAccessToken(ctx, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando nuevo access token"})
		return
	}

	newRefreshToken, err := h.service.GenerateRefreshToken(ctx, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando nuevo refresh token"})
	}

	delete(refreshTokens, request.RefreshToken)
	refreshTokens[newRefreshToken] = username

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
