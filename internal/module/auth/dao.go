//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package auth

import (
	"context"

	"github.com/smartmemos/memos/internal/module/auth/model"
)

type DAO interface {
	CreateAccessToken(ctx context.Context, m *model.AccessToken) error
	DeleteAccessToken(ctx context.Context, filter *model.FindAccessTokenFilter) error
}
