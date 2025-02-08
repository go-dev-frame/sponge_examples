package model

import (
	"time"
)

// Customer 客户信息表
type Customer struct {
	ID              uint64     `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 客户ID
	Name            string     `gorm:"column:name;type:varchar(50);not null" json:"name"`                          // 客户姓名
	Phone           string     `gorm:"column:phone;type:varchar(20);not null" json:"phone"`                        // 手机号
	Email           string     `gorm:"column:email;type:varchar(200)" json:"email"`                                // 客户邮箱
	Address         string     `gorm:"column:address;type:varchar(200)" json:"address"`                            // 地址
	Birthday        *time.Time `gorm:"column:birthday;type:date" json:"birthday"`                                  // 生日
	MembershipLevel int        `gorm:"column:membership_level;type:tinyint(4);default:0" json:"membershipLevel"`   // 会员等级(0普通客户)
	TotalPoints     uint       `gorm:"column:total_points;type:int(11) unsigned;default:0" json:"totalPoints"`     // 累计积分
	CreatedAt       *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt       *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Customer) TableName() string {
	return "customer"
}
