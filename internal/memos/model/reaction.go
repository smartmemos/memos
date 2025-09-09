package model

import (
	"github.com/smartmemos/memos/internal/pkg/db"
)

type Reaction struct {
	db.Model

	CreatorID    int32
	ContentID    string
	ReactionType string
}

func (Reaction) TableName() string {
	return TableReaction
}

type ReactionFilter struct {
	db.Query

	ID           db.F[int64]
	CreatorID    db.F[int32]
	ContentID    db.F[string]
	ContentIDs   db.F[[]string] `gorm:"content_id"`
	ReactionType db.F[string]
}

type UpsertReactionRequest struct {
	CreatorID    int32
	ContentID    string
	ReactionType string
}

type DeleteReactionRequest struct {
	ID int64
}

type ListReactionsRequest struct {
	db.Query

	ContentID  string
	ContentIDs []string
}
