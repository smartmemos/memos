package service

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) (user *model.User, err error) {
	user = &model.User{
		Username: req.Username,
		// Password: req.Password,
	}

	err = s.dao.CreateUser(ctx, user)
	if err != nil {
		return
	}
	return user, nil
}
