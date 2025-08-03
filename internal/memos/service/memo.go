package service

import (
	"context"

	"github.com/samber/lo"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (s *Service) CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (memo *model.Memo, err error) {

	return
}

func (s *Service) ListMemos(ctx context.Context, req *model.ListMemosRequest) (total int64, list []*model.Memo, err error) {
	filter := &model.FindMemoFilter{}

	total, err = s.dao.CountMemos(ctx, filter)
	if err != nil {
		return
	}

	memos, err := s.dao.FindMemos(ctx, filter)
	if err != nil {
		return
	}
	pids := lo.FlatMap(memos, func(item *model.Memo, _ int) []int64 {
		return lo.If(item.ParentID > 0, []int64{item.ParentID}).Else(nil)
	})
	var memosMap map[int64]*model.Memo
	if len(pids) > 0 {
		var parentList []*model.Memo
		parentList, err = s.dao.FindMemos(ctx, &model.FindMemoFilter{ParentIDs: db.In(pids)})
		if err != nil {
			return
		}
		memosMap = lo.Associate(parentList, func(item *model.Memo) (int64, *model.Memo) {
			return item.ID, item
		})
	}

	for _, memo := range memos {
		if v, ok := memosMap[memo.ID]; ok {
			memo.ParentID = v.ID
		}
		list = append(list, memo)
	}
	return
}

func (s *Service) GetMemo(ctx context.Context, req *model.GetMemoRequest) (memo *model.Memo, err error) {

	return
}
