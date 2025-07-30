package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateSession(ctx context.Context, m *model.Session) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateSessions(ctx context.Context, filter *model.FindSessionFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Session{}, filter, update)
}

func (d *Dao) CountSessions(ctx context.Context, f *model.FindSessionFilter) (total int64, err error) {
	return db.Count(ctx, &model.Session{}, f)
}

func (d *Dao) FindSessions(ctx context.Context, f *model.FindSessionFilter) (ms []*model.Session, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindSession(ctx context.Context, f *model.FindSessionFilter) (*model.Session, error) {
	var m model.Session
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) FindSessionByID(ctx context.Context, id int64) (m *model.Session, err error) {
	return d.FindSession(ctx, &model.FindSessionFilter{ID: db.Eq(id)})
}

func (d *Dao) DeleteSessions(ctx context.Context, filter *model.FindSessionFilter) error {
	_, err := db.Delete(ctx, &model.Session{}, filter)
	return err
}
