package handlers

import (
	"auth/internal/core/models"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var refreshTokens = make(map[string]string)

func (h *AuthHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var loginData models.LoginData
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if loginData.Username != "usuario" || loginData.Password != "contraseña" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales invalidas"})
		return
	}

	accessToken, err := h.service.GenerateAccessToken(ctx, loginData.Username)
	if err != nil {
		log.Fatalf("Error generando access token: %v", err)
	}

	refreshToken, err := h.service.GenerateRefreshToken(ctx, loginData.Username)
	if err != nil {
		log.Fatalf("Error generando refresh token: %v", err)
	}
	refreshTokens[refreshToken] = loginData.Username

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
