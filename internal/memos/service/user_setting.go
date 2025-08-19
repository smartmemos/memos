package service

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
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

func (s *Service) GetUserSessions(ctx context.Context, req *model.GetUserSessionsRequest) ([]*model.UserSession, error) {
	filter := &model.FindUserSettingFilter{
		UserID: db.Eq(req.UserID),
		Key:    db.Eq(model.UserSettingKeySessions),
	}
	setting, err := s.dao.FindUserSetting(ctx, filter)
	if err != nil {
		return nil, err
	}
	return lo.Map(setting.Value.Sessions, func(session *model.UserSession, _ int) *model.UserSession {
		session.Name = fmt.Sprintf("%s%d/sessions/%s", model.UserNamePrefix, setting.UserID, session.SessionID)
		return session
	}), nil
}

func (s *Service) RevokeUserSession(ctx context.Context, req *model.RevokeUserSessionRequest) (err error) {
	filter := &model.FindUserSettingFilter{
		UserID: db.Eq(req.UserID),
		Key:    db.Eq(model.UserSettingKeySessions),
	}
	setting, err := s.dao.FindUserSetting(ctx, filter)
	if err != nil {
		return err
	}
	sessions := lo.Filter(setting.Value.Sessions, func(session *model.UserSession, _ int) bool {
		return session.SessionID != req.SessionID
	})
	if _, err = s.dao.UpdateUserSettings(ctx, filter, map[string]any{"value": sessions}); err != nil {
		return err
	}
	return
}
