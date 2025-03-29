//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/smartmemos/memos/internal/module/memo/model"
)

type Service interface {
	CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (*model.Memo, error)
	UpdateMemo(ctx context.Context, req *model.UpdateMemoRequest) (*model.MemoInfo, error)
	ListMemos(ctx context.Context, req *model.ListMemosRequest) ([]*model.MemoInfo, error)
	DeleteMemo(ctx context.Context, id int64) error
}
