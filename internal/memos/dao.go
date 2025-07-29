//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memos

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

type DAO interface {
	CreateUser(ctx context.Context, m *model.User) error
}
