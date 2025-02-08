package model

import (
	"gorm.io/datatypes"
	"time"
)

// AuditLog 操作日志审计表
type AuditLog struct {
	ID          uint64          `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`                 // 日志ID
	OperatorID  uint            `gorm:"column:operator_id;type:int(11) unsigned;not null" json:"operatorID"`                     // 操作人ID
	ActionType  string          `gorm:"column:action_type;type:varchar(50);not null" json:"actionType"`                          // 操作类型
	TargetTable string          `gorm:"column:target_table;type:varchar(50);not null" json:"targetTable"`                        // 目标表名
	TargetID    string          `gorm:"column:target_id;type:varchar(50);not null" json:"targetID"`                              // 目标记录ID
	OldValue    *datatypes.JSON `gorm:"column:old_value;type:json" json:"oldValue"`                                              // 旧值
	NewValue    *datatypes.JSON `gorm:"column:new_value;type:json" json:"newValue"`                                              // 新值
	Remark      string          `gorm:"column:remark;type:varchar(255)" json:"remark"`                                           // 备注
	OperateTime *time.Time      `gorm:"column:operate_time;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"operateTime"` // 操作时间
	CreatedAt   *time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`              // 创建时间
}

// TableName table name
func (m *AuditLog) TableName() string {
	return "audit_log"
}
