package dao

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/workspace"
)

type Dao struct {
}

func New(i do.Injector) (workspace.DAO, error) {
	return &Dao{}, nil
}
