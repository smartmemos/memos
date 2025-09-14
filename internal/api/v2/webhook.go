package v2

import (
	"context"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartmemos/memos/internal/memos"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type WebhookService struct {
	v2pb.UnimplementedWebhookServiceHandler
	memosService memos.Service
}

func NewWebhookService(i do.Injector) (*WebhookService, error) {
	return &WebhookService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

func (s *WebhookService) ListWebhooks(ctx context.Context, request *connect.Request[v2pb.ListWebhooksRequest]) (response *connect.Response[v2pb.ListWebhooksResponse], err error) {
	return nil, nil
}

func (s *WebhookService) GetWebhook(ctx context.Context, request *connect.Request[v2pb.GetWebhookRequest]) (response *connect.Response[modelpb.Webhook], err error) {
	return nil, nil
}

func (s *WebhookService) CreateWebhook(ctx context.Context, request *connect.Request[v2pb.CreateWebhookRequest]) (response *connect.Response[modelpb.Webhook], err error) {
	return nil, nil
}

func (s *WebhookService) UpdateWebhook(ctx context.Context, request *connect.Request[v2pb.UpdateWebhookRequest]) (response *connect.Response[modelpb.Webhook], err error) {
	return nil, nil
}

func (s *WebhookService) DeleteWebhook(ctx context.Context, request *connect.Request[v2pb.DeleteWebhookRequest]) (response *connect.Response[emptypb.Empty], err error) {
	return nil, nil
}
