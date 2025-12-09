package models

import (
	"github.com/uptrace/bun"
)

type QuizTag struct {
	bun.BaseModel `bun:"table:quiz_tags,alias:qt"`

	QuizID int `bun:"quiz_id,pk"`
	TagID  int `bun:"tag_id,pk"`
}

// Getter methods
func (qt *QuizTag) GetQuizID() int {
	return qt.QuizID
}

func (qt *QuizTag) GetTagID() int {
	return qt.TagID
}
