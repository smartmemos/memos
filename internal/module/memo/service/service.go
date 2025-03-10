package service

import (
	"github.com/samber/do/v2"
	"github.com/smartmemos/memos/internal/module/memo"
	"github.com/smartmemos/memos/internal/module/memo/dao"
)

// Service struct
type Service struct {
	dao memo.DAO
}

// New 实例化
func New(i do.Injector) (memo.Service, error) {
	return &Service{
		dao: do.MustInvoke[memo.DAO](i),
	}, nil
}

// Init 注册服务
func Init(i do.Injector) {
	do.Provide(i, dao.New)
	do.Provide(i, New)
}
