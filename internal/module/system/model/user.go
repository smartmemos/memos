package model

import (
	"strings"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type User struct {
	db.Model

	// Domain specific fields
	Username     string
	Role         Role
	Email        string
	Nickname     string
	PasswordHash string
	AvatarURL    string
	Description  string
	Status       RowStatus
}

func (User) TableName() string {
	return TableUser
}

type FindUserFilter struct {
	db.BaseFilter

	ID       int64
	Username string
}

func (f FindUserFilter) GetQuery() (query string, args []any) {
	var where []string
	if f.ID > 0 {
		where = append(where, "id=?")
		args = append(args, f.ID)
	}
	if f.Username != "" {
		where = append(where, "username=?")
		args = append(args, f.Username)
	}
	query = strings.Join(where, " and ")
	return
}
