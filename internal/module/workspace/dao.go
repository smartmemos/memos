//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package workspace

import (
	"context"

	"github.com/smartmemos/memos/internal/module/workspace/model"
)

type DAO interface {
	CreateSetting(ctx context.Context, m *model.Setting) error
	UpdateSettings(ctx context.Context, filter *model.FindSettingFilter, update map[string]any) (int64, error)
	UpdateSetting(ctx context.Context, m *model.Setting, update map[string]any) error
	CountSettings(ctx context.Context, filter *model.FindSettingFilter) (int64, error)
	FindSettings(ctx context.Context, filter *model.FindSettingFilter) ([]*model.Setting, error)
	FindSetting(ctx context.Context, filter *model.FindSettingFilter) (*model.Setting, error)
	FindMemoRelatedSetting(ctx context.Context) (*model.MemoRelatedSetting, error)
}
