package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateAttachment(ctx context.Context, m *model.Attachment) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateAttachments(ctx context.Context, filter *model.AttachmentFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Attachment{}, filter, update)
}

func (d *Dao) UpdateAttachment(ctx context.Context, m *model.Attachment, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) CountAttachments(ctx context.Context, f *model.AttachmentFilter) (total int64, err error) {
	return db.Count(ctx, &model.Attachment{}, f)
}

func (d *Dao) FindAttachments(ctx context.Context, f *model.AttachmentFilter) (ms []*model.Attachment, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindAttachment(ctx context.Context, f *model.AttachmentFilter) (*model.Attachment, error) {
	var m model.Attachment
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) FindAttachmentByID(ctx context.Context, id int64) (m *model.Attachment, err error) {
	return d.FindAttachment(ctx, &model.AttachmentFilter{ID: db.Eq(id)})
}

func (d *Dao) DeleteAttachments(ctx context.Context, filter *model.AttachmentFilter) error {
	_, err := db.Delete(ctx, &model.Attachment{}, filter)
	return err
}
