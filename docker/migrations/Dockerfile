FROM alpine

ADD https://github.com/pressly/goose/releases/download/v3.15.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

# ENV GOOSE_DBSTRING defined in docker-compose.yml so goose understand how connect to Postgres
CMD ["/bin/sh", "-c", "/bin/goose --dir=/db/migrations --table=schema_migrations up && sleep infinity"]
