package organizer

import (
	"net/http"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
	template "github.com/readytotouch/readytotouch/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userFeatureWaitlistRepository *postgres.UserFeatureWaitlistRepository
	featureViewStatsRepository    *postgres.FeatureViewStatsRepository
}

func NewController(userFeatureWaitlistRepository *postgres.UserFeatureWaitlistRepository, featureViewStatsRepository *postgres.FeatureViewStatsRepository) *Controller {
	return &Controller{userFeatureWaitlistRepository: userFeatureWaitlistRepository, featureViewStatsRepository: featureViewStatsRepository}
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
	var (
		authUserID      = domain.ContextGetUserID(ctx)
		subscribedState = false
	)

	err := c.featureViewStatsRepository.Upsert(ctx, dbs.FeatureWaitOrganizerGolangVacancies, time.Now().UTC())
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	if authUserID > 0 {
		subscribedState, err = c.
			userFeatureWaitlistRepository.
			SubscribedState(ctx, authUserID, dbs.FeatureWaitOrganizerGolangVacancies)

		if err != nil {
			// @TODO logging

			// NOP, continue
		}
	}

	content := template.OrganizerVacanciesSubscribe(subscribedState)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) GolangCompaniesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerGolangCompanies)
}

func (c *Controller) GolangVacanciesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerGolangVacancies)
}

func (c *Controller) RustCompaniesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerRustCompanies)
}

func (c *Controller) RustVacanciesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerRustVacancies)
}

func (c *Controller) ZigCompaniesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerZigCompanies)
}

func (c *Controller) ZigVacanciesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerZigVacancies)
}

func (c *Controller) ScalaCompaniesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerScalaCompanies)
}

func (c *Controller) ScalaVacanciesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerScalaVacancies)
}

func (c *Controller) ElixirCompaniesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerElixirCompanies)
}

func (c *Controller) ElixirVacanciesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerElixirVacancies)
}

func (c *Controller) ClojureCompaniesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerClojureCompanies)
}

func (c *Controller) ClojureVacanciesDailyCountStats(ctx *gin.Context) {
	c.dailyCountStats(ctx, dbs.FeatureWaitOrganizerClojureVacancies)
}

func (c *Controller) dailyCountStats(ctx *gin.Context, feature dbs.FeatureWait) {
	stats, err := c.fetchStats(ctx, feature)
	if err != nil {
		// @TODO logging

		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(), // Yes, we are leaking the error message to the client, it's fine for now
		})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

func (c *Controller) fetchStats(
	ctx *gin.Context,
	feature dbs.FeatureWait,
) (*domain.UserFeatureWaitlistPageStats, error) {
	var (
		to   = time.Now().UTC().Truncate(24 * time.Hour)
		from = to.AddDate(0, -1, 0)
	)

	subscribersDailyStats, err := c.userFeatureWaitlistRepository.DailyCountStats(ctx, feature, from, to)
	if err != nil {
		return nil, err
	}

	subscribersTotalCount, err := c.userFeatureWaitlistRepository.TotalCount(ctx, feature)
	if err != nil {
		return nil, err
	}

	viewsDailyStats, err := c.featureViewStatsRepository.DailyCountStats(ctx, feature, from, to)
	if err != nil {
		return nil, err
	}

	viewsTotalCount, err := c.featureViewStatsRepository.TotalCount(ctx, feature)
	if err != nil {
		return nil, err
	}

	return &domain.UserFeatureWaitlistPageStats{
		Subscribers: domain.ChartStats{
			DailyStats: subscribersDailyStats,
			TotalCount: subscribersTotalCount,
		},
		Views: domain.ChartStats{
			DailyStats: viewsDailyStats,
			TotalCount: viewsTotalCount,
		},
	}, nil
}
