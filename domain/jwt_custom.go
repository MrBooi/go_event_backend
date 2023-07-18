package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Uuid string `json:"uuid"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
