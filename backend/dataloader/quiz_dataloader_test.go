package dataloader

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"quiz-log/graph/model"
	"quiz-log/models"
	"quiz-log/repository/mocks"
)

func TestBatchQuestionsByQuizID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	// Setup test data
	quizIDs := []int{1, 2, 3}
	now := time.Now()

	questionsMap := map[int][]*models.Question{
		1: {
			{
				ID:            1,
				QuizID:        1,
				Type:          "MULTIPLE_CHOICE",
				Content:       "What is Go?",
				Options:       []string{"A", "B", "C"},
				CorrectAnswer: "A",
				Explanation:   strPtr("Go is a programming language"),
				Difficulty:    "EASY",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			{
				ID:            2,
				QuizID:        1,
				Type:          "TRUE_FALSE",
				Content:       "Is Go compiled?",
				Options:       []string{"true", "false"},
				CorrectAnswer: "true",
				Explanation:   nil,
				Difficulty:    "MEDIUM",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
		},
		2: {
			{
				ID:            3,
				QuizID:        2,
				Type:          "SHORT_ANSWER",
				Content:       "What is a goroutine?",
				Options:       nil,
				CorrectAnswer: "A lightweight thread",
				Explanation:   nil,
				Difficulty:    "HARD",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
		},
		3: {}, // Empty questions for quiz 3
	}

	// Setup mock expectation
	mockRepo.EXPECT().
		FindQuestionsByQuizIDs(gomock.Any(), quizIDs).
		Return(questionsMap, nil)

	// Create batch function
	batchFunc := batchQuestionsByQuizID(mockRepo)

	// Execute batch function
	results := batchFunc(context.Background(), quizIDs)

	// Assertions
	assert.Len(t, results, 3)

	// Check quiz 1 results (2 questions)
	assert.NoError(t, results[0].Error)
	assert.Len(t, results[0].Data, 2)
	assert.Equal(t, "1", results[0].Data[0].ID)
	assert.Equal(t, "What is Go?", results[0].Data[0].Content)
	assert.Equal(t, model.QuestionTypeMultipleChoice, results[0].Data[0].Type)
	assert.Equal(t, model.DifficultyEasy, results[0].Data[0].Difficulty)
	assert.NotNil(t, results[0].Data[0].Explanation)
	assert.Equal(t, "Go is a programming language", *results[0].Data[0].Explanation)

	assert.Equal(t, "2", results[0].Data[1].ID)
	assert.Equal(t, "Is Go compiled?", results[0].Data[1].Content)
	assert.Equal(t, model.QuestionTypeTrueFalse, results[0].Data[1].Type)
	assert.Equal(t, model.DifficultyMedium, results[0].Data[1].Difficulty)
	assert.Nil(t, results[0].Data[1].Explanation)

	// Check quiz 2 results (1 question)
	assert.NoError(t, results[1].Error)
	assert.Len(t, results[1].Data, 1)
	assert.Equal(t, "3", results[1].Data[0].ID)
	assert.Equal(t, "What is a goroutine?", results[1].Data[0].Content)
	assert.Equal(t, model.QuestionTypeShortAnswer, results[1].Data[0].Type)
	assert.Equal(t, model.DifficultyHard, results[1].Data[0].Difficulty)

	// Check quiz 3 results (empty)
	assert.NoError(t, results[2].Error)
	assert.Len(t, results[2].Data, 0)
}

func TestBatchQuestionsByQuizID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	quizIDs := []int{1, 2, 3}
	expectedErr := errors.New("database error")

	// Setup mock expectation to return error
	mockRepo.EXPECT().
		FindQuestionsByQuizIDs(gomock.Any(), quizIDs).
		Return(nil, expectedErr)

	// Create batch function
	batchFunc := batchQuestionsByQuizID(mockRepo)

	// Execute batch function
	results := batchFunc(context.Background(), quizIDs)

	// Assertions - all results should have the same error
	assert.Len(t, results, 3)
	for i, result := range results {
		assert.Error(t, result.Error, "result %d should have error", i)
		assert.Equal(t, expectedErr, result.Error)
		assert.Nil(t, result.Data)
	}
}

func TestBatchQuestionsByQuizID_EmptyInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	quizIDs := []int{}

	// Setup mock expectation
	mockRepo.EXPECT().
		FindQuestionsByQuizIDs(gomock.Any(), quizIDs).
		Return(map[int][]*models.Question{}, nil)

	// Create batch function
	batchFunc := batchQuestionsByQuizID(mockRepo)

	// Execute batch function
	results := batchFunc(context.Background(), quizIDs)

	// Assertions
	assert.Len(t, results, 0)
}

