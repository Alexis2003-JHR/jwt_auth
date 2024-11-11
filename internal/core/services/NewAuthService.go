package services

import (
	"context"
	"crypto/rsa"
)

type AuthService interface {
	GenerateAccessToken(ctx context.Context, username string) (string, error)
	GenerateRefreshToken(ctx context.Context, username string) (string, error)
}

type authService struct {
	privateKey *rsa.PrivateKey
}

func NewUserService(privateKey *rsa.PrivateKey) AuthService {
	return &authService{privateKey: privateKey}
}
