package system

import (
	"NineTo5Server/global"
	BlogReq "NineTo5Server/model/system"
)

type BlogService struct{}

func (bs *BlogService) GetBlogList() ([]BlogReq.BlogModel, error) {
	var blogs []BlogReq.BlogModel
	err := global.NineTo5_DB.Preload("Tags").Find(&blogs).Error
	return blogs, err
}
