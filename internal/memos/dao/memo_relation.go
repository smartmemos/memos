package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateMemoRelation(ctx context.Context, m *model.MemoRelation) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateMemoRelations(ctx context.Context, filter *model.FindMemoRelationFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Memo{}, filter, update)
}

func (d *Dao) UpdateMemoRelation(ctx context.Context, m *model.MemoRelation, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) error {
	return db.GetDB(ctx).Where(model.MemoRelation{MemoID: m.MemoID, RelatedMemoID: m.RelatedMemoID, Type: m.Type}).FirstOrCreate(m).Error
}

func (d *Dao) CountMemoRelations(ctx context.Context, f *model.FindMemoRelationFilter) (total int64, err error) {
	return db.Count(ctx, &model.MemoRelation{}, f)
}

func (d *Dao) FindMemoRelations(ctx context.Context, f *model.FindMemoRelationFilter) (ms []*model.MemoRelation, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindMemoRelation(ctx context.Context, f *model.FindMemoRelationFilter) (*model.MemoRelation, error) {
	var m model.MemoRelation
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) DeleteMemoRelations(ctx context.Context, filter *model.FindMemoRelationFilter) error {
	_, err := db.Delete(ctx, &model.MemoRelation{}, filter)
	return err
}
