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
			database.GetDB(), // db driver is mysql
			cache.NewAfterSalesCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *afterSalesHandler) Create(ctx context.Context, req *storeV1.CreateAfterSalesRequest) (*storeV1.CreateAfterSalesReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	afterSales := &model.AfterSales{}
	err = copier.Copy(afterSales, req)
	if err != nil {
		return nil, ecode.ErrCreateAfterSales.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.afterSalesDao.Create(ctx, afterSales)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("afterSales", afterSales), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateAfterSalesReply{Id: afterSales.ID}, nil
}

// DeleteByID delete a record by id
func (h *afterSalesHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteAfterSalesByIDRequest) (*storeV1.DeleteAfterSalesByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.afterSalesDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteAfterSalesByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *afterSalesHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateAfterSalesByIDRequest) (*storeV1.UpdateAfterSalesByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	afterSales := &model.AfterSales{}
	err = copier.Copy(afterSales, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDAfterSales.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	afterSales.ID = req.Id

	err = h.afterSalesDao.UpdateByID(ctx, afterSales)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("afterSales", afterSales), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateAfterSalesByIDReply{}, nil
}

// GetByID get a record by id
func (h *afterSalesHandler) GetByID(ctx context.Context, req *storeV1.GetAfterSalesByIDRequest) (*storeV1.GetAfterSalesByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.afterSalesDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertAfterSales(record)
	if err != nil {
		logger.Warn("convertAfterSales error", logger.Err(err), logger.Any("afterSales", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDAfterSales.Err()
	}

	return &storeV1.GetAfterSalesByIDReply{
		AfterSales: data,
	}, nil
}

// List of records by query parameters
func (h *afterSalesHandler) List(ctx context.Context, req *storeV1.ListAfterSalesRequest) (*storeV1.ListAfterSalesReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListAfterSales.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.afterSalesDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	afterSaless := []*storeV1.AfterSales{}
	for _, record := range records {
		data, err := convertAfterSales(record)
		if err != nil {
			logger.Warn("convertAfterSales error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		afterSaless = append(afterSaless, data)
	}

	return &storeV1.ListAfterSalesReply{
		Total:       total,
		AfterSaless: afterSaless,
	}, nil
}

// CreateServiceOrderWithAssignment 创建维修工单
// 实现逻辑：
// 1. 验证客户购买记录有效性
// 2. 检查产品是否在保修期内
// 3. 分配就近服务网点
// 4. 生成带条形码的维修工单
func (h *afterSalesHandler) CreateServiceOrderWithAssignment(ctx context.Context, req *storeV1.CreateServiceOrderWithAssignmentRequest) (*storeV1.CreateServiceOrderWithAssignmentReply, error) {
	panic("prompt: 创建维修工单  实现逻辑：  1. 验证客户购买记录有效性  2. 检查产品是否在保修期内  3. 分配就近服务网点  4. 生成带条形码的维修工单")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.afterSalesDao.CreateServiceOrderWithAssignment(ctx, &model.AfterSales{
	//     	CustomerID: req.CustomerID,
	//     	ProductID: req.ProductID,
	//     	FaultDescription: req.FaultDescription,
	//     	PurchaseProof: req.PurchaseProof,
	//     })
	//	    if err != nil {
	//			logger.Warn("CreateServiceOrderWithAssignment error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.CreateServiceOrderWithAssignmentReply{
	//     	ServiceOrderID: reply.ServiceOrderID,
	//     	AssignedStore: reply.AssignedStore,
	//     	QrCodeURL: reply.QrCodeURL,
	//     }, nil
}

func convertAfterSales(record *model.AfterSales) (*storeV1.AfterSales, error) {
	value := &storeV1.AfterSales{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID
	value.CreatedAt = utils.FormatDateTimeRFC3339(*record.CreatedAt)
	value.UpdatedAt = utils.FormatDateTimeRFC3339(*record.UpdatedAt)

	return value, nil
}
