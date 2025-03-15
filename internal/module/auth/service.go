//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package auth

import (
	"context"

	"github.com/smartmemos/memos/internal/module/auth/model"
)

type Service interface {
	SignIn(ctx context.Context, req *model.SignInRequest) (accessToken *model.AccessToken, err error)
	Authenticate(ctx context.Context, token string) (accessToken *model.AccessToken, err error)
	DeleteAccessToken(ctx context.Context, userId int64, token string) error
	ValidateAccessToken(ctx context.Context, userId int64, token string) (bool, error)
}
