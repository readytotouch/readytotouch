package jwt

import (
	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type ReadyToTouchClaims struct {
	jwt.RegisteredClaims
	User domain.JwtUser `json:"user"`
}
