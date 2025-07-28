package model

import (
	"encoding/json"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type Setting struct {
	db.Model

	UserID int64
	Key    string
	Value  string `gorm:"serializer:json"`
}

func (Setting) TableName() string {
	return TableUserSetting
}

type SettingValue struct {
	json.RawMessage
}

type FindSettingFilter struct {
	db.BaseFilter

	UserID db.F[int64]
}

type GetSettingRequest struct {
}
