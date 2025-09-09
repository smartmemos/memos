package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateSystemSetting(ctx context.Context, m *model.SystemSetting) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateSystemSettings(ctx context.Context, filter *model.SystemSettingFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.SystemSetting{}, filter, update)
}

func (d *Dao) UpdateSystemSetting(ctx context.Context, m *model.SystemSetting, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) FindSystemSettings(ctx context.Context, f *model.SystemSettingFilter) (ms []*model.SystemSetting, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindSystemSetting(ctx context.Context, f *model.SystemSettingFilter) (*model.SystemSetting, error) {
	var m model.SystemSetting
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) DeleteSystemSettings(ctx context.Context, filter *model.SystemSettingFilter) error {
	_, err := db.Delete(ctx, &model.SystemSetting{}, filter)
	return err
}
