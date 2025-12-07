package services

import (
	"context"
	"quiz-log/db"
	"quiz-log/graph/model"
	"quiz-log/repository"
	"strconv"

	"github.com/uptrace/bun"
)

type QuizService struct {
	DB   *bun.DB
	Repo *repository.QuizRepository
}

func NewQuizService(database *bun.DB) *QuizService {
	return &QuizService{
		DB:   database,
		Repo: repository.NewQuizRepository(database),
	}
}

// CreateQuiz creates a new quiz with the given input
func (s *QuizService) CreateQuiz(ctx context.Context, input model.CreateQuizInput) (*model.Quiz, error) {
	quizID, err := s.Repo.Create(ctx, input.Title, input.Description)
	if err != nil {
		return nil, err
	}

	// Associate tags
	if len(input.TagIDs) > 0 {
		err = s.Repo.AssociateTags(ctx, quizID, input.TagIDs)
		if err != nil {
			return nil, err
		}
	}

	return s.GetQuizByID(ctx, strconv.Itoa(quizID))
}

// UpdateQuiz updates an existing quiz
func (s *QuizService) UpdateQuiz(ctx context.Context, id string, input model.UpdateQuizInput) (*model.Quiz, error) {
	quizID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	err = s.Repo.Update(ctx, quizID, input.Title, input.Description)
	if err != nil {
		return nil, err
	}

	// Update tags
	if input.TagIDs != nil {
		// Delete existing tags
		err = s.Repo.ClearTags(ctx, quizID)
		if err != nil {
			return nil, err
		}

		// Insert new tags
		if len(input.TagIDs) > 0 {
			err = s.Repo.AssociateTags(ctx, quizID, input.TagIDs)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.GetQuizByID(ctx, id)
}

// DeleteQuiz deletes a quiz by ID
func (s *QuizService) DeleteQuiz(ctx context.Context, id string) (bool, error) {
	quizID, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}

	err = s.Repo.Delete(ctx, quizID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetAllQuizzes retrieves all quizzes
func (s *QuizService) GetAllQuizzes(ctx context.Context) ([]*model.Quiz, error) {
	dbQuizzes, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var quizzes []*model.Quiz
	for _, dbQuiz := range dbQuizzes {
		quizzes = append(quizzes, db.QuizToGraphQL(dbQuiz))
	}

	return quizzes, nil
}

// GetQuizByID retrieves a quiz by its ID
func (s *QuizService) GetQuizByID(ctx context.Context, id string) (*model.Quiz, error) {
	quizID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	dbQuiz, err := s.Repo.FindByID(ctx, quizID)
	if err != nil {
		return nil, err
	}

	if dbQuiz == nil {
		return nil, nil
	}

	return db.QuizToGraphQL(dbQuiz), nil
}

// GetQuestionsByQuizID retrieves all questions for a quiz
func (s *QuizService) GetQuestionsByQuizID(ctx context.Context, quizID string) ([]*model.Question, error) {
	id, err := strconv.Atoi(quizID)
	if err != nil {
		return nil, err
	}

	dbQuestions, err := s.Repo.FindQuestionsByQuizID(ctx, id)
	if err != nil {
		return nil, err
	}

	var questions []*model.Question
	for _, dbQuestion := range dbQuestions {
		questions = append(questions, db.QuestionToGraphQL(dbQuestion))
	}

	return questions, nil
}

// GetTagsByQuizID retrieves all tags for a quiz
func (s *QuizService) GetTagsByQuizID(ctx context.Context, quizID string) ([]*model.Tag, error) {
	id, err := strconv.Atoi(quizID)
	if err != nil {
		return nil, err
	}

	dbTags, err := s.Repo.FindTagsByQuizID(ctx, id)
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	for _, dbTag := range dbTags {
		tags = append(tags, db.TagToGraphQL(dbTag))
	}

	return tags, nil
}
