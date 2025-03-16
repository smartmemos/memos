package model

import "github.com/smartmemos/memos/internal/pkg/db"

type Setting struct {
	db.Model

	UserId int64
	Key    string
	Value  string
}

func (Setting) TableName() string {
	return TableUserSetting
}

type GetSettingRequest struct {
}
