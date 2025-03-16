package model

import (
	"strings"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type Setting struct {
	db.Model

	Name        string
	Value       string
	Description string
}

func (Setting) TableName() string {
	return TableSetting
}

type FindSettingFilter struct {
	db.BaseFilter
	Name string
}

func (f FindSettingFilter) GetQuery() (query string, args []any) {
	var where []string
	if f.Name != "" {
		where = append(where, "name=?")
		args = append(args, f.Name)
	}
	query = strings.Join(where, " and ")
	return
}

type GetSettingRequest struct {
	Name string
}
