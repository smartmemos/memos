package server

import (
	"github.com/samber/do/v2"
)

type Profile struct {
	Container   do.Injector
	Addr        string
	Version     string
	Mode        string
	InstanceUrl string
}
