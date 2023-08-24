package system

import (
	"NineTo5Server/api"
	"github.com/gin-gonic/gin"
)

type BlogRouter struct{}

func (br *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	blogRouter := Router.Group("blog")
	blogApi := api.ApiGroupApp.SystemApiGroup.BlogApi
	{
		blogRouter.GET("getBlogList", blogApi.GetBlogList)
	}
	return blogRouter
}
