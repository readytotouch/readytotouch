package organizer

import (
	"encoding/xml"
	"fmt"
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
	testFullPublicSince = time.Date(2025, time.November, 1, 0, 0, 0, 0, time.UTC)
	testFullPublicUntil = time.Date(2026, time.January, 25, 0, 0, 0, 0, time.UTC)
	testStrictAuthUntil = time.Date(2025, time.December, 15, 0, 0, 0, 0, time.UTC)
)

func testStrictAuth(date time.Time) bool {
	return date.Before(testStrictAuthUntil)
}

func testFullPublic(date time.Time) bool {
	return date.After(testFullPublicSince) && date.Before(testFullPublicUntil)
}

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

func (c *Controller) IndexV1(ctx *gin.Context) {
	headerProfiles, err := c.getHeaderProfiles(ctx, domain.ContextGetUserID(ctx))
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	socialUserProfiles, err := c.userRepository.SocialUserProfiles(ctx, nil, domain.RegistrationHistoryLimit)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.IndexV1(headerProfiles, socialUserProfiles)))
}

func (c *Controller) IndexV2(ctx *gin.Context) {
	headerProfiles, err := c.getHeaderProfiles(ctx, domain.ContextGetUserID(ctx))
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	socialUserProfiles, err := c.userRepository.SocialUserProfiles(ctx, nil, domain.RegistrationHistoryLimit)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.OrganizersIndexV2(headerProfiles, socialUserProfiles)))
}

func (c *Controller) IndexV3(ctx *gin.Context) {
	headerProfiles, err := c.getHeaderProfiles(ctx, domain.ContextGetUserID(ctx))
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	socialUserProfiles, err := c.userRepository.SocialUserProfiles(
		ctx,
		[]dbs.SocialProvider{dbs.SocialProviderGithub},
		domain.RegistrationHistoryLimit,
	)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.OrganizersIndexV3(headerProfiles, socialUserProfiles)))
}

func (c *Controller) WelcomeV1(ctx *gin.Context) {
	organizer, ok := c.organizer(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Organizer not found"))

		return
	}

	content := template.OrganizersWelcomeV1(organizer, c.authQueryParams(ctx))

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) WelcomeV3(ctx *gin.Context) {
	organizer, ok := c.organizer(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Organizer not found"))

		return
	}

	content := template.OrganizersWelcomeV3(organizer, c.authQueryParams(ctx))

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) Organizers(ctx *gin.Context) {
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
	c.companiesAction(ctx, template.OrganizersCompaniesV1)
}

func (c *Controller) CompaniesV2(ctx *gin.Context) {
	c.companiesAction(ctx, template.OrganizersCompaniesV2)
}

func (c *Controller) CompaniesV3(ctx *gin.Context) {
	c.companiesAction(ctx, template.OrganizersCompaniesV3)
}

func (c *Controller) CompanyV1(ctx *gin.Context) {
	c.companyAction(ctx, func(organizerFeature domain.OrganizerFeature, headerProfiles []domain.SocialProviderUser, company domain.CompanyProfile, vacancies []domain.PreparedVacancy, ukrainianUniversities []domain.University, czechUniversities []domain.University, favorite bool, userVacancyFavoriteMap map[int64]bool, vacancyMonthlyViewsMap map[int64]int64, stats template.CompanyStats, authQueryParams string) string {
		return template.OrganizersCompanyV1(
			organizerFeature,
			headerProfiles,
			company,
			// vacancies,
			ukrainianUniversities,
			czechUniversities,
			favorite,
			// userVacancyFavoriteMap,
			// vacancyMonthlyViewsMap,
			stats,
			authQueryParams,
		)
	})
}

