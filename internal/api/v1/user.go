package v1

import (
	"context"
	"fmt"
	"time"

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
	// setting, err := s.userService.GetSetting(ctx, &model.GetSettingRequest{})
	// if err != nil {
	// 	return
	// }
	// logrus.Info(setting)
	resp = &userpb.Setting{
		Locale:         "en",
		Appearance:     "system",
		MemoVisibility: "PRIVATE",
	}
	return
}

func (s *UserService) ListAllUserStats(ctx context.Context, req *v1pb.ListAllUserStatsRequest) (resp *v1pb.ListAllUserStatsResponse, err error) {
	// s.userService.ListAllUserStats(ctx, &model.ListAllUserStatsRequest{})
	resp = &v1pb.ListAllUserStatsResponse{
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
	}
	return
}

// GetUserStats returns the stats of a user.
func (s *UserService) GetUserStats(context.Context, *v1pb.GetUserStatsRequest) (resp *userpb.Stats, err error) {
	return
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

func (s *UserService) CreateAccessToken(ctx context.Context, req *v1pb.CreateAccessTokenRequest) (resp *userpb.AccessToken, err error) {
	userID, err := ExtractUserIDFromName(req.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
	}
	token, err := s.userService.CreateAccessToken(ctx, &model.CreateAccessTokenRequest{
		UserId:      int64(userID),
		Description: req.Description,
		IssuedAt:    time.Now(),
		ExpiresAt:   req.ExpiresAt.AsTime(),
	})
	if err != nil {
		return
	}
	resp = &userpb.AccessToken{
		AccessToken: token.Token,
		Description: token.Description,
		IssuedAt:    timestamppb.New(token.IssuedAt),
		ExpiresAt:   timestamppb.New(token.ExpiresAt),
	}
	return
}

func (s *UserService) ListAccessTokens(ctx context.Context, req *v1pb.ListAccessTokensRequest) (resp *v1pb.ListAccessTokensResponse, err error) {
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
	resp = &v1pb.ListAccessTokensResponse{
		AccessTokens: list,
	}
	return
}

func (s *UserService) DeleteAccessToken(ctx context.Context, req *v1pb.DeleteAccessTokenRequest) (_ *emptypb.Empty, err error) {
	userID, err := ExtractUserIDFromName(req.Name)
	if err != nil {
		return
	}
	err = s.userService.DeleteAccessToken(ctx, &model.DeleteAccessTokenRequest{
		UserID:      int64(userID),
		AccessToken: req.AccessToken,
	})
	return
}
