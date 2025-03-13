package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/smartmemos/memos/internal/module/system/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) (user *model.User, err error) {
	if req.Password == "" {
		err = errors.New("password is empty")
		return
	}
	if !model.UsernameReg.MatchString(req.Username) {
		err = errors.Errorf("invalid username: %s", req.Username)
		return
	}
	total, err := s.dao.CountUsers(ctx, &model.FindUserFilter{Username: req.Username})
	if err != nil {
		return
	} else if total > 0 {
		err = errors.Errorf("username %s already exist", req.Username)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.Errorf("failed to generate password hash: %v", err)
		return
	}
	user = &model.User{
		Username:     req.Username,
		Role:         model.RoleAdmin,
		Status:       model.Normal,
		Email:        req.Email,
		Nickname:     req.Nickname,
		PasswordHash: string(passwordHash),
	}
	if err = s.dao.CreateUser(ctx, user); err != nil {
		err = errors.Errorf("failed to create user: %v", err)
		return
	}
	return
}
