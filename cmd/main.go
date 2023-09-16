package main

import (
	"context"

	"github.com/readytotouch-yaaws/yaaws-go/internal/env"
	"github.com/readytotouch-yaaws/yaaws-go/internal/server"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		dsn = env.Must("POSTGRES_DSN")
	)

	var pgConnection = postgres.MustConnection(context.Background(), dsn)
	defer pgConnection.Close()

	repository := postgres.NewRepository(pgConnection)

	var r = gin.New()

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	_ = repository

	r.StaticFile("/robots.txt", "./public/robots.txt")

	server.Run(r.Handler())
}
