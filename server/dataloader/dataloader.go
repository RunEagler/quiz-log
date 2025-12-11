package dataloader

import (
	"context"
	"net/http"
	"quiz-log/graph/model"
	"quiz-log/repository"
	"time"

	"github.com/graph-gophers/dataloader/v7"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders holds all dataloaders
type Loaders struct {
	QuestionsByQuizID *dataloader.Loader[int, []*model.Question]
	TagsByQuizID      *dataloader.Loader[int, []*model.Tag]
}

// NewLoaders creates new dataloaders
func NewLoaders(quizRepo repository.QuizRepository) *Loaders {
	return &Loaders{
		QuestionsByQuizID: dataloader.NewBatchedLoader(
			batchQuestionsByQuizID(quizRepo),
			dataloader.WithWait[int, []*model.Question](time.Millisecond),
		),
		TagsByQuizID: dataloader.NewBatchedLoader(
			batchTagsByQuizID(quizRepo),
			dataloader.WithWait[int, []*model.Tag](time.Millisecond),
		),
	}
}

// Middleware injects dataloaders into the context
func Middleware(loaders *Loaders) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, loaders)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// For returns the dataloaders from the context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
