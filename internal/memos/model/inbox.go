package model

import "github.com/smartmemos/memos/internal/pkg/db"

// Inbox 收件箱
type Inbox struct {
	db.Model

	SenderID   int64
	ReceiverID int64
	ActivityID int64
	Status     string
	Message    string
}

func (Inbox) TableName() string {
	return TableInbox
}

type InboxFilter struct {
	db.Query

	ID         db.F[int64]
	SenderID   db.F[int64]
	ReceiverID db.F[int64]
	Status     db.F[string]
}

type ListInboxesRequest struct {
	db.Query
}
