package model

import (
	"time"
)

// TransferDetail 库存调拨明细表
type TransferDetail struct {
	TransferID string     `gorm:"column:transfer_id;type:varchar(24);primary_key" json:"transferID"`          // 调拨单号
	SkuID      uint       `gorm:"column:sku_id;type:int(11) unsigned;not null" json:"skuID"`                  // SKU ID
	Quantity   uint       `gorm:"column:quantity;type:int(11) unsigned;not null" json:"quantity"`             // 调拨数量
	CreatedAt  *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *TransferDetail) TableName() string {
	return "transfer_detail"
}
