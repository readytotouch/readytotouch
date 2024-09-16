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

	pkgJWT "github.com/readytotouch/readytotouch/internal/jwt"
	pkgBitbucket "github.com/readytotouch/readytotouch/internal/oauth-providers/bitbucket"
	pkgGitHub "github.com/readytotouch/readytotouch/internal/oauth-providers/github"
	pkgGitLab "github.com/readytotouch/readytotouch/internal/oauth-providers/gitlab"
	pkgOnline "github.com/readytotouch/readytotouch/internal/online"
	pkgOrganizer "github.com/readytotouch/readytotouch/internal/organizer"
	pkgUsers "github.com/readytotouch/readytotouch/internal/users"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	var (
		dsn          = env.Required("POSTGRES_DSN")
		jwtSecretKey = env.Required("JWT_SECRET_KEY")
	)

	pgConnection := postgres.MustConnection(dsn)
	defer pgConnection.Close()

	database := postgres.MustDatabase(pgConnection)
	defer database.Queries().Close()

	redisClient := redis.MustClient(context.Background(), "redis:6379")

	var (
		userRepository                = postgres.NewUserRepository(database)
		onlineRepository              = postgres.NewOnlineRepository(database)
		userFeatureWaitlistRepository = postgres.NewUserFeatureWaitlistRepository(database)
		featureViewStatsRepository    = postgres.NewFeatureViewStatsRepository(database)
		userFavoriteCompanyRepository = postgres.NewUserFavoriteCompanyRepository(database)
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
		)
	)

	r := gin.New()
	r.Use(redirectFromWWW())
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
	r.GET("/", onlineController.Index)

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
	r.GET("/organizers/golang/companies", organizerController.Companies)
	r.GET("/organizers/golang/vacancies", organizerController.Waitlist)
	r.GET("/organizers/rust/companies", organizerController.Waitlist)
	r.GET("/organizers/rust/vacancies", organizerController.Waitlist)
	r.GET("/organizers/zig/companies", organizerController.Waitlist)
	r.GET("/organizers/zig/vacancies", organizerController.Waitlist)
	r.GET("/organizers/scala/companies", organizerController.Waitlist)
	r.GET("/organizers/scala/vacancies", organizerController.Waitlist)
	r.GET("/organizers/elixir/companies", organizerController.Waitlist)
	r.GET("/organizers/elixir/vacancies", organizerController.Waitlist)
	r.GET("/organizers/clojure/companies", organizerController.Waitlist)
	r.GET("/organizers/clojure/vacancies", organizerController.Waitlist)

	r.GET("/organizers/golang", found("/organizers/golang/companies"))
	r.GET("/organizers/rust", found("/organizers/rust/companies"))
	r.GET("/organizers/zig", found("/organizers/zig/companies"))
	r.GET("/organizers/scala", found("/organizers/scala/companies"))
	r.GET("/organizers/elixir", found("/organizers/elixir/companies"))
	r.GET("/organizers/clojure", found("/organizers/clojure/companies"))

	r.GET("/api/v1/features/auto/waitlist/stats.json", organizerController.WaitlistStats)
	r.POST("/api/v1/features/auto/waitlist/subscribe.json", organizerController.WaitlistSubscribe)
	r.PATCH("/api/v1/companies/:company_id/favorite.json", organizerController.FavoriteCompany)

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
		StaticFile("/wip/companies-and-connections", "./public/chatgpt-design/companies-and-connections.html").
		StaticFile("/wip/companies-and-connections/ukraine", "./public/chatgpt-design/companies-and-connections.html")

	r.
		StaticFile("/design", "./public/design/online.html").
		StaticFile("/design/online", "./public/design/online.html").
		StaticFile("/design/online-auth", "./public/design/online-auth.html").

		// Design from OrganizerFeature
		StaticFile("/design/organizers", "./public/design/organizer-main-page-auth.html").
		StaticFile("/design/organizers-auth", "./public/design/organizer-main-page.html").
		StaticFile("/design/organizers/golang/welcome", "./public/design/organizer-welcome.html").
		StaticFile("/design/organizers/golang/companies/ukraine", "./public/design/golang-companies-organizer.html").
		StaticFile("/design/organizers/golang/companies", "./public/design/organizer-companies.html").
		StaticFile("/design/organizers/golang/vacancies", "./public/design/organizer-vacancies-subscribe.html").
		StaticFile("/design/organizers/golang/vacancies/subscribe", "./public/design/organizer-vacancies-subscribe.html").
		StaticFile("/design/organizers/golang/vacancies/unsubscribe", "./public/design/organizer-vacancies-unsubscribe.html").

		// Design from ChatGPT
		StaticFile("/design/wip/companies-and-connections", "./public/chatgpt-design/companies-and-connections.html").
		StaticFile("/design/wip/companies-and-connections/ukraine", "./public/chatgpt-design/companies-and-connections.html")

	r.
		// Assets
		Static("/assets/images", "./public/assets/images").
		Static("/assets/fonts", "./public/assets/fonts").
		Static("/assets/js", "./public/assets/js").

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
