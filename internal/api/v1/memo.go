package v1

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/module/memo"
	"github.com/smartmemos/memos/internal/module/memo/model"
	"github.com/smartmemos/memos/internal/module/workspace"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	"github.com/smartmemos/memos/internal/proto/model/common"
	memopb "github.com/smartmemos/memos/internal/proto/model/memo"
)

type MemoService struct {
	v1pb.UnimplementedMemoServiceServer
	memoService      memo.Service
	workspaceService workspace.Service
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
	var pageSize, page int
	if req.PageToken != "" {
		page, pageSize, err = parsePageToken(req.PageToken)
		if err != nil {
			return
		}
	} else {
		pageSize = int(req.PageSize)
		if pageSize <= 0 {
			pageSize = DefaultPageSize
		}
	}
	memos, err := s.memoService.ListMemos(ctx, &model.ListMemosRequest{
		PageSize: pageSize + 1,
		Page:     page,
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to list memos: %v", err)
		return
	}
	nextPageToken := ""
	if len(memos) > pageSize {
		memos = memos[:pageSize]
		nextPageToken, err = getPageToken(pageSize, page+1)
		if err != nil {
			err = status.Errorf(codes.Internal, "failed to get next page token, error: %v", err)
			return
		}
	}
	var list []*memopb.Memo
	for _, memo := range memos {
		item, err := s.convertMemoToProto(ctx, memo)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert memo")
		}
		list = append(list, item)
	}

	resp = &v1pb.ListMemosResponse{
		Memos:         list,
		NextPageToken: nextPageToken,
	}
	return
}

func (s *MemoService) convertMemoToProto(ctx context.Context, memo *model.MemoInfo) (item *memopb.Memo, err error) {
	name := fmt.Sprintf("%s%s", MemoNamePrefix, memo.UID)
	item = &memopb.Memo{
		Name:        name,
		State:       common.State(common.State_value[memo.Status]),
		Creator:     fmt.Sprintf("%s%d", UserNamePrefix, memo.CreatorID),
		CreateTime:  timestamppb.New(memo.CreatedAt),
		UpdateTime:  timestamppb.New(memo.UpdatedAt),
		DisplayTime: timestamppb.New(memo.UpdatedAt),
		Content:     memo.Content,
		Visibility:  convertVisibilityToProto(memo.Visibility),
		Pinned:      memo.Pinned,
	}
	return
}
