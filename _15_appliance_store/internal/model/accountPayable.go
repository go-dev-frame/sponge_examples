package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// AccountPayable 供应商应付账款表
type AccountPayable struct {
	ID         uint64           `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 记录ID
	SupplierID uint             `gorm:"column:supplier_id;type:int(11) unsigned;not null" json:"supplierID"`           // 供应商ID
	PurchaseID string           `gorm:"column:purchase_id;type:varchar(24);not null" json:"purchaseID"`                // 采购单号
	DueAmount  *decimal.Decimal `gorm:"column:due_amount;type:decimal(12,2);not null" json:"dueAmount"`                // 应付款金额
	PaidAmount *decimal.Decimal `gorm:"column:paid_amount;type:decimal(12,2);default:0.00;not null" json:"paidAmount"` // 已付款金额
	DueDate    *time.Time       `gorm:"column:due_date;type:date;not null" json:"dueDate"`                             // 应付款日期
}

// TableName table name
func (m *AccountPayable) TableName() string {
	return "account_payable"
}
