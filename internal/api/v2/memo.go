package v2

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"
	"github.com/usememos/gomark/renderer"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/utils"
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
	info, err := convertMemoToProto(&model.Memo{})
	if err != nil {
		return
	}
	resp = connect.NewResponse(info)
	return
}

func (s *MemoService) ListMemos(ctx context.Context, req *connect.Request[v2pb.ListMemosRequest]) (resp *connect.Response[v2pb.ListMemosResponse], err error) {
	var req2 = &model.ListMemosRequest{}

	userInfo := utils.GetInfo(ctx)
	if userInfo == nil {
		req2.VisibilityList = []model.Visibility{model.Public}
	} else {
		// if req2.CreatorID == nil {
		// 	filter := fmt.Sprintf(`creator_id == %d || visibility in ["PUBLIC", "PROTECTED"]`, currentUser.ID)
		// 	req2.Filters = append(req2.Filters, filter)
		// } else if *req2.CreatorID != currentUser.ID {
		// 	req2.VisibilityList = []model.Visibility{model.Public, model.Protected}
		// }
	}

	if req.Msg.State == modelpb.State_ARCHIVED {
		req2.Status = model.Archived
	} else {
		req2.Status = model.Normal
	}
	total, memos, err := s.memosService.ListMemos(ctx, req2)
	if err != nil {
		return
	}
	nextPageToken, err := utils.GetPageToken(1, 0)
	if err != nil {
		return
	}

	var list []*modelpb.Memo
	for _, memo := range memos {
		var info *modelpb.Memo
		info, err = convertMemoToProto(memo)
		if err != nil {
			return
		}
		list = append(list, info)
	}

	resp = connect.NewResponse(&v2pb.ListMemosResponse{
		Memos:         list,
		TotalSize:     int32(total),
		NextPageToken: nextPageToken,
	})
	return
}

func (s *MemoService) GetMemo(ctx context.Context, req *connect.Request[v2pb.GetMemoRequest]) (resp *connect.Response[modelpb.Memo], err error) {
	info, err := convertMemoToProto(&model.Memo{})
	if err != nil {
		return
	}
	resp = connect.NewResponse(info)
	return
}

func convertMemoToProto(memo *model.Memo) (info *modelpb.Memo, err error) {
	info = &modelpb.Memo{
		Name:       fmt.Sprintf("memos/%d", memo.ID),
		Content:    memo.Content,
		CreateTime: timestamppb.New(memo.CreatedAt),
		UpdateTime: timestamppb.New(memo.UpdatedAt),
	}
	nodes, err := parser.Parse(tokenizer.Tokenize(memo.Content))
	if err != nil {
		err = errors.Wrap(err, "failed to parse content")
		return
	}
	plainText := renderer.NewStringRenderer().Render(nodes)
	info.Snippet = lo.If(len(plainText) > 64, lo.Substring(plainText, 0, 64)+"...").Else(plainText)
	info.Nodes = convertFromASTNodes(nodes)

	return
}
