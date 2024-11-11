package services

import "context"

type AuthService interface {
	GenerateAccessToken(ctx context.Context, username string) (string, error)
	GenerateRefreshToken(ctx context.Context) (string, error)
}

type authService struct {
	privateKey []byte
}

func NewUserService(privateKey []byte) AuthService {
	return &authService{privateKey: privateKey}
}
