package system

import "NineTo5Server/global"

// Category 分类
type CategoryModel struct {
	global.NineTo5_MODEL
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	//	其他字段
	Blogs []BlogModel `gorm:"foreignKey:CategoryID"`
}

func (CategoryModel) TableName() string {
	return "sys_category"
}
