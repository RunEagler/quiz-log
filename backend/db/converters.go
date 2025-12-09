package db

import (
	"strconv"
	
	"quiz-log/graph/model"
	"quiz-log/models"
)

// QuizToGraphQL converts a db.Quiz to a GraphQL model.Quiz
func QuizToGraphQL(q *models.Quiz) *model.Quiz {
	return &model.Quiz{
		ID:          strconv.Itoa(q.ID),
		Title:       q.Title,
		Description: q.Description,
		CreatedAt:   q.CreatedAt,
		UpdatedAt:   q.UpdatedAt,
	}
}

// TagToGraphQL converts a db.Tag to a GraphQL model.Tag
func TagToGraphQL(t *models.Tag) *model.Tag {
	return &model.Tag{
		ID:   strconv.Itoa(t.ID),
		Name: t.Name,
	}
}

// QuestionToGraphQL converts a db.Question to a GraphQL model.Question
func QuestionToGraphQL(q *models.Question) *model.Question {
	var quizID string
	if q.QuizID != nil {
		quizID = strconv.Itoa(*q.QuizID)
	}

	return &model.Question{
		ID:            strconv.Itoa(q.ID),
		QuizID:        quizID,
		Type:          model.QuestionType(q.Type),
		Content:       q.Content,
		Options:       q.Options,
		CorrectAnswer: q.CorrectAnswer,
		Explanation:   q.Explanation,
		Difficulty:    model.Difficulty(q.Difficulty),
		CreatedAt:     q.CreatedAt,
		UpdatedAt:     q.UpdatedAt,
	}
}

// AttemptToGraphQL converts a db.Attempt to a GraphQL model.Attempt
func AttemptToGraphQL(a *models.Attempt) *model.Attempt {
	var quizID string
	if a.QuizID != nil {
		quizID = strconv.Itoa(*a.QuizID)
	}

	return &model.Attempt{
		ID:             strconv.Itoa(a.ID),
		QuizID:         quizID,
		StartedAt:      a.StartedAt,
		CompletedAt:    a.CompletedAt,
		Score:          a.Score,
		TotalQuestions: a.TotalQuestions,
	}
}

// AnswerToGraphQL converts a db.Answer to a GraphQL model.Answer
func AnswerToGraphQL(a *models.Answer) *model.Answer {
	var attemptID, questionID string
	if a.AttemptID != nil {
		attemptID = strconv.Itoa(*a.AttemptID)
	}
	if a.QuestionID != nil {
		questionID = strconv.Itoa(*a.QuestionID)
	}

	return &model.Answer{
		ID:         strconv.Itoa(a.ID),
		AttemptID:  attemptID,
		QuestionID: questionID,
		UserAnswer: a.UserAnswer,
		IsCorrect:  a.IsCorrect,
	}
}
