package repository

import (
	"context"
	"quiz-log/db"
	"strconv"

	"github.com/lib/pq"
	"github.com/uptrace/bun"
)

// questionRepository defines the interface for question repository operations
//
//go:generate mockgen -destination=mocks/mock_question_repository.go -package=mocks quiz-log/repository QuestionRepository
type QuestionRepository interface {
	Create(ctx context.Context, quizID, questionType, content string, options []string, correctAnswer string, explanation *string, difficulty string) (int, error)
	Update(ctx context.Context, id int, questionType *string, content *string, options []string, correctAnswer *string, explanation *string, difficulty *string) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, quizID *int) ([]*db.Question, error)
	FindByID(ctx context.Context, id int) (*db.Question, error)
	FindWrongQuestions(ctx context.Context) ([]*db.Question, error)
	FindTagsByQuestionID(ctx context.Context, questionID int) ([]*db.Tag, error)
	AssociateTags(ctx context.Context, questionID int, tagIDs []string) error
	ClearTags(ctx context.Context, questionID int) error
}

type questionRepository struct {
	DB *bun.DB
}

func NewQuestionRepository(database *bun.DB) QuestionRepository {
	return &questionRepository{DB: database}
}

// Create creates a new question and returns its ID
func (r *questionRepository) Create(ctx context.Context, quizID, questionType, content string, options []string, correctAnswer string, explanation *string, difficulty string) (int, error) {
	var questionID int

	query := psql.Insert("questions").
		Columns("quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty").
		Values(quizID, questionType, content, pq.Array(options), correctAnswer, explanation, difficulty).
		Suffix("RETURNING id")

	err := ExecQueryWithReturning[int](ctx, r.DB, query, &questionID)
	if err != nil {
		return 0, err
	}

	return questionID, nil
}

// Update updates an existing question
func (r *questionRepository) Update(ctx context.Context, id int, questionType *string, content *string, options []string, correctAnswer *string, explanation *string, difficulty *string) error {
	query := psql.Update("questions").Where("id = ?", id)
	hasUpdates := false

	if questionType != nil {
		query = query.Set("type", *questionType)
		hasUpdates = true
	}

	if content != nil {
		query = query.Set("content", *content)
		hasUpdates = true
	}

	if options != nil {
		query = query.Set("options", pq.Array(options))
		hasUpdates = true
	}

	if correctAnswer != nil {
		query = query.Set("correct_answer", *correctAnswer)
		hasUpdates = true
	}

	if explanation != nil {
		query = query.Set("explanation", *explanation)
		hasUpdates = true
	}

	if difficulty != nil {
		query = query.Set("difficulty", *difficulty)
		hasUpdates = true
	}

	if !hasUpdates {
		return nil
	}

	query = query.Set("updated_at", psql.Select("NOW()"))

	_, err := ExecQuery(ctx, r.DB, query)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a question by ID
func (r *questionRepository) Delete(ctx context.Context, id int) error {
	query := psql.Delete("questions").
		Where("id = ?", id)

	_, err := ExecQuery(ctx, r.DB, query)
	if err != nil {
		return err
	}

	return nil
}

// FindAll retrieves all questions, optionally filtered by quiz ID
func (r *questionRepository) FindAll(ctx context.Context, quizID *int) ([]*db.Question, error) {
	queryBuilder := psql.Select("id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at").
		From("questions")

	if quizID != nil {
		queryBuilder = queryBuilder.Where("quiz_id = ?", *quizID)
	}

	queryBuilder = queryBuilder.OrderBy("created_at ASC")

	return FindAll[db.Question](ctx, r.DB, queryBuilder)
}

// FindByID retrieves a question by its ID
func (r *questionRepository) FindByID(ctx context.Context, id int) (*db.Question, error) {
	query := psql.Select("id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at").
		From("questions").
		Where("id = ?", id)

	return FindOne[db.Question](ctx, r.DB, query)
}

// FindWrongQuestions retrieves questions that were answered incorrectly
func (r *questionRepository) FindWrongQuestions(ctx context.Context) ([]*db.Question, error) {
	query := psql.Select("DISTINCT q.id", "q.quiz_id", "q.type", "q.content", "q.options", "q.correct_answer", "q.explanation", "q.difficulty", "q.created_at", "q.updated_at").
		From("questions q").
		Join("answers a ON q.id = a.question_id").
		Where("a.is_correct = false").
		OrderBy("q.created_at DESC")

	return FindAll[db.Question](ctx, r.DB, query)
}

// FindTagsByQuestionID retrieves all tags for a question
func (r *questionRepository) FindTagsByQuestionID(ctx context.Context, questionID int) ([]*db.Tag, error) {
	query := psql.Select("t.id", "t.name").
		From("tags t").
		Join("question_tags qt ON t.id = qt.tag_id").
		Where("qt.question_id = ?", questionID).
		OrderBy("t.name ASC")

	return FindAll[db.Tag](ctx, r.DB, query)
}

// AssociateTags associates tags with a question
func (r *questionRepository) AssociateTags(ctx context.Context, questionID int, tagIDs []string) error {
	if len(tagIDs) == 0 {
		return nil
	}

	query := psql.Insert("question_tags").Columns("question_id", "tag_id")
	for _, tagID := range tagIDs {
		id, _ := strconv.Atoi(tagID)
		query = query.Values(questionID, id)
	}

	_, err := ExecQuery(ctx, r.DB, query)
	if err != nil {
		return err
	}

	return nil
}

// ClearTags removes all tag associations for a question
func (r *questionRepository) ClearTags(ctx context.Context, questionID int) error {
	deleteSql, deleteArgs, err := psql.Delete("question_tags").
		Where("question_id = ?", questionID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, deleteSql, deleteArgs...)
	return err
}
