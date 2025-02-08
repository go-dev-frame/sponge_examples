package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// Coupon 优惠券管理表
type Coupon struct {
	ID              string           `gorm:"column:id;type:varchar(20);primary_key" json:"id"`                                        // 优惠券ID
	Type            int              `gorm:"column:type;type:tinyint(4);not null" json:"type"`                                        // 类型(1满减/2折扣/3代金券)
	Value           *decimal.Decimal `gorm:"column:value;type:decimal(10,2);default:0.00;not null" json:"value"`                      // 面值/折扣率
	AmountCondition *decimal.Decimal `gorm:"column:amount_condition;type:decimal(10,2);default:0.00;not null" json:"amountCondition"` // 使用条件(满多少元可用)
	CustomerID      uint             `gorm:"column:customer_id;type:int(11) unsigned" json:"customerID"`                              // 绑定客户ID
	Status          int              `gorm:"column:status;type:tinyint(4);not null" json:"status"`                                    // 状态(0未发放/1未使用/2已使用/3已过期)
	ExpireTime      *time.Time       `gorm:"column:expire_time;type:datetime;not null" json:"expireTime"`                             // 过期时间
}

// TableName table name
func (m *Coupon) TableName() string {
	return "coupon"
}
