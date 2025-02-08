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

var _ storeV1.PurchaseOrderItemLogicer = (*purchaseOrderItemHandler)(nil)
var _ time.Time

type purchaseOrderItemHandler struct {
	purchaseOrderItemDao dao.PurchaseOrderItemDao
}

// NewPurchaseOrderItemHandler create a handler
func NewPurchaseOrderItemHandler() storeV1.PurchaseOrderItemLogicer {
	return &purchaseOrderItemHandler{
		purchaseOrderItemDao: dao.NewPurchaseOrderItemDao(
			database.GetDB(), // db driver is mysql
			cache.NewPurchaseOrderItemCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *purchaseOrderItemHandler) Create(ctx context.Context, req *storeV1.CreatePurchaseOrderItemRequest) (*storeV1.CreatePurchaseOrderItemReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	purchaseOrderItem := &model.PurchaseOrderItem{}
	err = copier.Copy(purchaseOrderItem, req)
	if err != nil {
		return nil, ecode.ErrCreatePurchaseOrderItem.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.purchaseOrderItemDao.Create(ctx, purchaseOrderItem)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("purchaseOrderItem", purchaseOrderItem), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreatePurchaseOrderItemReply{Id: purchaseOrderItem.ID}, nil
}

// DeleteByID delete a record by id
func (h *purchaseOrderItemHandler) DeleteByID(ctx context.Context, req *storeV1.DeletePurchaseOrderItemByIDRequest) (*storeV1.DeletePurchaseOrderItemByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.purchaseOrderItemDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeletePurchaseOrderItemByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *purchaseOrderItemHandler) UpdateByID(ctx context.Context, req *storeV1.UpdatePurchaseOrderItemByIDRequest) (*storeV1.UpdatePurchaseOrderItemByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	purchaseOrderItem := &model.PurchaseOrderItem{}
	err = copier.Copy(purchaseOrderItem, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDPurchaseOrderItem.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	purchaseOrderItem.ID = req.Id

	err = h.purchaseOrderItemDao.UpdateByID(ctx, purchaseOrderItem)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("purchaseOrderItem", purchaseOrderItem), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdatePurchaseOrderItemByIDReply{}, nil
}

// GetByID get a record by id
func (h *purchaseOrderItemHandler) GetByID(ctx context.Context, req *storeV1.GetPurchaseOrderItemByIDRequest) (*storeV1.GetPurchaseOrderItemByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.purchaseOrderItemDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertPurchaseOrderItem(record)
	if err != nil {
		logger.Warn("convertPurchaseOrderItem error", logger.Err(err), logger.Any("purchaseOrderItem", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDPurchaseOrderItem.Err()
	}

	return &storeV1.GetPurchaseOrderItemByIDReply{
		PurchaseOrderItem: data,
	}, nil
}

// List of records by query parameters
func (h *purchaseOrderItemHandler) List(ctx context.Context, req *storeV1.ListPurchaseOrderItemRequest) (*storeV1.ListPurchaseOrderItemReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListPurchaseOrderItem.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.purchaseOrderItemDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	purchaseOrderItems := []*storeV1.PurchaseOrderItem{}
	for _, record := range records {
		data, err := convertPurchaseOrderItem(record)
		if err != nil {
			logger.Warn("convertPurchaseOrderItem error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		purchaseOrderItems = append(purchaseOrderItems, data)
	}

	return &storeV1.ListPurchaseOrderItemReply{
		Total:              total,
		PurchaseOrderItems: purchaseOrderItems,
	}, nil
}

func convertPurchaseOrderItem(record *model.PurchaseOrderItem) (*storeV1.PurchaseOrderItem, error) {
	value := &storeV1.PurchaseOrderItem{}
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
