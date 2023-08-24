package system

import (
	"NineTo5Server/model/common/response"
	"github.com/gin-gonic/gin"
)

type BlogApi struct{}

func (ba *BlogApi) GetBlogList(c *gin.Context) {
	// TODO: 获取已发布的博客列表
	if list, err := blogService.GetBlogList(); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
