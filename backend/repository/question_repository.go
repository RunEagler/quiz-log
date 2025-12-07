package repository

import (
	"context"
	"quiz-log/db"
	"strconv"

	"github.com/lib/pq"
	"github.com/uptrace/bun"
)

type QuestionRepository struct {
	DB *bun.DB
}

func NewQuestionRepository(database *bun.DB) *QuestionRepository {
	return &QuestionRepository{DB: database}
}

// Create creates a new question and returns its ID
func (r *QuestionRepository) Create(ctx context.Context, quizID, questionType, content string, options []string, correctAnswer string, explanation *string, difficulty string) (int, error) {
	var questionID int

	query := psql.Insert("questions").
		Columns("quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty").
		Values(quizID, questionType, content, pq.Array(options), correctAnswer, explanation, difficulty).
		Suffix("RETURNING id")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRowContext(ctx, sqlStr, args...).Scan(&questionID)
	if err != nil {
		return 0, err
	}

	return questionID, nil
}

// Update updates an existing question
func (r *QuestionRepository) Update(ctx context.Context, id int, questionType *string, content *string, options []string, correctAnswer *string, explanation *string, difficulty *string) error {
	updateBuilder := psql.Update("questions").Where("id = ?", id)
	hasUpdates := false

	if questionType != nil {
		updateBuilder = updateBuilder.Set("type", *questionType)
		hasUpdates = true
	}

	if content != nil {
		updateBuilder = updateBuilder.Set("content", *content)
		hasUpdates = true
	}

	if options != nil {
		updateBuilder = updateBuilder.Set("options", pq.Array(options))
		hasUpdates = true
	}

	if correctAnswer != nil {
		updateBuilder = updateBuilder.Set("correct_answer", *correctAnswer)
		hasUpdates = true
	}

	if explanation != nil {
		updateBuilder = updateBuilder.Set("explanation", *explanation)
		hasUpdates = true
	}

	if difficulty != nil {
		updateBuilder = updateBuilder.Set("difficulty", *difficulty)
		hasUpdates = true
	}

	if !hasUpdates {
		return nil
	}

	updateBuilder = updateBuilder.Set("updated_at", psql.Select("NOW()"))

	sqlStr, args, err := updateBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// Delete deletes a question by ID
func (r *QuestionRepository) Delete(ctx context.Context, id int) error {
	sqlStr, args, err := psql.Delete("questions").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// FindAll retrieves all questions, optionally filtered by quiz ID
func (r *QuestionRepository) FindAll(ctx context.Context, quizID *int) ([]*db.Question, error) {
	queryBuilder := psql.Select("id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at").
		From("questions")

	if quizID != nil {
		queryBuilder = queryBuilder.Where("quiz_id = ?", *quizID)
	}

	queryBuilder = queryBuilder.OrderBy("created_at ASC")

	sqlStr, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*db.Question
	for rows.Next() {
		var question db.Question
		err := r.DB.ScanRows(ctx, rows, &question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	return questions, nil
}

// FindByID retrieves a question by its ID
func (r *QuestionRepository) FindByID(ctx context.Context, id int) (*db.Question, error) {
	query := psql.Select("id", "quiz_id", "type", "content", "options", "correct_answer", "explanation", "difficulty", "created_at", "updated_at").
		From("questions").
		Where("id = ?", id)

	return FindOne[*db.Question](ctx, r.DB, query)
}

// FindWrongQuestions retrieves questions that were answered incorrectly
func (r *QuestionRepository) FindWrongQuestions(ctx context.Context) ([]*db.Question, error) {
	query := psql.Select("DISTINCT q.id", "q.quiz_id", "q.type", "q.content", "q.options", "q.correct_answer", "q.explanation", "q.difficulty", "q.created_at", "q.updated_at").
		From("questions q").
		Join("answers a ON q.id = a.question_id").
		Where("a.is_correct = false").
		OrderBy("q.created_at DESC")

	return FindAll[*db.Question](ctx, r.DB, query)
}

// FindTagsByQuestionID retrieves all tags for a question
func (r *QuestionRepository) FindTagsByQuestionID(ctx context.Context, questionID int) ([]*db.Tag, error) {
	query := psql.Select("t.id", "t.name").
		From("tags t").
		Join("question_tags qt ON t.id = qt.tag_id").
		Where("qt.question_id = ?", questionID).
		OrderBy("t.name ASC")

	return FindAll[*db.Tag](ctx, r.DB, query)
}

// AssociateTags associates tags with a question
func (r *QuestionRepository) AssociateTags(ctx context.Context, questionID int, tagIDs []string) error {
	if len(tagIDs) == 0 {
		return nil
	}

	insertBuilder := psql.Insert("question_tags").Columns("question_id", "tag_id")
	for _, tagID := range tagIDs {
		id, _ := strconv.Atoi(tagID)
		insertBuilder = insertBuilder.Values(questionID, id)
	}

	sqlStr, args, err := insertBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, sqlStr, args...)
	return err
}

// ClearTags removes all tag associations for a question
func (r *QuestionRepository) ClearTags(ctx context.Context, questionID int) error {
	deleteSql, deleteArgs, err := psql.Delete("question_tags").
		Where("question_id = ?", questionID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, deleteSql, deleteArgs...)
	return err
}
