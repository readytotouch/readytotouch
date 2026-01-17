-- name: UpsertGithubRepositoryStars :exec
INSERT INTO github_repository_stars (owner, repo, stargazers_count, updated_at)
VALUES ($1, $2, $3, $4)
ON CONFLICT (owner, repo) DO UPDATE
SET stargazers_count = EXCLUDED.stargazers_count,
    updated_at       = EXCLUDED.updated_at;

-- name: GetGithubRepositoryStars :many
SELECT owner, repo, stargazers_count, updated_at
FROM github_repository_stars;
