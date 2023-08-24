package system

import (
	"NineTo5Server/service"
)

type ApiGroup struct {
	BaseApi
	BlogApi
}

var (
	blogService = service.ServiceGroupApp.SystemServiceGroup.BlogService
)
