package online

import (
	"net/http"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	onlineRepository *postgres.OnlineRepository
}

func NewController(onlineRepository *postgres.OnlineRepository) *Controller {
	return &Controller{onlineRepository: onlineRepository}
}

func (c *Controller) DailyCountStats(ctx *gin.Context) {
	var (
		to   = time.Now().UTC().Truncate(24 * time.Hour)
		from = to.AddDate(0, -1, 0)
	)
	stats, err := c.onlineRepository.DailyCountStats(ctx, from, to)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}
