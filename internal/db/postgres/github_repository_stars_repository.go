package postgres

import (
	"context"
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

func (r *GithubRepositoryStarsRepository) Upsert(ctx context.Context, owner, repo string, stars int, updatedAt time.Time) error {
	return r.db.Queries().UpsertGithubRepositoryStars(ctx, dbs.UpsertGithubRepositoryStarsParams{
		Owner:           owner,
		Repo:            repo,
		StargazersCount: int32(stars),
		UpdatedAt:       updatedAt,
	})
}

func (r *GithubRepositoryStarsRepository) ListAll(ctx context.Context) ([]domain.GithubRepositoryStar, error) {
	rows, err := r.db.Queries().GetGithubRepositoryStars(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.GithubRepositoryStar, len(rows))
	for i, row := range rows {
		result[i] = domain.GithubRepositoryStar{
			Owner:           row.Owner,
			Repo:            row.Repo,
			StargazersCount: int(row.StargazersCount),
			UpdatedAt:       row.UpdatedAt,
		}
	}
	return result, nil
}
