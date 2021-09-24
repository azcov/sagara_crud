package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.getSessionByAccessTokenStmt, err = db.PrepareContext(ctx, getSessionByAccessToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetSessionByAccessToken: %w", err)
	}
	if q.insertSessionStmt, err = db.PrepareContext(ctx, insertSession); err != nil {
		return nil, fmt.Errorf("error preparing query InsertSession: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getSessionByAccessTokenStmt != nil {
		if cerr := q.getSessionByAccessTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSessionByAccessTokenStmt: %w", cerr)
		}
	}
	if q.insertSessionStmt != nil {
		if cerr := q.insertSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertSessionStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                          DBTX
	tx                          *sql.Tx
	getSessionByAccessTokenStmt *sql.Stmt
	insertSessionStmt           *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                          tx,
		tx:                          tx,
		getSessionByAccessTokenStmt: q.getSessionByAccessTokenStmt,
		insertSessionStmt:           q.insertSessionStmt,
	}
}
