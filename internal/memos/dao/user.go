package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateUser(ctx context.Context, m *model.User) error {
	return db.GetDB(ctx).Create(m).Error
}
