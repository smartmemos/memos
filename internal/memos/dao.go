//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memos

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

type DAO interface {
	// User
	CreateUser(ctx context.Context, m *model.User) error
	UpdateUsers(ctx context.Context, filter *model.FindUserFilter, update map[string]any) (int64, error)
	UpdateUser(ctx context.Context, m *model.User, update map[string]any) error
	CountUsers(ctx context.Context, filter *model.FindUserFilter) (int64, error)
	FindUsers(ctx context.Context, filter *model.FindUserFilter) ([]*model.User, error)
	FindUserByID(ctx context.Context, id int64) (*model.User, error)
	FindUser(ctx context.Context, filter *model.FindUserFilter) (*model.User, error)
	DeleteUsers(ctx context.Context, filter *model.FindUserFilter) error

	// Session
	CreateSession(ctx context.Context, m *model.Session) error

	// Memo
	CreateMemo(ctx context.Context, m *model.Memo) error
	CountMemos(ctx context.Context, filter *model.FindMemoFilter) (int64, error)
	FindMemos(ctx context.Context, filter *model.FindMemoFilter) ([]*model.Memo, error)
	FindMemoByID(ctx context.Context, id int64) (*model.Memo, error)
	FindMemo(ctx context.Context, filter *model.FindMemoFilter) (*model.Memo, error)
	DeleteMemos(ctx context.Context, filter *model.FindMemoFilter) error

	// SystemSetting
	CreateSystemSetting(ctx context.Context, m *model.SystemSetting) error
	UpdateSystemSettings(ctx context.Context, filter *model.FindSystemSettingFilter, update map[string]any) (int64, error)
	UpdateSystemSetting(ctx context.Context, m *model.SystemSetting, update map[string]any) error
	FindSystemSettings(ctx context.Context, filter *model.FindSystemSettingFilter) ([]*model.SystemSetting, error)
	FindSystemSetting(ctx context.Context, filter *model.FindSystemSettingFilter) (*model.SystemSetting, error)
	DeleteSystemSettings(ctx context.Context, filter *model.FindSystemSettingFilter) error
}
