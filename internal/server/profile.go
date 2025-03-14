package server

import (
	"github.com/samber/do/v2"
)

type Profile struct {
	Addr      string
	Container do.Injector
}
