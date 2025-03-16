package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/module/workspace/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateSetting(ctx context.Context, m *model.Setting) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateSettings(ctx context.Context, filter *model.FindSettingFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.Setting{}, filter, update)
}

func (d *Dao) UpdateSetting(ctx context.Context, m *model.Setting, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) CountSettings(ctx context.Context, f *model.FindSettingFilter) (total int64, err error) {
	return db.Count(ctx, &model.Setting{}, f)
}

func (d *Dao) FindSettings(ctx context.Context, f *model.FindSettingFilter) (ms []*model.Setting, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindSetting(ctx context.Context, f *model.FindSettingFilter) (*model.Setting, error) {
	var m model.Setting
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}