func TestBatchTagsByQuizID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	// Setup test data
	quizIDs := []int{1, 2, 3}

	tagsMap := map[int][]*models.Tag{
		1: {
			{ID: 1, Name: "Go"},
			{ID: 2, Name: "Programming"},
			{ID: 3, Name: "Backend"},
		},
		2: {
			{ID: 2, Name: "Programming"},
			{ID: 4, Name: "Advanced"},
		},
		3: {}, // Empty tags for quiz 3
	}

	// Setup mock expectation
	mockRepo.EXPECT().
		FindTagsByQuizIDs(gomock.Any(), quizIDs).
		Return(tagsMap, nil)

	// Create batch function
	batchFunc := batchTagsByQuizID(mockRepo)

	// Execute batch function
	results := batchFunc(context.Background(), quizIDs)

	// Assertions
	assert.Len(t, results, 3)

	// Check quiz 1 results (3 tags)
	assert.NoError(t, results[0].Error)
	assert.Len(t, results[0].Data, 3)
	assert.Equal(t, "1", results[0].Data[0].ID)
	assert.Equal(t, "Go", results[0].Data[0].Name)
	assert.Equal(t, "2", results[0].Data[1].ID)
	assert.Equal(t, "Programming", results[0].Data[1].Name)
	assert.Equal(t, "3", results[0].Data[2].ID)
	assert.Equal(t, "Backend", results[0].Data[2].Name)

	// Check quiz 2 results (2 tags)
	assert.NoError(t, results[1].Error)
	assert.Len(t, results[1].Data, 2)
	assert.Equal(t, "2", results[1].Data[0].ID)
	assert.Equal(t, "Programming", results[1].Data[0].Name)
	assert.Equal(t, "4", results[1].Data[1].ID)
	assert.Equal(t, "Advanced", results[1].Data[1].Name)

	// Check quiz 3 results (empty)
	assert.NoError(t, results[2].Error)
	assert.Len(t, results[2].Data, 0)
}

func TestBatchTagsByQuizID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	quizIDs := []int{1, 2, 3}
	expectedErr := errors.New("database connection failed")

	// Setup mock expectation to return error
	mockRepo.EXPECT().
		FindTagsByQuizIDs(gomock.Any(), quizIDs).
		Return(nil, expectedErr)

	// Create batch function
	batchFunc := batchTagsByQuizID(mockRepo)

	// Execute batch function
	results := batchFunc(context.Background(), quizIDs)

	// Assertions - all results should have the same error
	assert.Len(t, results, 3)
	for i, result := range results {
		assert.Error(t, result.Error, "result %d should have error", i)
		assert.Equal(t, expectedErr, result.Error)
		assert.Nil(t, result.Data)
	}
}

func TestBatchTagsByQuizID_EmptyInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	quizIDs := []int{}

	// Setup mock expectation
	mockRepo.EXPECT().
		FindTagsByQuizIDs(gomock.Any(), quizIDs).
		Return(map[int][]*models.Tag{}, nil)

	// Create batch function
	batchFunc := batchTagsByQuizID(mockRepo)

	// Execute batch function
	results := batchFunc(context.Background(), quizIDs)

	// Assertions
	assert.Len(t, results, 0)
}

func TestBatchQuestionsByQuizID_OrderPreservation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	// Test that results are returned in the same order as input
	quizIDs := []int{5, 2, 8, 1}
	now := time.Now()

	questionsMap := map[int][]*models.Question{
		1: {{ID: 100, QuizID: 1, Type: "MULTIPLE_CHOICE", Content: "Q1", CorrectAnswer: "A", Difficulty: "EASY", CreatedAt: now, UpdatedAt: now}},
		2: {{ID: 200, QuizID: 2, Type: "TRUE_FALSE", Content: "Q2", CorrectAnswer: "true", Difficulty: "MEDIUM", CreatedAt: now, UpdatedAt: now}},
		5: {{ID: 500, QuizID: 5, Type: "SHORT_ANSWER", Content: "Q5", CorrectAnswer: "Answer", Difficulty: "HARD", CreatedAt: now, UpdatedAt: now}},
		8: {{ID: 800, QuizID: 8, Type: "MULTIPLE_CHOICE", Content: "Q8", CorrectAnswer: "B", Difficulty: "EASY", CreatedAt: now, UpdatedAt: now}},
	}

	mockRepo.EXPECT().
		FindQuestionsByQuizIDs(gomock.Any(), quizIDs).
		Return(questionsMap, nil)

	batchFunc := batchQuestionsByQuizID(mockRepo)
	results := batchFunc(context.Background(), quizIDs)

	// Verify order matches input order: 5, 2, 8, 1
	assert.Len(t, results, 4)
	assert.Equal(t, "500", results[0].Data[0].ID, "First result should be for quiz 5")
	assert.Equal(t, "200", results[1].Data[0].ID, "Second result should be for quiz 2")
	assert.Equal(t, "800", results[2].Data[0].ID, "Third result should be for quiz 8")
	assert.Equal(t, "100", results[3].Data[0].ID, "Fourth result should be for quiz 1")
}

func TestBatchTagsByQuizID_OrderPreservation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuizRepository(ctrl)

	// Test that results are returned in the same order as input
	quizIDs := []int{3, 1, 4}

	tagsMap := map[int][]*models.Tag{
		1: {{ID: 10, Name: "Tag1"}},
		3: {{ID: 30, Name: "Tag3"}},
		4: {{ID: 40, Name: "Tag4"}},
	}

	mockRepo.EXPECT().
		FindTagsByQuizIDs(gomock.Any(), quizIDs).
		Return(tagsMap, nil)

	batchFunc := batchTagsByQuizID(mockRepo)
	results := batchFunc(context.Background(), quizIDs)

	// Verify order matches input order: 3, 1, 4
	assert.Len(t, results, 3)
	assert.Equal(t, "30", results[0].Data[0].ID, "First result should be for quiz 3")
	assert.Equal(t, "10", results[1].Data[0].ID, "Second result should be for quiz 1")
	assert.Equal(t, "40", results[2].Data[0].ID, "Third result should be for quiz 4")
}

// Helper function to create string pointers
func strPtr(s string) *string {
	return &s
}
