package cac

import (
	"fmt"
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
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)
	if authUserID == 0 {
		ctx.JSON(http.StatusUnauthorized, &domain.ErrorResponse{
			ErrorMessage: "Unauthorized",
		})
		return
	}

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
	type addRequestBody struct {
		Alias string `json:"alias"`
	}

	var (
		authUserID = domain.ContextGetUserID(ctx)
	)
	if authUserID == 0 {
		ctx.JSON(http.StatusUnauthorized, &domain.ErrorResponse{
			ErrorMessage: "Unauthorized",
		})
		return
	}

	var body addRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	fmt.Printf("Add company %s\n", body.Alias)

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
	type deleteRequestBody struct {
		ID int64 `json:"id"`
	}

	var (
		authUserID = domain.ContextGetUserID(ctx)
	)
	if authUserID == 0 {
		ctx.JSON(http.StatusUnauthorized, &domain.ErrorResponse{
			ErrorMessage: "Unauthorized",
		})
		return
	}

	var body deleteRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	fmt.Printf("Delete company %d\n", body.ID)

	ctx.JSON(http.StatusOK, nil)
}
