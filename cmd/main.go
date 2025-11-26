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
		onlineController    = pkgOnline.NewController(onlineRepository)
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
	r.Use(redirectTrimPrefix("/organizers/", "/"))
	r.Use(redirectTrimPrefix("/design/organizers/", "/design/"))
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
	r.GET("/v1", organizerController.IndexV1)
	r.GET("/v2", organizerController.IndexV2)
	r.GET("/v3", organizerController.IndexV3)
	r.GET("/", organizerController.IndexV3)

	r.GET("/api/v1/users/registration/stats/daily.json", userController.RegistrationDailyCountStats)
	r.GET("/api/v1/users/online/stats/daily.json", onlineController.DailyCountStats)

	r.GET("/organizers", organizerController.Organizers)

	r.GET("/golang/welcome", organizerController.WelcomeV3)
	r.GET("/rust/welcome", organizerController.WelcomeV3)
	r.GET("/zig/welcome", organizerController.WelcomeV3)
	r.GET("/scala/welcome", organizerController.WelcomeV3)
	r.GET("/elixir/welcome", organizerController.WelcomeV3)
	r.GET("/erlang/welcome", organizerController.WelcomeV3)
	r.GET("/clojure/welcome", organizerController.WelcomeV3)
	r.GET("/haskell/welcome", organizerController.WelcomeV3)
	r.GET("/fsharp/welcome", organizerController.WelcomeV3)
	r.GET("/ocaml/welcome", organizerController.WelcomeV3)
	r.GET("/gleam/welcome", organizerController.WelcomeV3)
	r.GET("/mojo/welcome", organizerController.WelcomeV3)

	r.GET("/golang/companies/ukraine", organizerController.GolangCompaniesUkraine)
	r.GET("/golang/companies", organizerController.CompaniesV3)
	r.GET("/golang/companies/v1", organizerController.CompaniesV1)
	r.GET("/golang/companies/v2", organizerController.CompaniesV2)
	r.GET("/golang/companies/v3", organizerController.CompaniesV3)
	r.GET("/golang/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/golang/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/golang/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/golang/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/golang/jobs", organizerController.JobsV3)
	r.GET("/golang/jobs/v2", organizerController.JobsV2)
	r.GET("/golang/jobs/v3", organizerController.JobsV3)
	r.GET("/golang/communities", organizerController.GolangCommunities)
	r.GET("/rust/companies", organizerController.CompaniesV3)
	r.GET("/rust/companies/v1", organizerController.CompaniesV1)
	r.GET("/rust/companies/v2", organizerController.CompaniesV2)
	r.GET("/rust/companies/v3", organizerController.CompaniesV3)
	r.GET("/rust/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/rust/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/rust/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/rust/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/rust/jobs", organizerController.JobsV3)
	r.GET("/rust/jobs/v2", organizerController.JobsV2)
	r.GET("/rust/jobs/v3", organizerController.JobsV3)
	r.GET("/rust/communities", organizerController.RustCommunities)
	r.GET("/zig/companies", organizerController.Waitlist)
	r.GET("/zig/companies/v1", organizerController.Waitlist)
	r.GET("/zig/companies/v2", organizerController.Waitlist)
	r.GET("/zig/companies/v3", organizerController.Waitlist)
	r.GET("/zig/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/zig/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/zig/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/zig/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/zig/jobs", organizerController.Waitlist)
	r.GET("/zig/jobs/v2", organizerController.Waitlist)
	r.GET("/zig/jobs/v3", organizerController.Waitlist)
	r.GET("/zig/communities", organizerController.ZigCommunities)
	r.GET("/scala/companies", organizerController.CompaniesV3)
	r.GET("/scala/companies/v1", organizerController.CompaniesV1)
	r.GET("/scala/companies/v2", organizerController.CompaniesV2)
	r.GET("/scala/companies/v3", organizerController.CompaniesV3)
	r.GET("/scala/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/scala/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/scala/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/scala/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/scala/jobs", organizerController.JobsV3)
	r.GET("/scala/jobs/v2", organizerController.JobsV2)
	r.GET("/scala/jobs/v3", organizerController.JobsV3)
	r.GET("/scala/communities", organizerController.ScalaCommunities)
	r.GET("/elixir/companies", organizerController.CompaniesV3)
	r.GET("/elixir/companies/v1", organizerController.CompaniesV1)
	r.GET("/elixir/companies/v2", organizerController.CompaniesV2)
	r.GET("/elixir/companies/v3", organizerController.CompaniesV3)
	r.GET("/elixir/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/elixir/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/elixir/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/elixir/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/elixir/jobs", organizerController.JobsV3)
	r.GET("/elixir/jobs/v2", organizerController.JobsV2)
	r.GET("/elixir/jobs/v3", organizerController.JobsV3)
	r.GET("/elixir/communities", organizerController.ElixirCommunities)
	r.GET("/erlang/companies", organizerController.Waitlist)
	r.GET("/erlang/companies/v1", organizerController.Waitlist)
	r.GET("/erlang/companies/v2", organizerController.Waitlist)
	r.GET("/erlang/companies/v3", organizerController.Waitlist)
	r.GET("/erlang/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/erlang/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/erlang/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/erlang/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/erlang/jobs", organizerController.Waitlist)
	r.GET("/erlang/jobs/v2", organizerController.JobsV2)
	r.GET("/erlang/jobs/v3", organizerController.JobsV3)
	r.GET("/erlang/communities", organizerController.TODO)
	r.GET("/clojure/companies", organizerController.CompaniesV3)
	r.GET("/clojure/companies/v1", organizerController.CompaniesV1)
	r.GET("/clojure/companies/v2", organizerController.CompaniesV2)
	r.GET("/clojure/companies/v3", organizerController.CompaniesV3)
	r.GET("/clojure/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/clojure/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/clojure/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/clojure/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/clojure/jobs", organizerController.JobsV3)
	r.GET("/clojure/jobs/v2", organizerController.JobsV2)
	r.GET("/clojure/jobs/v3", organizerController.JobsV3)
	r.GET("/clojure/communities", organizerController.ClojureCommunities)
	r.GET("/haskell/companies", organizerController.CompaniesV3)
	r.GET("/haskell/companies/v1", organizerController.CompaniesV1)
	r.GET("/haskell/companies/v2", organizerController.CompaniesV2)
	r.GET("/haskell/companies/v3", organizerController.CompaniesV3)
	r.GET("/haskell/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/haskell/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/haskell/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/haskell/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/haskell/jobs", organizerController.JobsV3)
	r.GET("/haskell/jobs/v2", organizerController.JobsV2)
	r.GET("/haskell/jobs/v3", organizerController.JobsV3)
	r.GET("/haskell/communities", organizerController.TODO)
	r.GET("/fsharp/companies", organizerController.Waitlist)
	r.GET("/fsharp/companies/v1", organizerController.Waitlist)
	r.GET("/fsharp/companies/v2", organizerController.Waitlist)
	r.GET("/fsharp/companies/v3", organizerController.Waitlist)
	r.GET("/fsharp/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/fsharp/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/fsharp/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/fsharp/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/fsharp/jobs", organizerController.Waitlist)
	r.GET("/fsharp/jobs/v2", organizerController.Waitlist)
	r.GET("/fsharp/jobs/v3", organizerController.Waitlist)
	r.GET("/fsharp/communities", organizerController.TODO)
	r.GET("/ocaml/companies", organizerController.Waitlist)
	r.GET("/ocaml/companies/v1", organizerController.Waitlist)
	r.GET("/ocaml/companies/v2", organizerController.Waitlist)
	r.GET("/ocaml/companies/v3", organizerController.Waitlist)
	r.GET("/ocaml/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/ocaml/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/ocaml/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/ocaml/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/ocaml/jobs", organizerController.Waitlist)
	r.GET("/ocaml/jobs/v2", organizerController.Waitlist)
	r.GET("/ocaml/jobs/v3", organizerController.Waitlist)
	r.GET("/ocaml/communities", organizerController.TODO)
	r.GET("/gleam/companies", organizerController.Waitlist)
	r.GET("/gleam/companies/v1", organizerController.Waitlist)
	r.GET("/gleam/companies/v2", organizerController.Waitlist)
	r.GET("/gleam/companies/v3", organizerController.Waitlist)
	r.GET("/gleam/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/gleam/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/gleam/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/gleam/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/gleam/jobs", organizerController.Waitlist)
	r.GET("/gleam/jobs/v2", organizerController.Waitlist)
	r.GET("/gleam/jobs/v3", organizerController.Waitlist)
	r.GET("/gleam/communities", organizerController.TODO)
	r.GET("/mojo/companies", organizerController.Waitlist)
	r.GET("/mojo/companies/v1", organizerController.Waitlist)
	r.GET("/mojo/companies/v2", organizerController.Waitlist)
	r.GET("/mojo/companies/v3", organizerController.Waitlist)
	r.GET("/mojo/companies/:company_alias", organizerController.CompanyV3)
	r.GET("/mojo/companies/:company_alias/v1", organizerController.CompanyV1)
	r.GET("/mojo/companies/:company_alias/v2", organizerController.CompanyV2)
	r.GET("/mojo/companies/:company_alias/v3", organizerController.CompanyV3)
	r.GET("/mojo/jobs", organizerController.Waitlist)
	r.GET("/mojo/jobs/v2", organizerController.Waitlist)
	r.GET("/mojo/jobs/v3", organizerController.Waitlist)
	r.GET("/mojo/communities", organizerController.TODO)
	r.GET("/v/:vacancy_id", organizerController.VacancyRedirect)

	r.GET("/golang", found("/golang/companies", true))
	r.GET("/rust", found("/rust/companies", true))
	r.GET("/zig", found("/zig/companies", true))
	r.GET("/scala", found("/scala/companies", true))
	r.GET("/elixir", found("/elixir/companies", true))
	r.GET("/erlang", found("/erlang/companies", true))
	r.GET("/clojure", found("/clojure/companies", true))
	r.GET("/haskell", found("/haskell/companies", true))
	r.GET("/fsharp", found("/fsharp/companies", true))
	r.GET("/ocaml", found("/ocaml/companies", true))
	r.GET("/gleam", found("/gleam/companies", true))
	r.GET("/mojo", found("/mojo/companies", true))

	r.GET("/golang/vacancies", found("/golang/jobs", true))
	r.GET("/rust/vacancies", found("/rust/jobs", true))
	r.GET("/zig/vacancies", found("/zig/jobs", true))
	r.GET("/scala/vacancies", found("/scala/jobs", true))
	r.GET("/elixir/vacancies", found("/elixir/jobs", true))
	r.GET("/clojure/vacancies", found("/clojure/jobs", true))
	r.GET("/erlang/vacancies", found("/erlang/jobs", true))
	r.GET("/fsharp/vacancies", found("/fsharp/jobs", true))
	r.GET("/ocaml/vacancies", found("/ocaml/jobs", true))
	r.GET("/gleam/vacancies", found("/gleam/jobs", true))
	r.GET("/mojo/vacancies", found("/mojo/jobs", true))

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
		GET("/companies-and-connections", found("/companies-and-connections/worldwide", true)).
		GET("/companies-and-connections/worldwide", cacController.Worldwide).
		GET("/companies-and-connections/ukraine", cacController.Ukraine).
		GET("/companies-and-connections/brazil", cacController.Brazil).
		GET("/wip/companies-and-connections", found("/companies-and-connections/worldwide", true)).
		GET("/wip/companies-and-connections/ukraine", found("/companies-and-connections/ukraine", true)).
		GET("/wip/companies-and-connections/brazil", found("/companies-and-connections/brazil", true)).
		GET("/api/v1/companies-and-connections/companies.json", cacController.Companies).
		POST("/api/v1/companies-and-connections/companies.json", cacController.AddCompany).
		DELETE("/api/v1/companies-and-connections/companies.json", cacController.DeleteCompany)

	r.
		// Unsafe API endpoints that can be changed without any notice.
		GET("/api/v1/unsafe/companies.json", organizerController.UnsafeCompanies).
		GET("/api/v1/unsafe/vacancies.json", organizerController.UnsafeVacancies)

	r.GET("/analytics", found("https://plausible.io/readytotouch.com", false))
	r.GET("/plausible", found("https://plausible.io/readytotouch.com", false))
	r.GET("/similarweb", found("https://www.similarweb.com/website/readytotouch.com/", false))

	r.
		StaticFile("/design", "./public/design-v3/main-responsive.html").
		StaticFile("/design/v1", "./public/design/online.html").
		StaticFile("/design/v1/auth", "./public/design/online-auth.html").
		StaticFile("/design/v2", "./public/design/online-new.html").
		StaticFile("/design/v2/auth", "./public/design/online-new-auth.html").
		StaticFile("/design/v3", "./public/design-v3/main-responsive.html").
		StaticFile("/design/v3/auth", "./public/design-v3/main-responsive-auth.html").

		// Organizers
		StaticFile("/design/organizers", "./public/design/organizer-main-page.html").
		StaticFile("/design/organizers-auth", "./public/design/organizer-main-page-auth.html").
		GET("/design/:language/welcome", s("./public/design/organizer-welcome.html")).
		GET("/design/:language/companies/ukraine", s("./public/design/golang-companies-organizer.html")).
		GET("/design/:language/companies/:company_alias", s("./public/design-v3/organizer-company-responsive-auth.html")).
		GET("/design/:language/companies/:company_alias/v1", s("./public/design-v1/organizer-statistics.html")).
		GET("/design/:language/companies/:company_alias/v2", s("./public/design/organizer-statistics.html")).
		GET("/design/:language/companies/:company_alias/v3", s("./public/design-v3/organizer-company-responsive-auth.html")).
		GET("/design/:language/companies", s("./public/design-v3/organizer-companies-responsive.html")).
		GET("/design/:language/companies/v1", s("./public/design-v1/organizer-companies.html")).
		GET("/design/:language/companies/v2", s("./public/design/organizer-companies.html")).
		GET("/design/:language/companies/v3", s("./public/design-v3/organizer-companies-responsive.html")).
		GET("/design/:language/companies/v3/auth", s("./public/design-v3/organizer-companies-responsive-auth.html")).
		GET("/design/:language/jobs", s("./public/design-v3/organizer-vacancies-responsive-auth.html")).
		// v1 has never existed for the jobs page
		// GET("/design/:language/jobs/v1", s("./public/design/organizer-vacancies.html")).
		GET("/design/:language/jobs/v2", s("./public/design/organizer-vacancies.html")).
		GET("/design/:language/jobs/v3", s("./public/design-v3/organizer-vacancies-responsive-auth.html")).
		GET("/design/:language/jobs/subscribe", s("./public/design/organizer-vacancies-subscribe.html")).
		GET("/design/:language/jobs/unsubscribe", s("./public/design/organizer-vacancies-unsubscribe.html")).
		GET("/design/:language/vacancies", s("./public/design-v3/organizer-vacancies-responsive-auth.html")).
		GET("/design/:language/vacancies/subscribe", s("./public/design/organizer-vacancies-subscribe.html")).
		GET("/design/:language/vacancies/unsubscribe", s("./public/design/organizer-vacancies-unsubscribe.html")).
		GET("/design/golang/communities", s("./public/design/organizer-go-communities.html")).
		GET("/design/rust/communities", s("./public/design/organizer-rust-communities.html")).
		GET("/design/scala/communities", s("./public/design/organizer-scala-communities.html")).
		GET("/design/elixir/communities", s("./public/design/organizer-elixir-communities.html")).
		GET("/design/clojure/communities", s("./public/design/organizer-clojure-communities.html")).
		GET("/design/:language/communities", s("./public/design/organizer-go-communities.html")).

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
		Static("/assets/unstable/logos/112x56/", "./public/logos/112x56/").
		Static("/assets/unstable/logos/72x72/", "./public/logos/72x72/").
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

	r.
		GET("/sitemap-golang-companies.xml", organizerController.SitemapCompanies(domain.OrganizerGolang)).
		GET("/sitemap-rust-companies.xml", organizerController.SitemapCompanies(domain.OrganizerRust)).
		GET("/sitemap-zig-companies.xml", organizerController.SitemapCompanies(domain.OrganizerZig)).
		GET("/sitemap-scala-companies.xml", organizerController.SitemapCompanies(domain.OrganizerScala)).
		GET("/sitemap-elixir-companies.xml", organizerController.SitemapCompanies(domain.OrganizerElixir)).
		GET("/sitemap-erlang-companies.xml", organizerController.SitemapCompanies(domain.OrganizerErlang)).
		GET("/sitemap-clojure-companies.xml", organizerController.SitemapCompanies(domain.OrganizerClojure)).
		GET("/sitemap-haskell-companies.xml", organizerController.SitemapCompanies(domain.OrganizerHaskell)).
		GET("/sitemap-fsharp-companies.xml", organizerController.SitemapCompanies(domain.OrganizerFSharp)).
		GET("/sitemap-ocaml-companies.xml", organizerController.SitemapCompanies(domain.OrganizerOCaml)).
		GET("/sitemap-gleam-companies.xml", organizerController.SitemapCompanies(domain.OrganizerGleam)).
		GET("/sitemap-mojo-companies.xml", organizerController.SitemapCompanies(domain.OrganizerMojo))

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

func redirectTrimPrefix(old, new string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, old) {
			found(new+strings.TrimPrefix(c.Request.URL.Path, old), true)(c)
			return
		}

		c.Next()
	}
}

func found(path string, keepQueryParams bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if keepQueryParams && c.Request.URL.RawQuery != "" {
			c.Redirect(http.StatusFound, path+"?"+c.Request.URL.RawQuery)
		} else {
			c.Redirect(http.StatusFound, path)
		}
	}
}

func s(path string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.File(path)
	}
}
