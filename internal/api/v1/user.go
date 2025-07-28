package v1

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/module/user"
	"github.com/smartmemos/memos/internal/module/user/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	userpb "github.com/smartmemos/memos/internal/proto/model/user"
)

type UserService struct {
	v1pb.UnimplementedUserServiceHandler
	userService user.Service
}

func NewUserService(i do.Injector) (*UserService, error) {
	return &UserService{
		userService: do.MustInvoke[user.Service](i),
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *connect.Request[v1pb.CreateUserRequest]) (resp *connect.Response[userpb.User], err error) {
	user, err := s.userService.CreateUser(ctx, &model.CreateUserRequest{
		Username: req.Msg.Username,
		Password: req.Msg.Password,
	})
	if err != nil {
		return
	}
	return connect.NewResponse(convertUserToProto(user)), nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *connect.Request[v1pb.UpdateUserRequest]) (resp *connect.Response[userpb.User], err error) {
	user, err := s.userService.UpdateUser(ctx, &model.UpdateUserRequest{})
	if err != nil {
		return
	}
	return connect.NewResponse(convertUserToProto(user)), nil
}

func (s *UserService) GetUserSetting(ctx context.Context, req *connect.Request[v1pb.GetUserSettingRequest]) (resp *connect.Response[userpb.Setting], err error) {
	// setting, err := s.userService.GetSetting(ctx, &model.GetSettingRequest{})
	// if err != nil {
	// 	return
	// }
	// logrus.Info(setting)
	return connect.NewResponse(&userpb.Setting{
		Locale:         "en",
		Appearance:     "system",
		MemoVisibility: "PRIVATE",
	}), nil
}

func (s *UserService) ListAllUserStats(ctx context.Context, req *connect.Request[v1pb.ListAllUserStatsRequest]) (resp *connect.Response[v1pb.ListAllUserStatsResponse], err error) {
	// s.userService.ListAllUserStats(ctx, &model.ListAllUserStatsRequest{})
	return connect.NewResponse(&v1pb.ListAllUserStatsResponse{
		UserStats: []*userpb.Stats{
			{
				Name:                  "",
				PinnedMemos:           []string{},
				TagCount:              map[string]int32{},
				MemoDisplayTimestamps: []*timestamppb.Timestamp{},
				MemoTypeStats:         &userpb.Stats_MemoTypeStats{},
				TotalMemoCount:        1,
			},
		},
	}), nil
}

// GetUserStats returns the stats of a user.
func (s *UserService) GetUserStats(ctx context.Context, req *connect.Request[v1pb.GetUserStatsRequest]) (resp *connect.Response[userpb.Stats], err error) {
	return connect.NewResponse(&userpb.Stats{}), nil
}

func convertUserToProto(user *model.User) *userpb.User {
	return &userpb.User{
		Id:          user.ID,
		Name:        fmt.Sprintf("users/%d", user.ID),
		Username:    user.Username,
		Nickname:    user.Nickname,
		Email:       user.Email,
		AvatarUrl:   user.AvatarURL,
		Description: user.Description,
		CreateAt:    timestamppb.New(user.CreatedAt),
		UpdateAt:    timestamppb.New(user.UpdatedAt),
		Role:        userpb.User_Role(userpb.User_Role_value[string(user.Role)]),
		State:       userpb.State(userpb.State_value[string(user.Status)]),
	}
}

func (s *UserService) CreateAccessToken(ctx context.Context, req *connect.Request[v1pb.CreateAccessTokenRequest]) (resp *connect.Response[userpb.AccessToken], err error) {
	userID, err := ExtractUserIDFromName(req.Msg.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
	}
	token, err := s.userService.CreateAccessToken(ctx, &model.CreateAccessTokenRequest{
		UserId:      int64(userID),
		Description: req.Msg.Description,
		IssuedAt:    time.Now(),
		ExpiresAt:   req.Msg.ExpiresAt.AsTime(),
	})
	if err != nil {
		return
	}
	return connect.NewResponse(&userpb.AccessToken{
		AccessToken: token.Token,
		Description: token.Description,
		IssuedAt:    timestamppb.New(token.IssuedAt),
		ExpiresAt:   timestamppb.New(token.ExpiresAt),
	}), nil
}

func (s *UserService) ListAccessTokens(ctx context.Context, req *connect.Request[v1pb.ListAccessTokensRequest]) (resp *connect.Response[v1pb.ListAccessTokensResponse], err error) {
	tokens, err := s.userService.ListAccessTokens(ctx, &model.ListAccessTokensRequest{})
	if err != nil {
		return
	}
	var list []*userpb.AccessToken
	for _, token := range tokens {
		item := &userpb.AccessToken{
			AccessToken: token.Token,
			Description: token.Description,
			IssuedAt:    timestamppb.New(token.IssuedAt),
			ExpiresAt:   timestamppb.New(token.ExpiresAt),
		}
		list = append(list, item)
	}
	return connect.NewResponse(&v1pb.ListAccessTokensResponse{
		AccessTokens: list,
	}), nil
}

func (s *UserService) DeleteAccessToken(ctx context.Context, req *connect.Request[v1pb.DeleteAccessTokenRequest]) (_ *connect.Response[emptypb.Empty], err error) {
	userID, err := ExtractUserIDFromName(req.Msg.Name)
	if err != nil {
		return
	}
	err = s.userService.DeleteAccessToken(ctx, &model.DeleteAccessTokenRequest{
		UserID:      int64(userID),
		AccessToken: req.Msg.AccessToken,
	})
	if err != nil {
		return
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
