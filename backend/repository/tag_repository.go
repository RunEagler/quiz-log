package repository

import (
	"context"
	"quiz-log/db"

	"github.com/uptrace/bun"
)

type TagRepository struct {
	DB *bun.DB
}

func NewTagRepository(database *bun.DB) *TagRepository {
	return &TagRepository{DB: database}
}

// Create creates a new tag or returns existing one by name
func (r *TagRepository) Create(ctx context.Context, name string) (int, error) {
	var tagID int

	query := psql.Insert("tags").
		Columns("name").
		Values(name).
		Suffix("ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING id")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&tagID)
	if err != nil {
		return 0, err
	}

	return tagID, nil
}

// FindAll retrieves all tags
func (r *TagRepository) FindAll(ctx context.Context) ([]*db.Tag, error) {
	query := psql.Select("id", "name").
		From("tags").
		OrderBy("name ASC")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*db.Tag
	for rows.Next() {
		var tag db.Tag
		err := r.DB.ScanRows(ctx, rows, &tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}
