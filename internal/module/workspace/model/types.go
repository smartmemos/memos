package model

const (
	TableSetting = "workspace_setting"
)

type SettingKey string

const (
	SettingKeyUnspecified SettingKey = "WORKSPACE_SETTING_KEY_UNSPECIFIED"
	// BASIC is the key for basic settings.
	SettingKeyBasic SettingKey = "BASIC"
	// GENERAL is the key for general settings.
	SettingKeyGeneral SettingKey = "GENERAL"
	// STORAGE is the key for storage settings.
	SettingKeyStorage SettingKey = "STORAGE"
	// MEMO_RELATED is the key for memo related settings.
	SettingKeyMemoRelated SettingKey = "MEMO_RELATED"
)

type StorageType string

const (
	StorageTypeUnspecified StorageType = "STORAGE_TYPE_UNSPECIFIED"
	// StorageTypeDatabase is the database storage type.
	StorageTypeDatabase StorageType = "DATABASE"
	// StorageTypeLocal is the local storage type.
	StorageTypeLocal StorageType = "LOCAL"
	// StorageTypeS3 is the S3 storage type.
	StorageTypeS3 StorageType = "S3"
)
