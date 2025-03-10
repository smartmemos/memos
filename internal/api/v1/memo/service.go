package memo

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/memo"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
)

type Service struct {
	v1pb.UnimplementedMemoServiceServer
	memo memo.Service
}

func New(i do.Injector) (*Service, error) {
	return &Service{
		memo: do.MustInvoke[memo.Service](i),
	}, nil
}
