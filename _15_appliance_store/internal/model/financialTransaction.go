package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// FinancialTransaction 财务收支流水表
type FinancialTransaction struct {
	ID              uint64           `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`    // 流水ID
	Type            int              `gorm:"column:type;type:tinyint(4);not null" json:"type"`                           // 类型(1销售收款/2采购付款/3其他收入/4其他支出)
	Amount          *decimal.Decimal `gorm:"column:amount;type:decimal(12,2);not null" json:"amount"`                    // 金额
	RelatedID       string           `gorm:"column:related_id;type:varchar(24)" json:"relatedID"`                        // 关联单据号(订单号/采购单号等)
	TransactionTime *time.Time       `gorm:"column:transaction_time;type:datetime;not null" json:"transactionTime"`      // 交易时间
	Operator        uint             `gorm:"column:operator;type:int(11) unsigned;not null" json:"operator"`             // 操作人ID
	Remark          string           `gorm:"column:remark;type:varchar(255)" json:"remark"`                              // 备注
	CreatedAt       *time.Time       `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt       *time.Time       `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *FinancialTransaction) TableName() string {
	return "financial_transaction"
}
