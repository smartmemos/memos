package model

import (
	"github.com/smartmemos/memos/internal/pkg/db"
)

type Resource struct {
	db.Model

	UID         string
	MemoID      int64
	CreatorID   int32
	Filename    string
	Blob        []byte
	Type        string
	Size        int64
	StorageType ResStorageType
	Reference   string
	Payload     *MemoPayload `gorm:"serializer:json"`
}

func (Resource) TableName() string {
	return TableResource
}
