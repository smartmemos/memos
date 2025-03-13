package v1

import (
	"context"

	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/system"
	"github.com/smartmemos/memos/internal/module/system/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	systempb "github.com/smartmemos/memos/internal/proto/model/system"
)

type AuthService struct {
	v1pb.UnimplementedAuthServiceServer
	system system.Service
}

func NewAuthService(i do.Injector) (*AuthService, error) {
	return &AuthService{
		system: do.MustInvoke[system.Service](i),
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req *v1pb.SignInRequest) (resp *systempb.User, err error) {
	user, err := s.system.SignIn(ctx, &model.SignInRequest{
		Username:    req.Username,
		Password:    req.Password,
		NeverExpire: req.NeverExpire,
	})
	if err != nil {
		return
	}
	resp = convertUserToProto(user)
	return resp, nil
}
