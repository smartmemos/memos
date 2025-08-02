package v2

import (
	"context"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/memos"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type MemoService struct {
	v2pb.UnimplementedMemoServiceHandler
	memosService memos.Service
}

func NewMemoService(i do.Injector) (*MemoService, error) {
	return &MemoService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

func (s *MemoService) CreateMemo(ctx context.Context, req *connect.Request[v2pb.CreateMemoRequest]) (resp *connect.Response[modelpb.Memo], err error) {

	return
}
