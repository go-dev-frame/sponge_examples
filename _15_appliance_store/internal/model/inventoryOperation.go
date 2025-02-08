package model

import (
	"time"
)

// InventoryOperation 库存操作记录表
type InventoryOperation struct {
	ID           uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`                 // 操作ID
	InventoryID  int64      `gorm:"column:inventory_id;type:bigint(20);not null" json:"inventoryID"`                         // 库存记录ID
	StoreID      uint       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"storeID"`                           // 门店ID
	SkuID        uint       `gorm:"column:sku_id;type:int(11) unsigned;not null" json:"skuID"`                               // SKU ID
	Type         int        `gorm:"column:type;type:tinyint(4);not null" json:"type"`                                        // 操作类型(1采购入库/2退货入库/3销售出库/4调拨出库/5调拨入库)
	Quantity     int        `gorm:"column:quantity;type:int(11);not null" json:"quantity"`                                   // 操作数量
	RelatedOrder string     `gorm:"column:related_order;type:varchar(24)" json:"relatedOrder"`                               // 关联订单号
	Remark       string     `gorm:"column:remark;type:varchar(500)" json:"remark"`                                           // 备注
	Operator     uint       `gorm:"column:operator;type:int(11) unsigned;not null" json:"operator"`                          // 操作人ID
	OperateTime  *time.Time `gorm:"column:operate_time;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"operateTime"` // 操作时间
}

// TableName table name
func (m *InventoryOperation) TableName() string {
	return "inventory_operation"
}
