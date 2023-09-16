package main

import (
	"context"
	"fmt"

	"github.com/readytotouch-yaaws/yaaws-go/internal/env"
	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"
)

func main() {
	var (
		dsn = env.Must("POSTGRES_DSN")
	)

	var pgConnection = postgres.MustConnection(context.Background(), dsn)
	defer pgConnection.Close()

	repository := postgres.NewRepository(pgConnection)
	_ = repository

	fmt.Println("success")
}
