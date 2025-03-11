package model

import (
	"time"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type AccessToken struct {
	db.Model

	UserId      int64
	Token       string
	Description string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}

type SignInRequest struct {
	Username    string
	Password    string
	NeverExpire bool
}
