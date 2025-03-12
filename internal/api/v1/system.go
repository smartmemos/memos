package v1

import (
	"context"

	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/system"
	"github.com/smartmemos/memos/internal/module/system/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	systempb "github.com/smartmemos/memos/internal/proto/model/system"
)

type SystemService struct {
	v1pb.UnimplementedSystemServiceServer
	system system.Service
}

func NewSystemService(i do.Injector) (*SystemService, error) {
	return &SystemService{
		system: do.MustInvoke[system.Service](i),
	}, nil
}

func (s *SystemService) CreateUser(ctx context.Context, req *v1pb.CreateUserRequest) (resp *systempb.User, err error) {
	user, err := s.system.CreateUser(ctx, &model.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	return convertUserToProto(user), nil
}

func convertUserToProto(user *model.User) *systempb.User {
	return &systempb.User{
		Id:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
}
