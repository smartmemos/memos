package v1

import (
	"context"

	"github.com/samber/do/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartmemos/memos/internal/module/auth"
	"github.com/smartmemos/memos/internal/module/auth/model"
	"github.com/smartmemos/memos/internal/module/user"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	userpb "github.com/smartmemos/memos/internal/proto/model/user"
)

type AuthService struct {
	v1pb.UnimplementedAuthServiceServer
	authService auth.Service
	userService user.Service
}

func NewAuthService(i do.Injector) (*AuthService, error) {
	return &AuthService{
		authService: do.MustInvoke[auth.Service](i),
		userService: do.MustInvoke[user.Service](i),
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req *v1pb.SignInRequest) (resp *userpb.User, err error) {
	accessToken, err := s.authService.SignIn(ctx, &model.SignInRequest{
		Username:    req.Username,
		Password:    req.Password,
		NeverExpire: req.NeverExpire,
	})
	if err != nil {
		return
	}
	if err = grpc_util.SetAccessTokenCookie(ctx, accessToken.Token, accessToken.ExpiresAt); err != nil {
		return
	}
	user, err := s.userService.GetUserByID(ctx, accessToken.UserId)
	if err != nil {
		return
	}
	resp = convertUserToProto(user)
	return resp, nil
}

func (s *AuthService) GetAuthStatus(ctx context.Context, req *v1pb.GetAuthStatusRequest) (resp *userpb.User, err error) {
	userID, err := grpc_util.GetUserID(ctx)
	if err != nil {
		return
	}
	user, err := s.userService.GetUserByID(ctx, userID)
	if err != nil {
		if err = grpc_util.ClearAccessTokenCookie(ctx); err != nil {
			err = status.Errorf(codes.Internal, "failed to set grpc header: %v", err)
		} else {
			err = status.Errorf(codes.Unauthenticated, "user not found")
		}
		return
	}
	resp = convertUserToProto(user)
	return
}
