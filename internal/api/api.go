package api

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/api/v1/auth"
	"github.com/smartmemos/memos/internal/api/v1/memo"
)

func Init(i do.Injector) {
	do.Provide(i, memo.New)
	do.Provide(i, auth.New)
}
