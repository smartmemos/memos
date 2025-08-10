package model

import (
	"time"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type UserSetting struct {
	UserID int64
	Key    string
	Value  UserSettingValue `gorm:"serializer:json"`
}

func (UserSetting) TableName() string {
	return TableUserSetting
}

type UserSettingKey string

const (
	UserSettingKeyGeneral      UserSettingKey = "GENERAL"
	UserSettingKeySessions     UserSettingKey = "SESSIONS"
	UserSettingKeyAccessTokens UserSettingKey = "ACCESS_TOKENS"
	UserSettingKeyWebhooks     UserSettingKey = "WEBHOOKS"
)

type UserSettingValue struct {
	*GeneralUserSetting
	*SessionsUserSetting
	*AccessTokensUserSetting
	*WebhooksUserSetting
}

type GeneralUserSetting struct {
	Locale         string
	Appearance     string
	MemoVisibility string
	Theme          string
}

type SessionsUserSetting struct {
	Sessions []*UserSession
}

type UserSession struct {
	Name             string
	SessionID        string
	CreateTime       time.Time
	LastAccessedTime time.Time
	ClientInfo       ClientInfo
}

type ClientInfo struct {
	UserAgent  string
	IPAddress  string
	DeviceType string
	OS         string
	Browser    string
}

type AccessTokensUserSetting struct {
	AccessTokens []*UserAccessToken
}

type UserAccessToken struct {
	Name        string
	AccessToken string
	Description string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}

type WebhooksUserSetting struct {
	Webhooks []*UserWebhook
}

type UserWebhook struct {
	Name        string
	URL         string
	DisplayName string
	CreateTime  time.Time
	UpdateTime  time.Time
}

type FindUserSettingFilter struct {
	db.BaseFilter

	UserID db.F[int64]
	Key    db.F[UserSettingKey]
}

type GetUserSettingRequest struct {
}

type GetUserSettingsRequest struct {
}
