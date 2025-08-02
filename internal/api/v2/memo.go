package v2

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/model"
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

	resp = connect.NewResponse(convertMemoToProto(&model.Memo{}))
	return
}

func (s *MemoService) ListMemos(ctx context.Context, req *connect.Request[v2pb.ListMemosRequest]) (resp *connect.Response[v2pb.ListMemosResponse], err error) {
	resp = connect.NewResponse(&v2pb.ListMemosResponse{
		Memos: []*modelpb.Memo{
			convertMemoToProto(&model.Memo{}),
		},
	})
	return
}

func (s *MemoService) GetMemo(ctx context.Context, req *connect.Request[v2pb.GetMemoRequest]) (resp *connect.Response[modelpb.Memo], err error) {

	resp = connect.NewResponse(convertMemoToProto(&model.Memo{}))
	return
}

func convertMemoToProto(memo *model.Memo) *modelpb.Memo {
	return &modelpb.Memo{
		Name:       fmt.Sprintf("memos/%d", memo.ID),
		Content:    memo.Content,
		CreateTime: timestamppb.New(memo.CreatedAt),
		UpdateTime: timestamppb.New(memo.UpdatedAt),
	}
}
