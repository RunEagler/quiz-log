package repository

import (
	"context"
	"quiz-log/db"
	"time"

	"github.com/uptrace/bun"
)

type AttemptRepository struct {
	DB *bun.DB
}

func NewAttemptRepository(database *bun.DB) *AttemptRepository {
	return &AttemptRepository{DB: database}
}

// Create creates a new attempt and returns its ID
func (r *AttemptRepository) Create(ctx context.Context, quizID int, startedAt, completedAt time.Time, score, totalQuestions int) (int, error) {
	var attemptID int

	query := psql.Insert("attempts").
		Columns("quiz_id", "started_at", "completed_at", "score", "total_questions").
		Values(quizID, startedAt, completedAt, score, totalQuestions).
		Suffix("RETURNING id")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&attemptID)
	if err != nil {
		return 0, err
	}

	return attemptID, nil
}

// UpdateScore updates the score of an attempt
func (r *AttemptRepository) UpdateScore(ctx context.Context, attemptID, score int) error {
	query := psql.Update("attempts").
		Set("score", score).
		Where("id = ?", attemptID)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// CountQuestionsByQuizID counts questions for a quiz
func (r *AttemptRepository) CountQuestionsByQuizID(ctx context.Context, quizID int) (int, error) {
	var count int

	query := psql.Select("COUNT(*)").
		From("questions").
		Where("quiz_id = ?", quizID)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&count)
	return count, err
}

// GetCorrectAnswer retrieves the correct answer for a question
func (r *AttemptRepository) GetCorrectAnswer(ctx context.Context, questionID int) (string, error) {
	var correctAnswer string

	query := psql.Select("correct_answer").
		From("questions").
		Where("id = ?", questionID)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return "", err
	}

	err = r.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&correctAnswer)
	return correctAnswer, err
}

// CreateAnswer creates a new answer record
func (r *AttemptRepository) CreateAnswer(ctx context.Context, attemptID, questionID int, userAnswer string, isCorrect bool) error {
	query := psql.Insert("answers").
		Columns("attempt_id", "question_id", "user_answer", "is_correct").
		Values(attemptID, questionID, userAnswer, isCorrect)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// FindByID retrieves an attempt by its ID
func (r *AttemptRepository) FindByID(ctx context.Context, attemptID int) (*db.Attempt, error) {
	query := psql.Select("id", "quiz_id", "started_at", "completed_at", "score", "total_questions").
		From("attempts").
		Where("id = ?", attemptID)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var attempt db.Attempt
	err = r.DB.NewRaw(sqlStr, args...).Scan(ctx, &attempt)
	if err != nil {
		return nil, err
	}

	return &attempt, nil
}

// FindAll retrieves attempts, optionally filtered by quiz ID
func (r *AttemptRepository) FindAll(ctx context.Context, quizID *int) ([]*db.Attempt, error) {
	queryBuilder := psql.Select("id", "quiz_id", "started_at", "completed_at", "score", "total_questions").
		From("attempts")

	if quizID != nil {
		queryBuilder = queryBuilder.Where("quiz_id = ?", *quizID)
	}

	queryBuilder = queryBuilder.OrderBy("started_at DESC")

	sqlStr, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attempts []*db.Attempt
	for rows.Next() {
		var attempt db.Attempt
		err := r.DB.ScanRows(ctx, rows, &attempt)
		if err != nil {
			return nil, err
		}
		attempts = append(attempts, &attempt)
	}

	return attempts, nil
}

// FindAnswersByAttemptID retrieves all answers for an attempt
func (r *AttemptRepository) FindAnswersByAttemptID(ctx context.Context, attemptID int) ([]*db.Answer, error) {
	query := psql.Select("id", "attempt_id", "question_id", "user_answer", "is_correct").
		From("answers").
		Where("attempt_id = ?", attemptID).
		OrderBy("id ASC")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var answers []*db.Answer
	for rows.Next() {
		var answer db.Answer
		err := r.DB.ScanRows(ctx, rows, &answer)
		if err != nil {
			return nil, err
		}
		answers = append(answers, &answer)
	}

	return answers, nil
}
