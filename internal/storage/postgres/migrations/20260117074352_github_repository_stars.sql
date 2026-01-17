-- +goose Up
-- +goose StatementBegin
CREATE TABLE github_repository_stars
(
    owner            TEXT                     NOT NULL,
    repo             TEXT                     NOT NULL,
    stargazers_count INTEGER                  NOT NULL DEFAULT 0,
    updated_at       TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (owner, repo)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE github_repository_stars;
-- +goose StatementEnd
