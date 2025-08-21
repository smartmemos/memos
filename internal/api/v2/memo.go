package v2

import (
	"context"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"
	"github.com/usememos/gomark/renderer"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *MemoService) CreateMemo(ctx context.Context, request *connect.Request[v2pb.CreateMemoRequest]) (response *connect.Response[modelpb.Memo], err error) {
	logrus.Infof("CreateMemo: %+v", request.Msg.Memo)

	userInfo := utils.GetInfo(ctx)
	if userInfo == nil {
		err = errors.New("failed to get user")
		return
	}

	req := &model.CreateMemoRequest{
		UserID:     userInfo.UserID,
		Content:    request.Msg.Memo.Content,
		Visibility: model.Visibility(modelpb.Visibility_name[int32(request.Msg.Memo.Visibility)]),
		RowStatus:  model.Normal,
		// RelationType: model.RelationType(request.Msg.Memo.RelationType),
	}
	memo, err := s.memosService.CreateMemo(ctx, req)
	if err != nil {
		return
	}
	info, err := convertMemoToProto(memo)
	if err != nil {
		return
	}
	response = connect.NewResponse(info)
	return
}

func (s *MemoService) ListMemos(ctx context.Context, request *connect.Request[v2pb.ListMemosRequest]) (response *connect.Response[v2pb.ListMemosResponse], err error) {
	var req = &model.ListMemosRequest{}

	userInfo := utils.GetInfo(ctx)
	if userInfo == nil {
		req.VisibilityList = []model.Visibility{model.Public}
	} else {
		// if req2.CreatorID == nil {
		// 	filter := fmt.Sprintf(`creator_id == %d || visibility in ["PUBLIC", "PROTECTED"]`, currentUser.ID)
		// 	req2.Filters = append(req2.Filters, filter)
		// } else if *req2.CreatorID != currentUser.ID {
		// 	req2.VisibilityList = []model.Visibility{model.Public, model.Protected}
		// }
	}

	if request.Msg.State == modelpb.State_ARCHIVED {
		req.Status = model.Archived
	} else {
		req.Status = model.Normal
	}
	total, memos, err := s.memosService.ListMemos(ctx, req)
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

	response = connect.NewResponse(&v2pb.ListMemosResponse{
		Memos:         list,
		TotalSize:     int32(total),
		NextPageToken: nextPageToken,
	})
	return
}

func (s *MemoService) GetMemo(ctx context.Context, request *connect.Request[v2pb.GetMemoRequest]) (response *connect.Response[modelpb.Memo], err error) {
	info, err := convertMemoToProto(&model.Memo{})
	if err != nil {
		return
	}
	response = connect.NewResponse(info)
	return
}

func (s *MemoService) UpdateMemo(ctx context.Context, request *connect.Request[v2pb.UpdateMemoRequest]) (response *connect.Response[modelpb.Memo], err error) {
	info, err := convertMemoToProto(&model.Memo{})
	if err != nil {
		return
	}
	response = connect.NewResponse(info)
	return
}

func (s *MemoService) DeleteMemo(ctx context.Context, request *connect.Request[v2pb.DeleteMemoRequest]) (response *connect.Response[emptypb.Empty], err error) {
	err = s.memosService.DeleteMemo(ctx, &model.DeleteMemoRequest{
		UID: strings.TrimPrefix(request.Msg.Name, model.MemoNamePrefix),
	})
	if err != nil {
		return
	}
	response = connect.NewResponse(&emptypb.Empty{})
	return
}

func convertMemoToProto(memo *model.Memo) (info *modelpb.Memo, err error) {
	displayTs := memo.CreatedAt

	info = &modelpb.Memo{
		Name:        fmt.Sprintf("%s%d", model.MemoNamePrefix, memo.ID),
		State:       modelpb.State(modelpb.State_value[string(memo.RowStatus)]),
		Creator:     fmt.Sprintf("%s%d", model.UserNamePrefix, memo.CreatorID),
		Content:     memo.Content,
		DisplayTime: timestamppb.New(displayTs),
		Visibility:  modelpb.Visibility(modelpb.Visibility_value[string(memo.Visibility)]),
		Pinned:      memo.Pinned,
		CreateTime:  timestamppb.New(memo.CreatedAt),
		UpdateTime:  timestamppb.New(memo.UpdatedAt),
	}
	if memo.Payload != nil {
		info.Tags = memo.Payload.Tags
		info.Property = convertMemoPropertyToProto(memo.Payload.Property)
		info.Location = convertLocationToProto(memo.Payload.Location)
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

func convertMemoPropertyToProto(property *model.MemoPayloadProperty) *modelpb.Memo_Property {
	if property == nil {
		return nil
	}
	return &modelpb.Memo_Property{
		HasLink:            property.HasLink,
		HasTaskList:        property.HasTaskList,
		HasCode:            property.HasCode,
		HasIncompleteTasks: property.HasIncompleteTasks,
	}
}

func convertLocationToProto(location *model.MemoPayloadLocation) *modelpb.Location {
	if location == nil {
		return nil
	}
	return &modelpb.Location{
		Placeholder: location.Placeholder,
		Latitude:    location.Latitude,
		Longitude:   location.Longitude,
	}
}
