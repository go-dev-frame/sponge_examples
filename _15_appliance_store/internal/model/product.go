package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// Product 产品基本信息表
type Product struct {
	ID             uint64           `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 产品ID
	Name           string           `gorm:"column:name;type:varchar(100);not null" json:"name"`                         // 产品名称
	CategoryID     uint             `gorm:"column:category_id;type:int(11) unsigned;not null" json:"categoryID"`        // 分类ID
	Brand          string           `gorm:"column:brand;type:varchar(50);not null" json:"brand"`                        // 品牌
	Model          string           `gorm:"column:model;type:varchar(50);not null" json:"model"`                        // 型号
	Spec           string           `gorm:"column:spec;type:varchar(100)" json:"spec"`                                  // 规格
	Description    string           `gorm:"column:description;type:text" json:"description"`                            // 详细描述
	Price          *decimal.Decimal `gorm:"column:price;type:decimal(10,2);not null" json:"price"`                      // 价格
	WarrantyMonths uint             `gorm:"column:warranty_months;type:smallint(6) unsigned" json:"warrantyMonths"`     // 保修月数
	ServiceTerms   string           `gorm:"column:service_terms;type:text" json:"serviceTerms"`                         // 售后服务条款
	CreatedAt      *time.Time       `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt      *time.Time       `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Product) TableName() string {
	return "product"
}
