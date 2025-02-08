package model

import (
	"time"
)

// Notification 系统消息通知表
type Notification struct {
	ID         uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`    // 通知ID
	Type       int        `gorm:"column:type;type:tinyint(4);not null" json:"type"`                           // 类型(1库存预警/2订单状态/3促销提醒)
	Recipient  string     `gorm:"column:recipient;type:varchar(100);not null" json:"recipient"`               // 接收人(手机号/邮箱/员工ID)
	Content    string     `gorm:"column:content;type:text;not null" json:"content"`                           // 通知内容
	SendStatus int        `gorm:"column:send_status;type:tinyint(4);not null" json:"sendStatus"`              // 发送状态(0待发送/1已发送/2失败)
	CreatedAt  *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Notification) TableName() string {
	return "notification"
}
