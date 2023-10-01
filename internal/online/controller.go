package online

import (
	"net/http"
	"time"

	"github.com/readytotouch-yaaws/yaaws-go/internal/db/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"
	template "github.com/readytotouch-yaaws/yaaws-go/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userRepository   *postgres.UserRepository
	onlineRepository *postgres.OnlineRepository
}

func NewController(userRepository *postgres.UserRepository, onlineRepository *postgres.OnlineRepository) *Controller {
	return &Controller{userRepository: userRepository, onlineRepository: onlineRepository}
}

func (c *Controller) Index(ctx *gin.Context) {
	socialUserProfiles, err := c.userRepository.SocialUserProfiles(ctx, domain.RegistrationHistoryLimit)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.Online(socialUserProfiles)))
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
