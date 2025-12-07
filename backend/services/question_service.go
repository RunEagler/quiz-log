package services

import (
	"context"
	"encoding/json"
	"quiz-log/db"
	"quiz-log/graph/model"
	"quiz-log/repository"
	"strconv"

	"github.com/uptrace/bun"
)

type QuestionService struct {
	DB   *bun.DB
	Repo *repository.QuestionRepository
}

func NewQuestionService(database *bun.DB) *QuestionService {
	return &QuestionService{
		DB:   database,
		Repo: repository.NewQuestionRepository(database),
	}
}

// CreateQuestion creates a new question
func (s *QuestionService) CreateQuestion(ctx context.Context, input model.CreateQuestionInput) (*model.Question, error) {
	questionID, err := s.Repo.Create(ctx, input.QuizID, string(input.Type), input.Content, input.Options, input.CorrectAnswer, input.Explanation, string(input.Difficulty))
	if err != nil {
		return nil, err
	}

	// Associate tags
	if len(input.TagIDs) > 0 {
		err = s.Repo.AssociateTags(ctx, questionID, input.TagIDs)
		if err != nil {
			return nil, err
		}
	}

	return s.GetQuestionByID(ctx, strconv.Itoa(questionID))
}

// UpdateQuestion updates an existing question
func (s *QuestionService) UpdateQuestion(ctx context.Context, id string, input model.UpdateQuestionInput) (*model.Question, error) {
	questionID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var qType, content, correctAnswer, difficulty *string
	if input.Type != nil {
		t := string(*input.Type)
		qType = &t
	}
	if input.Difficulty != nil {
		d := string(*input.Difficulty)
		difficulty = &d
	}
	content = input.Content
	correctAnswer = input.CorrectAnswer

	err = s.Repo.Update(ctx, questionID, qType, content, input.Options, correctAnswer, input.Explanation, difficulty)
	if err != nil {
		return nil, err
	}

	// Update tags
	if input.TagIDs != nil {
		err = s.Repo.ClearTags(ctx, questionID)
		if err != nil {
			return nil, err
		}

		if len(input.TagIDs) > 0 {
			err = s.Repo.AssociateTags(ctx, questionID, input.TagIDs)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.GetQuestionByID(ctx, id)
}

// DeleteQuestion deletes a question by ID
func (s *QuestionService) DeleteQuestion(ctx context.Context, id string) (bool, error) {
	questionID, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}

	err = s.Repo.Delete(ctx, questionID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetAllQuestions retrieves all questions, optionally filtered by quiz ID
func (s *QuestionService) GetAllQuestions(ctx context.Context, quizID *string) ([]*model.Question, error) {
	var qid *int
	if quizID != nil {
		id, _ := strconv.Atoi(*quizID)
		qid = &id
	}

	dbQuestions, err := s.Repo.FindAll(ctx, qid)
	if err != nil {
		return nil, err
	}

	var questions []*model.Question
	for _, dbQuestion := range dbQuestions {
		questions = append(questions, db.QuestionToGraphQL(dbQuestion))
	}

	return questions, nil
}

// GetQuestionByID retrieves a question by its ID
func (s *QuestionService) GetQuestionByID(ctx context.Context, id string) (*model.Question, error) {
	questionID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	dbQuestion, err := s.Repo.FindByID(ctx, questionID)
	if err != nil {
		return nil, err
	}

	if dbQuestion == nil {
		return nil, nil
	}

	return db.QuestionToGraphQL(dbQuestion), nil
}

// GetWrongQuestions retrieves questions that were answered incorrectly
func (s *QuestionService) GetWrongQuestions(ctx context.Context) ([]*model.Question, error) {
	dbQuestions, err := s.Repo.FindWrongQuestions(ctx)
	if err != nil {
		return nil, err
	}

	var questions []*model.Question
	for _, dbQuestion := range dbQuestions {
		questions = append(questions, db.QuestionToGraphQL(dbQuestion))
	}

	return questions, nil
}

// ImportQuestions imports questions from JSON data
func (s *QuestionService) ImportQuestions(ctx context.Context, data string) ([]*model.Question, error) {
	var questions []struct {
		QuizID        string   `json:"quizId"`
		Type          string   `json:"type"`
		Content       string   `json:"content"`
		Options       []string `json:"options"`
		CorrectAnswer string   `json:"correctAnswer"`
		Explanation   *string  `json:"explanation"`
		Difficulty    string   `json:"difficulty"`
		TagIDs        []string `json:"tagIds"`
	}

	err := json.Unmarshal([]byte(data), &questions)
	if err != nil {
		return nil, err
	}

	var result []*model.Question
	for _, q := range questions {
		qType := model.QuestionType(q.Type)
		qDifficulty := model.Difficulty(q.Difficulty)

		input := model.CreateQuestionInput{
			QuizID:        q.QuizID,
			Type:          qType,
			Content:       q.Content,
			Options:       q.Options,
			CorrectAnswer: q.CorrectAnswer,
			Explanation:   q.Explanation,
			Difficulty:    qDifficulty,
			TagIDs:        q.TagIDs,
		}

		question, err := s.CreateQuestion(ctx, input)
		if err != nil {
			return nil, err
		}
		result = append(result, question)
	}

	return result, nil
}

// ExportQuestions exports questions to JSON
func (s *QuestionService) ExportQuestions(ctx context.Context, quizID *string) (string, error) {
	questions, err := s.GetAllQuestions(ctx, quizID)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(questions)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// GetTagsByQuestionID retrieves all tags for a question
func (s *QuestionService) GetTagsByQuestionID(ctx context.Context, questionID string) ([]*model.Tag, error) {
	id, err := strconv.Atoi(questionID)
	if err != nil {
		return nil, err
	}

	dbTags, err := s.Repo.FindTagsByQuestionID(ctx, id)
	if err != nil {
		return nil, err
	}

	var tags []*model.Tag
	for _, dbTag := range dbTags {
		tags = append(tags, db.TagToGraphQL(dbTag))
	}

	return tags, nil
}
