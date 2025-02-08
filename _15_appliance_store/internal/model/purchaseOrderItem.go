package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// PurchaseOrderItem 采购订单明细表
type PurchaseOrderItem struct {
	ID        uint64           `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`     // 明细ID
	OrderID   string           `gorm:"column:order_id;type:varchar(24);not null" json:"orderID"`                    // 订单号
	SkuID     uint             `gorm:"column:sku_id;type:int(11) unsigned;not null" json:"skuID"`                   // SKU ID
	Quantity  uint             `gorm:"column:quantity;type:int(11) unsigned;not null" json:"quantity"`              // 购买数量
	Price     *decimal.Decimal `gorm:"column:price;type:decimal(10,2);default:0.00;not null" json:"price"`          // 产品单价
	UnitPrice *decimal.Decimal `gorm:"column:unit_price;type:decimal(10,2);default:0.00;not null" json:"unitPrice"` // 成交单价
	CreatedAt *time.Time       `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`  // 创建时间
	UpdatedAt *time.Time       `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"`  // 更新时间
}

// TableName table name
func (m *PurchaseOrderItem) TableName() string {
	return "purchase_order_item"
}
