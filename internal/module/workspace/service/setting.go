package service

import (
	"context"

	"github.com/smartmemos/memos/internal/module/workspace/model"
)

func (s *Service) GetSetting(ctx context.Context, req *model.GetSettingRequest) (setting *model.Setting, err error) {
	setting, err = s.dao.FindSetting(ctx, &model.FindSettingFilter{
		Name: req.Name,
	})
	return
}
