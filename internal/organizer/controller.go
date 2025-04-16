package organizer

import (
	"net/http"
	"net/url"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/generated/organizers"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
	template "github.com/readytotouch/readytotouch/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

var (
	testFullPublicUntil = time.Date(2025, time.April, 25, 0, 0, 0, 0, time.UTC)
)

type (
	companyAliasURI struct {
		CompanyAlias string `uri:"company_alias" binding:"required"`
	}

	vacancyURI struct {
		VacancyID int64 `uri:"vacancy_id" binding:"required"`
	}
)

type Controller struct {
	userRepository                  *postgres.UserRepository
	userFeatureWaitlistRepository   *postgres.UserFeatureWaitlistRepository
	featureViewStatsRepository      *postgres.FeatureViewStatsRepository
	userFavoriteCompanyRepository   *postgres.UserFavoriteCompanyRepository
	userFavoriteVacancyRepository   *postgres.UserFavoriteVacancyRepository
	companyViewDailyStatsRepository *postgres.CompanyViewDailyStatsRepository
	vacancyViewStatsRepository      *postgres.VacancyViewStatsRepository
}

func NewController(userRepository *postgres.UserRepository, userFeatureWaitlistRepository *postgres.UserFeatureWaitlistRepository, featureViewStatsRepository *postgres.FeatureViewStatsRepository, userFavoriteCompanyRepository *postgres.UserFavoriteCompanyRepository, userFavoriteVacancyRepository *postgres.UserFavoriteVacancyRepository, companyViewDailyStatsRepository *postgres.CompanyViewDailyStatsRepository, vacancyViewStatsRepository *postgres.VacancyViewStatsRepository) *Controller {
	return &Controller{userRepository: userRepository, userFeatureWaitlistRepository: userFeatureWaitlistRepository, featureViewStatsRepository: featureViewStatsRepository, userFavoriteCompanyRepository: userFavoriteCompanyRepository, userFavoriteVacancyRepository: userFavoriteVacancyRepository, companyViewDailyStatsRepository: companyViewDailyStatsRepository, vacancyViewStatsRepository: vacancyViewStatsRepository}
}

func (c *Controller) Index(ctx *gin.Context) {
	headerProfiles, err := c.getHeaderProfiles(ctx, domain.ContextGetUserID(ctx))
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	socialUserProfiles, err := c.userRepository.SocialUserProfiles(ctx, domain.RegistrationHistoryLimit)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.OrganizersOnline(headerProfiles, socialUserProfiles)))
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
	var (
		source    = db.CloneCompanies()
		companies = make([]domain.CompanyProfile, 0, len(source))
	)

	for _, company := range source {
		if len(company.Languages[domain.Go].Vacancies) == 0 {
			continue
		}

		companies = append(companies, company)
	}

	content := template.OrganizerStatic(companies, db.UkrainianUniversities())

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) CompaniesV1(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	organizerFeature, ok := c.organizerFeature(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	if c.softAuth(ctx, authUserID, organizerFeature.Organizer.Language) {
		ctx.Redirect(http.StatusFound, "/organizers/golang/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	companies := c.companies(organizerFeature)

	// Never first
	slices.Reverse(companies)

	userCompanyFavoriteMap, err := c.userFavoriteCompanyRepository.GetMap(ctx, authUserID, nil)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCompaniesV1(
		organizerFeature,
		headerProfiles,
		companies,
		db.UkrainianUniversities(),
		db.CzechUniversities(),
		userCompanyFavoriteMap,
		c.redirect(organizerFeature.Path),
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) CompaniesV2(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	organizerFeature, ok := c.organizerFeature(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	if c.softAuth(ctx, authUserID, organizerFeature.Organizer.Language) {
		ctx.Redirect(http.StatusFound, "/organizers/golang/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	companies := c.companies(organizerFeature)

	// Never first
	slices.Reverse(companies)

	userCompanyFavoriteMap, err := c.userFavoriteCompanyRepository.GetMap(ctx, authUserID, nil)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCompaniesV2(
		organizerFeature,
		headerProfiles,
		companies,
		db.UkrainianUniversities(),
		db.CzechUniversities(),
		userCompanyFavoriteMap,
		c.redirect(organizerFeature.Path),
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) CompanyV1(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	var (
		uri companyAliasURI
	)

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("Company alias is required"))

		return
	}

	var (
		featurePath = c.trimCompanyAlias(ctx)
	)

	// Redirect to lowercase company alias
	{
		redirectAlias := strings.ToLower(uri.CompanyAlias)
		if uri.CompanyAlias != redirectAlias {
			ctx.Redirect(http.StatusFound, featurePath+"/"+redirectAlias)

			return
		}
	}

	company, originalAlias, ok := c.findCompany(ctx, uri.CompanyAlias)
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Company not found"))

		return
	}

	if originalAlias != "" {
		ctx.Redirect(http.StatusFound, featurePath+"/"+originalAlias)

		return
	}

	company.ID = organizers.CompanyAliasToCodeMap[company.LinkedInProfile.Alias]
	if company.ID == 0 {
		// make generate-organizers

		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Company not found"))

		return
	}

	if company.Type == "" {
		company.Type = organizers.ToCompanyType(company.LinkedInProfile.Alias)
	}
	company.Logo = domain.CompanyLogo{
		V0: organizers.CompanyAliasToLogoMapV0[company.LinkedInProfile.Alias],
		V1: organizers.CompanyAliasToLogoMapV1[company.LinkedInProfile.Alias],
	}

	organizerFeature, ok := c.organizerFeature(featurePath)
	if !ok {
		// Should be unreachable
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	// Should be optimized
	userCompanyFavoriteMap, err := c.userFavoriteCompanyRepository.GetMap(ctx, authUserID, []int64{company.ID})
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	err = c.companyViewDailyStatsRepository.Upsert(ctx, company.ID, time.Now().UTC())
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCompanyV1(
		organizerFeature,
		headerProfiles,
		company,
		db.UkrainianUniversities(),
		db.CzechUniversities(),
		userCompanyFavoriteMap[company.ID],
		c.companyStats(ctx, company.ID),
		c.redirect(organizerFeature.Path+"/"+uri.CompanyAlias),
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) CompanyV2(ctx *gin.Context) {
	var (
		authUserID  = domain.ContextGetUserID(ctx)
		featurePath = c.trimCompanyAlias(ctx)
	)

	organizerFeature, ok := c.organizerFeature(featurePath)
	if !ok {
		// Should be unreachable
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	if c.softAuth(ctx, authUserID, organizerFeature.Organizer.Language) {
		ctx.Redirect(http.StatusFound, "/organizers/golang/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	var (
		uri companyAliasURI
	)

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("Company alias is required"))

		return
	}

	// Redirect to lowercase company alias
	{
		redirectAlias := strings.ToLower(uri.CompanyAlias)
		if uri.CompanyAlias != redirectAlias {
			ctx.Redirect(http.StatusFound, featurePath+"/"+redirectAlias)

			return
		}
	}

	company, originalAlias, ok := c.findCompany(ctx, uri.CompanyAlias)
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Company not found"))

		return
	}

	if originalAlias != "" {
		ctx.Redirect(http.StatusFound, featurePath+"/"+originalAlias)

		return
	}

	company.ID = organizers.CompanyAliasToCodeMap[company.LinkedInProfile.Alias]
	if company.ID == 0 {
		// make generate-organizers

		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Company not found"))

		return
	}

	if company.Type == "" {
		company.Type = organizers.ToCompanyType(company.LinkedInProfile.Alias)
	}
	company.Logo = domain.CompanyLogo{
		V0: organizers.CompanyAliasToLogoMapV0[company.LinkedInProfile.Alias],
		V1: organizers.CompanyAliasToLogoMapV1[company.LinkedInProfile.Alias],
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	// Should be optimized
	userCompanyFavoriteMap, err := c.userFavoriteCompanyRepository.GetMap(ctx, authUserID, []int64{company.ID})
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	err = c.companyViewDailyStatsRepository.Upsert(ctx, company.ID, time.Now().UTC())
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	var (
		preparedCompany   = c.toPrepareCompany(company)
		vacancies         = company.Languages[organizerFeature.Organizer.Language].Vacancies
		preparedVacancies = make([]domain.PreparedVacancy, 0, len(vacancies))
		vacancyIDs        = make([]int64, 0, len(vacancies))
	)

	for _, vacancy := range vacancies {
		id, ok := organizers.VacancyUrlMap[vacancy.URL]

		if ok {
			preparedVacancies = append(preparedVacancies, domain.PreparedVacancy{
				ID:                   id,
				Company:              preparedCompany,
				Title:                vacancy.Title,
				ShortDescription:     vacancy.ShortDescription,
				SwitchingOpportunity: vacancy.SwitchingOpportunity,
				URL:                  vacancy.URL,
				Date:                 vacancy.Date,
				WithSalary:           vacancy.WithSalary,
				Remote:               vacancy.Remote,
			})
			vacancyIDs = append(vacancyIDs, id)
		}
	}

	userVacancyFavoriteMap, err := c.userFavoriteVacancyRepository.GetMap(ctx, authUserID, vacancyIDs)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	month := time.Now().UTC().Truncate(time.Hour*24).AddDate(0, -1, 0)
	vacancyMonthlyViewsMap, err := c.vacancyViewStatsRepository.Stats(ctx, vacancyIDs, month)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	sort.Slice(preparedVacancies, func(i, j int) bool {
		return preparedVacancies[i].Date.After(preparedVacancies[j].Date)
	})

	content := template.OrganizersCompanyV2(
		organizerFeature,
		headerProfiles,
		company,
		preparedVacancies,
		db.UkrainianUniversities(),
		db.CzechUniversities(),
		userCompanyFavoriteMap[company.ID],
		userVacancyFavoriteMap,
		vacancyMonthlyViewsMap,
		c.companyStats(ctx, company.ID),
		c.redirect(organizerFeature.Path+"/"+uri.CompanyAlias),
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) Vacancies(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	organizerFeature, ok := c.organizerFeature(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	if c.softAuth(ctx, authUserID, organizerFeature.Organizer.Language) {
		ctx.Redirect(http.StatusFound, "/organizers/golang/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	companies := c.companies(organizerFeature)

	var (
		preparedVacancies = make([]domain.PreparedVacancy, 0, 1024)
		vacancyIDs        = make([]int64, 0, 1024)
	)

	for _, company := range companies {
		var (
			preparedCompany = c.toPrepareCompany(company)
			vacancies       = company.Languages[organizerFeature.Organizer.Language].Vacancies
		)

		for _, vacancy := range vacancies {
			id, ok := organizers.VacancyUrlMap[vacancy.URL]

			if ok {
				preparedVacancies = append(preparedVacancies, domain.PreparedVacancy{
					ID:                   id,
					Company:              preparedCompany,
					Title:                vacancy.Title,
					ShortDescription:     vacancy.ShortDescription,
					SwitchingOpportunity: vacancy.SwitchingOpportunity,
					URL:                  vacancy.URL,
					Date:                 vacancy.Date,
					WithSalary:           vacancy.WithSalary,
					Remote:               vacancy.Remote,
				})
				vacancyIDs = append(vacancyIDs, id)
			}
		}
	}

	userVacancyFavoriteMap, err := c.userFavoriteVacancyRepository.GetMap(ctx, authUserID, vacancyIDs)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	month := time.Now().UTC().Truncate(time.Hour*24).AddDate(0, -1, 0)
	vacancyMonthlyViewsMap, err := c.vacancyViewStatsRepository.Stats(ctx, vacancyIDs, month)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	sort.Slice(preparedVacancies, func(i, j int) bool {
		return preparedVacancies[i].Date.After(preparedVacancies[j].Date)
	})

	content := template.OrganizersVacanciesV2(
		organizerFeature,
		headerProfiles,
		companies,
		preparedVacancies,
		userVacancyFavoriteMap,
		vacancyMonthlyViewsMap,
		c.redirect(organizerFeature.Path),
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) trimCompanyAlias(ctx *gin.Context) (result string) {
	result = ctx.FullPath()
	result = strings.TrimSuffix(result, "/:company_alias")
	result = strings.TrimSuffix(result, "/:company_alias/v1")
	result = strings.TrimSuffix(result, "/:company_alias/v2")

	return result
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

func (c *Controller) GolangCommunities(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	if authUserID == 0 {
		ctx.Redirect(http.StatusFound, "/organizers/golang/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCommunitiesGolang(domain.OrganizerGolang, headerProfiles)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) RustCommunities(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCommunitiesRust(domain.OrganizerRust, headerProfiles)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) ZigCommunities(ctx *gin.Context) {
	c.TODO(ctx)
}

func (c *Controller) ScalaCommunities(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCommunitiesScala(domain.OrganizerScala, headerProfiles)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) ElixirCommunities(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCommunitiesElixir(domain.OrganizerElixir, headerProfiles)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) ClojureCommunities(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := template.OrganizersCommunitiesClojure(domain.OrganizerClojure, headerProfiles)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) VacancyRedirect(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
		now        = time.Now().UTC()
	)

	if authUserID == 0 && now.After(testFullPublicUntil) {
		ctx.Redirect(http.StatusFound, "/organizers/golang/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	var (
		uri vacancyURI
	)

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("Vacancy ID is required"))

		return
	}

	vacancyExternalURL, ok := organizers.VacancyIdMap[uri.VacancyID]
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Vacancy not found"))

		return
	}

	if authUserID > 0 {
		err := c.vacancyViewStatsRepository.Upsert(ctx, uri.VacancyID, authUserID, now)
		if err != nil {
			// @TODO logging

			// NOP, continue
		}
	}

	ctx.Redirect(http.StatusFound, vacancyExternalURL)
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

func (c *Controller) CompanyViewStats(ctx *gin.Context) {
	type (
		companyURI struct {
			CompanyID int64 `uri:"company_id" binding:"required"`
		}
	)

	var (
		uri companyURI
	)

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	var (
		to   = time.Now().UTC().Truncate(24 * time.Hour)
		from = to.AddDate(0, -1, 0)
	)

	viewsDailyStats, err := c.companyViewDailyStatsRepository.DailyCountStats(ctx, uri.CompanyID, from, to)
	if err != nil {
		// @TODO logging

		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(), // Yes, we are leaking the error message to the client, it's fine for now
		})
		return
	}

	ctx.JSON(http.StatusOK, viewsDailyStats)
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

func (c *Controller) FavoriteVacancy(ctx *gin.Context) {
	type (
		favoriteVacancyURI struct {
			VacancyID int64 `uri:"vacancy_id" binding:"required"`
		}
		favoriteVacancyRequestBody struct {
			Favorite bool `json:"favorite"`
		}
	)

	var (
		uri  favoriteVacancyURI
		body favoriteVacancyRequestBody
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

	err = c.userFavoriteVacancyRepository.Upsert(ctx, authUserID, uri.VacancyID, body.Favorite, time.Now().UTC())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(), // Yes, we are leaking the error message to the client, it's fine for now
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *Controller) UnsafeCompanies(ctx *gin.Context) {
	companies := db.CloneCompanies()

	result := make([]domain.UnsafeCompanyResponse, len(companies))
	for i, company := range companies {
		result[i] = domain.UnsafeCompanyResponse{
			ID:     company.LinkedInProfile.ID,
			Alias:  company.LinkedInProfile.Alias,
			Name:   company.LinkedInProfile.Name,
			Ignore: company.Ignore,
		}
	}

	ctx.JSON(http.StatusOK, &domain.UnsafeCompaniesResponse{
		Companies: result,
	})
}

func (c *Controller) UnsafeVacancies(ctx *gin.Context) {
	companies := db.CloneCompanies()

	result := make([]domain.UnsafeVacancyResponse, 0, 1024)
	for _, company := range companies {
		for _, language := range company.Languages {
			for _, vacancy := range language.Vacancies {
				result = append(result, domain.UnsafeVacancyResponse{
					URL:  vacancy.URL,
					Date: vacancy.Date,
				})
			}
		}
	}

	ctx.JSON(http.StatusOK, &domain.UnsafeVacanciesResponse{
		Vacancies: result,
	})
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
	path = strings.TrimSuffix(path, "/v1")
	path = strings.TrimSuffix(path, "/v2")

	featurePathMap := map[string]domain.OrganizerFeature{
		"/organizers/golang/companies": {
			Organizer: domain.OrganizerGolang,
			Feature:   dbs.FeatureWaitOrganizerGolangCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/golang/jobs": {
			Organizer: domain.OrganizerGolang,
			Feature:   dbs.FeatureWaitOrganizerGolangVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/organizers/rust/companies": {
			Organizer: domain.OrganizerRust,
			Feature:   dbs.FeatureWaitOrganizerRustCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/rust/jobs": {
			Organizer: domain.OrganizerRust,
			Feature:   dbs.FeatureWaitOrganizerRustVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/organizers/zig/companies": {
			Organizer: domain.OrganizerZig,
			Feature:   dbs.FeatureWaitOrganizerZigCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/zig/jobs": {
			Organizer: domain.OrganizerZig,
			Feature:   dbs.FeatureWaitOrganizerZigVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/organizers/scala/companies": {
			Organizer: domain.OrganizerScala,
			Feature:   dbs.FeatureWaitOrganizerScalaCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/scala/jobs": {
			Organizer: domain.OrganizerScala,
			Feature:   dbs.FeatureWaitOrganizerScalaVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/organizers/elixir/companies": {
			Organizer: domain.OrganizerElixir,
			Feature:   dbs.FeatureWaitOrganizerElixirCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/elixir/jobs": {
			Organizer: domain.OrganizerElixir,
			Feature:   dbs.FeatureWaitOrganizerElixirVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/organizers/clojure/companies": {
			Organizer: domain.OrganizerClojure,
			Feature:   dbs.FeatureWaitOrganizerClojureCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/organizers/clojure/jobs": {
			Organizer: domain.OrganizerClojure,
			Feature:   dbs.FeatureWaitOrganizerClojureVacancies,
			Path:      path,
			Title:     "Jobs",
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

func (c *Controller) companyStats(ctx *gin.Context, companyID int64) template.CompanyStats {
	var (
		to   = time.Now().UTC().Truncate(24 * time.Hour)
		from = to.AddDate(0, -1, 0)
	)

	totalViews, lastMonthViews, err := c.companyViewDailyStatsRepository.Stats(ctx, companyID, from)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	totalFavorites, lastMonthFavorites, err := c.userFavoriteCompanyRepository.Stats(ctx, companyID, from)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	return template.CompanyStats{
		TotalViews:         totalViews,
		LastMonthViews:     lastMonthViews,
		TotalFavorites:     totalFavorites,
		LastMonthFavorites: lastMonthFavorites,
	}
}

func (c *Controller) getHeaderProfiles(ctx *gin.Context, userID int64) ([]domain.SocialProviderUser, error) {
	if userID > 0 {
		return c.userRepository.SocialUserProfilesByUser(ctx, userID)
	}

	return nil, nil
}

func (c *Controller) companies(organizerFeature domain.OrganizerFeature) []domain.CompanyProfile {
	var (
		source    = db.CloneCompanies()
		companies = make([]domain.CompanyProfile, 0, len(source))
	)
	for _, company := range source {
		if company.Ignore {
			continue
		}

		company.ID = organizers.CompanyAliasToCodeMap[company.LinkedInProfile.Alias]
		if company.ID == 0 {
			// make generate-organizers

			continue
		}

		if company.Type == "" {
			company.Type = organizers.ToCompanyType(company.LinkedInProfile.Alias)
		}
		company.Logo = domain.CompanyLogo{
			V0: organizers.CompanyAliasToLogoMapV0[company.LinkedInProfile.Alias],
			V1: organizers.CompanyAliasToLogoMapV1[company.LinkedInProfile.Alias],
		}

		language := organizerFeature.Organizer.Language

		vacancies := company.Languages[language].Vacancies

		// Show companies only if they have vacancies or are Rust Foundation members
		if len(vacancies) == 0 && !(language == domain.Rust && company.RustFoundationMember) {
			continue
		}

		company.Remote = company.Remote || c.anyRemoteVacancy(vacancies)

		companies = append(companies, company)
	}
	return companies
}

func (c *Controller) findCompany(ctx *gin.Context, alias string) (domain.CompanyProfile, string, bool) {
	companies := db.CloneCompanies()

	// Yes, we are leaking the database implementation to the controller, it's fine for now
	// Yes, we use linear search, it's fine for now

	for _, company := range companies {
		if company.LinkedInProfile.Alias == alias {
			return company, "", true
		}
	}

	for _, company := range companies {
		for _, previousAlias := range company.LinkedInProfile.PreviousAliases {
			if previousAlias == alias {
				return company, company.LinkedInProfile.Alias, true
			}
		}
	}

	return domain.CompanyProfile{}, "", false
}

func (c *Controller) toPrepareCompany(company domain.CompanyProfile) domain.PreparedCompany {
	return domain.PreparedCompany{
		ID:                        company.ID,
		Type:                      company.Type,
		Logo:                      company.Logo,
		Name:                      company.Name,
		Alias:                     company.LinkedInProfile.Alias,
		Industries:                company.Industries,
		HasEmployeesFromCountries: company.HasEmployeesFromCountries,
		RustFoundationMember:      company.RustFoundationMember,
	}
}

// DataPopulationLists will be removed in the future.
func (c *Controller) DataPopulationLists(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationLists()))
}

// DataPopulationCompaniesCareersAndAboutAndBlog will be removed in the future.
func (c *Controller) DataPopulationCompaniesCareersAndAboutAndBlog(ctx *gin.Context) {
	var (
		companies = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			if c.skipSmallCompany(company) {
				return false
			}

			return company.Careers == "" || company.About == "" || company.Blog == ""
		})
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesCareersAndAboutAndBlog(companies, "Populate Careers & About & Blog")))
}

// DataPopulationCompaniesLinkedIn will be removed in the future.
func (c *Controller) DataPopulationCompaniesLinkedIn(ctx *gin.Context) {
	var (
		companies = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			if c.skipSmallCompany(company) {
				return false
			}

			return (company.LinkedInProfile.ID == 0 && len(company.LinkedInProfile.IDs) == 0) ||
				company.LinkedInProfile.Alias == "" ||
				company.LinkedInProfile.Name == "" ||
				company.LinkedInProfile.Followers == "" ||
				company.LinkedInProfile.Employees == "" ||
				company.LinkedInProfile.AssociatedMembers == ""
		})
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesLinkedIn(companies, "Populate LinkedIn")))
}

// DataPopulationCompaniesGitHub will be removed in the future.
func (c *Controller) DataPopulationCompaniesGitHub(ctx *gin.Context) {
	var (
		companies = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			if c.skipSmallCompany(company) {
				return false
			}

			return company.GitHubProfile.Followers == ""
		})
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesGitHub(companies, "Populate GitHub")))
}

// DataPopulationCompaniesGlassdoor will be removed in the future.
func (c *Controller) DataPopulationCompaniesGlassdoor(ctx *gin.Context) {
	var (
		companies = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			if c.skipSmallCompany(company) {
				return false
			}

			return company.GlassdoorProfile.OverviewURL == "" ||
				company.GlassdoorProfile.ReviewsURL == "" ||
				company.GlassdoorProfile.JobsURL == ""
		})
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesGlassdoor(companies, "Populate Glassdoor")))
}

// DataPopulationCompaniesBlind will be removed in the future.
func (c *Controller) DataPopulationCompaniesBlind(ctx *gin.Context) {
	var (
		companies = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			if c.skipSmallCompany(company) {
				return false
			}

			return company.BlindProfile.Alias == ""
		})
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesBlind(companies, "Populate Blind")))
}

