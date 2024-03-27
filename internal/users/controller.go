package users

import (
	"net/http"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository *postgres.UserRepository
}

func NewController(repository *postgres.UserRepository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) RegistrationDailyCountStats(ctx *gin.Context) {
	var (
		to   = time.Now().UTC().Truncate(24 * time.Hour)
		from = to.AddDate(0, -1, 0)
	)
	stats, err := c.repository.RegistrationDailyCountStats(ctx, from, to)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}
