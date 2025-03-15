package model

import (
	"time"
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

type CreateAccessTokenRequest struct {
	UserId      int64
	Token       string
	Description string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}
