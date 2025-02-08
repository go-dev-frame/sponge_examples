package model

import (
	"time"
)

// Employee 员工信息表
type Employee struct {
	ID           uint64     `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 员工ID
	Name         string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                          // 姓名
	Phone        string     `gorm:"column:phone;type:varchar(20);not null" json:"phone"`                        // 手机号
	Role         int        `gorm:"column:role;type:tinyint(4);not null" json:"role"`                           // 角色(1管理员/2店长/3销售员/4财务)
	StoreID      uint       `gorm:"column:store_id;type:int(11) unsigned" json:"storeID"`                       // 所属门店ID
	PasswordHash string     `gorm:"column:password_hash;type:varchar(255);not null" json:"passwordHash"`        // 密码哈希值
	CreatedAt    *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Employee) TableName() string {
	return "employee"
}
