package model

import (
	"time"
)

// ProductMedia 产品多媒体资源表
type ProductMedia struct {
	ID        uint64     `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 媒体ID
	ProductID uint       `gorm:"column:product_id;type:int(11) unsigned;not null" json:"productID"`          // 产品ID
	Type      int        `gorm:"column:type;type:tinyint(4);not null" json:"type"`                           // 类型(1图片/2视频/33D模型)
	Url       string     `gorm:"column:url;type:varchar(255);not null" json:"url"`                           // 资源地址
	SortOrder uint       `gorm:"column:sort_order;type:smallint(6) unsigned;default:0" json:"sortOrder"`     // 排序序号
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}
