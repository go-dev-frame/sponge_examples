package model

import (
	"time"
)

// ServiceOrder 售后服务维修工单表
type ServiceOrder struct {
	ID             string     `gorm:"column:id;type:varchar(24);primary_key" json:"id"`                           // 工单号
	CustomerID     uint       `gorm:"column:customer_id;type:int(11) unsigned;not null" json:"customerID"`        // 客户ID
	ProductID      uint       `gorm:"column:product_id;type:int(11) unsigned;not null" json:"productID"`          // 产品ID
	FaultDesc      string     `gorm:"column:fault_desc;type:text;not null" json:"faultDesc"`                      // 故障描述
	Status         int        `gorm:"column:status;type:tinyint(4);not null" json:"status"`                       // 状态(0待处理/1维修中/2已完成/3已关闭)
	TechnicianID   uint       `gorm:"column:technician_id;type:int(11) unsigned" json:"technicianID"`             // 维修人员ID
	CompletionTime *time.Time `gorm:"column:completion_time;type:datetime" json:"completionTime"`                 // 完成时间
	CreatedAt      *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt      *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *ServiceOrder) TableName() string {
	return "service_order"
}
