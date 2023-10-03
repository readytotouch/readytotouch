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

	"github.com/readytotouch-yaaws/yaaws-go/internal/auth"
	"github.com/readytotouch-yaaws/yaaws-go/internal/db/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/db/redis"
	"github.com/readytotouch-yaaws/yaaws-go/internal/domain"
	"github.com/readytotouch-yaaws/yaaws-go/internal/env"
	"github.com/readytotouch-yaaws/yaaws-go/internal/server"

	pkgJWT "github.com/readytotouch-yaaws/yaaws-go/internal/jwt"
	pkgBitbucket "github.com/readytotouch-yaaws/yaaws-go/internal/oauth-providers/bitbucket"
	pkgGitHub "github.com/readytotouch-yaaws/yaaws-go/internal/oauth-providers/github"
	pkgGitLab "github.com/readytotouch-yaaws/yaaws-go/internal/oauth-providers/gitlab"
	pkgOnline "github.com/readytotouch-yaaws/yaaws-go/internal/online"
	pkgUsers "github.com/readytotouch-yaaws/yaaws-go/internal/users"

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

	redisClient := redis.MustClient("redis:6379")

	var (
		userRepository   = postgres.NewUserRepository(database)
		onlineRepository = postgres.NewOnlineRepository(database)
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
		userController   = pkgUsers.NewController(userRepository)
		onlineController = pkgOnline.NewController(userRepository, onlineRepository)
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
		}).
		GET("/", onlineController.Index)

	r.GET("/api/v1/users/registration/stats/daily.json", userController.RegistrationDailyCountStats)
	r.GET("/api/v1/users/online/stats/daily.json", onlineController.DailyCountStats)

	r.
		GET("/auth/github", authController.GithubRedirect).
		GET("/auth/gitlab", authController.GitlabRedirect).
		GET("/auth/bitbucket", authController.BitbucketRedirect).
		GET("/auth/github/callback", authController.GithubCallback).
		GET("/auth/gitlab/callback", authController.GitlabCallback).
		GET("/auth/bitbucket/callback", authController.BitbucketCallback).
		GET("/logout", authController.Logout)

	r.
		StaticFile("/design", "./public/design/online.html").
		StaticFile("/design/online", "./public/design/online.html").
		StaticFile("/design/online-auth", "./public/design//online-auth.html")

	r.
		// Assets
		Static("/assets/images", "./public/assets/images").
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
		StaticFile("/browserconfig.xml", "./public/browserconfig.xml")

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
