package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/smartmemos/memos/internal/module/workspace/model"
)

func (s *Service) GetSetting(ctx context.Context, key model.SettingKey, value any) (err error) {
	switch key {
	case model.SettingKeyBasic:
	case model.SettingKeyGeneral:
	case model.SettingKeyMemoRelated:
	case model.SettingKeyStorage:
	default:
		errors.Errorf("unknown type: %s", key)
		return
	}
	setting, err := s.dao.FindSetting(ctx, &model.FindSettingFilter{Name: string(key)})
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(setting.Value), value)
	return
}
