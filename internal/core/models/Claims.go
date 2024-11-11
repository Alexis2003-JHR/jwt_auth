package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username  string `json:"username"`
	TokenType string `json:"token_type"`
	jwt.StandardClaims
}
