package cac

import (
	"net/http"
	"strings"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"
	template "github.com/readytotouch/readytotouch/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository
	service                         *Service
}

func NewController(userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository, service *Service) *Controller {
	return &Controller{userToLinkedInCompanyRepository: userToLinkedInCompanyRepository, service: service}
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

	c.companies(ctx, authUserID)
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

	var companyVanityName = strings.ToLower(strings.TrimSpace(body.Alias))
	if companyVanityName == "" {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: "Alias is required",
		})
		return
	}

	if err := c.service.Add(companyVanityName, authUserID, time.Now().UTC()); err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	c.companies(ctx, authUserID)
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
	if body.ID == 0 {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: "ID is required",
		})
		return
	}

	err := c.userToLinkedInCompanyRepository.Delete(ctx, authUserID, body.ID, time.Now().UTC())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *Controller) companies(ctx *gin.Context, authUserID int64) {
	companies, err := c.userToLinkedInCompanyRepository.List(ctx, authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
