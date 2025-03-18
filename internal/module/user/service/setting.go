package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
)

func (s *Service) GetSettings(ctx context.Context, req *model.GetSettingRequest) (setting *model.Setting, err error) {
	userID, err := grpc_util.GetUserID(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}
	settings, err := s.dao.FindSettings(ctx, &model.FindSettingFilter{UserID: userID})
	if err != nil {
		return
	}
	for _, setting := range settings {
		logrus.Info(setting)
	}
	return nil, nil
}
