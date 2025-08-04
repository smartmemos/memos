package v2

import (
	"context"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/memos"
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
		Version:     "version",
		Mode:        "mode",
		InstanceUrl: "instance_url",
	}
	return connect.NewResponse(info), nil
}
