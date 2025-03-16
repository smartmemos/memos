package model

import "github.com/smartmemos/memos/internal/pkg/db"

type Setting struct {
	db.Model

	Name        string
	Value       string
	Description string
}

func (Setting) TableName() string {
	return TableSetting
}

type GetSettingRequest struct {
}
