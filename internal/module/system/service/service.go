package service

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/system"
	"github.com/smartmemos/memos/internal/module/system/dao"
)

type Service struct {
	dao system.DAO
}

func New(i do.Injector) (system.Service, error) {
	return &Service{
		dao: do.MustInvoke[system.DAO](i),
	}, nil
}

func Init(i do.Injector) {
	do.Provide(i, dao.New)
	do.Provide(i, New)
}
