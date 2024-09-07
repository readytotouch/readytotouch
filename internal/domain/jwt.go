package domain

import (
	"github.com/gin-gonic/gin"
)

const (
	userContextKey = "rtt.user"
)

type JwtService interface {
	AddToken(ctx *gin.Context, user JwtUser) error
	ParseToken(ctx *gin.Context) (*JwtUser, error)
	RemoveToken(ctx *gin.Context)
}

type JwtUser struct {
	ID int64 `json:"id"`
}

func ContextSetUser(ctx *gin.Context, user *JwtUser) {
	ctx.Set(userContextKey, user)
}

func ContextGetUser(ctx *gin.Context) (*JwtUser, bool) {
	value, exists := ctx.Get(userContextKey)
	if !exists {
		return nil, false
	}

	user, ok := value.(*JwtUser)

	return user, ok
}

func ContextGetUserID(ctx *gin.Context) int64 {
	user, ok := ContextGetUser(ctx)
	if ok {
		return user.ID
	}

	return 0
}
