package cac

import (
	"net/http"

	"github.com/readytotouch/readytotouch/internal/domain"
	template "github.com/readytotouch/readytotouch/internal/templates/v1"

	"github.com/gin-gonic/gin"
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

func (c *Controller) Poland(ctx *gin.Context) {
	c.Index(ctx)
}

func (c *Controller) Companies(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []domain.LinkedInProfileResponse{
		{
			ID:    97909464,
			Alias: "readytotouch",
			Name:  "ReadyToTouch",
		},
		{
			ID:    14136494,
			Alias: "dochq",
			Name:  "DocHQ",
		},
		{
			ID:    1441,
			Alias: "google",
			Name:  "Google",
		},
	})
}

func (c *Controller) AddCompany(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []domain.LinkedInProfileResponse{
		{
			ID:    97909464,
			Alias: "readytotouch",
			Name:  "ReadyToTouch",
		},
		{
			ID:    14136494,
			Alias: "dochq",
			Name:  "DocHQ",
		},
		{
			ID:    1441,
			Alias: "google",
			Name:  "Google",
		},
	})
}

func (c *Controller) DeleteCompany(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
