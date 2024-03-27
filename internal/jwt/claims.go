package jwt

import (
	"github.com/readytotouch/readytotouch/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type ReadyToTouchClaims struct {
	jwt.RegisteredClaims
	User domain.JwtUser `json:"user"`
}
