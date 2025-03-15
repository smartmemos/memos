package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/module/auth/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateAccessToken(ctx context.Context, m *model.AccessToken) error {
	return db.GetDB(ctx).Create(m).Error
}