// DataPopulationCompaniesIndeed will be removed in the future.
func (c *Controller) DataPopulationCompaniesIndeed(ctx *gin.Context) {
	var (
		companies = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			if c.skipSmallCompany(company) {
				return false
			}

			return company.IndeedProfile.Alias == ""
		})
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesIndeed(companies, "Populate Indeed")))
}

// DataPopulationCompaniesLevelsFyi will be removed in the future.
func (c *Controller) DataPopulationCompaniesLevelsFyi(ctx *gin.Context) {
	var (
		companies = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			if c.skipSmallCompany(company) {
				return false
			}

			return company.LevelsFyiProfile.Alias == "" ||
				company.LevelsFyiProfile.Employees == ""
		})
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesLevelsFyi(companies, "Populate Levels.fyi")))
}

// DataPopulationCompaniesLogo will be removed in the future.
func (c *Controller) DataPopulationCompaniesLogo(ctx *gin.Context) {
	const (
		keepLatest = 50
	)

	var (
		source = c.dataPopulationCompanies(func(company domain.CompanyProfile) bool {
			return true
		})
		result = make([]domain.CompanyProfile, 0, len(source))
	)

	// Never first
	slices.Reverse(source)

	for i, company := range source {
		if i < keepLatest {
			result = append(result, company)

			continue
		}

		if c.skipSmallCompany(company) {
			continue
		}

		// The existing logos must be re-checked
		result = append(result, company)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.DataPopulationCompaniesLogo(result, "Populate Logo")))
}

