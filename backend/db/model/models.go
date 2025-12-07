package model

import "time"

// Quiz represents a quiz in the database
type Quiz struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description *string   `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Tag represents a tag in the database
type Tag struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// QuizTag represents the quiz_tags junction table
type QuizTag struct {
	QuizID int `db:"quiz_id"`
	TagID  int `db:"tag_id"`
}

// Question represents a question in the database
type Question struct {
	ID            int       `db:"id"`
	QuizID        int       `db:"quiz_id"`
	Type          string    `db:"type"`
	Content       string    `db:"content"`
	Options       []string  `db:"options"`
	CorrectAnswer string    `db:"correct_answer"`
	Explanation   *string   `db:"explanation"`
	Difficulty    string    `db:"difficulty"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// QuestionTag represents the question_tags junction table
type QuestionTag struct {
	QuestionID int `db:"question_id"`
	TagID      int `db:"tag_id"`
}

// Attempt represents a quiz attempt in the database
type Attempt struct {
	ID             int        `db:"id"`
	QuizID         int        `db:"quiz_id"`
	StartedAt      time.Time  `db:"started_at"`
	CompletedAt    *time.Time `db:"completed_at"`
	Score          int        `db:"score"`
	TotalQuestions int        `db:"total_questions"`
}

// Answer represents an answer to a question in the database
type Answer struct {
	ID         int    `db:"id"`
	AttemptID  int    `db:"attempt_id"`
	QuestionID int    `db:"question_id"`
	UserAnswer string `db:"user_answer"`
	IsCorrect  bool   `db:"is_correct"`
}
