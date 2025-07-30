//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package memos

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

type Service interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)

	CreateSession(ctx context.Context, req *model.CreateSessionRequest) (*model.Session, error)
}
