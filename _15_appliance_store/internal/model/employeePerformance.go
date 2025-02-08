package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// EmployeePerformance 员工月度绩效表
type EmployeePerformance struct {
	ID          uint64           `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`            // 记录ID
	EmployeeID  uint             `gorm:"column:employee_id;type:int(11) unsigned;not null" json:"employeeID"`             // 员工ID
	Month       string           `gorm:"column:month;type:char(7);not null" json:"month"`                                 // 统计月份(YYYY-MM)
	SalesAmount *decimal.Decimal `gorm:"column:sales_amount;type:decimal(12,2);default:0.00;not null" json:"salesAmount"` // 销售额
	Commission  *decimal.Decimal `gorm:"column:commission;type:decimal(10,2);default:0.00;not null" json:"commission"`    // 提成金额
	CreatedAt   *time.Time       `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`      // 创建时间
	UpdatedAt   *time.Time       `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"`      // 更新时间
}

// TableName table name
func (m *EmployeePerformance) TableName() string {
	return "employee_performance"
}
