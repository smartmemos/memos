package v1

import (
	"context"

	"github.com/samber/do/v2"

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

func (s *WorkspaceService) GetProfile(ctx context.Context, req *mpb.GetProfileRequest) (resp *mpb.Profile, err error) {
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
