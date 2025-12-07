package db

import (
	"time"

	"github.com/uptrace/bun"
)

type Quiz struct {
	bun.BaseModel `bun:"table:quizzes,alias:q"`

	ID          int       `bun:"id,pk,autoincrement"`
	Title       string    `bun:"title,notnull"`
	Description *string   `bun:"description"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:now()"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:now()"`
}

type Tag struct {
	bun.BaseModel `bun:"table:tags,alias:t"`

	ID   int    `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull,unique"`
}

type QuizTag struct {
	bun.BaseModel `bun:"table:quiz_tags,alias:qt"`

	QuizID int `bun:"quiz_id,pk"`
	TagID  int `bun:"tag_id,pk"`
}

type Question struct {
	bun.BaseModel `bun:"table:questions,alias:q"`

	ID            int       `bun:"id,pk,autoincrement"`
	QuizID        int       `bun:"quiz_id,notnull"`
	Type          string    `bun:"type,notnull"`
	Content       string    `bun:"content,notnull"`
	Options       []string  `bun:"options,array"`
	CorrectAnswer string    `bun:"correct_answer,notnull"`
	Explanation   *string   `bun:"explanation"`
	Difficulty    string    `bun:"difficulty,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:now()"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:now()"`
}

type QuestionTag struct {
	bun.BaseModel `bun:"table:question_tags,alias:qt"`

	QuestionID int `bun:"question_id,pk"`
	TagID      int `bun:"tag_id,pk"`
}

type Attempt struct {
	bun.BaseModel `bun:"table:attempts,alias:a"`

	ID             int        `bun:"id,pk,autoincrement"`
	QuizID         int        `bun:"quiz_id,notnull"`
	StartedAt      time.Time  `bun:"started_at,notnull,default:now()"`
	CompletedAt    *time.Time `bun:"completed_at"`
	Score          int        `bun:"score,notnull,default:0"`
	TotalQuestions int        `bun:"total_questions,notnull"`
}

type Answer struct {
	bun.BaseModel `bun:"table:answers,alias:a"`

	ID         int    `bun:"id,pk,autoincrement"`
	AttemptID  int    `bun:"attempt_id,notnull"`
	QuestionID int    `bun:"question_id,notnull"`
	UserAnswer string `bun:"user_answer,notnull"`
	IsCorrect  bool   `bun:"is_correct,notnull"`
}
