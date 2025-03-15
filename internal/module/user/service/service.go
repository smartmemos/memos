package service

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/user"
	"github.com/smartmemos/memos/internal/module/user/dao"
)

type Service struct {
	dao user.DAO
}

func New(i do.Injector) (user.Service, error) {
	return &Service{
		dao: do.MustInvoke[user.DAO](i),
	}, nil
}

func Init(i do.Injector) {
	do.Provide(i, dao.New)
	do.Provide(i, New)
}
