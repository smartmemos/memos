package v1

import (
	"context"
	"fmt"

	"github.com/samber/do/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/module/memo"
	"github.com/smartmemos/memos/internal/module/memo/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	memopb "github.com/smartmemos/memos/internal/proto/model/memo"
)

type MemoService struct {
	v1pb.UnimplementedMemoServiceServer
	memoService memo.Service
}

func NewMemoService(i do.Injector) (*MemoService, error) {
	return &MemoService{
		memoService: do.MustInvoke[memo.Service](i),
	}, nil
}

func (s *MemoService) CreateMemo(ctx context.Context, req *v1pb.CreateMemoRequest) (resp *memopb.Memo, err error) {
	memo, err := s.memoService.CreateMemo(ctx, &model.CreateMemoRequest{
		Content:    req.Memo.Content,
		Visibility: convertFromProtoVisibility(req.Memo.Visibility),
	})
	if err != nil {
		return
	}
	resp = convertMemoToProto(memo)
	return
}

func convertMemoToProto(memo *model.Memo) *memopb.Memo {
	resp := &memopb.Memo{
		Name: fmt.Sprintf("%d", memo.ID),
		// State:       convertStateFromStore(memo.RowStatus),
		// Creator:     fmt.Sprintf("%s%d", UserNamePrefix, memo.CreatorID),
		CreateTime: timestamppb.New(memo.CreatedAt),
		UpdateTime: timestamppb.New(memo.UpdatedAt),
		// DisplayTime: timestamppb.New(time.Unix(displayTs, 0)),
		Content:    memo.Content,
		Visibility: convertVisibilityToProto(memo.Visibility),
		// Pinned:     memo.Pinned,
	}
	return resp
}

func convertFromProtoVisibility(v memopb.Visibility) model.Visibility {
	return model.Visibility(memopb.Visibility_name[int32(v)])
}

func convertVisibilityToProto(v model.Visibility) memopb.Visibility {
	return memopb.Visibility(memopb.Visibility_value[string(v)])
}

func (s *MemoService) ListMemos(ctx context.Context, req *v1pb.ListMemosRequest) (resp *v1pb.ListMemosResponse, err error) {
	return
}
