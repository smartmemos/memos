package model

import (
	"github.com/smartmemos/memos/internal/pkg/db"
)

// Memo 笔记
type Memo struct {
	db.Model

	UID        string
	ParentID   int64
	CreatorID  int64
	Content    string
	Payload    *MemoPayload `gorm:"serializer:json"`
	Pinned     bool
	Visibility Visibility
	RowStatus  RowStatus
}

func (Memo) TableName() string {
	return TableMemo
}

type MemoPayload struct {
	Property *MemoPayloadProperty `json:"property"`
	Location *MemoPayloadLocation `json:"location"`
	Tags     []string             `json:"tags"`
}

type MemoPayloadLocation struct {
	Placeholder string  `json:"placeholder"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type MemoPayloadProperty struct {
	HasLink            bool     `json:"has_link"`
	HasTaskList        bool     `json:"has_task_list"`
	HasCode            bool     `json:"has_code"`
	HasIncompleteTasks bool     `json:"has_incomplete_tasks"`
	References         []string `json:"references"`
}

// MemoFilter is the filter for the memo.
type MemoFilter struct {
	db.Query

	ID              db.F[int64]
	IDs             db.F[[]int64] `gorm:"id"`
	UID             db.F[string]
	Pid             db.F[int64]
	ParentIDs       db.F[[]int64]
	CreatorID       db.F[int64]
	ExcludeComments db.F[bool]
	ExcludeContent  db.F[bool]
	RowStatus       db.F[RowStatus]
	VisibilityList  db.F[[]Visibility]
}

// MemoRequest is the request for the memo.
type MemoRequest struct {
	db.Query

	ID              int64
	IDs             []int64
	UID             string
	Status          RowStatus
	VisibilityList  []Visibility
	ExcludeContent  bool
	ExcludeComments bool
}

func (req *MemoRequest) ToFilter() *MemoFilter {
	filter := &MemoFilter{Query: req.Query}
	if req.ID != 0 {
		filter.ID = db.Eq(req.ID)
	}
	if len(req.IDs) > 0 {
		filter.IDs = db.In(req.IDs)
	}
	if req.Status != "" {
		filter.RowStatus = db.Eq(req.Status)
	}
	if req.UID != "" {
		filter.UID = db.Eq(req.UID)
	}
	if len(req.VisibilityList) > 0 {
		filter.VisibilityList = db.In(req.VisibilityList)
	}
	if req.ExcludeContent {
		filter.ExcludeContent = db.Eq(true)
	}
	if req.ExcludeComments {
		filter.ExcludeComments = db.Eq(true)
	}
	return filter
}

type CreateMemoRequest struct {
	Memo     *Memo
	Location *MemoPayloadLocation
}

type UpdateMemoRequest struct {
	UpdateMask []string
	Memo       *Memo
	Location   *MemoPayloadLocation
}
