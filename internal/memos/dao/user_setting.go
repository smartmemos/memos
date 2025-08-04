package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateUserSetting(ctx context.Context, m *model.UserSetting) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateUserSettings(ctx context.Context, filter *model.FindUserSettingFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.UserSetting{}, filter, update)
}

func (d *Dao) UpdateUserSetting(ctx context.Context, m *model.UserSetting, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) FindUserSettings(ctx context.Context, f *model.FindUserSettingFilter) (ms []*model.UserSetting, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindUserSetting(ctx context.Context, f *model.FindUserSettingFilter) (*model.UserSetting, error) {
	var m model.UserSetting
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) DeleteUserSettings(ctx context.Context, filter *model.FindUserSettingFilter) error {
	_, err := db.Delete(ctx, &model.UserSetting{}, filter)
	return err
}
