package model

import (
	"strings"
	"time"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type Profile struct {
	ID          int64 `gorm:"primary_key" json:"id,string"` // 主键ID
	UserId      int64
	Token       string
	Description string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}

func (Profile) TableName() string {
	return TableProfile
}

type FindProfileFilter struct {
	db.BaseFilter

	ID     int64
	UserId int64
	Token  string
}

func (f FindProfileFilter) GetQuery() (query string, args []any) {
	var where []string
	if f.UserId > 0 {
		where = append(where, "user_id=?")
		args = append(args, f.UserId)
	}
	if f.Token != "" {
		where = append(where, "token=?")
		args = append(args, f.Token)
	}
	query = strings.Join(where, " and ")
	return
}

type GetProfileRequest struct {
}
