//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package user

import (
	"context"

	"github.com/smartmemos/memos/internal/module/user/model"
)

type DAO interface {
	CreateAccessToken(ctx context.Context, m *model.AccessToken) error
	FindAccessToken(ctx context.Context, filter *model.FindAccessTokenFilter) (*model.AccessToken, error)
	DeleteAccessToken(ctx context.Context, filter *model.FindAccessTokenFilter) error
	CountAccessTokens(ctx context.Context, filter *model.FindAccessTokenFilter) (int64, error)

	CreateUser(ctx context.Context, m *model.User) error
	UpdateUsers(ctx context.Context, filter *model.FindUserFilter, update map[string]any) (int64, error)
	UpdateUser(ctx context.Context, m *model.User, update map[string]any) error
	CountUsers(ctx context.Context, filter *model.FindUserFilter) (int64, error)
	FindUsers(ctx context.Context, filter *model.FindUserFilter) ([]*model.User, error)
	FindUserByID(ctx context.Context, id int64) (*model.User, error)
	FindUser(ctx context.Context, filter *model.FindUserFilter) (*model.User, error)

	FindSettings(ctx context.Context, filter *model.FindSettingFilter) ([]*model.Setting, error)
	FindSetting(ctx context.Context, filter *model.FindSettingFilter) (*model.Setting, error)
}
