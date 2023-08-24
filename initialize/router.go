package initialize

import (
	"NineTo5Server/global"
	"NineTo5Server/middleware/log"
	"NineTo5Server/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(log.GinLogger(), log.GinRecovery(true))
	global.NineTo5_LOG.Info("register swagger handler")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由组
	systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		//	健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		// 注册基础功能路由 不做鉴权
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitBlogRouter(PublicGroup)
	}
	// 下方为需要鉴权的路由 采用jwt鉴权 TODO 未实现
	//PrivateGroup := Router.Group("")
	//PrivateGroup.Use()
	//{
	//	systemRouter.InitUserRouter(PrivateGroup) // 注册用户路由
	//}
	global.NineTo5_LOG.Info("router register success")
	return Router
}
