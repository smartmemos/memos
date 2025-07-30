package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/smartmemos/memos/internal/memos/model"
)

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) (user *model.User, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New("failed to generate password hash")
		return
	}

	user = &model.User{
		Role:         model.RoleHost,
		Status:       model.Normal,
		Username:     req.Username,
		PasswordHash: string(passwordHash),
	}
	if err = s.dao.CreateUser(ctx, user); err != nil {
		return
	}
	return user, nil
}
