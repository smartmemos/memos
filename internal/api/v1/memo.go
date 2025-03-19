package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/samber/do/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/module/memo"
	"github.com/smartmemos/memos/internal/module/memo/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	memopb "github.com/smartmemos/memos/internal/proto/model/memo"
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

func (s *MemoService) CreateMemo(ctx context.Context, req *v1pb.CreateMemoRequest) (resp *memopb.Memo, err error) {
	memo, err := s.memo.CreateMemo(ctx, &model.CreateMemoRequest{})
	if err != nil {
		return
	}

	return
}

func convertMemoToProto(memo *model.Memo) *memopb.Memo {
	resp := &memopb.Memo{
		Name: fmt.Sprintf("%d", memo.ID),
		// State:       convertStateFromStore(memo.RowStatus),
		// Creator:     fmt.Sprintf("%s%d", UserNamePrefix, memo.CreatorID),
		CreateTime:  timestamppb.New(time.Unix(memo.CreatedTs, 0)),
		UpdateTime:  timestamppb.New(time.Unix(memo.UpdatedTs, 0)),
		DisplayTime: timestamppb.New(time.Unix(displayTs, 0)),
		Content:     memo.Content,
		Visibility:  convertVisibilityFromStore(memo.Visibility),
		Pinned:      memo.Pinned,
	}
	return resp
}

func (s *MemoService) ListMemos(ctx context.Context, req *v1pb.ListMemosRequest) (resp *v1pb.ListMemosResponse, err error) {
	return
}
