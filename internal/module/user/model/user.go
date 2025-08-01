package model

import (
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

	ID       db.F[int64]
	Username db.F[string]
	Role     db.F[Role]
}

type CreateUserRequest struct {
	Username string
	Role     Role
	Email    string
	Nickname string
	Password string
}

type UpdateUserRequest struct {
	UpdateMask  []string
	ID          int64
	Username    string
	Role        Role
	Email       string
	Nickname    string
	Password    string
	AvatarURL   string
	Description string
	Status      RowStatus
}
