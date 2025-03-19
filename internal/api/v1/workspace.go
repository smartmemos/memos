package v1

import (
	"context"

	"github.com/samber/do/v2"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartmemos/memos/internal/module/workspace"
	"github.com/smartmemos/memos/internal/module/workspace/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	mpb "github.com/smartmemos/memos/internal/proto/model/workspace"
)

type WorkspaceService struct {
	v1pb.UnimplementedWorkspaceServiceServer
	workspaceService workspace.Service
}

func NewWorkspaceService(i do.Injector) (*WorkspaceService, error) {
	return &WorkspaceService{
		workspaceService: do.MustInvoke[workspace.Service](i),
	}, nil
}

func (s *WorkspaceService) GetWorkspaceProfile(ctx context.Context, req *v1pb.GetWorkspaceProfileRequest) (resp *mpb.Profile, err error) {
	profile, err := s.workspaceService.GetProfile(ctx, &model.GetProfileRequest{})
	if err != nil {
		return
	}
	return convertProfileToProto(profile), nil
}

func convertProfileToProto(profile *model.Profile) *mpb.Profile {
	return &mpb.Profile{
		Owner:       profile.Owner,
		Version:     profile.Version,
		Mode:        profile.Mode,
		InstanceUrl: profile.InstanceUrl,
	}
}

func (s *WorkspaceService) GetWorkspaceSetting(ctx context.Context, req *v1pb.GetWorkspaceSettingRequest) (resp *mpb.Setting, err error) {
	name, err := ExtractWorkspaceSettingKeyFromName(req.Name)
	if err != nil {
		return
	}
	key := model.SettingKey(name)
	switch key {
	case model.SettingKeyGeneral:
		var value model.GeneralSetting
		if err = s.workspaceService.GetSetting(ctx, key, &value); err != nil {
			return
		}
		resp = &mpb.Setting{
			Name: req.Name,
			Value: &mpb.Setting_GeneralSetting{
				GeneralSetting: convertWorkspaceGeneralSetting(&value),
			},
		}
	case model.SettingKeyMemoRelated:
		var value model.MemoRelatedSetting
		if err = s.workspaceService.GetSetting(ctx, key, &value); err != nil {
			return
		}
		resp = &mpb.Setting{
			Name: req.Name,
			Value: &mpb.Setting_MemoRelatedSetting{
				MemoRelatedSetting: convertWorkspaceMemoRelatedSetting(&value),
			},
		}
	default:
		return
	}
	return
}

func convertWorkspaceGeneralSetting(setting *model.GeneralSetting) *mpb.GeneralSetting {
	if setting == nil {
		return nil
	}
	generalSetting := &mpb.GeneralSetting{
		DisallowUserRegistration: setting.DisallowUserRegistration,
		DisallowPasswordAuth:     setting.DisallowPasswordAuth,
		AdditionalScript:         setting.AdditionalScript,
		AdditionalStyle:          setting.AdditionalStyle,
		WeekStartDayOffset:       setting.WeekStartDayOffset,
		DisallowChangeUsername:   setting.DisallowChangeUsername,
		DisallowChangeNickname:   setting.DisallowChangeNickname,
	}
	if setting.CustomProfile != nil {
		generalSetting.CustomProfile = &mpb.CustomProfile{
			Title:       setting.CustomProfile.Title,
			Description: setting.CustomProfile.Description,
			LogoUrl:     setting.CustomProfile.LogoUrl,
			Locale:      setting.CustomProfile.Locale,
			Appearance:  setting.CustomProfile.Appearance,
		}
	}
	return generalSetting
}

func convertWorkspaceMemoRelatedSetting(setting *model.MemoRelatedSetting) *mpb.MemoRelatedSetting {
	if setting == nil {
		return nil
	}
	return &mpb.MemoRelatedSetting{
		DisallowPublicVisibility: setting.DisallowPublicVisibility,
		DisplayWithUpdateTime:    setting.DisplayWithUpdateTime,
		ContentLengthLimit:       setting.ContentLengthLimit,
		EnableDoubleClickEdit:    setting.EnableDoubleClickEdit,
		EnableLinkPreview:        setting.EnableLinkPreview,
		EnableComment:            setting.EnableComment,
		EnableLocation:           setting.EnableLocation,
		Reactions:                setting.Reactions,
		DisableMarkdownShortcuts: setting.DisableMarkdownShortcuts,
		EnableBlurNsfwContent:    setting.EnableBlurNsfwContent,
		NsfwTags:                 setting.NsfwTags,
	}
}

func convertWorkspaceSetting(setting *model.Setting) (ret *mpb.Setting) {
	return
}

func (s *WorkspaceService) ListIdentityProviders(context.Context, *v1pb.ListIdentityProvidersRequest) (resp *v1pb.ListIdentityProvidersResponse, err error) {
	return
}

func (s *WorkspaceService) ListInboxes(context.Context, *v1pb.ListInboxesRequest) (resp *v1pb.ListInboxesResponse, err error) {
	return
}

func (s *WorkspaceService) UpdateInbox(context.Context, *v1pb.UpdateInboxRequest) (resp *mpb.Inbox, err error) {
	return
}

func (s *WorkspaceService) DeleteInbox(context.Context, *v1pb.DeleteInboxRequest) (_ *emptypb.Empty, err error) {
	return
}
