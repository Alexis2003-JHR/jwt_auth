package main

import (
	"fmt"

	"auth/internal/core/handlers"

	"github.com/gin-gonic/gin"
)

var (
	port = 8013
)

func main() {
	r := gin.Default()

	r.POST("/login", handlers.Login)
	r.GET("/refresh-token", handlers.RefreshToken)

	r.Run(fmt.Sprintf(":%d", port))
}
