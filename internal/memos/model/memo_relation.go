package model

import "github.com/smartmemos/memos/internal/pkg/db"

// MemoRelation 笔记关系
type MemoRelation struct {
	MemoID        int64
	RelatedMemoID int64
	Type          RelationType
}

func (MemoRelation) TableName() string {
	return TableMemoRelation
}

type MemoRelationFilter struct {
	db.Query

	MemoID        db.F[int64]
	MemoIDs       db.F[[]int64] `gorm:"column:memo_id"`
	RelatedMemoID db.F[int64]
	Type          db.F[RelationType]
}

type ListMemoRelationsRequest struct {
	db.Query

	MemoID  int64
	MemoIDs []int64
}

type UpsertMemoRelationRequest struct {
	MemoID        int64
	RelatedMemoID int64
	Type          RelationType
}
