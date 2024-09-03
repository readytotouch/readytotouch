package organizer

import (
	"net/http"

	"github.com/readytotouch/readytotouch/internal/organizer/db"
	template "github.com/readytotouch/readytotouch/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GolangCompaniesUkraine(ctx *gin.Context) {
	content := template.Organizer(db.GolangCompanies(), db.UkrainianUniversities())

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) GolangCompanies(ctx *gin.Context) {
	content := template.Organizer(db.GolangCompanies(), db.UkrainianUniversities())

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) GolangVacancies(ctx *gin.Context) {
	content := "" // @TODO

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}
