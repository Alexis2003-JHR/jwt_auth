package services

import (
	"auth/internal/core/models"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *authService) GenerateAccessToken(ctx context.Context, username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &models.Claims{
		Username:  username,
		TokenType: "access_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenSigned, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}
