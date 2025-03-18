package service

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/memo"
	"github.com/smartmemos/memos/internal/module/user"
	"github.com/smartmemos/memos/internal/module/user/dao"
	"github.com/smartmemos/memos/internal/module/workspace"
)

type Service struct {
	dao     user.DAO
	memoDao memo.DAO
	wsDao   workspace.DAO
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
