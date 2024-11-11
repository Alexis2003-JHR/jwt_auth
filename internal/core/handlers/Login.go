package handlers

import (
	"auth/internal/core/models"
	"auth/internal/core/services"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
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

	privateKey := []byte("private.key")
	authService := services.NewUserService(privateKey)

	accessToken, err := authService.GenerateAccessToken(ctx, loginData.Username)
	if err != nil {
		log.Fatalf("Error generando access token: %v", err)
	}

	refreshToken, err := authService.GenerateRefreshToken(ctx)
	if err != nil {
		log.Fatalf("Error generando refresh token: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
