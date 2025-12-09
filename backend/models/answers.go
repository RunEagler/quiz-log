package models

import (
	"github.com/uptrace/bun"
)

type Answer struct {
	bun.BaseModel `bun:"table:answers,alias:a"`

	ID         int    `bun:"id,pk,autoincrement"`
	AttemptID  *int   `bun:"attempt_id"`
	QuestionID *int   `bun:"question_id"`
	UserAnswer string `bun:"user_answer,notnull"`
	IsCorrect  bool   `bun:"is_correct,notnull"`
}

// Getter methods
func (a *Answer) GetID() int {
	return a.ID
}

func (a *Answer) GetAttemptID() *int {
	return a.AttemptID
}

func (a *Answer) GetQuestionID() *int {
	return a.QuestionID
}

func (a *Answer) GetUserAnswer() string {
	return a.UserAnswer
}

func (a *Answer) GetIsCorrect() bool {
	return a.IsCorrect
}
