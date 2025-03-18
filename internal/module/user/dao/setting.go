package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

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
