package system

import (
	"NineTo5Server/api"
	"github.com/gin-gonic/gin"
)

type BlogRouter struct{}

func (s *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	blogRouter := Router.Group("blog")
	blogApi := api.ApiGroupApp.SystemApiGroup.BlogApi
	{
		blogRouter.GET("getBlogInfo", blogApi.GetBlogList)
		// blogRouter.POST("getBlogs", blogApi.GetBlogs)
		// blogRouter.POST("deleteBlog", blogApi.DeleteBlog)
		// blogRouter.POST("addBlog", blogApi.AddBlog)
		// blogRouter.POST("editBlog", blogApi.EditBlog)
		// blogRouter.POST("switchActive", blogApi.SwitchActive)
	}
}
