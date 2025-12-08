package repository

import (
	"context"
	"quiz-log/db"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/uptrace/bun"
)

//go:generate mockgen -destination=mocks/mock_quiz_repository.go -package=mocks quiz-log/repository QuizRepository

// QuizRepository defines the interface for quiz repository operations
type QuizRepository interface {
	Create(ctx context.Context, title string, description *string) (int, error)
	Update(ctx context.Context, id int, title *string, description *string) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*db.Quiz, error)
	FindByID(ctx context.Context, id int) (*db.Quiz, error)
	FindQuestionsByQuizID(ctx context.Context, quizID int) ([]*db.Question, error)
	FindTagsByQuizID(ctx context.Context, quizID int) ([]*db.Tag, error)
	AssociateTags(ctx context.Context, quizID int, tagIDs []string) error
	ClearTags(ctx context.Context, quizID int) error
	FindQuestionsByQuizIDs(ctx context.Context, quizIDs []int) (map[int][]*db.Question, error)
	FindTagsByQuizIDs(ctx context.Context, quizIDs []int) (map[int][]*db.Tag, error)
}

type quizRepository struct {
	DB *bun.DB
}

func NewQuizRepository(database *bun.DB) QuizRepository {
	return &quizRepository{DB: database}
}

// Create creates a new quiz and returns its ID
func (r *quizRepository) Create(ctx context.Context, title string, description *string) (int, error) {
	var quizID int

	query := psql.Insert("quizzes").
		Columns("title", "description").
		Values(title, description).
		Suffix("RETURNING id")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&quizID)
	if err != nil {
		return 0, err
	}

	return quizID, nil
}

// Update updates an existing quiz
func (r *quizRepository) Update(ctx context.Context, id int, title *string, description *string) error {
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

	_, err := ExecQuery(ctx, r.DB, updateBuilder)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a quiz by ID
func (r *quizRepository) Delete(ctx context.Context, id int) error {
	query := psql.Delete("quizzes").
		Where("id = ?", id)

	_, err := ExecQuery(ctx, r.DB, query)
	if err != nil {
		return err
	}
	return nil
}

// FindAll retrieves all quizzes
func (r *quizRepository) FindAll(ctx context.Context) ([]*db.Quiz, error) {
	query := psql.Select("id", "title", "description", "created_at", "updated_at").
		From("quizzes").
		OrderBy("created_at DESC")

	return FindAll[db.Quiz](ctx, r.DB, query)
}

// FindByID retrieves a quiz by its ID
func (r *quizRepository) FindByID(ctx context.Context, id int) (*db.Quiz, error) {
	query := psql.Select("id", "title", "description", "created_at", "updated_at").
		From("quizzes").
		Where("id = ?", id)

	return FindOne[db.Quiz](ctx, r.DB, query)
}

// FindQuestionsByQuizID retrieves all questions for a quiz
func (r *quizRepository) FindQuestionsByQuizID(ctx context.Context, quizID int) ([]*db.Question, error) {
	query := psql.Select("id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at").
		From("questions").
		Where("quiz_id = ?", quizID).
		OrderBy("created_at ASC")

	return FindAll[db.Question](ctx, r.DB, query)
}

// FindTagsByQuizID retrieves all tags for a quiz
func (r *quizRepository) FindTagsByQuizID(ctx context.Context, quizID int) ([]*db.Tag, error) {
	query := psql.Select("t.id", "t.name").
		From("tags t").
		Join("quiz_tags qt ON t.id = qt.tag_id").
		Where("qt.quiz_id = ?", quizID).
		OrderBy("t.name ASC")

	return FindAll[db.Tag](ctx, r.DB, query)
}

// AssociateTags associates tags with a quiz
func (r *quizRepository) AssociateTags(ctx context.Context, quizID int, tagIDs []string) error {
	if len(tagIDs) == 0 {
		return nil
	}

	insertBuilder := psql.Insert("quiz_tags").Columns("quiz_id", "tag_id")
	for _, tagID := range tagIDs {
		id, _ := strconv.Atoi(tagID)
		insertBuilder = insertBuilder.Values(quizID, id)
	}
	_, err := ExecQuery(ctx, r.DB, insertBuilder)
	if err != nil {
		return err
	}
	return nil
}

// ClearTags removes all tag associations for a quiz
func (r *quizRepository) ClearTags(ctx context.Context, quizID int) error {
	query := psql.Delete("quiz_tags").
		Where("quiz_id = ?", quizID)

	_, err := ExecQuery(ctx, r.DB, query)
	if err != nil {
		return err
	}
	return nil
}

// FindQuestionsByQuizIDs retrieves all questions for multiple quizzes
func (r *quizRepository) FindQuestionsByQuizIDs(ctx context.Context, quizIDs []int) (map[int][]*db.Question, error) {
	if len(quizIDs) == 0 {
		return make(map[int][]*db.Question), nil
	}

	query := psql.Select("id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at").
		From("questions").
		Where(sq.Eq{"quiz_id": quizIDs}).
		OrderBy("quiz_id ASC", "created_at ASC")

	questions, err := FindAll[db.Question](ctx, r.DB, query)
	if err != nil {
		return nil, err
	}

	// Group questions by quiz_id
	result := make(map[int][]*db.Question)
	for _, q := range questions {
		result[q.QuizID] = append(result[q.QuizID], q)
	}

	return result, nil
}

// FindTagsByQuizIDs retrieves all tags for multiple quizzes
func (r *quizRepository) FindTagsByQuizIDs(ctx context.Context, quizIDs []int) (map[int][]*db.Tag, error) {
	if len(quizIDs) == 0 {
		return make(map[int][]*db.Tag), nil
	}

	query := psql.Select("t.id", "t.name", "qt.quiz_id").
		From("tags t").
		Join("quiz_tags qt ON t.id = qt.tag_id").
		Where(sq.Eq{"qt.quiz_id": quizIDs}).
		OrderBy("qt.quiz_id ASC", "t.name ASC")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	// Use raw query
	dbRows, err := r.DB.DB.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer dbRows.Close()

	// Group tags by quiz_id
	result := make(map[int][]*db.Tag)
	for dbRows.Next() {
		var tagID int
		var tagName string
		var quizID int
		err = dbRows.Scan(&tagID, &tagName, &quizID)
		if err != nil {
			return nil, err
		}
		tag := &db.Tag{
			ID:   tagID,
			Name: tagName,
		}
		result[quizID] = append(result[quizID], tag)
	}

	if err := dbRows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