func (c *Controller) CompanyV2(ctx *gin.Context) {
	c.companyAction(ctx, func(organizerFeature domain.OrganizerFeature, headerProfiles []domain.SocialProviderUser, company domain.CompanyProfile, vacancies []domain.PreparedVacancy, ukrainianUniversities []domain.University, czechUniversities []domain.University, favorite bool, userVacancyFavoriteMap map[int64]bool, vacancyMonthlyViewsMap map[int64]int64, stats template.CompanyStats, authQueryParams string) string {
		return template.OrganizersCompanyV2(
			organizerFeature,
			headerProfiles,
			company,
			vacancies,
			ukrainianUniversities,
			czechUniversities,
			favorite,
			userVacancyFavoriteMap,
			vacancyMonthlyViewsMap,
			stats,
			authQueryParams,
		)
	})
}

func (c *Controller) CompanyV3(ctx *gin.Context) {
	c.companyAction(ctx, func(organizerFeature domain.OrganizerFeature, headerProfiles []domain.SocialProviderUser, company domain.CompanyProfile, vacancies []domain.PreparedVacancy, ukrainianUniversities []domain.University, czechUniversities []domain.University, favorite bool, userVacancyFavoriteMap map[int64]bool, vacancyMonthlyViewsMap map[int64]int64, stats template.CompanyStats, authQueryParams string) string {
		return template.OrganizersCompanyV3(
			organizerFeature,
			headerProfiles,
			company,
			vacancies,
			ukrainianUniversities,
			czechUniversities,
			favorite,
			userVacancyFavoriteMap,
			vacancyMonthlyViewsMap,
			stats,
			authQueryParams,
		)
	})
}

