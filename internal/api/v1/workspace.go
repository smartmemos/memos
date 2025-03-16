package v1

import (
	"context"

	"github.com/samber/do/v2"
	"github.com/sirupsen/logrus"

	"github.com/smartmemos/memos/internal/module/workspace"
	"github.com/smartmemos/memos/internal/module/workspace/model"
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
	setting, err := s.workspaceService.GetSetting(ctx, &model.GetSettingRequest{})
	if err != nil {
		return
	}
	logrus.Info(setting)
	return
}

func (s *WorkspaceService) ListIdentityProviders(context.Context, *v1.ListIdentityProvidersRequest) (resp *v1.ListIdentityProvidersResponse, err error) {
	return
}
