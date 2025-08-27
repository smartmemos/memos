package model

const (
	TableMemo          = "memo"
	TableResource      = "resource"
	TableReaction      = "reaction"
	TableMemoRelation  = "memo_relation"
	TableUser          = "user"
	TableUserSetting   = "user_setting"
	TableSession       = "session"
	TableSystemSetting = "system_setting"
	TableInbox         = "inbox"
	TableAttachment    = "attachment"
)

const (
	WorkspaceSettingNamePrefix = "workspace/settings/"
	UserNamePrefix             = "users/"
	MemoNamePrefix             = "memos/"
	AttachmentNamePrefix       = "attachments/"
	ReactionNamePrefix         = "reactions/"
	InboxNamePrefix            = "inboxes/"
	IdentityProviderNamePrefix = "identityProviders/"
	ActivityNamePrefix         = "activities/"
	WebhookNamePrefix          = "webhooks/"
)

// Visibility is the type of a visibility.
type Visibility string

const (
	// Public is the PUBLIC visibility.
	Public Visibility = "PUBLIC"
	// Protected is the PROTECTED visibility.
	Protected Visibility = "PROTECTED"
	// Private is the PRIVATE visibility.
	Private Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	switch v {
	case Public:
		return "PUBLIC"
	case Protected:
		return "PROTECTED"
	case Private:
		return "PRIVATE"
	}
	return "PRIVATE"
}

type RelationType string

const (
	RelationReference RelationType = "REFERENCE"
	RelationComment   RelationType = "COMMENT"
)

type ResStorageType string

const (
	ResStorageTypeLocal         ResStorageType = "LOCAL"
	ResourceStorageTypeS3       ResStorageType = "S3"
	ResourceStorageTypeExternal ResStorageType = "EXTERNAL"
)

// Role is the type of a role.
type Role string

const (
	// RoleHost is the HOST role.
	RoleHost Role = "HOST"
	// RoleAdmin is the ADMIN role.
	RoleAdmin Role = "ADMIN"
	// RoleUser is the USER role.
	RoleUser Role = "USER"
)

type RowStatus string

const (
	// Normal is the status for a normal row.
	Normal RowStatus = "NORMAL"
	// Archived is the status for an archived row.
	Archived RowStatus = "ARCHIVED"
)
