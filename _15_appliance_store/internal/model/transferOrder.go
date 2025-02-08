package model

import (
	"time"
)

// TransferOrder 库存调拨单主表
type TransferOrder struct {
	ID            string     `gorm:"column:id;type:varchar(24);primary_key" json:"id"`                                    // 调拨单号
	FromStore     uint       `gorm:"column:from_store;type:int(11) unsigned;not null" json:"fromStore"`                   // 调出门店
	ToStore       uint       `gorm:"column:to_store;type:int(11) unsigned;not null" json:"toStore"`                       // 调入门店
	TotalQuantity int        `gorm:"column:total_quantity;type:int(11);not null" json:"totalQuantity"`                    // 总调拨数量
	Status        int        `gorm:"column:status;type:tinyint(4);not null" json:"status"`                                // 状态(0待处理/1已完成)
	CreatedAt     *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"createdAt"` // 创建时间
	CompleteTime  *time.Time `gorm:"column:complete_time;type:datetime" json:"completeTime"`                              // 完成时间
}

// TableName table name
func (m *TransferOrder) TableName() string {
	return "transfer_order"
}
