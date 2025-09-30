```go
package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"github.com/go-dev-frame/sponge/internal/cache"
	"github.com/go-dev-frame/sponge/internal/database"
	"github.com/go-dev-frame/sponge/internal/model"
)

var _ AfterSalesDao = (*afterSalesDao)(nil)

// AfterSalesDao defining the dao interface
type AfterSalesDao interface {
	Create(ctx context.Context, table *model.AfterSales) error
	UpdateByID(ctx context.Context, table *model.AfterSales) error
	GetByID(ctx context.Context, id uint64) (*model.AfterSales, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.AfterSales) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.AfterSales) error
	
	// 业务逻辑方法：创建维修工单并分配服务网点
	CreateServiceOrderWithAssignment(ctx context.Context, afterSales *model.AfterSales, customerID uint64, productID uint64, purchaseProof string) (string, string, string, error)
}

type afterSalesDao struct {
	db    *gorm.DB
	cache cache.AfterSalesCache // if nil, the cache is not used.
	sfg   *singleflight.Group    // if cache is nil, the sfg is not used.
}

// NewAfterSalesDao creating the dao interface
func NewAfterSalesDao(db *gorm.DB, xCache cache.AfterSalesCache) AfterSalesDao {
	if xCache == nil {
		return &afterSalesDao{db: db}
	}
	return &afterSalesDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

// CreateServiceOrderWithAssignment 创建维修工单并分配服务网点
func (d *afterSalesDao) CreateServiceOrderWithAssignment(ctx context.Context, afterSales *model.AfterSales, customerID uint64, productID uint64, purchaseProof string) (string, string, string, error) {
	// 步骤1: 验证客户购买记录有效性
	isValid, err := d.validatePurchaseRecord(ctx, customerID, productID, purchaseProof)
	if err != nil {
		return "", "", "", fmt.Errorf("验证购买记录失败: %w", err)
	}
	if !isValid {
		return "", "", "", errors.New("无效的购买记录")
	}

	// 步骤2: 检查产品是否在保修期内
	isUnderWarranty, err := d.checkWarrantyStatus(ctx, productID, purchaseProof)
	if err != nil {
		return "", "", "", fmt.Errorf("检查保修状态失败: %w", err)
	}
	if !isUnderWarranty {
		return "", "", "", errors.New("产品已过保修期")
	}

	// 步骤3: 分配就近服务网点
	assignedStore, err := d.assignNearestServiceStore(ctx, customerID)
	if err != nil {
		return "", "", "", fmt.Errorf("分配服务网点失败: %w", err)
	}

	// 步骤4: 生成服务订单ID和二维码
	serviceOrderID := generateServiceOrderID()
	qrCodeURL, err := d.generateQRCode(ctx, serviceOrderID)
	if err != nil {
		return "", "", "", fmt.Errorf("生成二维码失败: %w", err)
	}

	// 设置订单ID并保存到数据库
	afterSales.OrderID = serviceOrderID
	err = d.Create(ctx, afterSales)
	if err != nil {
		return "", "", "", fmt.Errorf("保存维修工单失败: %w", err)
	}

	return serviceOrderID, assignedStore, qrCodeURL, nil
}

// validatePurchaseRecord 验证购买记录有效性
func (d *afterSalesDao) validatePurchaseRecord(ctx context.Context, customerID uint64, productID uint64, purchaseProof string) (bool, error) {
	// 实现购买记录验证逻辑
	// 这里需要根据实际业务需求实现具体的验证逻辑
	// 示例：检查购买凭证是否存在且有效
	if purchaseProof == "" {
		return false, nil
	}
	
	// 模拟验证逻辑 - 实际项目中应替换为真实的验证代码
	return true, nil
}

// checkWarrantyStatus 检查产品保修状态
func (d *afterSalesDao) checkWarrantyStatus(ctx context.Context, productID uint64, purchaseProof string) (bool, error) {
	// 实现保修状态检查逻辑
	// 这里需要根据实际业务需求实现具体的检查逻辑
	// 示例：检查购买时间是否在保修期内
	
	// 模拟检查逻辑 - 实际项目中应替换为真实的检查代码
	return true, nil
}

// assignNearestServiceStore 分配就近服务网点
func (d *afterSalesDao) assignNearestServiceStore(ctx context.Context, customerID uint64) (string, error) {
	// 实现服务网点分配逻辑
	// 这里需要根据实际业务需求实现具体的分配逻辑
	// 示例：根据客户位置分配最近的服务网点
	
	// 模拟分配逻辑 - 实际项目中应替换为真实的分配代码
	return "服务网点A", nil
}

// generateQRCode 生成二维码
func (d *afterSalesDao) generateQRCode(ctx context.Context, serviceOrderID string) (string, error) {
	// 实现二维码生成逻辑
	// 这里需要根据实际业务需求实现具体的生成逻辑
	// 示例：调用二维码生成服务生成二维码URL
	
	// 模拟生成逻辑 - 实际项目中应替换为真实的生成代码
	return fmt.Sprintf("https://example.com/qrcode/%s", serviceOrderID), nil
}

// generateServiceOrderID 生成服务订单ID
func generateServiceOrderID() string {
	// 生成唯一的服务订单ID
	// 使用时间戳和随机数组合生成
	timestamp := time.Now().Unix()
	randomStr := utils.RandomString(6)
	return fmt.Sprintf("SO%d%s", timestamp, randomStr)
}
```
