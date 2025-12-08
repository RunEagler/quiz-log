package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAttemptRepository_Create(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	quizID := 1
	startedAt := time.Now()
	completedAt := time.Now().Add(5 * time.Minute)
	score := 8
	totalQuestions := 10
	expectedID := 1

	mock.ExpectQuery(`INSERT INTO attempts`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

	ctx := context.Background()
	id, err := repo.Create(ctx, quizID, startedAt, completedAt, score, totalQuestions)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if id != expectedID {
		t.Errorf("expected id %d, got %d", expectedID, id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_UpdateScore(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	attemptID := 1
	score := 9

	mock.ExpectExec(`UPDATE attempts`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := context.Background()
	err := repo.UpdateScore(ctx, attemptID, score)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_CountQuestionsByQuizID(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	quizID := 1
	expectedCount := 10

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM questions`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

	ctx := context.Background()
	count, err := repo.CountQuestionsByQuizID(ctx, quizID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if count != expectedCount {
		t.Errorf("expected count %d, got %d", expectedCount, count)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_GetCorrectAnswer(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	questionID := 1
	expectedAnswer := "Paris"

	mock.ExpectQuery(`SELECT correct_answer FROM questions`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"correct_answer"}).AddRow(expectedAnswer))

	ctx := context.Background()
	answer, err := repo.GetCorrectAnswer(ctx, questionID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if answer != expectedAnswer {
		t.Errorf("expected answer %s, got %s", expectedAnswer, answer)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_CreateAnswer(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	attemptID := 1
	questionID := 1
	userAnswer := "Paris"
	isCorrect := true

	mock.ExpectExec(`INSERT INTO answers`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	err := repo.CreateAnswer(ctx, attemptID, questionID, userAnswer, isCorrect)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_FindByID(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	attemptID := 1
	quizID := 1
	startedAt := time.Now()
	completedAt := time.Now().Add(5 * time.Minute)
	score := 8
	totalQuestions := 10

	rows := sqlmock.NewRows([]string{"id", "quiz_id", "started_at", "completed_at", "score", "total_questions"}).
		AddRow(attemptID, quizID, startedAt, completedAt, score, totalQuestions)

	mock.ExpectQuery(`SELECT (.+) FROM attempts`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	ctx := context.Background()
	attempt, err := repo.FindByID(ctx, attemptID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if attempt == nil {
		t.Fatal("expected attempt, got nil")
	}

	if attempt.ID != attemptID {
		t.Errorf("expected id %d, got %d", attemptID, attempt.ID)
	}

	if attempt.QuizID != quizID {
		t.Errorf("expected quiz_id %d, got %d", quizID, attempt.QuizID)
	}

	if attempt.Score != score {
		t.Errorf("expected score %d, got %d", score, attempt.Score)
	}

	if attempt.TotalQuestions != totalQuestions {
		t.Errorf("expected total_questions %d, got %d", totalQuestions, attempt.TotalQuestions)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_FindByID_NotFound(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	attemptID := 999

	mock.ExpectQuery(`SELECT (.+) FROM attempts`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnError(sql.ErrNoRows)

	ctx := context.Background()
	attempt, err := repo.FindByID(ctx, attemptID)

	if err != sql.ErrNoRows {
		t.Errorf("expected sql.ErrNoRows, got %v", err)
	}

	if attempt != nil {
		t.Errorf("expected nil attempt, got %v", attempt)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_FindAll(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	startedAt := time.Now()
	completedAt := time.Now().Add(5 * time.Minute)

	rows := sqlmock.NewRows([]string{"id", "quiz_id", "started_at", "completed_at", "score", "total_questions"}).
		AddRow(1, 1, startedAt, completedAt, 8, 10).
		AddRow(2, 1, startedAt, completedAt, 9, 10)

	mock.ExpectQuery(`SELECT (.+) FROM attempts`).
		WillReturnRows(rows)

	ctx := context.Background()
	attempts, err := repo.FindAll(ctx, nil)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(attempts) != 2 {
		t.Errorf("expected 2 attempts, got %d", len(attempts))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_FindAll_WithQuizID(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	quizID := 1
	startedAt := time.Now()
	completedAt := time.Now().Add(5 * time.Minute)

	rows := sqlmock.NewRows([]string{"id", "quiz_id", "started_at", "completed_at", "score", "total_questions"}).
		AddRow(1, quizID, startedAt, completedAt, 8, 10).
		AddRow(2, quizID, startedAt, completedAt, 9, 10)

	mock.ExpectQuery(`SELECT (.+) FROM attempts`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	ctx := context.Background()
	attempts, err := repo.FindAll(ctx, &quizID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(attempts) != 2 {
		t.Errorf("expected 2 attempts, got %d", len(attempts))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestAttemptRepository_FindAnswersByAttemptID(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewAttemptRepository(bunDB)

	attemptID := 1

	rows := sqlmock.NewRows([]string{"id", "attempt_id", "question_id", "user_answer", "is_correct"}).
		AddRow(1, attemptID, 1, "Paris", true).
		AddRow(2, attemptID, 2, "London", false)

	mock.ExpectQuery(`SELECT (.+) FROM answers`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	ctx := context.Background()
	answers, err := repo.FindAnswersByAttemptID(ctx, attemptID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(answers) != 2 {
		t.Errorf("expected 2 answers, got %d", len(answers))
	}

	if answers[0].UserAnswer != "Paris" {
		t.Errorf("expected user_answer 'Paris', got '%s'", answers[0].UserAnswer)
	}

	if !answers[0].IsCorrect {
		t.Error("expected first answer to be correct")
	}

	if answers[1].IsCorrect {
		t.Error("expected second answer to be incorrect")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
