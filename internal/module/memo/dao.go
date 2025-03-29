//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/smartmemos/memos/internal/module/memo/model"
)

type DAO interface {
	CreateMemo(ctx context.Context, m *model.Memo) error
	UpdateMemos(ctx context.Context, filter *model.FindMemoFilter, update map[string]any) (int64, error)
	UpdateMemo(ctx context.Context, m *model.Memo, update map[string]any) error
	CountMemos(ctx context.Context, filter *model.FindMemoFilter) (int64, error)
	FindMemos(ctx context.Context, filter *model.FindMemoFilter) ([]*model.Memo, error)
	FindMemoByID(ctx context.Context, id int64) (*model.Memo, error)
	FindMemo(ctx context.Context, filter *model.FindMemoFilter) (*model.Memo, error)
	DeleteMemos(ctx context.Context, filter *model.FindMemoFilter) (int64, error)
}
