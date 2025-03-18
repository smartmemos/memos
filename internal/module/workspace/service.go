//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package workspace

import (
	"context"

	"github.com/smartmemos/memos/internal/module/workspace/model"
)

type Service interface {
	GetProfile(ctx context.Context, req *model.GetProfileRequest) (*model.Profile, error)
	GetSetting(ctx context.Context, key model.SettingKey, value any) error
}
