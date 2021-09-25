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
	if q.deleteArticleStmt, err = db.PrepareContext(ctx, deleteArticle); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteArticle: %w", err)
	}
	if q.forceDeleteArticleStmt, err = db.PrepareContext(ctx, forceDeleteArticle); err != nil {
		return nil, fmt.Errorf("error preparing query ForceDeleteArticle: %w", err)
	}
	if q.getArticleByArticleIDStmt, err = db.PrepareContext(ctx, getArticleByArticleID); err != nil {
		return nil, fmt.Errorf("error preparing query GetArticleByArticleID: %w", err)
	}
	if q.getArticlesByUserIDStmt, err = db.PrepareContext(ctx, getArticlesByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query GetArticlesByUserID: %w", err)
	}
	if q.insertArticleStmt, err = db.PrepareContext(ctx, insertArticle); err != nil {
		return nil, fmt.Errorf("error preparing query InsertArticle: %w", err)
	}
	if q.updateArticleStmt, err = db.PrepareContext(ctx, updateArticle); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateArticle: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteArticleStmt != nil {
		if cerr := q.deleteArticleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteArticleStmt: %w", cerr)
		}
	}
	if q.forceDeleteArticleStmt != nil {
		if cerr := q.forceDeleteArticleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing forceDeleteArticleStmt: %w", cerr)
		}
	}
	if q.getArticleByArticleIDStmt != nil {
		if cerr := q.getArticleByArticleIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getArticleByArticleIDStmt: %w", cerr)
		}
	}
	if q.getArticlesByUserIDStmt != nil {
		if cerr := q.getArticlesByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getArticlesByUserIDStmt: %w", cerr)
		}
	}
	if q.insertArticleStmt != nil {
		if cerr := q.insertArticleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertArticleStmt: %w", cerr)
		}
	}
	if q.updateArticleStmt != nil {
		if cerr := q.updateArticleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateArticleStmt: %w", cerr)
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
	db                        DBTX
	tx                        *sql.Tx
	deleteArticleStmt         *sql.Stmt
	forceDeleteArticleStmt    *sql.Stmt
	getArticleByArticleIDStmt *sql.Stmt
	getArticlesByUserIDStmt   *sql.Stmt
	insertArticleStmt         *sql.Stmt
	updateArticleStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                        tx,
		tx:                        tx,
		deleteArticleStmt:         q.deleteArticleStmt,
		forceDeleteArticleStmt:    q.forceDeleteArticleStmt,
		getArticleByArticleIDStmt: q.getArticleByArticleIDStmt,
		getArticlesByUserIDStmt:   q.getArticlesByUserIDStmt,
		insertArticleStmt:         q.insertArticleStmt,
		updateArticleStmt:         q.updateArticleStmt,
	}
}
