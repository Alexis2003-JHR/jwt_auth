package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("Login Works!")
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint works!"})
}
