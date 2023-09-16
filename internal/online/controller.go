package online

import (
	"net/http"
	"time"

	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository *Repository
}

func NewController(repository *Repository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) Index(ctx *gin.Context) {
	var (
		to   = time.Now().UTC().Truncate(24 * time.Hour)
		from = to.AddDate(0, 0, -1)
	)
	stats, err := c.repository.DailyCountStats(ctx, from, to)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})

	}

	ctx.JSON(http.StatusOK, stats)
}
