package module

import (
	"github.com/samber/do/v2"

	memoimpl "github.com/smartmemos/memos/internal/module/memo/service"
	systemimpl "github.com/smartmemos/memos/internal/module/system/service"
)

func Init(i do.Injector) {
	memoimpl.Init(i)
	systemimpl.Init(i)
}
