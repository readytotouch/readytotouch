package jwt

import (
	"fmt"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/env"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secretKey string
	ttl       int
}

func NewService(secretKey string, ttl int) *Service {
	return &Service{secretKey: secretKey, ttl: ttl}
}

func (s *Service) AddToken(ctx *gin.Context, user domain.JwtUser) error {
	claims := &ReadyToTouchClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.ttl) * time.Second)),
		},
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return err
	}

	ctx.SetCookie(s.getCookieName(), tokenString, s.ttl, "/", "", true, true)

	return nil
}

func (s *Service) ParseToken(ctx *gin.Context) (*domain.JwtUser, error) {
	tokenString, err := ctx.Cookie(s.getCookieName())
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &ReadyToTouchClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &token.Claims.(*ReadyToTouchClaims).User, nil
}

func (s *Service) RemoveToken(ctx *gin.Context) {
	ctx.SetCookie(s.getCookieName(), "", -1, "/", "", true, true)
}

func (s *Service) getCookieName() string {
	if env.Environment() == env.Development {
		return "dev_jwt_auth_token"
	}

	return "jwt_auth_token"
}
