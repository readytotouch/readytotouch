INSERT INTO github_repository_stars (owner, repo, stargazers_count, updated_at)
VALUES ('readytotouch', 'readytotouch', 0, NOW())
ON CONFLICT DO NOTHING;
