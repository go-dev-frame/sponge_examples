package model

import (
	"time"
)

// AfterSales 售后服务表
type AfterSales struct {
	ID          uint64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`             // 售后服务ID
	OrderID     string     `gorm:"column:order_id;type:varchar(24);not null" json:"orderID"`                   // 订单ID
	Type        string     `gorm:"column:type;type:varchar(50);not null" json:"type"`                          // 售后类型（退款、换货、维修）
	Description string     `gorm:"column:description;type:text" json:"description"`                            // 售后描述
	Status      string     `gorm:"column:status;type:varchar(50)" json:"status"`                               // 售后状态
	CreatedAt   *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}
