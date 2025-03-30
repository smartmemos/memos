package dao

import (
	"context"

	"github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateAccessToken(ctx context.Context, m *model.AccessToken) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) DeleteAccessToken(ctx context.Context, filter *model.FindAccessTokenFilter) error {
	_, err := db.Delete(ctx, &model.AccessToken{}, filter)
	return err
}

func (d *Dao) FindAccessToken(ctx context.Context, filter *model.FindAccessTokenFilter) (*model.AccessToken, error) {
	var m model.AccessToken
	if err := db.FindOne(ctx, filter, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) CountAccessTokens(ctx context.Context, filter *model.FindAccessTokenFilter) (total int64, err error) {
	return db.Count(ctx, &model.AccessToken{}, filter)
}