func (c *Controller) companiesAction(
	ctx *gin.Context,
	render func(
		organizerFeature domain.OrganizerFeature,
		headerProfiles []domain.SocialProviderUser,
		companies []domain.CompanyProfile,
		ukrainianUniversities []domain.University,
		czechUniversities []domain.University,
		userCompanyFavoriteMap map[int64]bool,
		authQueryParams string,
	) string,
) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	organizerFeature, ok := c.organizerFeature(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	if c.softAuthRedirect(ctx, authUserID, organizerFeature.Organizer.Language) {
		ctx.Redirect(http.StatusFound, "/"+organizerFeature.Organizer.Alias+"/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	companies := c.companies(organizerFeature.Organizer.Language)

	c.sortCompanies(companies)

	userCompanyFavoriteMap, err := c.userFavoriteCompanyRepository.GetMap(ctx, authUserID, nil)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	content := render(
		organizerFeature,
		headerProfiles,
		c.pinnedFirst(companies),
		db.UkrainianUniversities(),
		db.CzechUniversities(),
		userCompanyFavoriteMap,
		c.redirect(organizerFeature.Path),
	)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}

func (c *Controller) companyAction(
	ctx *gin.Context,
	render func(
		organizerFeature domain.OrganizerFeature,
		headerProfiles []domain.SocialProviderUser,
		company domain.CompanyProfile,
		vacancies []domain.PreparedVacancy,
		ukrainianUniversities []domain.University,
		czechUniversities []domain.University,
		favorite bool,
		userVacancyFavoriteMap map[int64]bool,
		vacancyMonthlyViewsMap map[int64]int64,
		stats template.CompanyStats,
		authQueryParams string,
	) string,
) {
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

	if c.softAuthRedirect(ctx, authUserID, organizerFeature.Organizer.Language) {
		ctx.Redirect(http.StatusFound, "/"+organizerFeature.Organizer.Alias+"/welcome"+c.redirect(ctx.Request.URL.Path))

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
		V2: organizers.CompanyAliasToLogoMapV2[company.LinkedInProfile.Alias],
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
				Location:             vacancy.Location,
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

	company.LatestVacancyDate = c.maxLanguageDate(vacancies)

	content := render(
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

func (c *Controller) JobsV2(ctx *gin.Context) {
	c.jobsAction(ctx, template.OrganizersVacanciesV2)
}

func (c *Controller) JobsV3(ctx *gin.Context) {
	c.jobsAction(ctx, template.OrganizersVacanciesV3)
}

func (c *Controller) jobsAction(
	ctx *gin.Context,
	render func(
		organizerFeature domain.OrganizerFeature,
		headerProfiles []domain.SocialProviderUser,
		companies []domain.CompanyProfile,
		vacancies []domain.PreparedVacancy,
		userVacancyFavoriteMap map[int64]bool,
		vacancyMonthlyViewsMap map[int64]int64,
		authQueryParams string,
	) string,
) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)

	organizerFeature, ok := c.organizerFeature(ctx.FullPath())
	if !ok {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("Feature not found"))

		return
	}

	if c.softAuthRedirect(ctx, authUserID, organizerFeature.Organizer.Language) {
		ctx.Redirect(http.StatusFound, "/"+organizerFeature.Organizer.Alias+"/welcome"+c.redirect(ctx.Request.URL.Path))

		return
	}

	headerProfiles, err := c.getHeaderProfiles(ctx, authUserID)
	if err != nil {
		// @TODO logging

		// NOP, continue
	}

	companies := c.companies(organizerFeature.Organizer.Language)

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
					Location:             vacancy.Location,
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

	content := render(
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
	result = strings.TrimSuffix(result, "/:company_alias/v3")

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

	if authUserID == 0 && (testStrictAuth(now) || !testFullPublic(now)) {
		ctx.Redirect(http.StatusFound, "/golang/welcome"+c.redirect(ctx.Request.URL.Path))

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

func (c *Controller) SitemapCompanies(organizer domain.Organizer) gin.HandlerFunc {
	type Url struct {
		Loc     string `xml:"loc"`
		LastMod string `xml:"lastmod"`
	}

	type UrlSet struct {
		XMLName xml.Name `xml:"urlset"`
		Xmlns   string   `xml:"xmlns,attr"`
		Urls    []Url    `xml:"url"`
	}

	return func(ctx *gin.Context) {
		companies := c.companies(organizer.Language)

		urls := make([]Url, len(companies))

		for i, company := range companies {
			urls[i] = Url{
				Loc:     fmt.Sprintf("https://%s/%s/companies/%s", ctx.Request.Host, organizer.Alias, company.LinkedInProfile.Alias),
				LastMod: c.sitemapCompanyLastModified(company.Languages[organizer.Language].Vacancies),
			}
		}

		ctx.Header("Content-Type", "application/xml")
		ctx.XML(http.StatusOK, UrlSet{
			Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
			Urls:  urls,
		})
	}
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
		var languages []domain.UnsafeCompanyLanguageStats

		for i, language := range company.Languages {
			if len(language.Vacancies) == 0 {
				continue
			}

			languages = append(languages, domain.UnsafeCompanyLanguageStats{
				Language:       domain.Language(i).String(),
				MaxVacancyDate: c.maxLanguageDate(language.Vacancies),
			})
		}

		result[i] = domain.UnsafeCompanyResponse{
			ID:                 company.LinkedInProfile.ID,
			Alias:              company.LinkedInProfile.Alias,
			Name:               company.LinkedInProfile.Name,
			GitHubProfileAlias: company.GitHubProfile.Login,
			OttaProfileAlias:   company.OttaProfileSlug,
			Ignore:             company.Ignore,
			Languages:          languages,
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

func (c *Controller) maxLanguageDate(vacancies []domain.Vacancy) time.Time {
	var maxLanguageDate time.Time
	for _, vacancy := range vacancies {
		if vacancy.Date.After(maxLanguageDate) {
			maxLanguageDate = vacancy.Date
		}
	}
	return maxLanguageDate
}

func (c *Controller) maxDate(languages domain.Languages) time.Time {
	var maxLanguageDate time.Time
	for _, language := range languages {
		for _, vacancy := range language.Vacancies {
			if vacancy.Date.After(maxLanguageDate) {
				maxLanguageDate = vacancy.Date
			}
		}
	}
	return maxLanguageDate
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
	path = strings.TrimSuffix(path, "/v3")

	featurePathMap := map[string]domain.OrganizerFeature{
		"/golang/companies": {
			Organizer: domain.OrganizerGolang,
			Feature:   dbs.FeatureWaitOrganizerGolangCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/golang/jobs": {
			Organizer: domain.OrganizerGolang,
			Feature:   dbs.FeatureWaitOrganizerGolangVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/rust/companies": {
			Organizer: domain.OrganizerRust,
			Feature:   dbs.FeatureWaitOrganizerRustCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/rust/jobs": {
			Organizer: domain.OrganizerRust,
			Feature:   dbs.FeatureWaitOrganizerRustVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/zig/companies": {
			Organizer: domain.OrganizerZig,
			Feature:   dbs.FeatureWaitOrganizerZigCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/zig/jobs": {
			Organizer: domain.OrganizerZig,
			Feature:   dbs.FeatureWaitOrganizerZigVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/scala/companies": {
			Organizer: domain.OrganizerScala,
			Feature:   dbs.FeatureWaitOrganizerScalaCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/scala/jobs": {
			Organizer: domain.OrganizerScala,
			Feature:   dbs.FeatureWaitOrganizerScalaVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/elixir/companies": {
			Organizer: domain.OrganizerElixir,
			Feature:   dbs.FeatureWaitOrganizerElixirCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/elixir/jobs": {
			Organizer: domain.OrganizerElixir,
			Feature:   dbs.FeatureWaitOrganizerElixirVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/erlang/companies": {
			Organizer: domain.OrganizerErlang,
			Feature:   dbs.FeatureWaitOrganizerErlangCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/erlang/jobs": {
			Organizer: domain.OrganizerErlang,
			Feature:   dbs.FeatureWaitOrganizerErlangVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/clojure/companies": {
			Organizer: domain.OrganizerClojure,
			Feature:   dbs.FeatureWaitOrganizerClojureCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/clojure/jobs": {
			Organizer: domain.OrganizerClojure,
			Feature:   dbs.FeatureWaitOrganizerClojureVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/haskell/companies": {
			Organizer: domain.OrganizerHaskell,
			Feature:   dbs.FeatureWaitOrganizerHaskellCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/haskell/jobs": {
			Organizer: domain.OrganizerHaskell,
			Feature:   dbs.FeatureWaitOrganizerHaskellVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/fsharp/companies": {
			Organizer: domain.OrganizerFSharp,
			Feature:   dbs.FeatureWaitOrganizerFsharpCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/fsharp/jobs": {
			Organizer: domain.OrganizerFSharp,
			Feature:   dbs.FeatureWaitOrganizerFsharpVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/ocaml/companies": {
			Organizer: domain.OrganizerOCaml,
			Feature:   dbs.FeatureWaitOrganizerOcamlCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/ocaml/jobs": {
			Organizer: domain.OrganizerOCaml,
			Feature:   dbs.FeatureWaitOrganizerOcamlVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/gleam/companies": {
			Organizer: domain.OrganizerGleam,
			Feature:   dbs.FeatureWaitOrganizerGleamCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/gleam/jobs": {
			Organizer: domain.OrganizerGleam,
			Feature:   dbs.FeatureWaitOrganizerGleamVacancies,
			Path:      path,
			Title:     "Jobs",
		},
		"/mojo/companies": {
			Organizer: domain.OrganizerMojo,
			Feature:   dbs.FeatureWaitOrganizerMojoCompanies,
			Path:      path,
			Title:     "Companies",
		},
		"/mojo/jobs": {
			Organizer: domain.OrganizerMojo,
			Feature:   dbs.FeatureWaitOrganizerMojoVacancies,
			Path:      path,
			Title:     "Jobs",
		},
	}

	feature, ok := featurePathMap[path]

	return feature, ok
}

func (c *Controller) organizer(path string) (domain.Organizer, bool) {
	organizerPathMap := map[string]domain.Organizer{
		"/golang/welcome":  domain.OrganizerGolang,
		"/rust/welcome":    domain.OrganizerRust,
		"/zig/welcome":     domain.OrganizerZig,
		"/scala/welcome":   domain.OrganizerScala,
		"/elixir/welcome":  domain.OrganizerElixir,
		"/erlang/welcome":  domain.OrganizerErlang,
		"/clojure/welcome": domain.OrganizerClojure,
		"/haskell/welcome": domain.OrganizerHaskell,
		"/fsharp/welcome":  domain.OrganizerFSharp,
		"/ocaml/welcome":   domain.OrganizerOCaml,
		"/gleam/welcome":   domain.OrganizerGleam,
		"/mojo/welcome":    domain.OrganizerMojo,
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

func (c *Controller) companies(language domain.Language) []domain.CompanyProfile {
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
			V2: organizers.CompanyAliasToLogoMapV2[company.LinkedInProfile.Alias],
		}

		vacancies := company.Languages[language].Vacancies

		// Show companies only if they have vacancies or are Rust Foundation members
		if len(vacancies) == 0 && !(language == domain.Rust && company.RustFoundationMember) {
			continue
		}

		company.Remote = company.Remote || c.anyRemoteVacancy(vacancies)
		company.LatestVacancyDate = c.maxLanguageDate(vacancies)

		companies = append(companies, company)
	}
	return companies
}

func (c *Controller) sortCompanies(companies []domain.CompanyProfile) {
	slices.SortFunc(companies, func(a, b domain.CompanyProfile) int {
		if a.LatestVacancyDate.After(b.LatestVacancyDate) {
			return -1
		}

		if a.LatestVacancyDate.Before(b.LatestVacancyDate) {
			return 1
		}

		return int(b.ID - a.ID) // Sort by ID descending
	})
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

	for i, company := range source {
		if i < keepLatest {
			result = append(result, company)

			continue
		}

		if c.skipSmallCompany(company) {
			continue
		}

		company.LatestVacancyDate = c.maxDate(company.Languages)

		// The existing logos must be re-checked
		result = append(result, company)
	}

	c.sortCompanies(result)

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
			V2: organizers.CompanyAliasToLogoMapV2[company.LinkedInProfile.Alias],
		}

		if !c.hasVacancies(company) {
			continue
		}

		company.LatestVacancyDate = c.maxDate(company.Languages)

		companies = append(companies, company)
	}

	c.sortCompanies(companies)

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

func (c *Controller) softAuthRedirect(ctx *gin.Context, authUserID int64, language domain.Language) bool {
	if authUserID > 0 {
		return false
	}

	if ctx.Request.Host == "localhost" {
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

	if testStrictAuth(now) {
		return true
	}

	if testFullPublic(now) {
		return false
	}

	const (
		often = 8  // once in 8 minutes
		rare  = 16 // once in 16 minutes
	)

	switch language {
	case domain.Go:
		return minute%often == 0
	case domain.Rust:
		return minute%often == 0
	case domain.Zig:
		return minute%rare == 0
	case domain.Scala:
		return minute%rare == 0
	case domain.Elixir:
		return minute%often == 0
	case domain.Clojure:
		return minute%rare == 0
	case domain.Haskell:
		return minute%rare == 0
	}

	return false
}

func (c *Controller) pinnedFirst(companies []domain.CompanyProfile) []domain.CompanyProfile {
	var (
		pinnedCompanies   = make([]domain.CompanyProfile, 0, len(companies))
		unpinnedCompanies = make([]domain.CompanyProfile, 0, len(companies))
		now               = time.Now()
	)
	for _, company := range companies {
		if company.PinnedUntil.After(now) && company.LinkedInProfile.Verified {
			pinnedCompanies = append(pinnedCompanies, company)
		} else {
			unpinnedCompanies = append(unpinnedCompanies, company)
		}
	}

	return append(pinnedCompanies, unpinnedCompanies...)
}

func (c *Controller) skipSmallCompany(company domain.CompanyProfile) bool {
	switch company.LinkedInProfile.Employees {
	case "10K+", "5K-10K", "1K-5K", "501-1K", "201-500", "51-200":
		return false
	}

	return true
}

func (c *Controller) sitemapCompanyLastModified(vacancies []domain.Vacancy) string {
	if len(vacancies) == 0 {
		return "2025-01-01"
	}

	var maxVacancyDate time.Time

	for _, vacancy := range vacancies {
		if vacancy.Date.After(maxVacancyDate) {
			maxVacancyDate = vacancy.Date
		}
	}

	return maxVacancyDate.Format(time.DateOnly)
}
