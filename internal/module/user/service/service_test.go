package service

import (
	"testing"

	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/config"
	"github.com/smartmemos/memos/internal/module/user"
)

func TestMain(m *testing.M) {

	config.Init("../../../../configs/config.toml")

	do.Provide(nil, func(i do.Injector) (*Service, error) {
		return &Service{
			dao: do.MustInvoke[user.DAO](nil),
		}, nil
	})
	do.Provide(nil, func(i do.Injector) (user.DAO, error) {
		return nil, nil
	})

	_ = m.Run()
}
