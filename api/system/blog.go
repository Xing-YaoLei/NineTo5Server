package system

import "github.com/gin-gonic/gin"

type BlogApi struct{}

func (ba *BlogApi) GetBlogList(c *gin.Context) {
	// TODO: 获取已发布的博客列表
}
