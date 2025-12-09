package repository

import (
	"context"
	"quiz-log/models"

	"github.com/uptrace/bun"
)

//go:generate mockgen -destination=mocks/mock_tag_repository.go -package=mocks quiz-log/repository TagRepository

// TagRepository defines the interface for tag repository operations
type TagRepository interface {
	Create(ctx context.Context, name string) (int, error)
	FindAll(ctx context.Context) ([]*models.Tag, error)
}

type tagRepository struct {
	DB *bun.DB
}

func NewTagRepository(database *bun.DB) TagRepository {
	return &tagRepository{DB: database}
}

// Create creates a new tag or returns existing one by name
func (r *tagRepository) Create(ctx context.Context, name string) (int, error) {
	var tagID int

	query := psql.Insert("tags").
		Columns("name").
		Values(name).
		Suffix("ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING id")

	err := ExecQueryWithReturning[int](ctx, r.DB, query, &tagID)
	if err != nil {
		return 0, err
	}

	return tagID, nil
}

// FindAll retrieves all tags
func (r *tagRepository) FindAll(ctx context.Context) ([]*models.Tag, error) {
	query := psql.Select("id", "name").
		From("tags").
		OrderBy("name ASC")

	return FindAll[models.Tag](ctx, r.DB, query)
}
