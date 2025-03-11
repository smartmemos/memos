package auth

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/system"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
)

type Service struct {
	v1pb.UnimplementedAuthServiceServer
	system system.Service
}

func New(i do.Injector) (*Service, error) {
	return &Service{
		system: do.MustInvoke[system.Service](i),
	}, nil
}
