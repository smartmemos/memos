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

type FindMemoRelationFilter struct {
	db.Query

	MemoID        db.F[int64]
	RelatedMemoID db.F[int64]
	Type          db.F[RelationType]
}

type UpsertMemoRelationRequest struct {
	MemoID        int64
	RelatedMemoID int64
	Type          RelationType
}
