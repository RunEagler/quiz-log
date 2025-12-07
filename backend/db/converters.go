package db

import (
	"quiz-log/graph/model"
	"strconv"
)

// QuizToGraphQL converts a db.Quiz to a GraphQL model.Quiz
func QuizToGraphQL(q *Quiz) *model.Quiz {
	return &model.Quiz{
		ID:          strconv.Itoa(q.ID),
		Title:       q.Title,
		Description: q.Description,
		CreatedAt:   q.CreatedAt,
		UpdatedAt:   q.UpdatedAt,
	}
}

// TagToGraphQL converts a db.Tag to a GraphQL model.Tag
func TagToGraphQL(t *Tag) *model.Tag {
	return &model.Tag{
		ID:   strconv.Itoa(t.ID),
		Name: t.Name,
	}
}

// QuestionToGraphQL converts a db.Question to a GraphQL model.Question
func QuestionToGraphQL(q *Question) *model.Question {
	return &model.Question{
		ID:            strconv.Itoa(q.ID),
		QuizID:        strconv.Itoa(q.QuizID),
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
func AttemptToGraphQL(a *Attempt) *model.Attempt {
	return &model.Attempt{
		ID:             strconv.Itoa(a.ID),
		QuizID:         strconv.Itoa(a.QuizID),
		StartedAt:      a.StartedAt,
		CompletedAt:    a.CompletedAt,
		Score:          a.Score,
		TotalQuestions: a.TotalQuestions,
	}
}

// AnswerToGraphQL converts a db.Answer to a GraphQL model.Answer
func AnswerToGraphQL(a *Answer) *model.Answer {
	return &model.Answer{
		ID:         strconv.Itoa(a.ID),
		AttemptID:  strconv.Itoa(a.AttemptID),
		QuestionID: strconv.Itoa(a.QuestionID),
		UserAnswer: a.UserAnswer,
		IsCorrect:  a.IsCorrect,
	}
}
