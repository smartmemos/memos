package v2

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/model"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type InboxService struct {
	v2pb.UnimplementedInboxServiceHandler
	memosService memos.Service
}

func NewInboxService(i do.Injector) (*InboxService, error) {
	return &InboxService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

func (s *InboxService) ListInboxes(ctx context.Context, request *connect.Request[v2pb.ListInboxesRequest]) (response *connect.Response[v2pb.ListInboxesResponse], err error) {
	logrus.Info("req: ", request.Msg)

	total, inboxes, err := s.memosService.ListInboxes(ctx, &model.ListInboxesRequest{})
	if err != nil {
		return
	}

	response = connect.NewResponse(&v2pb.ListInboxesResponse{
		TotalSize:     int32(total),
		NextPageToken: "",
		Inboxes: lo.Map(inboxes, func(inbox *model.Inbox, _ int) *modelpb.Inbox {
			return convertInboxToProto(inbox)
		}),
	})
	return
}

func convertInboxToProto(inbox *model.Inbox) *modelpb.Inbox {
	info := &modelpb.Inbox{
		Name:       fmt.Sprintf("inboxes/%d", inbox.ID),
		Sender:     fmt.Sprintf("users/%d", inbox.SenderID),
		Receiver:   fmt.Sprintf("users/%d", inbox.ReceiverID),
		Status:     modelpb.Inbox_Status(modelpb.Inbox_Status_value[inbox.Status]),
		CreateTime: timestamppb.New(inbox.CreatedAt),
		Type:       modelpb.Inbox_Type(modelpb.Inbox_Type_value[inbox.Message]),
	}
	if activityID := int32(inbox.ActivityID); activityID != 0 {
		info.ActivityId = &activityID
	}
	return info
}
