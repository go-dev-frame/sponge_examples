package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// SalesOrder 销售订单主表
type SalesOrder struct {
	ID            string           `gorm:"column:id;type:varchar(24);primary_key" json:"id"`                           // 订单号（规则：YYYYMMDDHHMMSSmmm+6位序列）
	StoreID       uint             `gorm:"column:store_id;type:int(11) unsigned;not null" json:"storeID"`              // 门店ID
	CustomerID    uint             `gorm:"column:customer_id;type:int(11) unsigned" json:"customerID"`                 // 客户ID
	TotalAmount   *decimal.Decimal `gorm:"column:total_amount;type:decimal(12,2);not null" json:"totalAmount"`         // 订单总额
	Status        int              `gorm:"column:status;type:tinyint(4);not null" json:"status"`                       // 状态(0待支付/1已支付/2配送中/3已完成/4已退货)
	PaymentMethod int              `gorm:"column:payment_method;type:tinyint(4)" json:"paymentMethod"`                 // 支付方式(1现金/2微信/3支付宝/4刷卡)
	CreatedAt     *time.Time       `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt     *time.Time       `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *SalesOrder) TableName() string {
	return "sales_order"
}
