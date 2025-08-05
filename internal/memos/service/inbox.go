package service

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

func (s *Service) ListInboxes(ctx context.Context, req *model.ListInboxesRequest) (total int64, list []*model.Inbox, err error) {
	filter := &model.FindInboxFilter{}

	total, err = s.dao.CountInboxes(ctx, filter)
	if err != nil {
		return
	}
	list, err = s.dao.FindInboxes(ctx, filter)
	if err != nil {
		return
	}
	return
}
