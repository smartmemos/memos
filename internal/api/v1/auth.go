package v1

import (
	"context"
	"log/slog"

	"github.com/samber/do/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartmemos/memos/internal/module/auth"
	"github.com/smartmemos/memos/internal/module/auth/model"
	"github.com/smartmemos/memos/internal/module/user"
	usermd "github.com/smartmemos/memos/internal/module/user/model"
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

func (s *AuthService) SignUp(ctx context.Context, req *v1pb.SignUpRequest) (resp *userpb.User, err error) {
	user, err := s.userService.CreateUser(ctx, &usermd.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	resp = convertUserToProto(user)
	return resp, nil
}

func (s *AuthService) SignOut(ctx context.Context, _ *v1pb.SignOutRequest) (*emptypb.Empty, error) {
	token, _ := grpc_util.GetAccessToken(ctx)
	if token != "" {
		userID, _ := grpc_util.GetUserID(ctx)
		if userID > 0 {
			err := s.authService.DeleteAccessToken(ctx, userID, token)
			if err != nil {
				slog.Error("failed to delete access token", "error", err)
			}
		}
	}

	if err := grpc_util.ClearAccessTokenCookie(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set grpc header, error: %v", err)
	}
	return &emptypb.Empty{}, nil
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
