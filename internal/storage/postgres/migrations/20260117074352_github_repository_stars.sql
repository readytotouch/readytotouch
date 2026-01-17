-- +goose Up
-- +goose StatementBegin
CREATE TABLE github_repository_stars
(
    owner            VARCHAR                  NOT NULL,
    repo             VARCHAR                  NOT NULL,
    stargazers_count INT                      NOT NULL,
    created_at       TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at       TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (owner, repo)
);

INSERT INTO github_repository_stars (owner, repo, stargazers_count, created_at, updated_at)
VALUES ('readytotouch', 'readytotouch', 1325, '2026-01-17 12:00:00', '2026-01-17 12:00:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE github_repository_stars;
-- +goose StatementEnd
