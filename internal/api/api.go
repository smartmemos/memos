package api

import (
	"github.com/samber/do/v2"

	apiv2 "github.com/smartmemos/memos/internal/api/v2"
)

func Init(i do.Injector) {
	// api v2
	do.Provide(i, apiv2.NewAuthService)

}
