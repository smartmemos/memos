package v1

import (
	"context"

	"github.com/samber/do/v2"
	"github.com/sirupsen/logrus"

	"github.com/smartmemos/memos/internal/module/user"
	"github.com/smartmemos/memos/internal/module/user/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	userpb "github.com/smartmemos/memos/internal/proto/model/user"
)

type UserService struct {
	v1pb.UnimplementedUserServiceServer
	userService user.Service
}

func NewUserService(i do.Injector) (*UserService, error) {
	return &UserService{
		userService: do.MustInvoke[user.Service](i),
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *v1pb.CreateUserRequest) (resp *userpb.User, err error) {
	user, err := s.userService.CreateUser(ctx, &model.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	return convertUserToProto(user), nil
}

func (s *UserService) GetUserSetting(ctx context.Context, req *v1pb.GetUserSettingRequest) (resp *userpb.Setting, err error) {
	setting, err := s.userService.GetSetting(ctx, &model.GetSettingRequest{})
	if err != nil {
		return
	}
	logrus.Info(setting)
	return
}

func convertUserToProto(user *model.User) *userpb.User {
	return &userpb.User{
		Id:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
}
