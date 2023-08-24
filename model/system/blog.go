package system

import "NineTo5Server/global"

// Blog 博客
type BlogModel struct {
	global.NineTo5_MODEL
	Title       string `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Content     string `json:"content" gorm:"column:content;type:longtext;not null"`
	Description string `json:"description" gorm:"column:description;type:longtext;not null"`
	//	其他字段
	Slug       string          `json:"slug" gorm:"column:slug;type:varchar(255);not null"`
	Tags       []*TagModel     `gorm:"many2many:sys_blog_tags;"`
	CategoryID uint            `gorm:"index"`
	Category   CategoryModel   `gorm:"foreignKey:CategoryID;associationForeignKey:ID"`
	Comments   []*CommentModel `gorm:"foreignKey:BlogID"`
	// 缩略图
	Thumbnail string `json:"thumbnail" gorm:"column:thumbnail;type:varchar(512)"`
	// 元数据
	MetaDescription string `json:"meta_description" gorm:"column:meta_description;type:varchar(512)"`
}

func (BlogModel) TableName() string {
	return "sys_blog"
}
