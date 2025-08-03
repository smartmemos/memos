package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateMemo(ctx context.Context, m *model.Memo) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateMemos(ctx context.Context, filter *model.FindMemoFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Memo{}, filter, update)
}

func (d *Dao) UpdateMemo(ctx context.Context, m *model.Memo, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) CountMemos(ctx context.Context, f *model.FindMemoFilter) (total int64, err error) {
	return db.Count(ctx, &model.Memo{}, f)
}

func (d *Dao) FindMemos(ctx context.Context, f *model.FindMemoFilter) (ms []*model.Memo, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindMemo(ctx context.Context, f *model.FindMemoFilter) (*model.Memo, error) {
	var m model.Memo
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) FindMemoByID(ctx context.Context, id int64) (m *model.Memo, err error) {
	return d.FindMemo(ctx, &model.FindMemoFilter{ID: db.Eq(id)})
}

func (d *Dao) DeleteMemos(ctx context.Context, filter *model.FindMemoFilter) error {
	_, err := db.Delete(ctx, &model.Memo{}, filter)
	return err
}
