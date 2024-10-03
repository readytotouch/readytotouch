package cac

import (
	"github.com/gin-gonic/gin"
	"net/http"

	template "github.com/readytotouch/readytotouch/internal/templates/v1"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.WipCompaniesAndConnections()))
}

func (c *Controller) Ukraine(ctx *gin.Context) {
	c.Index(ctx)
}

func (c *Controller) Czechia(ctx *gin.Context) {
	c.Index(ctx)
}
