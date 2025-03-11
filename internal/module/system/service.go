//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package system

import (
	"context"

	"github.com/smartmemos/memos/internal/module/system/model"
)

type Service interface {
	SignIn(ctx context.Context, req *model.SignInRequest) (*model.User, error)
}
