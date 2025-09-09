//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memos

import (
	"context"

	"github.com/smartmemos/memos/internal/memos/model"
)

type DAO interface {
	// User
	CreateUser(ctx context.Context, m *model.User) error
	UpdateUsers(ctx context.Context, filter *model.FindUserFilter, update map[string]any) (int64, error)
	UpdateUser(ctx context.Context, m *model.User, update map[string]any) error
	CountUsers(ctx context.Context, filter *model.FindUserFilter) (int64, error)
	FindUsers(ctx context.Context, filter *model.FindUserFilter) ([]*model.User, error)
	FindUserByID(ctx context.Context, id int64) (*model.User, error)
	FindUser(ctx context.Context, filter *model.FindUserFilter) (*model.User, error)
	DeleteUsers(ctx context.Context, filter *model.FindUserFilter) error

	// Inbox
	CreateInbox(ctx context.Context, m *model.Inbox) error
	UpdateInboxes(ctx context.Context, filter *model.FindInboxFilter, update map[string]any) (int64, error)
	UpdateInbox(ctx context.Context, m *model.Inbox, update map[string]any) error
	CountInboxes(ctx context.Context, filter *model.FindInboxFilter) (int64, error)
	FindInboxes(ctx context.Context, filter *model.FindInboxFilter) ([]*model.Inbox, error)
	FindInboxByID(ctx context.Context, id int64) (*model.Inbox, error)
	FindInbox(ctx context.Context, filter *model.FindInboxFilter) (*model.Inbox, error)
	DeleteInboxes(ctx context.Context, filter *model.FindInboxFilter) error

	// Session
	CreateSession(ctx context.Context, m *model.Session) error

	// Memo
	CreateMemo(ctx context.Context, m *model.Memo) error
	CountMemos(ctx context.Context, filter *model.MemoFilter) (int64, error)
	FindMemos(ctx context.Context, filter *model.MemoFilter) ([]*model.Memo, error)
	FindMemoByID(ctx context.Context, id int64) (*model.Memo, error)
	FindMemo(ctx context.Context, filter *model.MemoFilter) (*model.Memo, error)
	DeleteMemos(ctx context.Context, filter *model.MemoFilter) error
	UpdateMemo(ctx context.Context, m *model.Memo, update map[string]any) error

	// MemoRelation
	CreateMemoRelation(ctx context.Context, m *model.MemoRelation) error
	CountMemoRelations(ctx context.Context, filter *model.FindMemoRelationFilter) (int64, error)
	FindMemoRelations(ctx context.Context, filter *model.FindMemoRelationFilter) ([]*model.MemoRelation, error)
	FindMemoRelation(ctx context.Context, filter *model.FindMemoRelationFilter) (*model.MemoRelation, error)
	DeleteMemoRelations(ctx context.Context, filter *model.FindMemoRelationFilter) error
	UpdateMemoRelation(ctx context.Context, m *model.MemoRelation, update map[string]any) error
	UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) error

	// Reaction
	CreateReaction(ctx context.Context, m *model.Reaction) error
	CountReactions(ctx context.Context, filter *model.FindReactionFilter) (int64, error)
	FindReactions(ctx context.Context, filter *model.FindReactionFilter) ([]*model.Reaction, error)
	FindReactionByID(ctx context.Context, id int64) (*model.Reaction, error)
	FindReaction(ctx context.Context, filter *model.FindReactionFilter) (*model.Reaction, error)
	DeleteReactions(ctx context.Context, filter *model.FindReactionFilter) error
	UpdateReaction(ctx context.Context, m *model.Reaction, update map[string]any) error

	// Attachment
	CreateAttachment(ctx context.Context, m *model.Attachment) error
	CountAttachments(ctx context.Context, filter *model.FindAttachmentFilter) (int64, error)
	FindAttachments(ctx context.Context, filter *model.FindAttachmentFilter) ([]*model.Attachment, error)
	FindAttachmentByID(ctx context.Context, id int64) (*model.Attachment, error)
	FindAttachment(ctx context.Context, filter *model.FindAttachmentFilter) (*model.Attachment, error)
	DeleteAttachments(ctx context.Context, filter *model.FindAttachmentFilter) error

	// SystemSetting
	CreateSystemSetting(ctx context.Context, m *model.SystemSetting) error
	UpdateSystemSettings(ctx context.Context, filter *model.FindSystemSettingFilter, update map[string]any) (int64, error)
	UpdateSystemSetting(ctx context.Context, m *model.SystemSetting, update map[string]any) error
	FindSystemSettings(ctx context.Context, filter *model.FindSystemSettingFilter) ([]*model.SystemSetting, error)
	FindSystemSetting(ctx context.Context, filter *model.FindSystemSettingFilter) (*model.SystemSetting, error)
	DeleteSystemSettings(ctx context.Context, filter *model.FindSystemSettingFilter) error

	// UserSetting
	CreateUserSetting(ctx context.Context, m *model.UserSetting) error
	UpdateUserSettings(ctx context.Context, filter *model.FindUserSettingFilter, update map[string]any) (int64, error)
	UpdateUserSetting(ctx context.Context, m *model.UserSetting, update map[string]any) error
	FindUserSettings(ctx context.Context, filter *model.FindUserSettingFilter) ([]*model.UserSetting, error)
	FindUserSetting(ctx context.Context, filter *model.FindUserSettingFilter) (*model.UserSetting, error)
	DeleteUserSettings(ctx context.Context, filter *model.FindUserSettingFilter) error
}
