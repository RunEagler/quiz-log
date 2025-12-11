package dataloader

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"

	"quiz-log/db"
	"quiz-log/graph/model"
	"quiz-log/repository"
)

// batchQuestionsByQuizID batches questions by quiz IDs
func batchQuestionsByQuizID(quizRepo repository.QuizRepository) dataloader.BatchFunc[int, []*model.Question] {
	return func(ctx context.Context, quizIDs []int) []*dataloader.Result[[]*model.Question] {
		questionsMap, err := quizRepo.FindQuestionsByQuizIDs(ctx, quizIDs)
		if err != nil {
			// Return error for all keys
			results := make([]*dataloader.Result[[]*model.Question], len(quizIDs))
			for i := range quizIDs {
				results[i] = &dataloader.Result[[]*model.Question]{Error: err}
			}
			return results
		}

		// Create results in the same order as requested keys
		results := make([]*dataloader.Result[[]*model.Question], len(quizIDs))
		for i, quizID := range quizIDs {
			dbQuestions := questionsMap[quizID]
			questions := make([]*model.Question, len(dbQuestions))
			for j, dbQ := range dbQuestions {
				questions[j] = db.QuestionToGraphQL(dbQ)
			}
			results[i] = &dataloader.Result[[]*model.Question]{Data: questions}
		}
		return results
	}
}

// batchTagsByQuizID batches tags by quiz IDs
func batchTagsByQuizID(quizRepo repository.QuizRepository) dataloader.BatchFunc[int, []*model.Tag] {
	return func(ctx context.Context, quizIDs []int) []*dataloader.Result[[]*model.Tag] {
		tagsMap, err := quizRepo.FindTagsByQuizIDs(ctx, quizIDs)
		if err != nil {
			// Return error for all keys
			results := make([]*dataloader.Result[[]*model.Tag], len(quizIDs))
			for i := range quizIDs {
				results[i] = &dataloader.Result[[]*model.Tag]{Error: err}
			}
			return results
		}

		// Create results in the same order as requested keys
		results := make([]*dataloader.Result[[]*model.Tag], len(quizIDs))
		for i, quizID := range quizIDs {
			dbTags := tagsMap[quizID]
			tags := make([]*model.Tag, len(dbTags))
			for j, dbT := range dbTags {
				tags[j] = db.TagToGraphQL(dbT)
			}
			results[i] = &dataloader.Result[[]*model.Tag]{Data: tags}
		}
		return results
	}
}
