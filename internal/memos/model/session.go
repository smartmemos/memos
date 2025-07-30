package model

import "github.com/smartmemos/memos/internal/pkg/db"

type Session struct {
	db.Model

	UserID      int64
	AccessToken string
}

type FindSessionFilter struct {
	db.BaseFilter

	ID     db.F[int64]
	UserID db.F[int64]
}

func (Session) TableName() string {
	return TableSession
}

type CreateSessionRequest struct {
	Username string
	Password string
}
