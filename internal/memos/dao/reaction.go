package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateReaction(ctx context.Context, m *model.Reaction) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateReactions(ctx context.Context, filter *model.FindReactionFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Reaction{}, filter, update)
}

func (d *Dao) UpdateReaction(ctx context.Context, m *model.Reaction, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) CountReactions(ctx context.Context, f *model.FindReactionFilter) (total int64, err error) {
	return db.Count(ctx, &model.Reaction{}, f)
}

func (d *Dao) FindReactions(ctx context.Context, f *model.FindReactionFilter) (ms []*model.Reaction, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindReaction(ctx context.Context, f *model.FindReactionFilter) (*model.Reaction, error) {
	var m model.Reaction
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) FindReactionByID(ctx context.Context, id int64) (m *model.Reaction, err error) {
	return d.FindReaction(ctx, &model.FindReactionFilter{ID: db.Eq(id)})
}

func (d *Dao) DeleteReactions(ctx context.Context, filter *model.FindReactionFilter) error {
	_, err := db.Delete(ctx, &model.Reaction{}, filter)
	return err
}
