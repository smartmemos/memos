package model

import (
	"time"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type AccessToken struct {
	ID          int64 `gorm:"primary_key" json:"id,string"` // 主键ID
	UserId      int64
	Token       string
	Description string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}

func (AccessToken) TableName() string {
	return TableAccessToken
}

type FindAccessTokenFilter struct {
	db.BaseFilter

	UserId db.F[int64]
	Token  db.F[string]
}

type CreateAccessTokenRequest struct {
	UserId      int64
	Token       string
	Description string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}

type ListAccessTokensRequest struct {
}

type DeleteAccessTokenRequest struct {
	UserID      int64
	AccessToken string
}
