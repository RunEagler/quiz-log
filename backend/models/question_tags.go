package models

import (
	"github.com/uptrace/bun"
)

type QuestionTag struct {
	bun.BaseModel `bun:"table:question_tags,alias:qt"`

	QuestionID int `bun:"question_id,pk"`
	TagID      int `bun:"tag_id,pk"`
}

// Getter methods
func (qt *QuestionTag) GetQuestionID() int {
	return qt.QuestionID
}

func (qt *QuestionTag) GetTagID() int {
	return qt.TagID
}
