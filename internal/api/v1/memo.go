package v1

import (
	"context"

	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/memo"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
)

type MemoService struct {
	v1pb.UnimplementedMemoServiceServer
	memo memo.Service
}

func NewMemoService(i do.Injector) (*MemoService, error) {
	return &MemoService{
		memo: do.MustInvoke[memo.Service](i),
	}, nil
}

func (s *MemoService) ListMemos(context.Context, *v1pb.ListMemosRequest) (*v1pb.ListMemosResponse, error) {
	return nil, nil
}
