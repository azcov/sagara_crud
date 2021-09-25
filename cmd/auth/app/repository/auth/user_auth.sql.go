package repository

import (
	"context"
	"database/sql"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
	errorCommon "github.com/azcov/sagara_crud/errors"
	"go.uber.org/zap"
)

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    user_id, 
    role_id, 
    name, 
    PGP_SYM_DECRYPT(email, $1::text)::text as email,
    password
FROM user_authentications
WHERE LOWER(PGP_SYM_DECRYPT(email, $1)) = LOWER($2)
`

func (q *Queries) GetUserByEmail(ctx context.Context, arg *domain.GetUserByEmailParams) (*domain.UserRow, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, arg.EmailSecret, arg.Email)
	var i domain.UserRow
	err := row.Scan(
		&i.UserID,
		&i.RoleID,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	if err == sql.ErrNoRows {
		err = errorCommon.ErrEmailNotFound
		return nil, err
	}
	return &i, err
}

const getUserByUserId = `-- name: GetUserByUserId :one
SELECT
    user_id, 
    role_id, 
    name, 
    PGP_SYM_DECRYPT(email, $1::text)::text as email,
    password
FROM user_authentications
WHERE 
    user_id = $2
`

func (q *Queries) GetUserByUserId(ctx context.Context, arg *domain.GetUserByUserIdParams) (*domain.UserRow, error) {
	row := q.queryRow(ctx, q.getUserByUserIdStmt, getUserByUserId, arg.Email, arg.EmailSecret, arg.UserID)
	var i domain.UserRow
	err := row.Scan(
		&i.UserID,
		&i.RoleID,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	return &i, err
}

const insertAuthData = `-- name: InsertAuthData :exec
INSERT INTO user_authentications (role_id, name, email, password, created_at)
VALUES (
    $1, 
    $2, 
    PGP_SYM_ENCRYPT(lower($3::text), $4::text)::bytea,
    $5,
    $6
)
`

func (q *Queries) InsertAuthData(ctx context.Context, arg *domain.InsertAuthDataParams) error {
	zap.S().Infof("%+v", arg)
	_, err := q.exec(ctx, q.insertAuthDataStmt, insertAuthData,
		arg.RoleID,
		arg.Name,
		arg.Email,
		arg.EmailSecret,
		arg.Password,
		arg.CreatedAt,
	)
	return err
}
