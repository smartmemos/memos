package v1

import (
	"context"

	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/workspace"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
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

func (s *WorkspaceService) GetWorkspaceProfile(context.Context, *v1pb.GetWorkspaceProfileRequest) (*v1pb.WorkspaceProfile, error) {
	return nil, nil
}
