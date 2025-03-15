package model

import (
	"strings"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type Memo struct {
	db.Model

	Pid        int
	Content    string
	Status     string
	Pinned     bool
	Visibility Visibility
}

func (Memo) TableName() string {
	return TableMemo
}

type FindMemoFilter struct {
	db.BaseFilter

	ID  int64
	Pid int64
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

type CreateMemoRequest struct {
}

type UpdateMemoRequest struct {
}
