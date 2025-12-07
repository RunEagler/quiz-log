package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/uptrace/bun"
)

type StatisticsRepository struct {
	DB *bun.DB
}

func NewStatisticsRepository(database *bun.DB) *StatisticsRepository {
	return &StatisticsRepository{DB: database}
}

// CountTotalAttempts counts the total number of attempts
func (r *StatisticsRepository) CountTotalAttempts(ctx context.Context) (int, error) {
	countSql, countArgs, err := psql.Select("COUNT(*)").
		From("attempts").
		ToSql()
	if err != nil {
		return 0, err
	}

	var count int
	err = r.DB.QueryRowContext(ctx, countSql, countArgs...).Scan(&count)
	return count, err
}

// CalculateAverageScore calculates the average score
func (r *StatisticsRepository) CalculateAverageScore(ctx context.Context) (float64, error) {
	avgSql, avgArgs, err := psql.Select("AVG(CAST(score AS FLOAT) / CAST(total_questions AS FLOAT) * 100)").
		From("attempts").
		Where(sq.Gt{"total_questions": 0}).
		ToSql()
	if err != nil {
		return 0, err
	}

	var avgScore sql.NullFloat64
	err = r.DB.QueryRowContext(ctx, avgSql, avgArgs...).Scan(&avgScore)
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
func (r *StatisticsRepository) GetCategoryStats(ctx context.Context) ([]*CategoryStat, error) {
	categorySql, categoryArgs, err := psql.Select(
		"t.name",
		"AVG(CASE WHEN a.is_correct THEN 1.0 ELSE 0.0 END) * 100 as correct_rate",
		"COUNT(*) as total",
	).
		From("tags t").
		Join("question_tags qt ON t.id = qt.tag_id").
		Join("answers a ON qt.question_id = a.question_id").
		GroupBy("t.name").
		OrderBy("t.name").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, categorySql, categoryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*CategoryStat
	for rows.Next() {
		var stat CategoryStat
		err := rows.Scan(&stat.TagName, &stat.CorrectRate, &stat.TotalQuestions)
		if err != nil {
			return nil, err
		}
		stats = append(stats, &stat)
	}

	return stats, nil
}
