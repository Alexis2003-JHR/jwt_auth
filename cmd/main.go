package main

import (
	"fmt"
	"log"

	"auth/internal/core/handlers"
	"auth/internal/core/services"
	"auth/internal/utils"

	"github.com/gin-gonic/gin"
)

var (
	port = 8013
)

func main() {
	r := gin.Default()

	privateKey, err := utils.LoadPrivateKey("private.key")
	if err != nil {
		log.Print("Error al obtener private.key ", err)
	}
	authService := services.NewUserService(privateKey)

	authHandlers := handlers.NewAuthHandler(authService)

	r.POST("/login", authHandlers.Login)
	r.GET("/refresh-token", authHandlers.RefreshToken)

	r.Run(fmt.Sprintf(":%d", port))
}
