package service

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/smartmemos/memos/internal/module/workspace/model"
)

func (s *Service) GetSetting(ctx context.Context, key model.SettingKey, value any) (err error) {
	setting, err := s.dao.FindSetting(ctx, &model.FindSettingFilter{Name: string(key)})
	if err != nil {
		return
	}
	if err = json.Unmarshal(setting.Value.RawMessage, value); err != nil {
		logrus.WithContext(ctx).Errorf("解析json失败, err:%v", err)
		return
	}
	return
}
