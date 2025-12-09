package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Attempt struct {
	bun.BaseModel `bun:"table:attempts,alias:a"`

	ID             int        `bun:"id,pk,autoincrement"`
	QuizID         *int       `bun:"quiz_id"`
	StartedAt      time.Time  `bun:"started_at,notnull,nullzero,default:now()"`
	CompletedAt    *time.Time `bun:"completed_at"`
	Score          int        `bun:"score,notnull,default:0"`
	TotalQuestions int        `bun:"total_questions,notnull"`
}

// Getter methods
func (a *Attempt) GetID() int {
	return a.ID
}

func (a *Attempt) GetQuizID() *int {
	return a.QuizID
}

func (a *Attempt) GetStartedAt() time.Time {
	return a.StartedAt
}

func (a *Attempt) GetCompletedAt() *time.Time {
	return a.CompletedAt
}

func (a *Attempt) GetScore() int {
	return a.Score
}

func (a *Attempt) GetTotalQuestions() int {
	return a.TotalQuestions
}
