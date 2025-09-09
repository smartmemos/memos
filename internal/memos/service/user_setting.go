package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (s *Service) GetUserSetting(ctx context.Context, req *model.GetUserSettingRequest) (*model.UserSetting, error) {
	filter := &model.UserSettingFilter{}
	setting, err := s.dao.FindUserSetting(ctx, filter)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (s *Service) GetUserSettings(ctx context.Context, req *model.GetUserSettingsRequest) ([]*model.UserSetting, error) {
	filter := &model.UserSettingFilter{}
	settings, err := s.dao.FindUserSettings(ctx, filter)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (s *Service) UpdateUserSetting(ctx context.Context, req *model.UpdateUserSettingRequest) (setting *model.UserSetting, err error) {
	filter := &model.UserSettingFilter{
		UserID: db.Eq(req.UserID),
		Key:    db.Eq(req.Key),
	}
	setting, err = s.dao.FindUserSetting(ctx, filter)
	if err != nil {
		return
	}
	switch req.Key {
	case model.UserSettingKeyGeneral:
		if setting.Value.GeneralUserSetting == nil {
			setting.Value.GeneralUserSetting = &model.GeneralUserSetting{}
		}
		if lo.Contains(req.UpdateMask, "locale") {
			setting.Value.GeneralUserSetting.Locale = req.Value.GeneralUserSetting.Locale
		}
		if lo.Contains(req.UpdateMask, "appearance") {
			setting.Value.GeneralUserSetting.Appearance = req.Value.GeneralUserSetting.Appearance
		}
		if lo.Contains(req.UpdateMask, "memo_visibility") {
			setting.Value.GeneralUserSetting.MemoVisibility = req.Value.GeneralUserSetting.MemoVisibility
		}
		if lo.Contains(req.UpdateMask, "theme") {
			setting.Value.GeneralUserSetting.Theme = req.Value.GeneralUserSetting.Theme
		}
	case model.UserSettingKeySessions:
		setting.Value.SessionsUserSetting = req.Value.SessionsUserSetting
	}
	valueBytes, err := json.Marshal(setting.Value)
	if err != nil {
		return
	}
	if _, err = s.dao.UpdateUserSettings(ctx, filter, map[string]any{"value": valueBytes}); err != nil {
		return
	}
	return setting, nil
}

func (s *Service) GetUserSessions(ctx context.Context, req *model.GetUserSessionsRequest) ([]*model.UserSession, error) {
	filter := &model.UserSettingFilter{
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
	filter := &model.UserSettingFilter{
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
	valueBytes, err := json.Marshal(model.UserSettingValue{
		SessionsUserSetting: &model.SessionsUserSetting{Sessions: sessions},
	})
	if err != nil {
		return err
	}
	if _, err = s.dao.UpdateUserSettings(ctx, filter, map[string]any{"value": string(valueBytes)}); err != nil {
		return err
	}
	return
}
