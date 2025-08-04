package v2

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/model"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type WorkspaceService struct {
	v2pb.UnimplementedWorkspaceServiceHandler
	memosService memos.Service
}

func NewWorkspaceService(i do.Injector) (*WorkspaceService, error) {
	return &WorkspaceService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

func (s *WorkspaceService) GetWorkspaceProfile(ctx context.Context, req *connect.Request[v2pb.GetWorkspaceProfileRequest]) (resp *connect.Response[modelpb.WorkspaceProfile], err error) {
	info := &modelpb.WorkspaceProfile{
		Owner:       "owner",
		Version:     "0.1.0",
		Mode:        "dev",
		InstanceUrl: "http://localhost:8080",
	}
	return connect.NewResponse(info), nil
}

func (s *WorkspaceService) GetWorkspaceSetting(ctx context.Context, req *connect.Request[v2pb.GetWorkspaceSettingRequest]) (resp *connect.Response[modelpb.WorkspaceSetting], err error) {
	parts := strings.Split(req.Msg.Name, "/")
	if len(parts) != 3 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid request"))
	}

	var info *modelpb.WorkspaceSetting
	switch strings.ToLower(parts[2]) {
	case "general":
		setting, err := s.memosService.GetGeneralSetting(ctx)
		if err != nil {
			return nil, err
		}
		info = &modelpb.WorkspaceSetting{
			Name: parts[2],
			Value: &modelpb.WorkspaceSetting_GeneralSetting_{
				GeneralSetting: convertGeneralSettingToProto(setting),
			},
		}
	case "storage":
		info := &modelpb.WorkspaceSetting{
			Name: "name",
		}
		return connect.NewResponse(info), nil
	case "memo_related":
		setting, err := s.memosService.GetMemoRelatedSetting(ctx)
		if err != nil {
			return nil, err
		}
		info = &modelpb.WorkspaceSetting{
			Name: parts[2],
			Value: &modelpb.WorkspaceSetting_MemoRelatedSetting_{
				MemoRelatedSetting: convertMemoRelatedSettingToProto(setting),
			},
		}
	}
	return connect.NewResponse(info), nil
}

func convertGeneralSettingToProto(setting *model.GeneralSetting) *modelpb.WorkspaceSetting_GeneralSetting {
	info := &modelpb.WorkspaceSetting_GeneralSetting{
		Theme:                    setting.Theme,
		DisallowUserRegistration: setting.DisallowUserRegistration,
		DisallowPasswordAuth:     setting.DisallowPasswordAuth,
		AdditionalScript:         setting.AdditionalScript,
		AdditionalStyle:          setting.AdditionalStyle,
		WeekStartDayOffset:       int32(setting.WeekStartDayOffset),
		DisallowChangeUsername:   setting.DisallowChangeUsername,
		DisallowChangeNickname:   setting.DisallowChangeNickname,
	}
	if setting.CustomProfile != nil {
		info.CustomProfile = &modelpb.WorkspaceSetting_GeneralSetting_CustomProfile{
			Title:       setting.CustomProfile.Title,
			Description: setting.CustomProfile.Description,
			LogoUrl:     setting.CustomProfile.LogoURL,
			Locale:      setting.CustomProfile.Locale,
			Appearance:  setting.CustomProfile.Appearance,
		}
	}
	return info
}

func convertMemoRelatedSettingToProto(setting *model.MemoRelatedSetting) *modelpb.WorkspaceSetting_MemoRelatedSetting {
	info := &modelpb.WorkspaceSetting_MemoRelatedSetting{
		DisallowPublicVisibility: setting.DisallowPublicVisibility,
		DisplayWithUpdateTime:    setting.DisplayWithUpdateTime,
		ContentLengthLimit:       int32(setting.ContentLengthLimit),
		EnableDoubleClickEdit:    setting.EnableDoubleClickEdit,
		EnableLinkPreview:        setting.EnableLinkPreview,
		Reactions:                setting.Reactions,
		DisableMarkdownShortcuts: setting.DisableMarkdownShortcuts,
		EnableBlurNsfwContent:    setting.EnableBlurNsfwContent,
		NsfwTags:                 setting.NsfwTags,
	}
	return info
}
