package services

import (
	"context"
	"quiz-log/db"
	"quiz-log/graph/model"
	"quiz-log/repository"
	"strconv"
	"time"

	"github.com/uptrace/bun"
)

type AttemptService struct {
	DB              *bun.DB
	Repo            repository.AttemptRepository
	QuestionService *QuestionService
}

func NewAttemptService(database *bun.DB, questionService *QuestionService) *AttemptService {
	return &AttemptService{
		DB:              database,
		Repo:            repository.NewAttemptRepository(database),
		QuestionService: questionService,
	}
}

// SubmitAttempt creates a new attempt and processes answers
func (s *AttemptService) SubmitAttempt(ctx context.Context, input model.SubmitAttemptInput) (*model.AttemptResult, error) {
	quizID, _ := strconv.Atoi(input.QuizID)

	// Get all questions for the quiz to calculate total
	totalQuestions, err := s.Repo.CountQuestionsByQuizID(ctx, quizID)
	if err != nil {
		return nil, err
	}

	// Create attempt record
	attemptID, err := s.Repo.Create(ctx, quizID, time.Now(), time.Now(), 0, totalQuestions)
	if err != nil {
		return nil, err
	}

	// Process answers and calculate score
	correctCount := 0
	var wrongQuestionIDs []int

	for _, answer := range input.Answers {
		questionID, _ := strconv.Atoi(answer.QuestionID)

		// Get correct answer
		correctAnswer, err := s.Repo.GetCorrectAnswer(ctx, questionID)
		if err != nil {
			return nil, err
		}

		// Check if answer is correct
		isCorrect := answer.UserAnswer == correctAnswer
		if isCorrect {
			correctCount++
		} else {
			wrongQuestionIDs = append(wrongQuestionIDs, questionID)
		}

		// Save answer
		err = s.Repo.CreateAnswer(ctx, attemptID, questionID, answer.UserAnswer, isCorrect)
		if err != nil {
			return nil, err
		}
	}

	// Calculate score percentage
	score := 0
	if totalQuestions > 0 {
		score = (correctCount * 100) / totalQuestions
	}

	// Update attempt with final score
	err = s.Repo.UpdateScore(ctx, attemptID, score)
	if err != nil {
		return nil, err
	}

	// Get attempt details
	dbAttempt, err := s.Repo.FindByID(ctx, attemptID)
	if err != nil {
		return nil, err
	}

	// Get wrong questions
	var wrongQuestions []*model.Question
	if len(wrongQuestionIDs) > 0 {
		for _, qid := range wrongQuestionIDs {
			q, err := s.QuestionService.GetQuestionByID(ctx, strconv.Itoa(qid))
			if err != nil {
				return nil, err
			}
			if q != nil {
				wrongQuestions = append(wrongQuestions, q)
			}
		}
	}

	return &model.AttemptResult{
		Attempt:        db.AttemptToGraphQL(dbAttempt),
		Score:          score,
		TotalQuestions: totalQuestions,
		CorrectCount:   correctCount,
		WrongQuestions: wrongQuestions,
	}, nil
}

// GetAttempts retrieves attempts, optionally filtered by quiz ID
func (s *AttemptService) GetAttempts(ctx context.Context, quizID *string) ([]*model.Attempt, error) {
	var qid *int
	if quizID != nil {
		id, _ := strconv.Atoi(*quizID)
		qid = &id
	}

	dbAttempts, err := s.Repo.FindAll(ctx, qid)
	if err != nil {
		return nil, err
	}

	var attempts []*model.Attempt
	for _, dbAttempt := range dbAttempts {
		attempts = append(attempts, db.AttemptToGraphQL(dbAttempt))
	}

	return attempts, nil
}

// GetAnswersByAttemptID retrieves all answers for an attempt
func (s *AttemptService) GetAnswersByAttemptID(ctx context.Context, attemptID string) ([]*model.Answer, error) {
	id, err := strconv.Atoi(attemptID)
	if err != nil {
		return nil, err
	}

	dbAnswers, err := s.Repo.FindAnswersByAttemptID(ctx, id)
	if err != nil {
		return nil, err
	}

	var answers []*model.Answer
	for _, dbAnswer := range dbAnswers {
		answers = append(answers, db.AnswerToGraphQL(dbAnswer))
	}

	return answers, nil
}
