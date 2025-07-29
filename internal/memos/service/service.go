package service

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/dao"
)

// Service struct
type Service struct {
	dao memos.DAO
}

// New 实例化
func New(i do.Injector) (memos.Service, error) {
	return &Service{
		dao: do.MustInvoke[memos.DAO](i),
	}, nil
}

// Init 注册服务
func Init(i do.Injector) {
	do.Provide(i, dao.New)
	do.Provide(i, New)
}
