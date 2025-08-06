package model

import "github.com/smartmemos/memos/internal/pkg/db"

// Memo 笔记
type Memo struct {
	db.Model

	UID          string
	ParentID     int64
	RelationType RelationType
	CreatorID    int64
	Content      string
	Payload      *MemoPayload `gorm:"serializer:json"`
	Tags         []string     `gorm:"serializer:json"`
	Pinned       bool
	Visibility   Visibility
	RowStatus    string
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
	db.BaseFilter

	ID              db.F[int64]
	Pid             db.F[int64]
	ParentIDs       db.F[[]int64]
	CreatorID       db.F[int64]
	ExcludeComments db.F[bool]
	ExcludeContent  db.F[bool]
	Status          db.F[string]

	VisibilityList db.F[[]Visibility]
}

type CreateMemoRequest struct {
	ParentID     int64
	RelationType RelationType
	CreatorID    int64
	Content      string
	Payload      MemoPayload `gorm:"serializer:json"`
}

type ListMemosRequest struct {
	Status          RowStatus
	VisibilityList  []Visibility
	ExcludeContent  bool
	ExcludeComments bool
}

type GetMemoRequest struct {
	ID  int64
	UID string
}
