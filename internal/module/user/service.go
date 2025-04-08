//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package user

import (
	"context"

	"github.com/smartmemos/memos/internal/module/user/model"
)

type Service interface {
	CreateUser(context.Context, *model.CreateUserRequest) (*model.User, error)
	UpdateUser(context.Context, *model.UpdateUserRequest) (*model.User, error)
	GetUserByID(context.Context, int64) (*model.User, error)
	// GetSetting(ctx context.Context, req *model.GetSettingRequest) (*model.Setting, error)
	ListAllUserStats(ctx context.Context, req *model.ListAllUserStatsRequest) (stats *model.Stats, err error)
	CreateAccessToken(ctx context.Context, req *model.CreateAccessTokenRequest) (*model.AccessToken, error)
	ListAccessTokens(ctx context.Context, req *model.ListAccessTokensRequest) ([]*model.AccessToken, error)
	DeleteAccessToken(ctx context.Context, req *model.DeleteAccessTokenRequest) error
}
