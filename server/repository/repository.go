package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/uptrace/bun"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// DBExecutor defines the interface for database operations
//
//go:generate mockgen -destination=mocks/mock_db_executor.go -package=mocks quiz-log/repository DBExecutor
type DBExecutor interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func ExecQuery(ctx context.Context, db *bun.DB, sqlBuilder squirrel.Sqlizer) (sql.Result, error) {
	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	ret, err := db.DB.ExecContext(ctx, sqlStr, args...)
	return ret, err
}

func ExecQueryWithReturning[T any](ctx context.Context, db *bun.DB, sqlBuilder squirrel.Sqlizer, returningValue *T) error {
	sqlStr, args, err := sqlBuilder.ToSql()
	if err != nil {
		return err
	}

	return db.DB.QueryRowContext(ctx, sqlStr, args...).Scan(returningValue)
}

func FindOne[T any](ctx context.Context, db *bun.DB, query squirrel.SelectBuilder) (*T, error) {
	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.DB.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return nil, nil
	}

	var record T
	err = db.ScanRow(ctx, rows, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func FindAll[T any](ctx context.Context, db *bun.DB, query squirrel.SelectBuilder) ([]*T, error) {
	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.DB.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*T
	for rows.Next() {
		var r T
		err := db.ScanRow(ctx, rows, &r)
		if err != nil {
			return nil, err
		}
		records = append(records, &r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
