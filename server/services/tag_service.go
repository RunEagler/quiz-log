package services

import (
	"context"
	"quiz-log/db"
	"strconv"

	"github.com/uptrace/bun"

	"quiz-log/graph/model"
	"quiz-log/repository"
)

type TagService struct {
	DB   *bun.DB
	Repo repository.TagRepository
}

func NewTagService(database *bun.DB) *TagService {
	return &TagService{
		DB:   database,
		Repo: repository.NewTagRepository(database),
	}
}

// CreateTag creates a new tag or returns existing one
func (s *TagService) CreateTag(ctx context.Context, name string) (*model.Tag, error) {
	tagID, err := s.Repo.Create(ctx, name)
	if err != nil {
		return nil, err
	}

	return &model.Tag{
		ID:   strconv.Itoa(tagID),
		Name: name,
	}, nil
}

// GetAllTags retrieves all tags
func (s *TagService) GetAllTags(ctx context.Context) ([]*model.Tag, error) {
	dbTags, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	for _, dbTag := range dbTags {
		tags = append(tags, db.TagToGraphQL(dbTag))
	}

	return tags, nil
}
