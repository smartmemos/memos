package model

import "github.com/smartmemos/memos/internal/pkg/db"

type SystemSetting struct {
	Name        string
	Value       SystemSettingValue `gorm:"serializer:json"`
	Description string
}

func (SystemSetting) TableName() string {
	return TableSystemSetting
}

type FindSystemSettingFilter struct {
	db.Query

	Name db.F[string]
}

type SystemSettingValue struct {
	*GeneralSetting
	*StorageSetting
	*MemoRelatedSetting
}

type GeneralSetting struct {
	// theme is the name of the selected theme.
	// This references a CSS file in the web/public/themes/ directory.
	Theme string
	// disallow_user_registration disallows user registration.
	DisallowUserRegistration bool
	// disallow_password_auth disallows password authentication.
	DisallowPasswordAuth bool
	// additional_script is the additional script.
	AdditionalScript string
	// additional_style is the additional style.
	AdditionalStyle string
	// custom_profile is the custom profile.
	CustomProfile *CustomProfile
	// week_start_day_offset is the week start day offset from Sunday.
	// 0: Sunday, 1: Monday, 2: Tuesday, 3: Wednesday, 4: Thursday, 5: Friday, 6: Saturday
	// Default is Sunday.
	WeekStartDayOffset int
	// disallow_change_username disallows changing username.
	DisallowChangeUsername bool
	// disallow_change_nickname disallows changing nickname.
	DisallowChangeNickname bool
}

type CustomProfile struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	Locale      string `json:"locale"`
	Appearance  string `json:"appearance"`
}

type StorageSetting struct {
}

type MemoRelatedSetting struct {
	// disallow_public_visibility disallows set memo as public visibility.
	DisallowPublicVisibility bool
	// display_with_update_time orders and displays memo with update time.
	DisplayWithUpdateTime bool
	// content_length_limit is the limit of content length. Unit is byte.
	ContentLengthLimit int
	// enable_double_click_edit enables editing on double click.
	EnableDoubleClickEdit bool
	// enable_link_preview enables links preview.
	EnableLinkPreview bool
	// reactions is the list of reactions.
	Reactions []string
	// disable_markdown_shortcuts disallow the registration of markdown shortcuts.
	DisableMarkdownShortcuts bool
	// enable_blur_nsfw_content enables blurring of content marked as not safe for work (NSFW).
	EnableBlurNsfwContent bool
	// nsfw_tags is the list of tags that mark content as NSFW for blurring.
	NsfwTags []string
}
