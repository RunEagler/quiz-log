package models

import (
	"github.com/uptrace/bun"
)

type Tag struct {
	bun.BaseModel `bun:"table:tags,alias:t"`

	ID   int    `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull,unique"`
}

// Getter methods
func (t *Tag) GetID() int {
	return t.ID
}

func (t *Tag) GetName() string {
	return t.Name
}
