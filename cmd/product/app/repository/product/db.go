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
	if q.deleteProductStmt, err = db.PrepareContext(ctx, deleteProduct); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProduct: %w", err)
	}
	if q.forceDeleteProductStmt, err = db.PrepareContext(ctx, forceDeleteProduct); err != nil {
		return nil, fmt.Errorf("error preparing query ForceDeleteProduct: %w", err)
	}
	if q.getProductByProductIDStmt, err = db.PrepareContext(ctx, getProductByProductID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductByProductID: %w", err)
	}
	if q.getProductsByUserIDStmt, err = db.PrepareContext(ctx, getProductsByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductsByUserID: %w", err)
	}
	if q.insertProductStmt, err = db.PrepareContext(ctx, insertProduct); err != nil {
		return nil, fmt.Errorf("error preparing query InsertProduct: %w", err)
	}
	if q.updateProductStmt, err = db.PrepareContext(ctx, updateProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProduct: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteProductStmt != nil {
		if cerr := q.deleteProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductStmt: %w", cerr)
		}
	}
	if q.forceDeleteProductStmt != nil {
		if cerr := q.forceDeleteProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing forceDeleteProductStmt: %w", cerr)
		}
	}
	if q.getProductByProductIDStmt != nil {
		if cerr := q.getProductByProductIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductByProductIDStmt: %w", cerr)
		}
	}
	if q.getProductsByUserIDStmt != nil {
		if cerr := q.getProductsByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductsByUserIDStmt: %w", cerr)
		}
	}
	if q.insertProductStmt != nil {
		if cerr := q.insertProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertProductStmt: %w", cerr)
		}
	}
	if q.updateProductStmt != nil {
		if cerr := q.updateProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductStmt: %w", cerr)
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
	deleteProductStmt         *sql.Stmt
	forceDeleteProductStmt    *sql.Stmt
	getProductByProductIDStmt *sql.Stmt
	getProductsByUserIDStmt   *sql.Stmt
	insertProductStmt         *sql.Stmt
	updateProductStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                        tx,
		tx:                        tx,
		deleteProductStmt:         q.deleteProductStmt,
		forceDeleteProductStmt:    q.forceDeleteProductStmt,
		getProductByProductIDStmt: q.getProductByProductIDStmt,
		getProductsByUserIDStmt:   q.getProductsByUserIDStmt,
		insertProductStmt:         q.insertProductStmt,
		updateProductStmt:         q.updateProductStmt,
	}
}
