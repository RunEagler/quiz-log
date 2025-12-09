package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Quiz struct {
	bun.BaseModel `bun:"table:quizzes,alias:q"`

	ID          int       `bun:"id,pk,autoincrement"`
	Title       string    `bun:"title,notnull"`
	Description *string   `bun:"description"`
	CreatedAt   time.Time `bun:"created_at,notnull,nullzero,default:now()"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,nullzero,default:now()"`
}

// Getter methods
func (q *Quiz) GetID() int {
	return q.ID
}

func (q *Quiz) GetTitle() string {
	return q.Title
}

func (q *Quiz) GetDescription() *string {
	return q.Description
}

func (q *Quiz) GetCreatedAt() time.Time {
	return q.CreatedAt
}

func (q *Quiz) GetUpdatedAt() time.Time {
	return q.UpdatedAt
}
