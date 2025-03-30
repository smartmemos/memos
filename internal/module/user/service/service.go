package service

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/auth"
	"github.com/smartmemos/memos/internal/module/memo"
	"github.com/smartmemos/memos/internal/module/user"
	"github.com/smartmemos/memos/internal/module/user/dao"
	"github.com/smartmemos/memos/internal/module/workspace"
)

type Service struct {
	dao         user.DAO
	memoDao     memo.DAO
	wsDao       workspace.DAO
	authService auth.Service
}

func New(i do.Injector) (user.Service, error) {
	return &Service{
		dao:         do.MustInvoke[user.DAO](i),
		memoDao:     do.MustInvoke[memo.DAO](i),
		wsDao:       do.MustInvoke[workspace.DAO](i),
		authService: do.MustInvoke[auth.Service](i),
	}, nil
}

func Init(i do.Injector) {
	do.Provide(i, dao.New)
	do.Provide(i, New)
}
