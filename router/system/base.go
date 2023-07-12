package system

import (
	"NineTo5Server/api"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (br *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")

	baseApi := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
	}

	return baseRouter
}
