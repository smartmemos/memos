package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateInbox(ctx context.Context, m *model.Inbox) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateInboxes(ctx context.Context, filter *model.InboxFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Inbox{}, filter, update)
}

func (d *Dao) UpdateInbox(ctx context.Context, m *model.Inbox, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) CountInboxes(ctx context.Context, f *model.InboxFilter) (total int64, err error) {
	return db.Count(ctx, &model.Inbox{}, f)
}

func (d *Dao) FindInboxes(ctx context.Context, f *model.InboxFilter) (ms []*model.Inbox, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindInbox(ctx context.Context, f *model.InboxFilter) (*model.Inbox, error) {
	var m model.Inbox
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) FindInboxByID(ctx context.Context, id int64) (m *model.Inbox, err error) {
	return d.FindInbox(ctx, &model.InboxFilter{ID: db.Eq(id)})
}

func (d *Dao) DeleteInboxes(ctx context.Context, filter *model.InboxFilter) error {
	_, err := db.Delete(ctx, &model.Inbox{}, filter)
	return err
}
