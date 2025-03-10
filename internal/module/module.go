package module

import (
	"github.com/samber/do/v2"

	memoimpl "github.com/smartmemos/memos/internal/module/memo/service"
)

func Init(i do.Injector) {
	memoimpl.Init(i)
}
