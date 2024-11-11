package utils

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt"
)

func LoadPrivateKey(keyPath string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
