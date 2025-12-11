package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func setupMockDB(t *testing.T) (*bun.DB, sqlmock.Sqlmock, func()) {
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
		sqlmock.MonitorPingsOption(false),
	)
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	// Use QueryRowContext through bun.DB.DB to properly interact with sqlmock
	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	cleanup := func() {
		bunDB.Close()
	}

	return bunDB, mock, cleanup
}

func TestQuizRepository_Create(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	title := "Test Quiz"
	description := "Test Description"
	expectedID := 1

	mock.ExpectQuery(`INSERT INTO quizzes`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

	ctx := context.Background()
	id, err := repo.Create(ctx, title, &description)

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

func TestQuizRepository_Create_WithNullDescription(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	title := "Test Quiz"
	expectedID := 1

	mock.ExpectQuery(`INSERT INTO quizzes`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

	ctx := context.Background()
	id, err := repo.Create(ctx, title, nil)

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

func TestQuizRepository_Update(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	id := 1
	title := "Updated Quiz"
	description := "Updated Description"

	mock.ExpectExec(`UPDATE quizzes`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := context.Background()
	err := repo.Update(ctx, id, &title, &description)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_Update_NoUpdates(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	id := 1

	ctx := context.Background()
	err := repo.Update(ctx, id, nil, nil)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_Delete(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	id := 1

	mock.ExpectExec(`DELETE FROM quizzes`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := context.Background()
	err := repo.Delete(ctx, id)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_FindByID(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	id := 1
	title := "Test Quiz"
	description := "Test Description"
	createdAt := time.Now()
	updatedAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "updated_at"}).
		AddRow(id, title, description, createdAt, updatedAt)

	mock.ExpectQuery(`SELECT (.+) FROM quizzes`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	ctx := context.Background()
	quiz, err := repo.FindByID(ctx, id)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if quiz == nil {
		t.Fatal("expected quiz, got nil")
	}

	if quiz.ID != id {
		t.Errorf("expected id %d, got %d", id, quiz.ID)
	}

	if quiz.Title != title {
		t.Errorf("expected title %s, got %s", title, quiz.Title)
	}

	if quiz.Description == nil || *quiz.Description != description {
		t.Errorf("expected description %s, got %v", description, quiz.Description)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_FindByID_NotFound(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	id := 999

	mock.ExpectQuery(`SELECT (.+) FROM quizzes`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnError(sql.ErrNoRows)

	ctx := context.Background()
	quiz, err := repo.FindByID(ctx, id)

	if err != sql.ErrNoRows {
		t.Errorf("expected sql.ErrNoRows, got %v", err)
	}

	if quiz != nil {
		t.Errorf("expected nil quiz, got %v", quiz)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_FindAll(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	createdAt := time.Now()
	updatedAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "updated_at"}).
		AddRow(1, "Quiz 1", "Description 1", createdAt, updatedAt).
		AddRow(2, "Quiz 2", "Description 2", createdAt, updatedAt)

	mock.ExpectQuery(`SELECT (.+) FROM quizzes`).
		WillReturnRows(rows)

	ctx := context.Background()
	quizzes, err := repo.FindAll(ctx)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(quizzes) != 2 {
		t.Errorf("expected 2 quizzes, got %d", len(quizzes))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_FindQuestionsByQuizID(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	quizID := 1
	createdAt := time.Now()
	updatedAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at"}).
		AddRow(1, quizID, "MULTIPLE_CHOICE", "Question 1", pq.Array([]string{"A", "B", "C"}), "A", "Explanation", "EASY", createdAt, updatedAt).
		AddRow(2, quizID, "TRUE_FALSE", "Question 2", pq.Array([]string{"True", "False"}), "True", "Explanation", "MEDIUM", createdAt, updatedAt)

	mock.ExpectQuery(`SELECT (.+) FROM questions`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	ctx := context.Background()
	questions, err := repo.FindQuestionsByQuizID(ctx, quizID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(questions) != 2 {
		t.Errorf("expected 2 questions, got %d", len(questions))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_AssociateTags(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	quizID := 1
	tagIDs := []string{"1", "2", "3"}

	mock.ExpectExec(`INSERT INTO quiz_tags`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 3))

	ctx := context.Background()
	err := repo.AssociateTags(ctx, quizID, tagIDs)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_AssociateTags_Empty(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	quizID := 1
	tagIDs := []string{}

	ctx := context.Background()
	err := repo.AssociateTags(ctx, quizID, tagIDs)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_ClearTags(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	quizID := 1

	mock.ExpectExec(`DELETE FROM quiz_tags`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 3))

	ctx := context.Background()
	err := repo.ClearTags(ctx, quizID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestQuizRepository_FindTagsByQuizID(t *testing.T) {
	bunDB, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewQuizRepository(bunDB)

	quizID := 1

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Tag 1").
		AddRow(2, "Tag 2")

	mock.ExpectQuery(`SELECT (.+) FROM tags`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	ctx := context.Background()
	tags, err := repo.FindTagsByQuizID(ctx, quizID)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(tags) != 2 {
		t.Errorf("expected 2 tags, got %d", len(tags))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
