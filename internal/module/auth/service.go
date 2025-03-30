//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package auth

import (
	"context"
	"time"

	"github.com/smartmemos/memos/internal/module/auth/model"
	usermd "github.com/smartmemos/memos/internal/module/user/model"
)

type Service interface {
	SignIn(ctx context.Context, req *model.SignInRequest) (accessToken *usermd.AccessToken, err error)
	Authenticate(ctx context.Context, token string) (accessToken *usermd.AccessToken, err error)
	DeleteAccessToken(ctx context.Context, userId int64, token string) error
	ValidateAccessToken(ctx context.Context, userId int64, token string) (bool, error)
	GenerateAccessToken(_ context.Context, userID int64, issuedAt, expirationTime time.Time) (tokenStr string, err error)
}
