package dao

import (
	"context"
	"errors"

	"github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (d *Dao) CreateUser(ctx context.Context, m *model.User) error {
	return db.GetDB(ctx).Create(m).Error
}

func (d *Dao) UpdateUsers(ctx context.Context, filter *model.FindUserFilter, update map[string]any) (int64, error) {
	return db.Updates(ctx, &model.User{}, filter, update)
}

func (d *Dao) UpdateUser(ctx context.Context, m *model.User, update map[string]any) error {
	return db.Update(ctx, m, update)
}

func (d *Dao) CountUsers(ctx context.Context, f *model.FindUserFilter) (total int64, err error) {
	return db.Count(ctx, &model.User{}, f)
}

func (d *Dao) FindUsers(ctx context.Context, f *model.FindUserFilter) (ms []*model.User, err error) {
	err = db.Find(ctx, f, &ms)
	return
}

func (d *Dao) FindUser(ctx context.Context, f *model.FindUserFilter) (*model.User, error) {
	var m model.User
	if err := db.FindOne(ctx, f, &m); err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (d *Dao) FindUserByID(ctx context.Context, id int64) (m *model.User, err error) {
	if id <= 0 {
		err = errors.New("id必须大于0")
		return
	}
	return d.FindUser(ctx, &model.FindUserFilter{ID: id})
}
