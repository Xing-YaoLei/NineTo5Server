package system

// Blog 博客
type Blog struct {
	ID      int    `json:"id" gorm:"primary_key:true;column:id;type:int(11);not null"`
	Title   string `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Content string `json:"content" gorm:"column:content;type:longtext;not null"`
	//	其他字段
	Tags     []Tag     `gorm:"many2many:blog_tags;"`
	Category Category  `gorm:"ForeignKey:CategoryID;AssociationForeignKey:ID"`
	Comments []Comment `gorm:"ForeignKey:BlogID;AssociationForeignKey:ID"`
	//	发布状态
	Status int `json:"status" gorm:"column:status;type:int(11);not null;default:0"`
}

// Tag 标签
type Tag struct {
	ID   int    `json:"id" gorm:"primary_key:true;column:id;type:int(11);not null"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	//	其他字段
	Blogs []Blog `gorm:"many2many:blog_tags;"`
}

// Category 分类
type Category struct {
	ID   int    `json:"id" gorm:"primary_key:true;column:id;type:int(11);not null"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	//	其他字段
	Blogs []Blog `gorm:"ForeignKey:CategoryID;AssociationForeignKey:ID"`
}

// Comment 评论
type Comment struct {
	ID      int    `json:"id" gorm:"primary_key:true;column:id;type:int(11);not null"`
	BlogID  int    `json:"blog_id" gorm:"column:blog_id;type:int(11);not null"`
	Email   string `json:"email" gorm:"column:email;type:varchar(255);not null"`
	Content string `json:"content" gorm:"column:content;type:longtext;not null"`
	//	其他字段
	Author string `json:"author" gorm:"column:author;type:varchar(255);not null"`
}
