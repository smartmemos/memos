package v2

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/memos"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
)

type MemoService struct {
	v2pb.UnimplementedMemoServiceHandler
	memosService memos.Service
}

func NewMemoService(i do.Injector) (*MemoService, error) {
	return &MemoService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}
