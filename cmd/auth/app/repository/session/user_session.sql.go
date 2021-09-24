package repository

import (
	"context"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
)

const getSessionByAccessToken = `-- name: GetSessionByAccessToken :one
SELECT 
    user_id, 
    last_login, 
    access_token, 
    refresh_token, 
    expires, 
    last_ip_address
FROM user_sessions
WHERE 
    access_token = $1
`

func (q *Queries) GetSessionByAccessToken(ctx context.Context, accessToken string) (*domain.GetSessionRow, error) {
	row := q.queryRow(ctx, q.getSessionByAccessTokenStmt, getSessionByAccessToken, accessToken)
	var i domain.GetSessionRow
	err := row.Scan(
		&i.UserID,
		&i.LastLogin,
		&i.AccessToken,
		&i.RefreshToken,
		&i.Expires,
		&i.LastIpAddress,
	)
	return &i, err
}

const insertSession = `-- name: InsertSession :exec
INSERT INTO user_sessions (user_id, last_login, access_token, refresh_token, expires, last_ip_address)
VALUES (
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6
)
`

func (q *Queries) InsertSession(ctx context.Context, arg *domain.InsertSessionParams) error {
	_, err := q.exec(ctx, q.insertSessionStmt, insertSession,
		arg.UserID,
		arg.LastLogin,
		arg.AccessToken,
		arg.RefreshToken,
		arg.Expires,
		arg.LastIpAddress,
	)
	return err
}
