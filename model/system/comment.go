package system

import "NineTo5Server/global"

// // Comment 评论
type CommentModel struct {
	global.NineTo5_MODEL
	Blog    BlogModel `gorm:"foreignKey:BlogID"`
	BlogID  uint      `gorm:"index"`
	Email   string    `json:"email" gorm:"column:email;type:varchar(255);not null"`
	Content string    `json:"content" gorm:"column:content;type:longtext;not null"`
	//	其他字段
	Author   string `json:"author" gorm:"column:author;type:varchar(255);not null"`
	ParentID *int   `json:"parent_id" gorm:"column:parent_id;type:int(11);"`
}

func (CommentModel) TableName() string {
	return "sys_comment"
}
