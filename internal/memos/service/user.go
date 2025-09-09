package service

import (
	"context"
	"errors"

	"github.com/samber/lo"
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

func (s *Service) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (user *model.User, err error) {
	user, err = s.dao.FindUserByID(ctx, req.ID)
	if err != nil {
		return
	}
	update := make(map[string]any)
	if lo.Contains(req.UpdateMask, "username") {
		update["username"] = req.Username
	}
	if lo.Contains(req.UpdateMask, "email") {
		update["email"] = req.Email
	}
	if lo.Contains(req.UpdateMask, "display_name") {
		update["nickname"] = req.Nickname
	}
	if lo.Contains(req.UpdateMask, "avatar_url") {
		update["avatar_url"] = req.AvatarURL
	}
	if lo.Contains(req.UpdateMask, "password") {
		var passwordHash []byte
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			err = errors.New("failed to generate password hash")
			return
		}
		update["password_hash"] = string(passwordHash)
	}
	if lo.Contains(req.UpdateMask, "role") {
		update["role"] = req.Role
	}
	if lo.Contains(req.UpdateMask, "description") {
		update["description"] = req.Description
	}
	if lo.Contains(req.UpdateMask, "status") {
		update["status"] = req.Status
	}
	if len(update) == 0 {
		return user, nil
	}
	if err = s.dao.UpdateUser(ctx, user, update); err != nil {
		return
	}
	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int64) (user *model.User, err error) {
	return s.dao.FindUserByID(ctx, id)
}

func (s *Service) ListUsers(ctx context.Context, req *model.ListUsersRequest) (total int64, users []*model.User, err error) {
	filter := &model.UserFilter{
		Query: req.Query,
	}

	total, err = s.dao.CountUsers(ctx, filter)
	if err != nil {
		return
	}
	users, err = s.dao.FindUsers(ctx, filter)
	return
}

func (s *Service) SearchUsers(ctx context.Context, req *model.SearchUsersRequest) (total int64, users []*model.User, err error) {
	filter := &model.UserFilter{
		Query: req.Query,
	}
	total, err = s.dao.CountUsers(ctx, filter)
	if err != nil {
		return
	}
	users, err = s.dao.FindUsers(ctx, filter)
	return
}
