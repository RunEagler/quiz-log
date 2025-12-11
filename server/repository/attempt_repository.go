package repository

import (
	"context"
	"quiz-log/models"
	"time"

	"github.com/uptrace/bun"
)

//go:generate mockgen -destination=mocks/mock_attempt_repository.go -package=mocks quiz-log/repository AttemptRepository

// AttemptRepository defines the interface for attempt repository operations
type AttemptRepository interface {
	Create(ctx context.Context, quizID int, startedAt, completedAt time.Time, score, totalQuestions int) (int, error)
	UpdateScore(ctx context.Context, attemptID, score int) error
	CountQuestionsByQuizID(ctx context.Context, quizID int) (int, error)
	GetCorrectAnswer(ctx context.Context, questionID int) (string, error)
	CreateAnswer(ctx context.Context, attemptID, questionID int, userAnswer string, isCorrect bool) error
	FindByID(ctx context.Context, attemptID int) (*models.Attempt, error)
	FindAll(ctx context.Context, quizID *int) ([]*models.Attempt, error)
	FindAnswersByAttemptID(ctx context.Context, attemptID int) ([]*models.Answer, error)
}

type attemptRepository struct {
	DB *bun.DB
}

func NewAttemptRepository(database *bun.DB) AttemptRepository {
	return &attemptRepository{DB: database}
}

// Create creates a new attempt and returns its ID
func (r *attemptRepository) Create(ctx context.Context, quizID int, startedAt, completedAt time.Time, score, totalQuestions int) (int, error) {
	var attemptID int

	query := psql.Insert("attempts").
		Columns("quiz_id", "started_at", "completed_at", "score", "total_questions").
		Values(quizID, startedAt, completedAt, score, totalQuestions).
		Suffix("RETURNING id")

	err := ExecQueryWithReturning[int](ctx, r.DB, query, &attemptID)
	if err != nil {
		return 0, err
	}

	return attemptID, nil
}

// UpdateScore updates the score of an attempt
func (r *attemptRepository) UpdateScore(ctx context.Context, attemptID, score int) error {
	query := psql.Update("attempts").
		Set("score", score).
		Where("id = ?", attemptID)

	_, err := ExecQuery(ctx, r.DB, query)
	if err != nil {
		return err
	}
	return nil
}

// CountQuestionsByQuizID counts questions for a quiz
func (r *attemptRepository) CountQuestionsByQuizID(ctx context.Context, quizID int) (int, error) {
	var count int

	query := psql.Select("COUNT(*)").
		From("questions").
		Where("quiz_id = ?", quizID)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetCorrectAnswer retrieves the correct answer for a question
func (r *attemptRepository) GetCorrectAnswer(ctx context.Context, questionID int) (string, error) {
	var correctAnswer string

	query := psql.Select("correct_answer").
		From("questions").
		Where("id = ?", questionID)

	err := ExecQueryWithReturning[string](ctx, r.DB, query, &correctAnswer)
	if err != nil {
		return "", err
	}

	return correctAnswer, nil
}

// CreateAnswer creates a new answer record
func (r *attemptRepository) CreateAnswer(ctx context.Context, attemptID, questionID int, userAnswer string, isCorrect bool) error {
	query := psql.Insert("answers").
		Columns("attempt_id", "question_id", "user_answer", "is_correct").
		Values(attemptID, questionID, userAnswer, isCorrect)

	_, err := ExecQuery(ctx, r.DB, query)
	if err != nil {
		return err
	}

	return nil
}

// FindByID retrieves an attempt by its ID
func (r *attemptRepository) FindByID(ctx context.Context, attemptID int) (*models.Attempt, error) {
	query := psql.Select("id", "quiz_id", "started_at", "completed_at", "score", "total_questions").
		From("attempts").
		Where("id = ?", attemptID)

	return FindOne[models.Attempt](ctx, r.DB, query)
}

// FindAll retrieves attempts, optionally filtered by quiz ID
func (r *attemptRepository) FindAll(ctx context.Context, quizID *int) ([]*models.Attempt, error) {
	queryBuilder := psql.Select("id", "quiz_id", "started_at", "completed_at", "score", "total_questions").
		From("attempts")

	if quizID != nil {
		queryBuilder = queryBuilder.Where("quiz_id = ?", *quizID)
	}

	queryBuilder = queryBuilder.OrderBy("started_at DESC")

	return FindAll[models.Attempt](ctx, r.DB, queryBuilder)
}

// FindAnswersByAttemptID retrieves all answers for an attempt
func (r *attemptRepository) FindAnswersByAttemptID(ctx context.Context, attemptID int) ([]*models.Answer, error) {
	query := psql.Select("id", "attempt_id", "question_id", "user_answer", "is_correct").
		From("answers").
		Where("attempt_id = ?", attemptID).
		OrderBy("id ASC")

	return FindAll[models.Answer](ctx, r.DB, query)
}
