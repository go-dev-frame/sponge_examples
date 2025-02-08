package model

import (
	"time"
)

// Inventory 门店库存表（分店独立库存）
type Inventory struct {
	ID          uint64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`             // 库存记录ID
	StoreID     uint       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"storeID"`              // 门店ID
	SkuID       uint       `gorm:"column:sku_id;type:int(11) unsigned;not null" json:"skuID"`                  // SKU ID
	Quantity    int        `gorm:"column:quantity;type:int(11);default:0;not null" json:"quantity"`            // 当前库存
	SafetyStock int        `gorm:"column:safety_stock;type:int(11);default:0;not null" json:"safetyStock"`     // 安全库存
	CreatedAt   *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Inventory) TableName() string {
	return "inventory"
}
