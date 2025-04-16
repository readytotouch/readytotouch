package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/bitbucket"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/gitlab"

	"github.com/readytotouch/readytotouch/internal/auth"
	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/db/redis"
	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/env"
	"github.com/readytotouch/readytotouch/internal/server"

	pkgCAC "github.com/readytotouch/readytotouch/internal/cac"
	pkgJWT "github.com/readytotouch/readytotouch/internal/jwt"
	pkgLinkedIn "github.com/readytotouch/readytotouch/internal/linkedin"
	pkgBitbucket "github.com/readytotouch/readytotouch/internal/oauth-providers/bitbucket"
	pkgGitHub "github.com/readytotouch/readytotouch/internal/oauth-providers/github"
	pkgGitLab "github.com/readytotouch/readytotouch/internal/oauth-providers/gitlab"
	pkgOnline "github.com/readytotouch/readytotouch/internal/online"
	pkgOrganizer "github.com/readytotouch/readytotouch/internal/organizer"
	pkgUsers "github.com/readytotouch/readytotouch/internal/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	var (
		dsn                = env.Required("POSTGRES_DSN")
		jwtSecretKey       = env.Required("JWT_SECRET_KEY")
		linkedinOAuthToken = os.Getenv("LINKEDIN_OAUTH2_TOKEN")
	)

	pgConnection := postgres.MustConnection(dsn)
	defer pgConnection.Close()

	database := postgres.MustDatabase(pgConnection)
	defer database.Queries().Close()

	redisClient := redis.MustClient(context.Background(), "redis:6379")

	var (
		userRepository                  = postgres.NewUserRepository(database)
		onlineRepository                = postgres.NewOnlineRepository(database)
		userFeatureWaitlistRepository   = postgres.NewUserFeatureWaitlistRepository(database)
		featureViewStatsRepository      = postgres.NewFeatureViewStatsRepository(database)
		userFavoriteCompanyRepository   = postgres.NewUserFavoriteCompanyRepository(database)
		userFavoriteVacancyRepository   = postgres.NewUserFavoriteVacancyRepository(database)
		companyViewDailyStatsRepository = postgres.NewCompanyViewDailyStatsRepository(database)
		vacancyViewStatsRepository      = postgres.NewVacancyViewStatsRepository(database)
		userToLinkedInCompanyRepository = postgres.NewUserToLinkedInCompanyRepository(database)
	)

	var (
		onlineRedisStorage    = redis.NewHashOnlineStorage(redisClient)
		onlinePostgresStorage = postgres.NewBatchUpsertOnlineStorage(database)
	)

	var (
		githubOAuthProvider = pkgGitHub.NewGithubOAuthProvider(&oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
			Endpoint:     github.Endpoint,
			Scopes:       []string{"user:email"},
		})

		gitlabOAuthProvider = pkgGitLab.NewGitlabOAuthProvider(&oauth2.Config{
			ClientID:     os.Getenv("GITLAB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITLAB_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GITLAB_REDIRECT_URL"),
			Endpoint:     gitlab.Endpoint,
			Scopes:       []string{"read_user"},
		})

		bitbucketOAuthProvider = pkgBitbucket.NewBitbucketOAuthProvider(&oauth2.Config{
			ClientID:     os.Getenv("BITBUCKET_CLIENT_ID"),
			ClientSecret: os.Getenv("BITBUCKET_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("BITBUCKET_REDIRECT_URL"),
			Endpoint:     bitbucket.Endpoint,
			Scopes:       nil, // Scopes are defined on the client/consumer instance.
		})
	)

	var (
		jwtService = pkgJWT.NewService(jwtSecretKey, 2*30*24*3600)
	)

	var (
		authController = auth.NewController(
			userRepository,
			githubOAuthProvider,
			gitlabOAuthProvider,
			bitbucketOAuthProvider,
			jwtService,
		)
		userController      = pkgUsers.NewController(userRepository)
		onlineController    = pkgOnline.NewController(userRepository, onlineRepository)
		organizerController = pkgOrganizer.NewController(
			userRepository,
			userFeatureWaitlistRepository,
			featureViewStatsRepository,
			userFavoriteCompanyRepository,
			userFavoriteVacancyRepository,
			companyViewDailyStatsRepository,
			vacancyViewStatsRepository,
		)
		cacController = pkgCAC.NewController(
			userRepository,
			userToLinkedInCompanyRepository,
			pkgCAC.NewService(
				userToLinkedInCompanyRepository,
				pkgLinkedIn.NewClient(linkedinOAuthToken),
			),
		)
	)

	r := gin.New()
	r.Use(redirectFromWWW())
	r.Use(cors.Default())
	r.Use(func(ctx *gin.Context) {
		user, err := jwtService.ParseToken(ctx)
		if err != nil {
			// NOP,

			ctx.Next()
		}

		domain.ContextSetUser(ctx, user)

		ctx.Next()
	})
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.
		Use(func(ctx *gin.Context) {
			userID := domain.ContextGetUserID(ctx)
			if userID > 0 {
				err := onlineRedisStorage.Store(context.Background(), domain.UserOnlinePair{
					UserID:    userID,
					Timestamp: time.Now().Unix(),
				})
				if err != nil {
					// NOP
				}
			}

			ctx.Next()
		})
	r.GET("/", organizerController.Index)

	r.GET("/api/v1/users/registration/stats/daily.json", userController.RegistrationDailyCountStats)
	r.GET("/api/v1/users/online/stats/daily.json", onlineController.DailyCountStats)

	r.GET("/organizers", organizerController.Main)

	r.GET("/organizers/golang/welcome", organizerController.Welcome)
	r.GET("/organizers/rust/welcome", organizerController.Welcome)
	r.GET("/organizers/zig/welcome", organizerController.Welcome)
	r.GET("/organizers/scala/welcome", organizerController.Welcome)
	r.GET("/organizers/elixir/welcome", organizerController.Welcome)
	r.GET("/organizers/clojure/welcome", organizerController.Welcome)

	r.GET("/organizers/golang/companies/ukraine", organizerController.GolangCompaniesUkraine)
	r.GET("/organizers/golang/companies", organizerController.CompaniesV2)
	r.GET("/organizers/golang/companies/v1", organizerController.CompaniesV1)
	r.GET("/organizers/golang/companies/v2", organizerController.CompaniesV2)
	r.GET("/organizers/golang/companies/:company_alias", organizerController.CompanyV2)
	r.GET("/organizers/golang/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/organizers/golang/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/organizers/golang/jobs", organizerController.Vacancies)
	r.GET("/organizers/golang/communities", organizerController.GolangCommunities)
	r.GET("/organizers/rust/companies", organizerController.CompaniesV2)
	r.GET("/organizers/rust/companies/v1", organizerController.CompaniesV1)
	r.GET("/organizers/rust/companies/v2", organizerController.CompaniesV2)
	r.GET("/organizers/rust/companies/:company_alias", organizerController.CompanyV2)
	r.GET("/organizers/rust/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/organizers/rust/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/organizers/rust/jobs", organizerController.Vacancies)
	r.GET("/organizers/rust/communities", organizerController.RustCommunities)
	r.GET("/organizers/zig/companies", organizerController.Waitlist)
	r.GET("/organizers/zig/companies/v1", organizerController.Waitlist)
	r.GET("/organizers/zig/companies/v2", organizerController.Waitlist)
	r.GET("/organizers/zig/companies/:company_alias", organizerController.CompanyV2)
	r.GET("/organizers/zig/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/organizers/zig/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/organizers/zig/jobs", organizerController.Waitlist)
	r.GET("/organizers/zig/communities", organizerController.ZigCommunities)
	r.GET("/organizers/scala/companies", organizerController.CompaniesV2)
	r.GET("/organizers/scala/companies/v1", organizerController.CompaniesV1)
	r.GET("/organizers/scala/companies/v2", organizerController.CompaniesV2)
	r.GET("/organizers/scala/companies/:company_alias", organizerController.CompanyV2)
	r.GET("/organizers/scala/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/organizers/scala/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/organizers/scala/jobs", organizerController.Vacancies)
	r.GET("/organizers/scala/communities", organizerController.ScalaCommunities)
	r.GET("/organizers/elixir/companies", organizerController.CompaniesV2)
	r.GET("/organizers/elixir/companies/v1", organizerController.CompaniesV1)
	r.GET("/organizers/elixir/companies/v2", organizerController.CompaniesV2)
	r.GET("/organizers/elixir/companies/:company_alias", organizerController.CompanyV2)
	r.GET("/organizers/elixir/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/organizers/elixir/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/organizers/elixir/jobs", organizerController.Vacancies)
	r.GET("/organizers/elixir/communities", organizerController.ElixirCommunities)
	r.GET("/organizers/clojure/companies", organizerController.CompaniesV2)
	r.GET("/organizers/clojure/companies/v1", organizerController.CompaniesV1)
	r.GET("/organizers/clojure/companies/v2", organizerController.CompaniesV2)
	r.GET("/organizers/clojure/companies/:company_alias", organizerController.CompanyV2)
	r.GET("/organizers/clojure/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/organizers/clojure/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/organizers/clojure/jobs", organizerController.Vacancies)
	r.GET("/organizers/clojure/communities", organizerController.ClojureCommunities)
	r.GET("/organizers/v/:vacancy_id", organizerController.VacancyRedirect)

	r.GET("/organizers/golang", found("/organizers/golang/companies"))
	r.GET("/organizers/rust", found("/organizers/rust/companies"))
	r.GET("/organizers/zig", found("/organizers/zig/companies"))
	r.GET("/organizers/scala", found("/organizers/scala/companies"))
	r.GET("/organizers/elixir", found("/organizers/elixir/companies"))
	r.GET("/organizers/clojure", found("/organizers/clojure/companies"))

	r.GET("/organizers/golang/vacancies", found("/organizers/golang/jobs"))
	r.GET("/organizers/rust/vacancies", found("/organizers/rust/jobs"))
	r.GET("/organizers/zig/vacancies", found("/organizers/zig/jobs"))
	r.GET("/organizers/scala/vacancies", found("/organizers/scala/jobs"))
	r.GET("/organizers/elixir/vacancies", found("/organizers/elixir/jobs"))
	r.GET("/organizers/clojure/vacancies", found("/organizers/clojure/jobs"))

	/*
		Will be removed in the future.
	*/
	r.GET("/private/unstable/wip/organizers/data-population-lists", organizerController.DataPopulationLists)
	r.GET("/private/unstable/wip/organizers/data-population-lists/careers-and-about-and-blog", organizerController.DataPopulationCompaniesCareersAndAboutAndBlog)
	r.GET("/private/unstable/wip/organizers/data-population-lists/linkedin", organizerController.DataPopulationCompaniesLinkedIn)
	r.GET("/private/unstable/wip/organizers/data-population-lists/github", organizerController.DataPopulationCompaniesGitHub)
	r.GET("/private/unstable/wip/organizers/data-population-lists/glassdoor", organizerController.DataPopulationCompaniesGlassdoor)
	r.GET("/private/unstable/wip/organizers/data-population-lists/blind", organizerController.DataPopulationCompaniesBlind)
	r.GET("/private/unstable/wip/organizers/data-population-lists/indeed", organizerController.DataPopulationCompaniesIndeed)
	r.GET("/private/unstable/wip/organizers/data-population-lists/levels", organizerController.DataPopulationCompaniesLevelsFyi)
	r.GET("/private/unstable/wip/organizers/data-population-lists/logo", organizerController.DataPopulationCompaniesLogo)

	r.GET("/api/v1/features/auto/waitlist/stats.json", organizerController.WaitlistStats)
	r.POST("/api/v1/features/auto/waitlist/subscribe.json", organizerController.WaitlistSubscribe)
	r.PATCH("/api/v1/companies/:company_id/favorite.json", organizerController.FavoriteCompany)
	r.GET("/api/v1/companies/:company_id/views/stats/daily.json", organizerController.CompanyViewStats)
	r.PATCH("/api/v1/vacancies/:vacancy_id/favorite.json", organizerController.FavoriteVacancy)

	r.
		GET("/auth/github", authController.GithubRedirect).
		GET("/auth/gitlab", authController.GitlabRedirect).
		GET("/auth/bitbucket", authController.BitbucketRedirect).
		GET("/auth/github/callback", authController.GithubCallback).
		GET("/auth/gitlab/callback", authController.GitlabCallback).
		GET("/auth/bitbucket/callback", authController.BitbucketCallback).
		GET("/logout", authController.Logout)

	r.
		// WIP
		GET("/companies-and-connections", found("/companies-and-connections/worldwide")).
		GET("/companies-and-connections/worldwide", cacController.Worldwide).
		GET("/companies-and-connections/ukraine", cacController.Ukraine).
		GET("/companies-and-connections/brazil", cacController.Brazil).
		GET("/wip/companies-and-connections", found("/companies-and-connections/worldwide")).
		GET("/wip/companies-and-connections/ukraine", found("/companies-and-connections/ukraine")).
		GET("/wip/companies-and-connections/brazil", found("/companies-and-connections/brazil")).
		GET("/api/v1/companies-and-connections/companies.json", cacController.Companies).
		POST("/api/v1/companies-and-connections/companies.json", cacController.AddCompany).
		DELETE("/api/v1/companies-and-connections/companies.json", cacController.DeleteCompany)

	r.
		// Unsafe API endpoints that can be changed without any notice.
		GET("/api/v1/unsafe/companies.json", organizerController.UnsafeCompanies).
		GET("/api/v1/unsafe/vacancies.json", organizerController.UnsafeVacancies)

	r.GET("/analytics", found("https://plausible.io/readytotouch.com"))
	r.GET("/plausible", found("https://plausible.io/readytotouch.com"))
	r.GET("/similarweb", found("https://www.similarweb.com/website/readytotouch.com/"))

	r.
		StaticFile("/design", "./public/design/online.html").
		StaticFile("/design/online", "./public/design/online.html").
		StaticFile("/design/online-auth", "./public/design/online-auth.html").

		// Organizers
		StaticFile("/design/organizers", "./public/design/organizer-main-page-auth.html").
		StaticFile("/design/organizers-auth", "./public/design/organizer-main-page.html").
		GET("/design/organizers/:language/welcome", s("./public/design/organizer-welcome.html")).
		GET("/design/organizers/:language/companies/ukraine", s("./public/design/golang-companies-organizer.html")).
		GET("/design/organizers/:language/companies/:company_alias", s("./public/design/organizer-statistics.html")).
		GET("/design/organizers/:language/companies/:company_alias/v1", s("./public/design-v1/organizer-statistics.html")).
		GET("/design/organizers/:language/companies/:company_alias/v2", s("./public/design/organizer-statistics.html")).
		GET("/design/organizers/:language/companies", s("./public/design/organizer-companies.html")).
		GET("/design/organizers/:language/companies/v1", s("./public/design-v1/organizer-companies.html")).
		GET("/design/organizers/:language/companies/v2", s("./public/design/organizer-companies.html")).
		GET("/design/organizers/:language/jobs", s("./public/design/organizer-vacancies.html")).
		GET("/design/organizers/:language/jobs/subscribe", s("./public/design/organizer-vacancies-subscribe.html")).
		GET("/design/organizers/:language/jobs/unsubscribe", s("./public/design/organizer-vacancies-unsubscribe.html")).
		GET("/design/organizers/:language/vacancies", s("./public/design/organizer-vacancies.html")).
		GET("/design/organizers/:language/vacancies/subscribe", s("./public/design/organizer-vacancies-subscribe.html")).
		GET("/design/organizers/:language/vacancies/unsubscribe", s("./public/design/organizer-vacancies-unsubscribe.html")).
		GET("/design/organizers/golang/communities", s("./public/design/organizer-go-communities.html")).
		GET("/design/organizers/rust/communities", s("./public/design/organizer-rust-communities.html")).
		GET("/design/organizers/scala/communities", s("./public/design/organizer-scala-communities.html")).
		GET("/design/organizers/elixir/communities", s("./public/design/organizer-elixir-communities.html")).
		GET("/design/organizers/clojure/communities", s("./public/design/organizer-clojure-communities.html")).
		GET("/design/organizers/:language/communities", s("./public/design/organizer-go-communities.html")).

		// Companies and connections
		StaticFile("/design/companies-and-connections", "./public/design/connections.html").
		StaticFile("/design/companies-and-connections/worldwide", "./public/design/connections.html").
		StaticFile("/design/companies-and-connections/ukraine", "./public/design/connections.html").
		StaticFile("/design/companies-and-connections/brazil", "./public/design/connections.html").

		// Design from ChatGPT
		StaticFile("/design/wip/companies-and-connections", "./public/chatgpt-design/companies-and-connections.html").
		StaticFile("/design/wip/companies-and-connections/ukraine", "./public/chatgpt-design/companies-and-connections.html").
		StaticFile("/design/wip/companies-and-connections/brazil", "./public/chatgpt-design/companies-and-connections.html")

	r.
		// Assets
		Static("/assets/images", "./public/assets/images").
		Static("/assets/fonts", "./public/assets/fonts").
		Static("/assets/js", "./public/assets/js").
		Static("/assets/unstable/logos/", "./public/logos-v0/adapted/").
		Static("/assets/unstable/logos-v0/", "./public/logos-v0/adapted/").
		Static("/assets/unstable/logos-v1/", "./public/logos-v1/adapted/").

		// Favicons
		StaticFile("/android-icon-144x144.png", "./public/android-icon-144x144.png").
		StaticFile("/android-icon-192x192.png", "./public/android-icon-192x192.png").
		StaticFile("/android-icon-36x36.png", "./public/android-icon-36x36.png").
		StaticFile("/android-icon-48x48.png", "./public/android-icon-48x48.png").
		StaticFile("/android-icon-72x72.png", "./public/android-icon-72x72.png").
		StaticFile("/android-icon-96x96.png", "./public/android-icon-96x96.png").
		StaticFile("/apple-icon-114x114.png", "./public/apple-icon-114x114.png").
		StaticFile("/apple-icon-120x120.png", "./public/apple-icon-120x120.png").
		StaticFile("/apple-icon-144x144.png", "./public/apple-icon-144x144.png").
		StaticFile("/apple-icon-152x152.png", "./public/apple-icon-152x152.png").
		StaticFile("/apple-icon-180x180.png", "./public/apple-icon-180x180.png").
		StaticFile("/apple-icon-57x57.png", "./public/apple-icon-57x57.png").
		StaticFile("/apple-icon-60x60.png", "./public/apple-icon-60x60.png").
		StaticFile("/apple-icon-72x72.png", "./public/apple-icon-72x72.png").
		StaticFile("/apple-icon-76x76.png", "./public/apple-icon-76x76.png").
		StaticFile("/apple-icon.png", "./public/apple-icon.png").
		StaticFile("/apple-icon-precomposed.png", "./public/apple-icon-precomposed.png").
		StaticFile("/favicon-16x16.png", "./public/favicon-16x16.png").
		StaticFile("/favicon-32x32.png", "./public/favicon-32x32.png").
		StaticFile("/favicon-96x96.png", "./public/favicon-96x96.png").
		StaticFile("/favicon.ico", "./public/favicon.ico").
		StaticFile("/ms-icon-144x144.png", "./public/ms-icon-144x144.png").
		StaticFile("/ms-icon-150x150.png", "./public/ms-icon-150x150.png").
		StaticFile("/ms-icon-310x310.png", "./public/ms-icon-310x310.png").
		StaticFile("/ms-icon-70x70.png", "./public/ms-icon-70x70.png").

		// Sitemaps
		StaticFile("/sitemap.xml", "./public/sitemap.xml").
		StaticFile("/sitemap-main.xml", "./public/sitemap-main.xml").
		StaticFile("/sitemap-companies.xml", "./public/sitemap-companies.xml").
		StaticFile("/sitemap-vacancies.xml", "./public/sitemap-vacancies.xml").
		StaticFile("/sitemap-projects.xml", "./public/sitemap-projects.xml").

		// System
		StaticFile("/humans.txt", "./public/humans.txt").
		StaticFile("/robots.txt", "./public/robots.txt").
		StaticFile("/manifest.json", "./public/manifest.json").
		StaticFile("/browserconfig.xml", "./public/browserconfig.xml").

		// Privacy policy aliases
		// https://dou.ua/privacy/
		// https://privacy.xing.com/en/privacy-policy
		// https://www.linkedin.com/legal/privacy-policy
		// https://www.facebook.com/privacy/policy/
		StaticFile("privacy-policy", "./public/privacy-policy.html")

	{
		// very fast solution
		go func() {
			for range time.Tick(time.Minute) {
				pairs, err := onlineRedisStorage.GetAndClear(context.Background())
				if err != nil {
					// NOP

					continue
				}

				err = onlinePostgresStorage.BatchStore(context.Background(), pairs)
				if err != nil {
					// NOP

					continue
				}
			}
		}()
	}

	server.Run(r.Handler())
}

func redirectFromWWW() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.Host, "www.") {
			newHost := "https://" + c.Request.Host[len("www."):]
			c.Redirect(http.StatusMovedPermanently, newHost+c.Request.URL.String())
			return
		}
		c.Next()
	}
}

func found(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusFound, path)
	}
}

func s(path string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.File(path)
	}
}
