package organizer

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/protos/organizers"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
	template "github.com/readytotouch/readytotouch/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userRepository                *postgres.UserRepository
	userFeatureWaitlistRepository *postgres.UserFeatureWaitlistRepository
	featureViewStatsRepository    *postgres.FeatureViewStatsRepository
	userFavoriteCompanyRepository *postgres.UserFavoriteCompanyRepository
}

func NewController(userRepository *postgres.UserRepository, userFeatureWaitlistRepository *postgres.UserFeatureWaitlistRepository, featureViewStatsRepository *postgres.FeatureViewStatsRepository, userFavoriteCompanyRepository *postgres.UserFavoriteCompanyRepository) *Controller {
	return &Controller{userRepository: userRepository, userFeatureWaitlistRepository: userFeatureWaitlistRepository, featureViewStatsRepository: featureViewStatsRepository, userFavoriteCompanyRepository: userFavoriteCompanyRepository}
}

func (c *Controller) Welcome(ctx *gin.Context) {
	organizer, ok := c.organizer(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Organizer not found"))

		return
	}

	content := template.OrganizersWelcome(organizer, c.authQueryParams(ctx))

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) Main(ctx *gin.Context) {
	headerProfiles, err := c.getHeaderProfiles(ctx, domain.ContextGetUserID(ctx))
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersMain(headerProfiles, c.redirect("/organizers"))

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) GolangCompaniesUkraine(ctx *gin.Context) {
	content := template.OrganizerStatic(db.Companies(), db.UkrainianUniversities())

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) Companies(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	organizerFeature, ok := c.organizerFeature(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	var (
		source    = db.Companies()
		companies = make([]domain.Company, 0, len(source))
	)
	for _, company := range source {
		company.ID = organizers.CompanyAliasMap[company.LinkedInProfile.Alias]
		if company.ID == 0 {
			// make generate-company-code

			continue
		}

		// nil slice mean skip company for the language
		if company.Vacancies[organizerFeature.Organizer.Language] == nil {
			continue
		}

		companies = append(companies, company)
	}

	userCompanyFavoriteMap, err := c.userFavoriteCompanyRepository.GetMap(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCompanies(
		organizerFeature,
		headerProfiles,
		companies,
		userCompanyFavoriteMap,
		c.redirect(organizerFeature.Path),
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) Waitlist(ctx *gin.Context) {
	var (
		authUserID      = domain.ContextGetUserID(ctx)
		subscribedState = false
	)

	organizerFeature, ok := c.organizerFeature(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	err := c.featureViewStatsRepository.Upsert(ctx, organizerFeature.Feature, time.Now().UTC())
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	if authUserID > 0 {
		subscribedState, err = c.
			userFeatureWaitlistRepository.
			SubscribedState(ctx, authUserID, organizerFeature.Feature)

		if err != nil {
			// @TODO logging

			// NOP, continue
		}
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersWaitlist(organizerFeature, headerProfiles, c.redirect(organizerFeature.Path), subscribedState)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) TODO(ctx *gin.Context) {
	ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("TODO"))
}

func (c *Controller) WaitlistStats(ctx *gin.Context) {
	feature, done := c.parseFeatureFromReferer(ctx)
	if done {
		return
	}

	c.waitlistStats(ctx, feature)
}

func (c *Controller) WaitlistSubscribe(ctx *gin.Context) {
	type subscribeRequestBody struct {
		Active bool `json:"active"`
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

	var body subscribeRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	feature, done := c.parseFeatureFromReferer(ctx)
	if done {
		return
	}

	err := c.userFeatureWaitlistRepository.Upsert(ctx, authUserID, feature, body.Active, time.Now().UTC())
	if err != nil {
		// @TODO logging

		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(), // Yes, we are leaking the error message to the client, it's fine for now
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *Controller) FavoriteCompany(ctx *gin.Context) {
	type (
		favoriteCompanyURI struct {
			CompanyID int64 `uri:"company_id" binding:"required"`
		}
		favoriteCompanyRequestBody struct {
			Favorite bool `json:"favorite"`
		}
	)

	var (
		uri  favoriteCompanyURI
		body favoriteCompanyRequestBody
	)

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
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

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	err = c.userFavoriteCompanyRepository.Upsert(ctx, authUserID, uri.CompanyID, body.Favorite, time.Now().UTC())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(), // Yes, we are leaking the error message to the client, it's fine for now
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *Controller) parseFeatureFromReferer(ctx *gin.Context) (dbs.FeatureWait, bool) {
	var referer = ctx.Request.Referer()
	if referer == "" {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: "Referer is empty",
		})
		return "", true
	}

	refererURL, err := url.Parse(referer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return "", true
	}

	organizer, ok := c.organizerFeature(refererURL.Path)
	if !ok {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: "Feature not found",
		})
		return "", true
	}

	return organizer.Feature, false
}

func (c *Controller) organizerFeature(path string) (domain.OrganizerFeature, bool) {
	featurePathMap := map[string]domain.OrganizerFeature{
		"/organizers/golang/companies": {
			Organizer: domain.OrganizerGolang,
			Feature:   dbs.FeatureWaitOrganizerGolangCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/golang/vacancies": {
			Organizer: domain.OrganizerGolang,
			Feature:   dbs.FeatureWaitOrganizerGolangVacancies,
			Path:      path,
			Title:     "Vacancies",
		},
		"/organizers/rust/companies": {
			Organizer: domain.OrganizerRust,
			Feature:   dbs.FeatureWaitOrganizerRustCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/rust/vacancies": {
			Organizer: domain.OrganizerRust,
			Feature:   dbs.FeatureWaitOrganizerRustVacancies,
			Path:      path,
			Title:     "Vacancies",
		},
		"/organizers/zig/companies": {
			Organizer: domain.OrganizerZig,
			Feature:   dbs.FeatureWaitOrganizerZigCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/zig/vacancies": {
			Organizer: domain.OrganizerZig,
			Feature:   dbs.FeatureWaitOrganizerZigVacancies,
			Path:      path,
			Title:     "Vacancies",
		},
		"/organizers/scala/companies": {
			Organizer: domain.OrganizerScala,
			Feature:   dbs.FeatureWaitOrganizerScalaCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/scala/vacancies": {
			Organizer: domain.OrganizerScala,
			Feature:   dbs.FeatureWaitOrganizerScalaVacancies,
			Path:      path,
			Title:     "Vacancies",
		},
		"/organizers/elixir/companies": {
			Organizer: domain.OrganizerElixir,
			Feature:   dbs.FeatureWaitOrganizerElixirCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/elixir/vacancies": {
			Organizer: domain.OrganizerElixir,
			Feature:   dbs.FeatureWaitOrganizerElixirVacancies,
			Path:      path,
			Title:     "Vacancies",
		},
		"/organizers/clojure/companies": {
			Organizer: domain.OrganizerClojure,
			Feature:   dbs.FeatureWaitOrganizerClojureCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/clojure/vacancies": {
			Organizer: domain.OrganizerClojure,
			Feature:   dbs.FeatureWaitOrganizerClojureVacancies,
			Path:      path,
			Title:     "Vacancies",
		},
	}

	feature, ok := featurePathMap[path]

	return feature, ok
}

func (c *Controller) organizer(path string) (domain.Organizer, bool) {
	organizerPathMap := map[string]domain.Organizer{
		"/organizers/golang/welcome":  domain.OrganizerGolang,
		"/organizers/rust/welcome":    domain.OrganizerRust,
		"/organizers/zig/welcome":     domain.OrganizerZig,
		"/organizers/scala/welcome":   domain.OrganizerScala,
		"/organizers/elixir/welcome":  domain.OrganizerElixir,
		"/organizers/clojure/welcome": domain.OrganizerClojure,
	}

	feature, ok := organizerPathMap[path]

	return feature, ok
}

func (c *Controller) authQueryParams(ctx *gin.Context) string {
	redirect := ctx.Query("redirect")
	if strings.HasPrefix(redirect, "/") {
		return c.redirect(redirect)
	}

	return ""
}

func (c *Controller) redirect(redirect string) string {
	return "?" + url.Values{"redirect": []string{redirect}}.Encode()
}

func (c *Controller) waitlistStats(ctx *gin.Context, feature dbs.FeatureWait) {
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
) (*domain.UserFeatureWaitlistStats, error) {
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

	return &domain.UserFeatureWaitlistStats{
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

func (c *Controller) getHeaderProfiles(ctx *gin.Context, userID int64) ([]domain.SocialProviderUser, error) {
	if userID > 0 {
		return c.userRepository.SocialUserProfilesByUser(ctx, userID)
	}

	return nil, nil
}
