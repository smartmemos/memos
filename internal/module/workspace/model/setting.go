package model

import (
	"strings"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type Setting struct {
	db.Model

	Name        string
	Value       string
	Description string
}

func (Setting) TableName() string {
	return TableSetting
}

type FindSettingFilter struct {
	db.BaseFilter
	Name string
}

func (f FindSettingFilter) GetQuery() (query string, args []any) {
	var where []string
	if f.Name != "" {
		where = append(where, "name=?")
		args = append(args, f.Name)
	}
	query = strings.Join(where, " and ")
	return
}

type BasicSetting struct {
	SecretKey string `json:"secret_key"`
}

type GeneralSetting struct {
	// additional_script is the additional script.
	AdditionalScript string `json:"additional_script"`
	// additional_style is the additional style.
	AdditionalStyle string `json:"additional_style"`
	// custom_profile is the custom profile.
	CustomProfile *CustomProfile `json:"custom_profile"`
}

type CustomProfile struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LogoUrl     string `json:"logo_url"`
	Locale      string `json:"locale"`
	Appearance  string `json:"appearance"`
}

type StorageSetting struct {
	StorageType       StorageType      `json:"storage_type"`
	FilepathTemplate  string           `json:"filepath_template"`
	UploadSizeLimitMb int64            `json:"upload_size_limit_mb"`
	S3Config          *StorageS3Config `json:"s3_config"`
}

type StorageS3Config struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Endpoint        string `json:"endpoint"`
	Region          string `json:"region"`
	Bucket          string `json:"bucket"`
}

type MemoRelatedSetting struct {
	DisallowPublicVisibility bool  `json:"disallow_public_visibility"`
	DisplayWithUpdateTime    bool  `json:"display_with_update_time"`
	ContentLengthLimit       int32 `json:"content_length_limit"`
	EnableAutoCompact        bool  `json:"enable_auto_compact"`
	EnableDoubleClickEdit    bool  `json:"enable_double_click_edit"`
	EnableLinkPreview        bool  `json:"enable_link_preview"`
}

type GetSettingRequest struct {
	Name string
}