func (c *Controller) dataPopulationCompanies(match func(company domain.CompanyProfile) bool) []domain.CompanyProfile {
	var (
		source    = db.CloneCompanies()
		companies = make([]domain.CompanyProfile, 0, len(source))
	)
	for _, company := range source {
		if company.Ignore {
			continue
		}

		company.ID = organizers.CompanyAliasToCodeMap[company.LinkedInProfile.Alias]
		if company.ID == 0 {
			// make generate-organizers

			continue
		}

		if !match(company) {
			continue
		}

		if company.Type == "" {
			company.Type = organizers.ToCompanyType(company.LinkedInProfile.Alias)
		}
		company.Logo = domain.CompanyLogo{
			V0: organizers.CompanyAliasToLogoMapV0[company.LinkedInProfile.Alias],
			V1: organizers.CompanyAliasToLogoMapV1[company.LinkedInProfile.Alias],
		}

		if !c.hasVacancies(company) {
			continue
		}

		companies = append(companies, company)
	}
	return companies
}

func (c *Controller) hasVacancies(company domain.CompanyProfile) bool {
	for _, language := range company.Languages {
		if len(language.Vacancies) > 0 {
			return true
		}
	}

	return false
}

func (c *Controller) anyRemoteVacancy(vacancies []domain.Vacancy) bool {
	for _, vacancy := range vacancies {
		if vacancy.Remote {
			return true
		}
	}

	return false
}

func (c *Controller) softAuth(ctx *gin.Context, authUserID int64, language domain.Language) bool {
	if authUserID > 0 {
		return false
	}

	if c.random(language) {
		if strings.Contains(ctx.GetHeader("User-Agent"), "Googlebot") {
			return false
		}

		return true
	}

	return false
}

func (c *Controller) random(language domain.Language) bool {
	var (
		now    = time.Now()
		minute = now.Minute()
	)

	if now.Before(testFullPublicUntil) {
		return false
	}

	switch language {
	case domain.Go:
		return minute%2 == 0
	case domain.Rust:
		return minute%4 == 0
	case domain.Zig:
		return minute%16 == 0
	case domain.Scala:
		return minute%4 == 0
	case domain.Elixir:
		return minute%8 == 0
	case domain.Clojure:
		return minute%16 == 0
	case domain.Haskell:
		return minute%16 == 0
	}

	return false
}

func (c *Controller) skipSmallCompany(company domain.CompanyProfile) bool {
	switch company.LinkedInProfile.Employees {
	case "10K+", "5K-10K", "1K-5K", "501-1K", "201-500", "51-200":
		return false
	}

	return true
}
