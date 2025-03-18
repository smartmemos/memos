package model

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Stats struct {
	Name                  string                   `json:"name,"`
	MemoDisplayTimestamps []*timestamppb.Timestamp `json:"memo_display_timestamps"`
	MemoTypeStats         MemoTypeStats            `json:"memo_type_stats"`
	TagCount              map[string]int32         `json:"tag_count"`
	PinnedMemos           []string                 `json:"pinned_memos"`
	TotalMemoCount        int32                    `json:"total_memo_count"`
}

type MemoTypeStats struct {
	LinkCount int32 `json:"link_count"`
	CodeCount int32 `json:"code_count"`
	TodoCount int32 `json:"todo_count"`
	UndoCount int32 `json:"undo_count"`
}

type ListAllUserStatsRequest struct {
}
