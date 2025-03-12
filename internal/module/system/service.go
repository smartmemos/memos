//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package system

import (
	"context"

	"github.com/smartmemos/memos/internal/module/system/model"
)

type Service interface {
	CreateUser(context.Context, *model.CreateUserRequest) (*model.User, error)
	SignIn(context.Context, *model.SignInRequest) (*model.User, error)
}
