package service

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

func (s *Service) GetUserSetting(ctx context.Context, req *model.GetUserSettingRequest) (*model.UserSetting, error) {
	filter := &model.FindUserSettingFilter{}
	setting, err := s.dao.FindUserSetting(ctx, filter)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (s *Service) GetUserSettings(ctx context.Context, req *model.GetUserSettingsRequest) ([]*model.UserSetting, error) {
	filter := &model.FindUserSettingFilter{}
	settings, err := s.dao.FindUserSettings(ctx, filter)
	if err != nil {
		return nil, err
	}
	return settings, nil
}
