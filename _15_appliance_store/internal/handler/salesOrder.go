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

var _ storeV1.SalesOrderLogicer = (*salesOrderHandler)(nil)
var _ time.Time

type salesOrderHandler struct {
	salesOrderDao dao.SalesOrderDao
}

// NewSalesOrderHandler create a handler
func NewSalesOrderHandler() storeV1.SalesOrderLogicer {
	return &salesOrderHandler{
		salesOrderDao: dao.NewSalesOrderDao(
			database.GetDB(), // db driver is mysql
			cache.NewSalesOrderCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *salesOrderHandler) Create(ctx context.Context, req *storeV1.CreateSalesOrderRequest) (*storeV1.CreateSalesOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	salesOrder := &model.SalesOrder{}
	err = copier.Copy(salesOrder, req)
	if err != nil {
		return nil, ecode.ErrCreateSalesOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.salesOrderDao.Create(ctx, salesOrder)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("salesOrder", salesOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateSalesOrderReply{Id: salesOrder.ID}, nil
}

// DeleteByID delete a record by id
func (h *salesOrderHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteSalesOrderByIDRequest) (*storeV1.DeleteSalesOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.salesOrderDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteSalesOrderByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *salesOrderHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateSalesOrderByIDRequest) (*storeV1.UpdateSalesOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	salesOrder := &model.SalesOrder{}
	err = copier.Copy(salesOrder, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDSalesOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	salesOrder.ID = req.Id

	err = h.salesOrderDao.UpdateByID(ctx, salesOrder)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("salesOrder", salesOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateSalesOrderByIDReply{}, nil
}

// GetByID get a record by id
func (h *salesOrderHandler) GetByID(ctx context.Context, req *storeV1.GetSalesOrderByIDRequest) (*storeV1.GetSalesOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.salesOrderDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertSalesOrderPb(record)
	if err != nil {
		logger.Warn("convertSalesOrder error", logger.Err(err), logger.Any("salesOrder", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDSalesOrder.Err()
	}

	return &storeV1.GetSalesOrderByIDReply{
		SalesOrder: data,
	}, nil
}

// List of records by query parameters
func (h *salesOrderHandler) List(ctx context.Context, req *storeV1.ListSalesOrderRequest) (*storeV1.ListSalesOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListSalesOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.salesOrderDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	salesOrders := []*storeV1.SalesOrder{}
	for _, record := range records {
		data, err := convertSalesOrderPb(record)
		if err != nil {
			logger.Warn("convertSalesOrder error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		salesOrders = append(salesOrders, data)
	}

	return &storeV1.ListSalesOrderReply{
		Total:       total,
		SalesOrders: salesOrders,
	}, nil
}

// GenerateHotSalesReport 生成热销商品分析报告
// 实现逻辑：
// 1. 按指定时间范围统计销量
// 2. 计算商品销售增长率
// 3. 关联库存数据标记滞销风险
// 4. 生成可视化数据格式（JSON/CSV）
func (h *salesOrderHandler) GenerateHotSalesReport(ctx context.Context, req *storeV1.GenerateHotSalesReportRequest) (*storeV1.GenerateHotSalesReportReply, error) {
	panic("prompt: 生成热销商品分析报告  实现逻辑：  1. 按指定时间范围统计销量  2. 计算商品销售增长率  3. 关联库存数据标记滞销风险  4. 生成可视化数据格式（JSON/CSV）")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.salesOrderDao.GenerateHotSalesReport(ctx, &model.SalesOrder{
	//     	StartTime: req.StartTime,
	//     	EndTime: req.EndTime,
	//     	TopN: req.TopN,
	//     })
	//	    if err != nil {
	//			logger.Warn("GenerateHotSalesReport error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.GenerateHotSalesReportReply{
	//     	Items: reply.Items,
	//     	ReportTime: reply.ReportTime,
	//     }, nil
}

func convertSalesOrderPb(record *model.SalesOrder) (*storeV1.SalesOrder, error) {
	value := &storeV1.SalesOrder{}
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
