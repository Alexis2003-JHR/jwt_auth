package services

import (
	"auth/internal/core/models"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *authService) GenerateRefreshToken(ctx context.Context, username string) (string, error) {
	expirationTime := time.Now().Add(15 * 24 * time.Hour)
	claims := models.Claims{
		Username:  username,
		TokenType: "refresh_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(s.privateKey)
}
