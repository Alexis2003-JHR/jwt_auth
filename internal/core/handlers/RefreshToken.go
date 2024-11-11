package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	fmt.Println("RefreshToken works!")
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint works!"})
}
