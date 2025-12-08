package services

import (
	"context"
	"testing"
	"time"

	"quiz-log/db"
	"quiz-log/graph/model"
	mocks "quiz-log/repository/mocks"

	"go.uber.org/mock/gomock"
)

func TestQuizService_CreateQuiz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)
	service := &QuizService{
		Repo: mockRepo,
	}

	ctx := context.Background()
	input := model.CreateQuizInput{
		Title:       "Test Quiz",
		Description: stringPtr("Test Description"),
		TagIDs:      []string{"1", "2"},
	}

	expectedQuizID := 1
	createdAt := time.Now()
	updatedAt := time.Now()

	// Expect Create to be called
	mockRepo.EXPECT().
		Create(ctx, input.Title, input.Description).
		Return(expectedQuizID, nil)

	// Expect AssociateTags to be called
	mockRepo.EXPECT().
		AssociateTags(ctx, expectedQuizID, input.TagIDs).
		Return(nil)

	// Expect FindByID to be called
	expectedQuiz := &db.Quiz{
		ID:          expectedQuizID,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	mockRepo.EXPECT().
		FindByID(ctx, expectedQuizID).
		Return(expectedQuiz, nil)

	// Execute
	result, err := service.CreateQuiz(ctx, input)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.ID != "1" {
		t.Errorf("expected ID '1', got '%s'", result.ID)
	}

	if result.Title != input.Title {
		t.Errorf("expected title '%s', got '%s'", input.Title, result.Title)
	}
}

func TestQuizService_UpdateQuiz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)
	service := &QuizService{
		Repo: mockRepo,
	}

	ctx := context.Background()
	quizID := "1"
	input := model.UpdateQuizInput{
		Title:       stringPtr("Updated Quiz"),
		Description: stringPtr("Updated Description"),
		TagIDs:      []string{"1", "3"},
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	// Expect Update to be called
	mockRepo.EXPECT().
		Update(ctx, 1, input.Title, input.Description).
		Return(nil)

	// Expect ClearTags to be called
	mockRepo.EXPECT().
		ClearTags(ctx, 1).
		Return(nil)

	// Expect AssociateTags to be called
	mockRepo.EXPECT().
		AssociateTags(ctx, 1, input.TagIDs).
		Return(nil)

	// Expect FindByID to be called
	expectedQuiz := &db.Quiz{
		ID:          1,
		Title:       *input.Title,
		Description: input.Description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	mockRepo.EXPECT().
		FindByID(ctx, 1).
		Return(expectedQuiz, nil)

	// Execute
	result, err := service.UpdateQuiz(ctx, quizID, input)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.Title != *input.Title {
		t.Errorf("expected title '%s', got '%s'", *input.Title, result.Title)
	}
}

func TestQuizService_DeleteQuiz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)
	service := &QuizService{
		Repo: mockRepo,
	}

	ctx := context.Background()
	quizID := "1"

	// Expect Delete to be called
	mockRepo.EXPECT().
		Delete(ctx, 1).
		Return(nil)

	// Execute
	result, err := service.DeleteQuiz(ctx, quizID)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !result {
		t.Error("expected true, got false")
	}
}

func TestQuizService_GetAllQuizzes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)
	service := &QuizService{
		Repo: mockRepo,
	}

	ctx := context.Background()
	createdAt := time.Now()
	updatedAt := time.Now()

	expectedQuizzes := []*db.Quiz{
		{
			ID:        1,
			Title:     "Quiz 1",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		{
			ID:        2,
			Title:     "Quiz 2",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	// Expect FindAll to be called
	mockRepo.EXPECT().
		FindAll(ctx).
		Return(expectedQuizzes, nil)

	// Execute
	result, err := service.GetAllQuizzes(ctx)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 quizzes, got %d", len(result))
	}

	if result[0].Title != "Quiz 1" {
		t.Errorf("expected title 'Quiz 1', got '%s'", result[0].Title)
	}
}

func TestQuizService_GetQuizByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)
	service := &QuizService{
		Repo: mockRepo,
	}

	ctx := context.Background()
	quizID := "1"
	createdAt := time.Now()
	updatedAt := time.Now()

	expectedQuiz := &db.Quiz{
		ID:          1,
		Title:       "Test Quiz",
		Description: stringPtr("Test Description"),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	// Expect FindByID to be called
	mockRepo.EXPECT().
		FindByID(ctx, 1).
		Return(expectedQuiz, nil)

	// Execute
	result, err := service.GetQuizByID(ctx, quizID)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.Title != "Test Quiz" {
		t.Errorf("expected title 'Test Quiz', got '%s'", result.Title)
	}
}

func TestQuizService_GetQuestionsByQuizID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)
	service := &QuizService{
		Repo: mockRepo,
	}

	ctx := context.Background()
	quizID := "1"
	createdAt := time.Now()
	updatedAt := time.Now()

	expectedQuestions := []*db.Question{
		{
			ID:            1,
			QuizID:        1,
			Type:          "MULTIPLE_CHOICE",
			Content:       "Question 1",
			Options:       []string{"A", "B", "C"},
			CorrectAnswer: "A",
			Difficulty:    "EASY",
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		},
	}

	// Expect FindQuestionsByQuizID to be called
	mockRepo.EXPECT().
		FindQuestionsByQuizID(ctx, 1).
		Return(expectedQuestions, nil)

	// Execute
	result, err := service.GetQuestionsByQuizID(ctx, quizID)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 question, got %d", len(result))
	}

	if result[0].Content != "Question 1" {
		t.Errorf("expected content 'Question 1', got '%s'", result[0].Content)
	}
}

func TestQuizService_GetTagsByQuizID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)
	service := &QuizService{
		Repo: mockRepo,
	}

	ctx := context.Background()
	quizID := "1"

	expectedTags := []*db.Tag{
		{ID: 1, Name: "Tag1"},
		{ID: 2, Name: "Tag2"},
	}

	// Expect FindTagsByQuizID to be called
	mockRepo.EXPECT().
		FindTagsByQuizID(ctx, 1).
		Return(expectedTags, nil)

	// Execute
	result, err := service.GetTagsByQuizID(ctx, quizID)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 tags, got %d", len(result))
	}

	if result[0].Name != "Tag1" {
		t.Errorf("expected name 'Tag1', got '%s'", result[0].Name)
	}
}

func stringPtr(s string) *string {
	return &s
}
