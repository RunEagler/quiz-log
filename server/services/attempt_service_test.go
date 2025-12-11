package services

import (
	"context"
	"quiz-log/models"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	"quiz-log/graph/model"
	mocks "quiz-log/repository/mocks"
)

// Helper function to convert int to *int
func intPtr(i int) *int {
	return &i
}

func TestAttemptService_SubmitAttempt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAttemptRepo := mocks.NewMockAttemptRepository(ctrl)
	mockQuestionRepo := mocks.NewMockQuestionRepository(ctrl)

	mockQuestionService := &QuestionService{
		Repo: mockQuestionRepo,
	}

	service := &AttemptService{
		Repo:            mockAttemptRepo,
		QuestionService: mockQuestionService,
	}

	ctx := context.Background()
	input := model.SubmitAttemptInput{
		QuizID: "1",
		Answers: []*model.AnswerInput{
			{QuestionID: "1", UserAnswer: "Paris"},
			{QuestionID: "2", UserAnswer: "London"},
		},
	}

	totalQuestions := 2
	attemptID := 1
	startTime := time.Now()

	// Expect CountQuestionsByQuizID to be called
	mockAttemptRepo.EXPECT().
		CountQuestionsByQuizID(ctx, 1).
		Return(totalQuestions, nil)

	// Expect Create to be called
	mockAttemptRepo.EXPECT().
		Create(ctx, 1, gomock.Any(), gomock.Any(), 0, totalQuestions).
		Return(attemptID, nil)

	// Expect GetCorrectAnswer for question 1
	mockAttemptRepo.EXPECT().
		GetCorrectAnswer(ctx, 1).
		Return("Paris", nil)

	// Expect CreateAnswer for question 1 (correct)
	mockAttemptRepo.EXPECT().
		CreateAnswer(ctx, attemptID, 1, "Paris", true).
		Return(nil)

	// Expect GetCorrectAnswer for question 2
	mockAttemptRepo.EXPECT().
		GetCorrectAnswer(ctx, 2).
		Return("Tokyo", nil)

	// Expect CreateAnswer for question 2 (incorrect)
	mockAttemptRepo.EXPECT().
		CreateAnswer(ctx, attemptID, 2, "London", false).
		Return(nil)

	// Expect UpdateScore to be called with 50% score (1 out of 2 correct)
	mockAttemptRepo.EXPECT().
		UpdateScore(ctx, attemptID, 50).
		Return(nil)

	// Expect FindByID to be called
	mockAttemptRepo.EXPECT().
		FindByID(ctx, attemptID).
		Return(&models.Attempt{
			ID:             attemptID,
			QuizID:         intPtr(1),
			StartedAt:      startTime,
			CompletedAt:    &startTime,
			Score:          50,
			TotalQuestions: totalQuestions,
		}, nil)

	// Expect GetQuestionByID for wrong question
	mockQuestionRepo.EXPECT().
		FindByID(ctx, 2).
		Return(&models.Question{
			ID:            2,
			QuizID:        intPtr(1),
			Type:          "MULTIPLE_CHOICE",
			Content:       "What is the capital of Japan?",
			Options:       []string{"Tokyo", "London", "Paris"},
			CorrectAnswer: "Tokyo",
			Difficulty:    "EASY",
			CreatedAt:     startTime,
			UpdatedAt:     startTime,
		}, nil)

	// Execute
	result, err := service.SubmitAttempt(ctx, input)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.Score != 50 {
		t.Errorf("expected score 50, got %d", result.Score)
	}

	if result.TotalQuestions != 2 {
		t.Errorf("expected total_questions 2, got %d", result.TotalQuestions)
	}

	if result.CorrectCount != 1 {
		t.Errorf("expected correct_count 1, got %d", result.CorrectCount)
	}

	if len(result.WrongQuestions) != 1 {
		t.Errorf("expected 1 wrong question, got %d", len(result.WrongQuestions))
	}
}

func TestAttemptService_GetAttempts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAttemptRepo := mocks.NewMockAttemptRepository(ctrl)
	service := &AttemptService{
		Repo: mockAttemptRepo,
	}

	ctx := context.Background()
	quizID := "1"
	quizIDInt := 1
	startTime := time.Now()

	expectedAttempts := []*models.Attempt{
		{
			ID:             1,
			QuizID:         quizIDInt,
			StartedAt:      startTime,
			CompletedAt:    &startTime,
			Score:          80,
			TotalQuestions: 10,
		},
		{
			ID:             2,
			QuizID:         quizIDInt,
			StartedAt:      startTime,
			CompletedAt:    &startTime,
			Score:          90,
			TotalQuestions: 10,
		},
	}

	// Expect FindAll to be called with quiz ID
	mockAttemptRepo.EXPECT().
		FindAll(ctx, &quizIDInt).
		Return(expectedAttempts, nil)

	// Execute
	result, err := service.GetAttempts(ctx, &quizID)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 attempts, got %d", len(result))
	}

	if result[0].Score != 80 {
		t.Errorf("expected score 80, got %d", result[0].Score)
	}

	if result[1].Score != 90 {
		t.Errorf("expected score 90, got %d", result[1].Score)
	}
}

func TestAttemptService_GetAttempts_WithoutFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAttemptRepo := mocks.NewMockAttemptRepository(ctrl)
	service := &AttemptService{
		Repo: mockAttemptRepo,
	}

	ctx := context.Background()
	startTime := time.Now()

	expectedAttempts := []*models.Attempt{
		{
			ID:             1,
			QuizID:         1,
			StartedAt:      startTime,
			CompletedAt:    &startTime,
			Score:          80,
			TotalQuestions: 10,
		},
	}

	// Expect FindAll to be called without quiz ID filter
	mockAttemptRepo.EXPECT().
		FindAll(ctx, nil).
		Return(expectedAttempts, nil)

	// Execute
	result, err := service.GetAttempts(ctx, nil)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 attempt, got %d", len(result))
	}
}

func TestAttemptService_GetAnswersByAttemptID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAttemptRepo := mocks.NewMockAttemptRepository(ctrl)
	service := &AttemptService{
		Repo: mockAttemptRepo,
	}

	ctx := context.Background()
	attemptID := "1"

	expectedAnswers := []*models.Answer{
		{
			ID:         1,
			AttemptID:  1,
			QuestionID: 1,
			UserAnswer: "Paris",
			IsCorrect:  true,
		},
		{
			ID:         2,
			AttemptID:  1,
			QuestionID: 2,
			UserAnswer: "London",
			IsCorrect:  false,
		},
	}

	// Expect FindAnswersByAttemptID to be called
	mockAttemptRepo.EXPECT().
		FindAnswersByAttemptID(ctx, 1).
		Return(expectedAnswers, nil)

	// Execute
	result, err := service.GetAnswersByAttemptID(ctx, attemptID)

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 answers, got %d", len(result))
	}

	if result[0].UserAnswer != "Paris" {
		t.Errorf("expected user_answer 'Paris', got '%s'", result[0].UserAnswer)
	}

	if !result[0].IsCorrect {
		t.Error("expected first answer to be correct")
	}

	if result[1].IsCorrect {
		t.Error("expected second answer to be incorrect")
	}
}
