//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package user

import (
	"context"

	"github.com/smartmemos/memos/internal/module/user/model"
)

type Service interface {
	CreateUser(context.Context, *model.CreateUserRequest) (*model.User, error)
	GetUserByID(context.Context, int64) (*model.User, error)
	GetSetting(ctx context.Context, req *model.GetSettingRequest) (*model.Setting, error)
}
