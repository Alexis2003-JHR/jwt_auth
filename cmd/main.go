package main

import (
	"fmt"

	"auth/internal/core/handlers"
	"auth/internal/core/services"

	"github.com/gin-gonic/gin"
)

var (
	port = 8013
)

func main() {
	r := gin.Default()

	privateKey := []byte("private.key")
	authService := services.NewUserService(privateKey)

	authHandlers := handlers.NewAuthHandler(authService)

	r.POST("/login", authHandlers.Login)
	r.GET("/refresh-token", authHandlers.RefreshToken)

	r.Run(fmt.Sprintf(":%d", port))
}
