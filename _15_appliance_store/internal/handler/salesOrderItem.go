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

var _ storeV1.SalesOrderItemLogicer = (*salesOrderItemHandler)(nil)
var _ time.Time

type salesOrderItemHandler struct {
	salesOrderItemDao dao.SalesOrderItemDao
}

// NewSalesOrderItemHandler create a handler
func NewSalesOrderItemHandler() storeV1.SalesOrderItemLogicer {
	return &salesOrderItemHandler{
		salesOrderItemDao: dao.NewSalesOrderItemDao(
			database.GetDB(), // db driver is mysql
			cache.NewSalesOrderItemCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *salesOrderItemHandler) Create(ctx context.Context, req *storeV1.CreateSalesOrderItemRequest) (*storeV1.CreateSalesOrderItemReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	salesOrderItem := &model.SalesOrderItem{}
	err = copier.Copy(salesOrderItem, req)
	if err != nil {
		return nil, ecode.ErrCreateSalesOrderItem.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.salesOrderItemDao.Create(ctx, salesOrderItem)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("salesOrderItem", salesOrderItem), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateSalesOrderItemReply{Id: salesOrderItem.ID}, nil
}

// DeleteByID delete a record by id
func (h *salesOrderItemHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteSalesOrderItemByIDRequest) (*storeV1.DeleteSalesOrderItemByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.salesOrderItemDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteSalesOrderItemByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *salesOrderItemHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateSalesOrderItemByIDRequest) (*storeV1.UpdateSalesOrderItemByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	salesOrderItem := &model.SalesOrderItem{}
	err = copier.Copy(salesOrderItem, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDSalesOrderItem.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	salesOrderItem.ID = req.Id

	err = h.salesOrderItemDao.UpdateByID(ctx, salesOrderItem)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("salesOrderItem", salesOrderItem), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateSalesOrderItemByIDReply{}, nil
}

// GetByID get a record by id
func (h *salesOrderItemHandler) GetByID(ctx context.Context, req *storeV1.GetSalesOrderItemByIDRequest) (*storeV1.GetSalesOrderItemByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.salesOrderItemDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertSalesOrderItem(record)
	if err != nil {
		logger.Warn("convertSalesOrderItem error", logger.Err(err), logger.Any("salesOrderItem", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDSalesOrderItem.Err()
	}

	return &storeV1.GetSalesOrderItemByIDReply{
		SalesOrderItem: data,
	}, nil
}

// List of records by query parameters
func (h *salesOrderItemHandler) List(ctx context.Context, req *storeV1.ListSalesOrderItemRequest) (*storeV1.ListSalesOrderItemReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListSalesOrderItem.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.salesOrderItemDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	salesOrderItems := []*storeV1.SalesOrderItem{}
	for _, record := range records {
		data, err := convertSalesOrderItem(record)
		if err != nil {
			logger.Warn("convertSalesOrderItem error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		salesOrderItems = append(salesOrderItems, data)
	}

	return &storeV1.ListSalesOrderItemReply{
		Total:           total,
		SalesOrderItems: salesOrderItems,
	}, nil
}

func convertSalesOrderItem(record *model.SalesOrderItem) (*storeV1.SalesOrderItem, error) {
	value := &storeV1.SalesOrderItem{}
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
