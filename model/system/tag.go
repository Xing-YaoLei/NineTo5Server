package system

import "NineTo5Server/global"

// // Tag 标签
type TagModel struct {
	global.NineTo5_MODEL
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	//	其他字段
	Blogs []*BlogModel `gorm:"many2many:sys_blog_tags;"`
}

func (TagModel) TableName() string {
	return "sys_tag"
}
