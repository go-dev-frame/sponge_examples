package model

import (
	"time"
)

// InventoryCheck 库存盘点主表
type InventoryCheck struct {
	ID         string     `gorm:"column:id;type:varchar(20);primary_key" json:"id"`                                    // 盘点单号
	StoreID    uint       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"storeID"`                       // 门店ID
	OperatorID uint       `gorm:"column:operator_id;type:int(11) unsigned;not null" json:"operatorID"`                 // 操作人ID
	CheckTime  *time.Time `gorm:"column:check_time;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"checkTime"` // 盘点时间
	TotalDiff  int        `gorm:"column:total_diff;type:int(11);not null" json:"totalDiff"`                            // 总差异数量
}

// TableName table name
func (m *InventoryCheck) TableName() string {
	return "inventory_check"
}
