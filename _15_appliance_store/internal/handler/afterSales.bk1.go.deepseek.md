```go
package handler

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/copier"

	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
	"github.com/go-dev-frame/sponge/pkg/utils"

	storeV1 "store/api/store/v1"
	"store/internal/cache"
	"store/internal/dao"
	"store/internal/database"
	"store/internal/ecode"
	"store/internal/model"
)

var _ storeV1.AfterSalesLogicer = (*afterSalesHandler)(nil)
var _ time.Time

type afterSalesHandler struct {
	afterSalesDao dao.AfterSalesDao
}

// NewAfterSalesHandler create a handler
func NewAfterSalesHandler() storeV1.AfterSalesLogicer {
	return &afterSalesHandler{
		afterSalesDao: dao.NewAfterSalesDao(
			database.GetDB(),	// db driver is mysql
			cache.NewAfterSalesCache(database.GetCacheType()),
		),
	}
}

// CreateServiceOrderWithAssignment 创建维修工单
// 实现逻辑：
// 1. 验证客户购买记录有效性
// 2. 检查产品是否在保修期内
// 3. 分配就近服务网点
// 4. 生成带条形码的维修工单
func (h *afterSalesHandler) CreateServiceOrderWithAssignment(ctx context.Context, req *storeV1.CreateServiceOrderWithAssignmentRequest) (*storeV1.CreateServiceOrderWithAssignmentReply, error) {
	// 验证请求参数
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	// 创建售后服务模型对象
	afterSales := &model.AfterSales{
		OrderID:     req.OrderID,
		Type:        "维修", // 固定为维修类型
		Description: req.FaultDescription,
		Status:      "待处理",
	}

	// 调用DAO层创建维修工单
	serviceOrderID, assignedStore, qrCodeURL, err := h.afterSalesDao.CreateServiceOrderWithAssignment(ctx, afterSales, req.CustomerID, req.ProductID, req.PurchaseProof)
	if err != nil {
		logger.Warn("CreateServiceOrderWithAssignment error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	// 返回创建结果
	return &storeV1.CreateServiceOrderWithAssignmentReply{
		ServiceOrderID: serviceOrderID,
		AssignedStore:  assignedStore,
		QrCodeURL:      qrCodeURL,
	}, nil
}
```
