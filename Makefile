env-docker-compose-development:
	rm -f docker-compose.yml
	ln -s ./docker/compose/development/docker-compose.yml docker-compose.yml

env-docker-compose-production:
	rm -f docker-compose.yml
	ln -s ./docker/compose/production/docker-compose.yml docker-compose.yml

env-up:
	docker-compose -f docker-compose.yml --env-file .env up -d

postgres-test-run:
	docker exec readytotouch_postgres_db psql -U yaaws_user -d yaaws -c "SELECT VERSION();"

logs:
	docker logs readytotouch_go_app

pg:
	docker exec -it readytotouch_postgres_db bash

pg-users:
	cat ./fixtures/postgres/users.sql | docker exec -i readytotouch_postgres_db psql -d yaaws -U yaaws_user

env-down:
	docker-compose -f docker-compose.yml --env-file .env down

env-down-with-clear:
	docker-compose -f docker-compose.yml --env-file .env down --remove-orphans -v # --rmi=all

app-stop:
	docker exec readytotouch_go_app pkill go

app-start:
	docker exec readytotouch_go_app go run ./cmd/main.go

test:
	docker exec readytotouch_go_app go test ./... -v -count=1

bench:
	docker exec readytotouch_go_app go test ./... -v -run=$$^ -bench=. -benchmem -benchtime=1000x

# make migrate-pgsql-create NAME=init
migrate-pgsql-create:
	goose -dir ./internal/storage/postgres/migrations -table schema_migrations postgres create $(NAME) sql

migrate-pgsql-up:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		up
migrate-pgsql-redo:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		redo
migrate-pgsql-down:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		down-to 20230815012800
migrate-pgsql-reset:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		reset
migrate-pgsql-status:
	docker exec readytotouch_postgres_goose_migrations \
		goose --dir=/db/migrations --table=schema_migrations \
		status

generate-template:
	# go get -u github.com/valyala/quicktemplate/qtc
	qtc -dir=./internal/templates/v1 -skipLineComments
	git add .

generate-sqlc:
	sqlc generate

esbuild-minify:
	npm --prefix=client i
	MINIFY=true npm run --prefix=client esbuild
	tree -h ./public/assets/js

esbuild:
	MINIFY=false npm run --prefix=client esbuild
	tree -h ./public/assets/js

go-mod-update:
	go mod tidy
	go mod vendor

# make design DESIGN="~/go/src/github.com/readytotouch-yaaws/readytotouch-yaaws.github.io"
design:
	$(eval DESIGN := ~/go/src/github.com/readytotouch-yaaws/readytotouch-yaaws.github.io)
	rm -rf ./public/assets/images ./public/design
	mkdir -p ./public/assets/images ./public/design

	cp -r $(DESIGN)/public/assets/images/* ./public/assets/images
	cp -r $(DESIGN)/public/*.html ./public/design

	git add .
