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
}
