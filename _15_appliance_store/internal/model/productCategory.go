package model

import (
	"time"
)

// ProductCategory 产品分类表（支持三级分类）
type ProductCategory struct {
	ID        uint64     `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 分类ID
	ParentID  uint       `gorm:"column:parent_id;type:int(11) unsigned;default:0;not null" json:"parentID"`  // 父分类ID(0表示顶级)
	Name      string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                          // 分类名称
	Level     uint       `gorm:"column:level;type:tinyint(4) unsigned;not null" json:"level"`                // 分类层级(1/2/3)
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *ProductCategory) TableName() string {
	return "product_category"
}
