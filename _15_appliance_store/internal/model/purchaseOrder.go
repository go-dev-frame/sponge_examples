package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// PurchaseOrder 采购订单主表
type PurchaseOrder struct {
	ID           string           `gorm:"column:id;type:varchar(24);primary_key" json:"id"`                           // 采购单号（规则：YYYYMMDDHHMMSSmmm+6位序列）
	SupplierID   uint             `gorm:"column:supplier_id;type:int(11) unsigned;not null" json:"supplierID"`        // 供应商ID
	TotalAmount  *decimal.Decimal `gorm:"column:total_amount;type:decimal(12,2);not null" json:"totalAmount"`         // 总金额
	Status       int              `gorm:"column:status;type:tinyint(4);not null" json:"status"`                       // 状态(0待审批/1已批准/2已到货)
	ExpectedDate *time.Time       `gorm:"column:expected_date;type:date;not null" json:"expectedDate"`                // 预计到货日期
	CreatedAt    *time.Time       `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt    *time.Time       `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *PurchaseOrder) TableName() string {
	return "purchase_order"
}
