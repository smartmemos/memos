package service

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
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

func (s *Service) ListMemoRelations(ctx context.Context, req *model.ListMemoRelationsRequest) (total int64, relations []*model.MemoRelation, err error) {
	filter := &model.MemoRelationFilter{
		MemoIDs: db.In(req.MemoIDs),
		Query:   req.Query,
	}
	total, err = s.dao.CountMemoRelations(ctx, filter)
	if err != nil {
		return
	}
	if !filter.HasNextPage(total) {
		return
	}
	relations, err = s.dao.FindMemoRelations(ctx, filter)
	return
}
