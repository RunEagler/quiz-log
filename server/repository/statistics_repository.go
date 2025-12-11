package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/uptrace/bun"
)

//go:generate mockgen -destination=mocks/mock_statistics_repository.go -package=mocks quiz-log/repository StatisticsRepository

// StatisticsRepository defines the interface for statistics repository operations
type StatisticsRepository interface {
	CountTotalAttempts(ctx context.Context) (int, error)
	CalculateAverageScore(ctx context.Context) (float64, error)
	GetCategoryStats(ctx context.Context) ([]*CategoryStat, error)
}

type statisticsRepository struct {
	DB *bun.DB
}

func NewStatisticsRepository(database *bun.DB) StatisticsRepository {
	return &statisticsRepository{DB: database}
}

// CountTotalAttempts counts the total number of attempts
func (r *statisticsRepository) CountTotalAttempts(ctx context.Context) (int, error) {
	query := psql.Select("COUNT(*)").
		From("attempts")

	var count int
	err := ExecQueryWithReturning[int](ctx, r.DB, query, &count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// CalculateAverageScore calculates the average score
func (r *statisticsRepository) CalculateAverageScore(ctx context.Context) (float64, error) {
	query := psql.Select("AVG(CAST(score AS FLOAT) / CAST(total_questions AS FLOAT) * 100)").
		From("attempts").
		Where(sq.Gt{"total_questions": 0})

	var avgScore sql.NullFloat64

	err := ExecQueryWithReturning[sql.NullFloat64](ctx, r.DB, query, &avgScore)
	if err != nil {
		return 0, err
	}

	if !avgScore.Valid {
		return 0, nil
	}

	return avgScore.Float64, nil
}

// CategoryStat represents category statistics
type CategoryStat struct {
	TagName        string
	CorrectRate    float64
	TotalQuestions int
}

// GetCategoryStats retrieves statistics by category (tag)
func (r *statisticsRepository) GetCategoryStats(ctx context.Context) ([]*CategoryStat, error) {
	queryBuilder := psql.Select(
		"t.name",
		"AVG(CASE WHEN a.is_correct THEN 1.0 ELSE 0.0 END) * 100 as correct_rate",
		"COUNT(*) as total",
	).
		From("tags t").
		Join("question_tags qt ON t.id = qt.tag_id").
		Join("answers a ON qt.question_id = a.question_id").
		GroupBy("t.name").
		OrderBy("t.name")

	return FindAll[CategoryStat](ctx, r.DB, queryBuilder)
}
