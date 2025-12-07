package repository

import (
	"context"
	"database/sql"
	"quiz-log/db"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/uptrace/bun"
)

type QuizRepository struct {
	DB *bun.DB
}

func NewQuizRepository(database *bun.DB) *QuizRepository {
	return &QuizRepository{DB: database}
}

// Create creates a new quiz and returns its ID
func (r *QuizRepository) Create(ctx context.Context, title string, description *string) (int, error) {
	var quizID int

	query := psql.Insert("quizzes").
		Columns("title", "description").
		Values(title, description).
		Suffix("RETURNING id")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&quizID)
	if err != nil {
		return 0, err
	}

	return quizID, nil
}

// Update updates an existing quiz
func (r *QuizRepository) Update(ctx context.Context, id int, title *string, description *string) error {
	updateBuilder := psql.Update("quizzes").Where("id = ?", id)
	hasUpdates := false

	if title != nil {
		updateBuilder = updateBuilder.Set("title", *title)
		hasUpdates = true
	}

	if description != nil {
		updateBuilder = updateBuilder.Set("description", *description)
		hasUpdates = true
	}

	if !hasUpdates {
		return nil
	}

	updateBuilder = updateBuilder.Set("updated_at", sq.Expr("NOW()"))

	sqlStr, args, err := updateBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// Delete deletes a quiz by ID
func (r *QuizRepository) Delete(ctx context.Context, id int) error {
	sqlStr, args, err := psql.Delete("quizzes").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// FindAll retrieves all quizzes
func (r *QuizRepository) FindAll(ctx context.Context) ([]db.Quiz, error) {
	query := psql.Select("id", "title", "description", "created_at", "updated_at").
		From("quizzes").
		OrderBy("created_at DESC")

	return FindAll[db.Quiz](ctx, r.DB, query)
}

// FindByID retrieves a quiz by its ID
func (r *QuizRepository) FindByID(ctx context.Context, id int) (*db.Quiz, error) {
	query := psql.Select("id", "title", "description", "created_at", "updated_at").
		From("quizzes").
		Where("id = ?", id)

	return FindOne[db.Quiz](ctx, r.DB, query)
}

// FindQuestionsByQuizID retrieves all questions for a quiz
func (r *QuizRepository) FindQuestionsByQuizID(ctx context.Context, quizID int) ([]*db.Question, error) {
	query := psql.Select("id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at").
		From("questions").
		Where("quiz_id = ?", quizID).
		OrderBy("created_at ASC")

	return FindAll[*db.Question](ctx, r.DB, query)
}

// FindTagsByQuizID retrieves all tags for a quiz
func (r *QuizRepository) FindTagsByQuizID(ctx context.Context, quizID int) ([]*db.Tag, error) {
	query := psql.Select("t.id", "t.name").
		From("tags t").
		Join("quiz_tags qt ON t.id = qt.tag_id").
		Where("qt.quiz_id = ?", quizID).
		OrderBy("t.name ASC")

	return FindAll[db.Tag](ctx, r.DB, query)
}

// AssociateTags associates tags with a quiz
func (r *QuizRepository) AssociateTags(ctx context.Context, quizID int, tagIDs []string) error {
	if len(tagIDs) == 0 {
		return nil
	}

	insertBuilder := psql.Insert("quiz_tags").Columns("quiz_id", "tag_id")
	for _, tagID := range tagIDs {
		id, _ := strconv.Atoi(tagID)
		insertBuilder = insertBuilder.Values(quizID, id)
	}

	sqlStr, args, err := insertBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// ClearTags removes all tag associations for a quiz
func (r *QuizRepository) ClearTags(ctx context.Context, quizID int) error {
	deleteSql, deleteArgs, err := psql.Delete("quiz_tags").
		Where("quiz_id = ?", quizID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, deleteSql, deleteArgs...)
	return err
}
