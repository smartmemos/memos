package dao

import (
	"context"
	"errors"

	"github.com/smartmemos/memos/internal/module/workspace/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateProfile(ctx context.Context, m *model.Profile) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateProfiles(ctx context.Context, filter *model.FindProfileFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Profile{}, filter, update)
}

func (d *Dao) UpdateProfile(ctx context.Context, m *model.Profile, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) CountProfiles(ctx context.Context, f *model.FindProfileFilter) (total int64, err error) {
	return db.Count(ctx, &model.Profile{}, f)
}

func (d *Dao) FindProfiles(ctx context.Context, f *model.FindProfileFilter) (ms []*model.Profile, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindProfile(ctx context.Context, f *model.FindProfileFilter) (*model.Profile, error) {
	var m model.Profile
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) FindProfileByID(ctx context.Context, id int64) (m *model.Profile, err error) {
	if id <= 0 {
		err = errors.New("id必须大于0")
		return
	}
	return d.FindProfile(ctx, &model.FindProfileFilter{ID: id})
}
