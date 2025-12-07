package resolvers

import (
	"quiz-log/services"

	sq "github.com/Masterminds/squirrel"
	"github.com/uptrace/bun"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	DB                *bun.DB
	QuizService       *services.QuizService
	QuestionService   *services.QuestionService
	TagService        *services.TagService
	AttemptService    *services.AttemptService
	StatisticsService *services.StatisticsService
}

// PostgreSQL query builder
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
