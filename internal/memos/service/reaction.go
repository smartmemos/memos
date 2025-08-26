package service

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (s *Service) ListReactions(ctx context.Context, req *model.ListReactionsRequest) (total int64, list []*model.Reaction, err error) {
	filter := &model.FindReactionFilter{}
	if len(req.ContentIDs) > 0 {
		filter.ContentIDs = db.In(req.ContentIDs)
	}

	total, err = s.dao.CountReactions(ctx, filter)
	if err != nil {
		return
	}
	if !db.HasNextPage(total, req.Page, req.PageSize) {
		return
	}
	list, err = s.dao.FindReactions(ctx, filter)
	return
}

func (s *Service) UpsertReaction(ctx context.Context, req *model.UpsertReactionRequest) (reaction *model.Reaction, err error) {
	reaction, err = s.dao.FindReaction(ctx, &model.FindReactionFilter{
		CreatorID:    db.Eq(req.CreatorID),
		ContentID:    db.Eq(req.ContentID),
		ReactionType: db.Eq(req.ReactionType),
	})
	if db.IsDbError(err) {
		return
	}
	if db.IsRecordNotFound(err) {
		reaction = &model.Reaction{
			CreatorID:    req.CreatorID,
			ContentID:    req.ContentID,
			ReactionType: req.ReactionType,
		}
		err = s.dao.CreateReaction(ctx, reaction)
		return
	}
	return
}

func (s *Service) DeleteReaction(ctx context.Context, req *model.DeleteReactionRequest) (err error) {
	err = s.dao.DeleteReactions(ctx, &model.FindReactionFilter{
		ID: db.Eq(req.ID),
	})
	return
}
