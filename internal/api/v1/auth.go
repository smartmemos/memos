package v1

import (
	"context"

	"github.com/samber/do/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartmemos/memos/internal/api/base"
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
	accessToken, err := s.system.SignIn(ctx, &model.SignInRequest{
		Username:    req.Username,
		Password:    req.Password,
		NeverExpire: req.NeverExpire,
	})
	if err != nil {
		return
	}
	if err = base.SetAccessTokenCookie(ctx, accessToken.Token, accessToken.ExpiresAt); err != nil {
		return
	}
	user, err := s.system.GetUserByID(ctx, accessToken.UserId)
	if err != nil {
		return
	}
	resp = convertUserToProto(user)
	return resp, nil
}

func (s *AuthService) GetAuthStatus(ctx context.Context, req *v1pb.GetAuthStatusRequest) (resp *systempb.User, err error) {
	userID, err := base.GetUserID(ctx)
	if err != nil {
		return
	}
	user, err := s.system.GetUserByID(ctx, userID)
	if err != nil {
		if err = base.ClearAccessTokenCookie(ctx); err != nil {
			err = status.Errorf(codes.Internal, "failed to set grpc header: %v", err)
		} else {
			err = status.Errorf(codes.Unauthenticated, "user not found")
		}
		return
	}
	resp = convertUserToProto(user)
	return
}
