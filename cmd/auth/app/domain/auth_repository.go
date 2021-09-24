package domain

import (
	"github.com/google/uuid"
)

type RoleID int32

const (
	_ RoleID = iota
	RoleIDAdmin
	RoleIDMember
)

type GetUserByEmailParams struct {
	Email       []byte `json:"email"`
	EmailSecret string `json:"email_secret"`
}

type UserRow struct {
	UserID   uuid.UUID `json:"user_id"`
	RoleID   int32     `json:"role_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type GetUserByUserIdParams struct {
	Email       string    `json:"email"`
	EmailSecret string    `json:"email_secret"`
	UserID      uuid.UUID `json:"user_id"`
}

type InsertAuthDataParams struct {
	RoleID      int32  `json:"role_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	EmailSecret string `json:"email_secret"`
	Password    string `json:"password"`
	CreatedAt   int64  `json:"created_at"`
}

type GetSessionRow struct {
	UserID        uuid.UUID `json:"user_id"`
	LastLogin     int64     `json:"last_login"`
	AccessToken   string    `json:"access_token"`
	RefreshToken  string    `json:"refresh_token"`
	Expires       int64     `json:"expires"`
	LastIpAddress string    `json:"last_ip_address"`
}

type InsertSessionParams struct {
	UserID        uuid.UUID `json:"user_id"`
	LastLogin     int64     `json:"last_login"`
	AccessToken   string    `json:"access_token"`
	RefreshToken  string    `json:"refresh_token"`
	Expires       int64     `json:"expires"`
	LastIpAddress string    `json:"last_ip_address"`
}
