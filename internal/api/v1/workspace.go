package v1

import (
	"context"
	"fmt"

	"github.com/samber/do/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usermd "github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/module/workspace"
	"github.com/smartmemos/memos/internal/module/workspace/model"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
	v1 "github.com/smartmemos/memos/internal/proto/api/v1"
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
	setting, err := s.workspaceService.GetSetting(ctx, &model.GetSettingRequest{
		Name: name,
	})
	if err != nil {
		return
	}
	if setting == nil {
		return nil, status.Errorf(codes.NotFound, "workspace setting not found")
	}
	// For storage setting, only host can get it.
	if setting.Key == mpb.SettingKey_STORAGE {
		user, err := grpc_util.GetUserID(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get current user: %v", err)
		}
		if user == nil || user.Role != usermd.RoleHost {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
	}

	return convertWorkspaceSettingFromStore(workspaceSetting), nil
}

func convertWorkspaceSettingFromStore(setting *model.Setting) *mpb.Setting {
	workspaceSetting := &mpb.Setting{
		Name: fmt.Sprintf("%s%s", WorkspaceSettingNamePrefix, setting.Key.String()),
	}
	switch setting.Value.(type) {
	case *mpb.Setting_GeneralSetting:
		workspaceSetting.Value = &mpb.Setting_GeneralSetting{
			GeneralSetting: convertWorkspaceGeneralSettingFromStore(setting.GetGeneralSetting()),
		}
	case *mpb.Setting_StorageSetting:
		workspaceSetting.Value = &mpb.Setting_StorageSetting{
			StorageSetting: convertWorkspaceStorageSettingFromStore(setting.GetStorageSetting()),
		}
	case *mpb.Setting_MemoRelatedSetting:
		workspaceSetting.Value = &mpb.Setting_MemoRelatedSetting{
			MemoRelatedSetting: convertWorkspaceMemoRelatedSettingFromStore(setting.GetMemoRelatedSetting()),
		}
	}
	return workspaceSetting
}

func convertWorkspaceGeneralSettingFromStore(setting *model.GeneralSetting) *mpb.GeneralSetting {
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

func convertWorkspaceStorageSettingFromStore(settingpb *model.StorageSetting) *mpb.StorageSetting {
	if settingpb == nil {
		return nil
	}
	setting := &mpb.StorageSetting{
		StorageType:       mpb.StorageSetting_StorageType(settingpb.StorageType),
		FilepathTemplate:  settingpb.FilepathTemplate,
		UploadSizeLimitMb: settingpb.UploadSizeLimitMb,
	}
	if settingpb.S3Config != nil {
		setting.S3Config = &mpb.StorageSetting_S3Config{
			AccessKeyId:     settingpb.S3Config.AccessKeyId,
			AccessKeySecret: settingpb.S3Config.AccessKeySecret,
			Endpoint:        settingpb.S3Config.Endpoint,
			Region:          settingpb.S3Config.Region,
			Bucket:          settingpb.S3Config.Bucket,
			UsePathStyle:    settingpb.S3Config.UsePathStyle,
		}
	}
	return setting
}

func convertWorkspaceMemoRelatedSettingFromStore(setting *model.MemoRelatedSetting) *mpb.MemoRelatedSetting {
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

func (s *WorkspaceService) ListIdentityProviders(context.Context, *v1.ListIdentityProvidersRequest) (resp *v1.ListIdentityProvidersResponse, err error) {
	return
}
