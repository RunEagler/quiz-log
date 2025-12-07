package services

import (
	"context"
	"quiz-log/graph/model"
	"quiz-log/repository"

	"github.com/uptrace/bun"
)

type StatisticsService struct {
	DB             *bun.DB
	Repo           *repository.StatisticsRepository
	AttemptService *AttemptService
}

func NewStatisticsService(database *bun.DB, attemptService *AttemptService) *StatisticsService {
	return &StatisticsService{
		DB:             database,
		Repo:           repository.NewStatisticsRepository(database),
		AttemptService: attemptService,
	}
}

// GetStatistics retrieves overall statistics
func (s *StatisticsService) GetStatistics(ctx context.Context) (*model.Statistics, error) {
	stats := &model.Statistics{}

	// Total attempts
	totalAttempts, err := s.Repo.CountTotalAttempts(ctx)
	if err == nil {
		stats.TotalAttempts = totalAttempts
	}

	// Average score
	avgScore, err := s.Repo.CalculateAverageScore(ctx)
	if err == nil {
		stats.AverageScore = avgScore
	}

	// Category stats
	categoryStats, err := s.Repo.GetCategoryStats(ctx)
	if err == nil {
		for _, stat := range categoryStats {
			stats.CategoryStats = append(stats.CategoryStats, &model.CategoryStat{
				TagName:        stat.TagName,
				CorrectRate:    stat.CorrectRate,
				TotalQuestions: stat.TotalQuestions,
			})
		}
	}

	// Recent attempts
	stats.RecentAttempts, _ = s.AttemptService.GetAttempts(ctx, nil)
	if len(stats.RecentAttempts) > 10 {
		stats.RecentAttempts = stats.RecentAttempts[:10]
	}

	return stats, nil
}
