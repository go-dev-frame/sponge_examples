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

var _ storeV1.PurchaseOrderLogicer = (*purchaseOrderHandler)(nil)
var _ time.Time

type purchaseOrderHandler struct {
	purchaseOrderDao dao.PurchaseOrderDao
}

// NewPurchaseOrderHandler create a handler
func NewPurchaseOrderHandler() storeV1.PurchaseOrderLogicer {
	return &purchaseOrderHandler{
		purchaseOrderDao: dao.NewPurchaseOrderDao(
			database.GetDB(), // db driver is mysql
			cache.NewPurchaseOrderCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *purchaseOrderHandler) Create(ctx context.Context, req *storeV1.CreatePurchaseOrderRequest) (*storeV1.CreatePurchaseOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	purchaseOrder := &model.PurchaseOrder{}
	err = copier.Copy(purchaseOrder, req)
	if err != nil {
		return nil, ecode.ErrCreatePurchaseOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.purchaseOrderDao.Create(ctx, purchaseOrder)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("purchaseOrder", purchaseOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreatePurchaseOrderReply{Id: purchaseOrder.ID}, nil
}

// DeleteByID delete a record by id
func (h *purchaseOrderHandler) DeleteByID(ctx context.Context, req *storeV1.DeletePurchaseOrderByIDRequest) (*storeV1.DeletePurchaseOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.purchaseOrderDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeletePurchaseOrderByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *purchaseOrderHandler) UpdateByID(ctx context.Context, req *storeV1.UpdatePurchaseOrderByIDRequest) (*storeV1.UpdatePurchaseOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	purchaseOrder := &model.PurchaseOrder{}
	err = copier.Copy(purchaseOrder, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDPurchaseOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	purchaseOrder.ID = req.Id

	err = h.purchaseOrderDao.UpdateByID(ctx, purchaseOrder)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("purchaseOrder", purchaseOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdatePurchaseOrderByIDReply{}, nil
}

// GetByID get a record by id
func (h *purchaseOrderHandler) GetByID(ctx context.Context, req *storeV1.GetPurchaseOrderByIDRequest) (*storeV1.GetPurchaseOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.purchaseOrderDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertPurchaseOrderPb(record)
	if err != nil {
		logger.Warn("convertPurchaseOrder error", logger.Err(err), logger.Any("purchaseOrder", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDPurchaseOrder.Err()
	}

	return &storeV1.GetPurchaseOrderByIDReply{
		PurchaseOrder: data,
	}, nil
}

// List of records by query parameters
func (h *purchaseOrderHandler) List(ctx context.Context, req *storeV1.ListPurchaseOrderRequest) (*storeV1.ListPurchaseOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListPurchaseOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.purchaseOrderDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	purchaseOrders := []*storeV1.PurchaseOrder{}
	for _, record := range records {
		data, err := convertPurchaseOrderPb(record)
		if err != nil {
			logger.Warn("convertPurchaseOrder error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		purchaseOrders = append(purchaseOrders, data)
	}

	return &storeV1.ListPurchaseOrderReply{
		Total:          total,
		PurchaseOrders: purchaseOrders,
	}, nil
}

func convertPurchaseOrderPb(record *model.PurchaseOrder) (*storeV1.PurchaseOrder, error) {
	value := &storeV1.PurchaseOrder{}
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
