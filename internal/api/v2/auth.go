package v2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/utils"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type AuthService struct {
	v2pb.UnimplementedAuthServiceHandler
	memosService memos.Service
}

func NewAuthService(i do.Injector) (*AuthService, error) {
	return &AuthService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

func (s *AuthService) GetCurrentSession(ctx context.Context, req *connect.Request[v2pb.GetCurrentSessionRequest]) (resp *connect.Response[v2pb.GetCurrentSessionResponse], err error) {
	info := utils.GetInfo(ctx)
	if info == nil {
		// Clear auth cookies
		// if err := s.clearAuthCookies(ctx); err != nil {
		// 	return nil, connect.NewError(connect.CodeInternal, errors.New("failed to clear auth cookies"))
		// }
		// return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("user not found"))
	}
	user, err := s.memosService.GetUserByID(ctx, info.UserID)
	if err != nil {
		err = connect.NewError(connect.CodeUnauthenticated, errors.New("failed to get current user"))
		return
	}

	var lastAccessedAt *timestamppb.Timestamp
	if info.SessionID != "" {
		now := timestamppb.Now()
		// if err := s.Store.UpdateUserSessionLastAccessed(ctx, user.ID, sessionID, now); err != nil {
		// 	slog.Error("failed to update session last accessed time", "error", err)
		// }
		lastAccessedAt = now
	}

	return connect.NewResponse(&v2pb.GetCurrentSessionResponse{
		User:           convertUserToProto(user),
		LastAccessedAt: lastAccessedAt,
	}), nil
}

func (s *AuthService) CreateSession(ctx context.Context, req *connect.Request[v2pb.CreateSessionRequest]) (resp *connect.Response[v2pb.CreateSessionResponse], err error) {
	passwordCredentials := req.Msg.GetPasswordCredentials()
	if passwordCredentials == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("password credentials are required"))
	}

	session, err := s.memosService.CreateSession(ctx, &model.CreateSessionRequest{
		Username: passwordCredentials.Username,
		Password: passwordCredentials.Password,
	})
	if err != nil {
		return
	}

	user, err := s.memosService.GetUserByID(ctx, session.UserID)
	if err != nil {
		return
	}

	resp = connect.NewResponse(&v2pb.CreateSessionResponse{
		User:           convertUserToProto(user),
		LastAccessedAt: timestamppb.New(session.CreatedAt),
	})
	cookie, err := utils.BuildCookie(ctx, "memos.access-token", "abc123", "", time.Now().Add(time.Hour*24*30))
	if err != nil {
		return
	}
	resp.Header().Set("Set-Cookie", cookie)
	return
}

// SignUp creates a new user.
func (s *AuthService) SignUp(ctx context.Context, req *connect.Request[v2pb.SignUpRequest]) (resp *connect.Response[modelpb.User], err error) {
	user, err := s.memosService.CreateUser(ctx, &model.CreateUserRequest{
		Username: req.Msg.Username,
		Password: req.Msg.Password,
	})
	if err != nil {
		return
	}
	resp = connect.NewResponse(convertUserToProto(user))
	return
}

func convertUserToProto(user *model.User) *modelpb.User {
	return &modelpb.User{
		Id:          user.ID,
		Name:        fmt.Sprintf("users/%d", user.ID),
		Username:    user.Username,
		Nickname:    user.Nickname,
		Email:       user.Email,
		AvatarUrl:   user.AvatarURL,
		Description: user.Description,
		CreateAt:    timestamppb.New(user.CreatedAt),
		UpdateAt:    timestamppb.New(user.UpdatedAt),
		Role:        modelpb.User_Role(modelpb.User_Role_value[string(user.Role)]),
		State:       modelpb.State(modelpb.State_value[string(user.Status)]),
	}
}
