package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"
)

type GithubRepositoryStarsRepository struct {
	db *Database
}

func NewGithubRepositoryStarsRepository(db *Database) *GithubRepositoryStarsRepository {
	return &GithubRepositoryStarsRepository{db: db}
}

func (r *GithubRepositoryStarsRepository) Upsert(ctx context.Context, owner, repo string, stargazersCount int32, createdAt time.Time) error {
	return r.db.Queries().GithubRepositoryStarsUpsert(ctx, dbs.GithubRepositoryStarsUpsertParams{
		Owner:           owner,
		Repo:            repo,
		StargazersCount: stargazersCount,
		CreatedAt:       createdAt,
	})
}

func (r *GithubRepositoryStarsRepository) All(ctx context.Context) ([]domain.GithubRepositoryStar, error) {
	rows, err := r.db.Queries().GithubRepositoryStarsAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.GithubRepositoryStar, len(rows))
	for i, row := range rows {
		result[i] = domain.GithubRepositoryStar(row)
	}
	return result, nil
}

func (r *GithubRepositoryStarsRepository) Get(ctx context.Context, owner, repo string) (int32, error) {
	stars, err := r.db.Queries().GithubRepositoryStars(ctx, dbs.GithubRepositoryStarsParams{
		Owner: owner,
		Repo:  repo,
	})
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return stars, nil
}

func (r *GithubRepositoryStarsRepository) Default(ctx context.Context) (int32, error) {
	return r.Get(ctx, "readytotouch", "readytotouch")
}
