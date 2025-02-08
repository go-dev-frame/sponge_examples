package model

// CheckDetail 库存盘点明细表
type CheckDetail struct {
	CheckID   string `gorm:"column:check_id;type:varchar(20);primary_key" json:"checkID"` // 盘点单号
	SkuID     uint   `gorm:"column:sku_id;type:int(11) unsigned;not null" json:"skuID"`   // SKU ID
	SystemQty int    `gorm:"column:system_qty;type:int(11);not null" json:"systemQty"`    // 系统库存
	ActualQty int    `gorm:"column:actual_qty;type:int(11);not null" json:"actualQty"`    // 实际库存
}

// TableName table name
func (m *CheckDetail) TableName() string {
	return "check_detail"
}
