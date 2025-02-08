package model

import (
	"time"
)

// Store 门店信息表
type Store struct {
	ID           uint64     `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 门店ID
	Name         string     `gorm:"column:name;type:varchar(100);not null" json:"name"`                         // 门店名称
	Address      string     `gorm:"column:address;type:varchar(255);not null" json:"address"`                   // 详细地址
	ContactPhone string     `gorm:"column:contact_phone;type:varchar(20);not null" json:"contactPhone"`         // 联系电话
	ManagerID    uint       `gorm:"column:manager_id;type:int(11) unsigned" json:"managerID"`                   // 店长ID
	CreatedAt    *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Store) TableName() string {
	return "store"
}
