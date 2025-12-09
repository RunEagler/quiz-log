package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Question struct {
	bun.BaseModel `bun:"table:questions,alias:q"`

	ID            int       `bun:"id,pk,autoincrement"`
	QuizID        *int      `bun:"quiz_id"`
	Type          string    `bun:"type,notnull"`
	Content       string    `bun:"content,notnull"`
	Options       []string  `bun:"options,array"`
	CorrectAnswer string    `bun:"correct_answer,notnull"`
	Explanation   *string   `bun:"explanation"`
	Difficulty    string    `bun:"difficulty,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,nullzero,default:now()"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,nullzero,default:now()"`
}

// Getter methods
func (q *Question) GetID() int {
	return q.ID
}

func (q *Question) GetQuizID() *int {
	return q.QuizID
}

func (q *Question) GetType() string {
	return q.Type
}

func (q *Question) GetContent() string {
	return q.Content
}

func (q *Question) GetOptions() []string {
	return q.Options
}

func (q *Question) GetCorrectAnswer() string {
	return q.CorrectAnswer
}

func (q *Question) GetExplanation() *string {
	return q.Explanation
}

func (q *Question) GetDifficulty() string {
	return q.Difficulty
}

func (q *Question) GetCreatedAt() time.Time {
	return q.CreatedAt
}

func (q *Question) GetUpdatedAt() time.Time {
	return q.UpdatedAt
}
