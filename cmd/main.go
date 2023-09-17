package main

import (
	"context"

	"github.com/readytotouch-yaaws/yaaws-go/internal/env"
	"github.com/readytotouch-yaaws/yaaws-go/internal/server"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"

	pkgOnline "github.com/readytotouch-yaaws/yaaws-go/internal/online"
	pkgUser "github.com/readytotouch-yaaws/yaaws-go/internal/user"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		dsn = env.Must("POSTGRES_DSN")
	)

	pgConnection := postgres.MustConnection(context.Background(), dsn)
	defer pgConnection.Close()

	database := postgres.NewDatabase(pgConnection)

	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	var (
		userController   = pkgUser.NewController(pkgUser.NewRepository(database))
		onlineController = pkgOnline.NewController(pkgOnline.NewRepository(database))
	)

	r.GET("/api/v1/users/registration/stats/daily.json", userController.RegistrationDailyCountStats)
	r.GET("/api/v1/users/online/stats/daily.json", onlineController.DailyCountStats)

	r.StaticFile("/robots.txt", "./public/robots.txt")

	server.Run(r.Handler())
}
