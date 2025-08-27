//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package memos

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

type Service interface {
	// CreateUser creates a user.
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	// UpdateUser updates a user.
	UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (*model.User, error)
	// GetUserByID gets a user by ID.
	GetUserByID(ctx context.Context, id int64) (*model.User, error)

	// CreateSession creates a session.
	CreateSession(ctx context.Context, req *model.CreateSessionRequest) (*model.Session, error)

	// CreateMemo creates a memo.
	CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (*model.Memo, error)
	// ListMemos lists the memos.
	ListMemos(ctx context.Context, req *model.ListMemosRequest) (int64, []*model.Memo, error)
	// GetMemo gets a memo by ID.
	GetMemo(ctx context.Context, req *model.GetMemoRequest) (*model.Memo, error)
	// UpdateMemo updates a memo.
	UpdateMemo(ctx context.Context, req *model.UpdateMemoRequest) (*model.Memo, error)
	// DeleteMemo deletes a memo.
	DeleteMemo(ctx context.Context, req *model.DeleteMemoRequest) error

	// UpsertReaction upserts a memo reaction.
	UpsertReaction(ctx context.Context, req *model.UpsertReactionRequest) (*model.Reaction, error)
	// ListReactions lists the memo reactions.
	ListReactions(ctx context.Context, req *model.ListReactionsRequest) (int64, []*model.Reaction, error)
	// DeleteReaction deletes a memo reaction.
	DeleteReaction(ctx context.Context, req *model.DeleteReactionRequest) error

	// UpsertMemoRelation upserts a memo relation.
	UpsertMemoRelation(ctx context.Context, req *model.UpsertMemoRelationRequest) (*model.MemoRelation, error)

	// GetGeneralSetting gets the general setting.
	GetGeneralSetting(ctx context.Context) (*model.GeneralSetting, error)
	// GetStorageSetting gets the storage setting.
	GetStorageSetting(ctx context.Context) (*model.StorageSetting, error)
	// GetMemoRelatedSetting gets the memo related setting.
	GetMemoRelatedSetting(ctx context.Context) (*model.MemoRelatedSetting, error)

	// GetUserSetting gets the user setting.
	GetUserSetting(ctx context.Context, req *model.GetUserSettingRequest) (*model.UserSetting, error)
	// UpdateUserSetting updates the user setting.
	UpdateUserSetting(ctx context.Context, req *model.UpdateUserSettingRequest) (*model.UserSetting, error)
	// GetUserSettings gets the user settings.
	GetUserSettings(ctx context.Context, req *model.GetUserSettingsRequest) ([]*model.UserSetting, error)

	// GetUserSessions lists the user sessions.
	GetUserSessions(ctx context.Context, req *model.GetUserSessionsRequest) ([]*model.UserSession, error)
	// RevokeUserSession revokes a user session.
	RevokeUserSession(ctx context.Context, req *model.RevokeUserSessionRequest) error

	// ListInboxes lists the inboxes.
	ListInboxes(ctx context.Context, req *model.ListInboxesRequest) (int64, []*model.Inbox, error)
}
