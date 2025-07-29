package dao

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/memos"
)

type Dao struct{}

func New(i do.Injector) (memos.DAO, error) {
	return &Dao{}, nil
}
