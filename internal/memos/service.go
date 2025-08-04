//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package memos

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

type Service interface {
	// CreateUser creates a user.
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	// GetUserByID gets a user by ID.
	GetUserByID(ctx context.Context, id int64) (*model.User, error)

	// CreateSession creates a session.
	CreateSession(ctx context.Context, req *model.CreateSessionRequest) (*model.Session, error)

	// CreateMemo creates a memo.
	CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (*model.Memo, error)
	// ListMemos lists the memos.
	ListMemos(ctx context.Context, req *model.ListMemosRequest) (int64, []*model.Memo, error)
	// GetMemo gets a memo by ID.
	GetMemo(ctx context.Context, req *model.GetMemoRequest) (*model.Memo, error)

	// GetGeneralSetting gets the general setting.
	GetGeneralSetting(ctx context.Context) (*model.GeneralSetting, error)
	// GetStorageSetting gets the storage setting.
	GetStorageSetting(ctx context.Context) (*model.StorageSetting, error)
	// GetMemoRelatedSetting gets the memo related setting.
	GetMemoRelatedSetting(ctx context.Context) (*model.MemoRelatedSetting, error)
}
