package global

import (
	"time"
)

type NineTo5_MODEL struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt time.Time `gorm:"index" json:"-"` // 删除时间
}
