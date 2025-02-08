package model

import (
	"gorm.io/datatypes"
	"time"
)

// Promotion 促销活动表
type Promotion struct {
	ID        uint64          `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`       // 活动ID
	Name      string          `gorm:"column:name;type:varchar(100);not null" json:"name"`                         // 活动名称
	Type      int             `gorm:"column:type;type:tinyint(4);not null" json:"type"`                           // 类型(1满减/2折扣/3赠品/4组合优惠)
	Rule      *datatypes.JSON `gorm:"column:rule;type:json;not null" json:"rule"`                                 // 活动规则（JSON格式）
	StartTime *time.Time      `gorm:"column:start_time;type:datetime;not null" json:"startTime"`                  // 开始时间
	EndTime   *time.Time      `gorm:"column:end_time;type:datetime;not null" json:"endTime"`                      // 结束时间
	CreatedAt *time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt *time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Promotion) TableName() string {
	return "promotion"
}
