package v1

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
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
	v1pb.UnimplementedAuthServiceHandler
	authService auth.Service
	userService user.Service
}

func NewAuthService(i do.Injector) (*AuthService, error) {
	return &AuthService{
		authService: do.MustInvoke[auth.Service](i),
		userService: do.MustInvoke[user.Service](i),
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req *connect.Request[v1pb.SignInRequest]) (resp *connect.Response[userpb.User], err error) {
	accessToken, err := s.authService.SignIn(ctx, &model.SignInRequest{
		Username:    req.Msg.Username,
		Password:    req.Msg.Password,
		NeverExpire: req.Msg.NeverExpire,
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
	return connect.NewResponse(convertUserToProto(user)), nil
}

func (s *AuthService) SignUp(ctx context.Context, req *connect.Request[v1pb.SignUpRequest]) (resp *connect.Response[userpb.User], err error) {
	user, err := s.userService.CreateUser(ctx, &usermd.CreateUserRequest{
		Username: req.Msg.Username,
		Password: req.Msg.Password,
	})
	if err != nil {
		return
	}
	return connect.NewResponse(convertUserToProto(user)), nil
}

func (s *AuthService) SignOut(ctx context.Context, _ *connect.Request[v1pb.SignOutRequest]) (*connect.Response[emptypb.Empty], error) {
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
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *AuthService) GetAuthStatus(ctx context.Context, req *connect.Request[v1pb.GetAuthStatusRequest]) (resp *connect.Response[userpb.User], err error) {
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
	return connect.NewResponse(convertUserToProto(user)), nil
}
