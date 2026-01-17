-- name: GithubRepositoryStarsUpsert :exec
INSERT INTO github_repository_stars (owner, repo, stargazers_count, created_at, updated_at)
VALUES (@owner, @repo, @stargazers_count, @created_at, @created_at)
ON CONFLICT (owner, repo) DO UPDATE
    SET stargazers_count = excluded.stargazers_count,
        updated_at       = excluded.updated_at;

-- name: GithubRepositoryStarsAll :many
SELECT owner, repo, stargazers_count
FROM github_repository_stars;

-- name: GithubRepositoryStars :one
SELECT stargazers_count
FROM github_repository_stars
WHERE owner = @owner
  AND repo = @repo;
