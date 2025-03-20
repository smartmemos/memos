package model

import (
	"strings"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type Memo struct {
	db.Model

	UID        string
	CreatorID  int64
	Content    string
	Payload    MemoPayload `gorm:"serializer:json"`
	Tags       []string    `gorm:"serializer:json"`
	Pinned     bool
	Visibility Visibility
	Status     string
}

func (Memo) TableName() string {
	return TableMemo
}

type FindMemoFilter struct {
	db.BaseFilter

	ID              int64
	Pid             int64
	CreatorID       int64
	ExcludeComments bool
	ExcludeContent  bool
	Status          string

	VisibilityList []Visibility
}

func (f FindMemoFilter) GetQuery() (query string, args []any) {
	var where []string
	if f.ID > 0 {
		where = append(where, "id=?")
		args = append(args, f.ID)
	}
	if f.Pid > 0 {
		where = append(where, "pid=?")
		args = append(args, f.Pid)
	}
	query = strings.Join(where, " and ")
	return
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

type CreateMemoRequest struct {
	Content    string
	Visibility Visibility
	// Resource
}

type UpdateMemoRequest struct {
}

type ListMemosRequest struct {
}

type GetMemosRequest struct {
	Page     int
	PageSize int
}
