package service

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

// UpsertMemoRelation upserts a memo relation.
func (s *Service) UpsertMemoRelation(ctx context.Context, req *model.UpsertMemoRelationRequest) (relation *model.MemoRelation, err error) {
	relation = &model.MemoRelation{
		MemoID:        req.MemoID,
		RelatedMemoID: req.RelatedMemoID,
		Type:          req.Type,
	}
	err = s.dao.UpsertMemoRelation(ctx, relation)
	return
}
