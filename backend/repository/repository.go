package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/uptrace/bun"
)

func FindOne[T any](ctx context.Context, db *bun.DB, query squirrel.SelectBuilder) (*T, error) {
	var record T

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = db.NewRaw(sqlStr, args...).Scan(ctx, &record)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
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

	rows, err := db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*T
	for rows.Next() {
		var r T
		err := db.ScanRows(ctx, rows, &r)
		if err != nil {
			return nil, err
		}
		records = append(records, &r)
	}
	return records, nil
}
