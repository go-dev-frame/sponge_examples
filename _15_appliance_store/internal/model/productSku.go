package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"time"
)

// ProductSku SKU库存单元表
type ProductSku struct {
	ID         uint64           `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // SKU ID
	ProductID  uint             `gorm:"column:product_id;type:int(11) unsigned;not null" json:"productID"`          // 产品ID
	SkuCode    string           `gorm:"column:sku_code;type:varchar(50);not null" json:"skuCode"`                   // SKU编码
	Attributes *datatypes.JSON  `gorm:"column:attributes;type:json;not null" json:"attributes"`                     // 规格属性(JSON格式，如{"color":"白","capacity":"10L"})
	Price      *decimal.Decimal `gorm:"column:price;type:decimal(10,2);not null" json:"price"`                      // 实际售价
	Stock      int              `gorm:"column:stock;type:int(11);default:0" json:"stock"`                           // 库存数量
	CreatedAt  *time.Time       `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt  *time.Time       `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *ProductSku) TableName() string {
	return "product_sku"
}
