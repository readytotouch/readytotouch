POSTGRES_DSN="postgresql://yaaws_user:yaaws_password@localhost:5432/yaaws?sslmode=disable"

env-up:
	docker-compose -f docker-compose.yml --env-file .env up -d

logs:
	docker logs readytotouch_go_app

pg:
	docker exec -it readytotouch_postgres_db bash

env-down:
	docker-compose -f docker-compose.yml --env-file .env down

env-down-with-clear:
	docker-compose -f docker-compose.yml --env-file .env down --remove-orphans -v # --rmi=all

# make migrate-pgsql-create NAME=init
migrate-pgsql-create:
	# mkdir -p ./internal/storage/postgres/migrations
	$(eval NAME ?= todo)
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations postgres $(POSTGRES_DSN) create $(NAME) sql

migrate-pgsql-up:
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations postgres $(POSTGRES_DSN) up
migrate-pgsql-redo:
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations postgres $(POSTGRES_DSN) redo
migrate-pgsql-down:
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations postgres $(POSTGRES_DSN) down
migrate-pgsql-reset:
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations postgres $(POSTGRES_DSN) reset
migrate-pgsql-status:
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations postgres $(POSTGRES_DSN) status

generate-sqlc:
	# go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	sqlc generate
	# alternative
	# docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

go-mod-update:
	go mod tidy
	go mod vendor
