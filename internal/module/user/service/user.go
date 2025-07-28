package service

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
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

func (s *Service) GetUserByID(ctx context.Context, userID int64) (user *model.User, err error) {
	return s.dao.FindUserByID(ctx, userID)
}

func (s *Service) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (user *model.User, err error) {
	userID, err := grpc_util.GetUserID(ctx)
	if err != nil {
		err = errors.Errorf("failed to get user: %v", err)
		return
	}
	currentUser, err := s.dao.FindUserByID(ctx, req.ID)
	if err != nil {
		return
	}
	if currentUser.ID != userID && currentUser.Role != model.RoleAdmin && currentUser.Role != model.RoleHost {
		err = errors.New("permission denied")
		return
	}
	if len(req.UpdateMask) == 0 {
		err = errors.New("update mask is empty")
		return
	}
	setting, err := s.wsDao.FindGeneralSetting(ctx)
	if err != nil {
		return
	}
	update := make(map[string]any)
	if lo.Contains(req.UpdateMask, "username") {
		if setting.DisallowChangeUsername {
			err = errors.New("permission denied: disallow change username")
			return
		}

		if !model.UsernameReg.MatchString(strings.ToLower(req.Username)) {
			err = errors.Errorf("invalid username: %s", req.Username)
			return
		}
		if !util.UIDMatcher.MatchString(strings.ToLower(req.Username)) {
			err = "invalid username: %s", request.User.Username)
		}
		update.Username = &request.User.Username
	}

	user, err = s.dao.FindUserByID(ctx, req.ID)
	if err != nil {
		return
	}

	return
}
