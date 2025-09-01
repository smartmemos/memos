package model

import "github.com/smartmemos/memos/internal/pkg/db"

type Session struct {
	db.Model

	UserID      int64
	AccessToken string
}

type FindSessionFilter struct {
	db.Query

	ID     db.F[int64]
	UserID db.F[int64]
}

func (Session) TableName() string {
	return TableSession
}

type CreateSessionRequest struct {
	db.Query

	Username string
	Password string
}

type RevokeUserSessionRequest struct {
	UserID    int64
	SessionID string
}
