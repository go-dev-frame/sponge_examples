package model

import (
	"time"
)

// Supplier 供应商信息表
type Supplier struct {
	ID            uint64     `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 供应商ID
	Name          string     `gorm:"column:name;type:varchar(100);not null" json:"name"`                         // 供应商名称
	ContactPerson string     `gorm:"column:contact_person;type:varchar(50);not null" json:"contactPerson"`       // 联系人
	Phone         string     `gorm:"column:phone;type:varchar(20);not null" json:"phone"`                        // 联系电话
	PaymentTerms  string     `gorm:"column:payment_terms;type:text" json:"paymentTerms"`                         // 结算条款
	CreatedAt     *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Supplier) TableName() string {
	return "supplier"
}
