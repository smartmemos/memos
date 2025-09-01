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

type FindMemoFilter struct {
	db.Query

	ID              db.F[int64]
	UID             db.F[string]
	Pid             db.F[int64]
	ParentIDs       db.F[[]int64]
	CreatorID       db.F[int64]
	ExcludeComments db.F[bool]
	ExcludeContent  db.F[bool]
	Status          db.F[string]

	VisibilityList db.F[[]Visibility]
}

type CreateMemoRequest struct {
	UserID       int64
	ParentID     int64
	RelationType RelationType
	Visibility   Visibility
	Content      string
	RowStatus    RowStatus
	Location     *MemoPayloadLocation
}

type ListMemosRequest struct {
	db.Query

	IDs             []int64
	Status          RowStatus
	VisibilityList  []Visibility
	ExcludeContent  bool
	ExcludeComments bool
}

type UpdateMemoRequest struct {
	UpdateMask   []string
	ID           int64
	UID          string
	Pinned       bool
	UserID       int64
	ParentID     int64
	RelationType RelationType
	Visibility   Visibility
	Content      string
	RowStatus    RowStatus
	MemoPayload  *MemoPayload
}

type GetMemoRequest struct {
	ID  int64
	UID string
}

type DeleteMemoRequest struct {
	ID  int64
	UID string
}
