package model

import (
	"encoding/json"
	"strings"

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

	UserID int64
}

func (f FindSettingFilter) GetQuery() (query string, args []any) {
	var where []string
	if f.UserID > 0 {
		where = append(where, "user_id=?")
		args = append(args, f.UserID)
	}
	query = strings.Join(where, " and ")
	return
}

type GetSettingRequest struct {
}
