package model

import (
	"github.com/smartmemos/memos/internal/pkg/db"
)

type Reaction struct {
	db.Model

	Type      string
	CreatorID int32
	ContentID string
}

func (Reaction) TableName() string {
	return TableMemo
}
