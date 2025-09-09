package service

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (s *Service) GetGeneralSetting(ctx context.Context) (*model.GeneralSetting, error) {
	filter := &model.SystemSettingFilter{
		Name: db.Eq("general"),
	}
	setting, err := s.dao.FindSystemSetting(ctx, filter)
	if err != nil {
		return nil, err
	}
	return setting.Value.GeneralSetting, nil
}

func (s *Service) GetStorageSetting(ctx context.Context) (*model.StorageSetting, error) {
	filter := &model.SystemSettingFilter{
		Name: db.Eq("storage"),
	}
	setting, err := s.dao.FindSystemSetting(ctx, filter)
	if err != nil {
		return nil, err
	}
	return setting.Value.StorageSetting, nil
}

func (s *Service) GetMemoRelatedSetting(ctx context.Context) (*model.MemoRelatedSetting, error) {
	filter := &model.SystemSettingFilter{
		Name: db.Eq("memo_related"),
	}
	setting, err := s.dao.FindSystemSetting(ctx, filter)
	if err != nil {
		return nil, err
	}
	return setting.Value.MemoRelatedSetting, nil
}
