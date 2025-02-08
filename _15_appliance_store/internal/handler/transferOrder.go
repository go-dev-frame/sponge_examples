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

var _ storeV1.TransferOrderLogicer = (*transferOrderHandler)(nil)
var _ time.Time

type transferOrderHandler struct {
	transferOrderDao dao.TransferOrderDao
}

// NewTransferOrderHandler create a handler
func NewTransferOrderHandler() storeV1.TransferOrderLogicer {
	return &transferOrderHandler{
		transferOrderDao: dao.NewTransferOrderDao(
			database.GetDB(), // db driver is mysql
			cache.NewTransferOrderCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *transferOrderHandler) Create(ctx context.Context, req *storeV1.CreateTransferOrderRequest) (*storeV1.CreateTransferOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	transferOrder := &model.TransferOrder{}
	err = copier.Copy(transferOrder, req)
	if err != nil {
		return nil, ecode.ErrCreateTransferOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.transferOrderDao.Create(ctx, transferOrder)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("transferOrder", transferOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateTransferOrderReply{Id: transferOrder.ID}, nil
}

// DeleteByID delete a record by id
func (h *transferOrderHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteTransferOrderByIDRequest) (*storeV1.DeleteTransferOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.transferOrderDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteTransferOrderByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *transferOrderHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateTransferOrderByIDRequest) (*storeV1.UpdateTransferOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	transferOrder := &model.TransferOrder{}
	err = copier.Copy(transferOrder, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDTransferOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	transferOrder.ID = req.Id

	err = h.transferOrderDao.UpdateByID(ctx, transferOrder)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("transferOrder", transferOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateTransferOrderByIDReply{}, nil
}

// GetByID get a record by id
func (h *transferOrderHandler) GetByID(ctx context.Context, req *storeV1.GetTransferOrderByIDRequest) (*storeV1.GetTransferOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.transferOrderDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertTransferOrderPb(record)
	if err != nil {
		logger.Warn("convertTransferOrder error", logger.Err(err), logger.Any("transferOrder", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDTransferOrder.Err()
	}

	return &storeV1.GetTransferOrderByIDReply{
		TransferOrder: data,
	}, nil
}

// List of records by query parameters
func (h *transferOrderHandler) List(ctx context.Context, req *storeV1.ListTransferOrderRequest) (*storeV1.ListTransferOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListTransferOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.transferOrderDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	transferOrders := []*storeV1.TransferOrder{}
	for _, record := range records {
		data, err := convertTransferOrderPb(record)
		if err != nil {
			logger.Warn("convertTransferOrder error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		transferOrders = append(transferOrders, data)
	}

	return &storeV1.ListTransferOrderReply{
		Total:          total,
		TransferOrders: transferOrders,
	}, nil
}

// PrecheckTransfer 库存调拨预检查
func (h *transferOrderHandler) PrecheckTransfer(ctx context.Context, req *storeV1.PrecheckTransferRequest) (*storeV1.PrecheckTransferReply, error) {
	panic("prompt: 库存调拨预检查")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.transferOrderDao.PrecheckTransfer(ctx, &model.TransferOrder{
	//     	FromStore: req.FromStore,
	//     	ToStore: req.ToStore,
	//     	Items: req.Items,
	//     })
	//	    if err != nil {
	//			logger.Warn("PrecheckTransfer error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.PrecheckTransferReply{
	//     	Status: reply.Status,
	//     }, nil
}

func convertTransferOrderPb(record *model.TransferOrder) (*storeV1.TransferOrder, error) {
	value := &storeV1.TransferOrder{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID
	value.CreatedAt = utils.FormatDateTimeRFC3339(*record.CreatedAt)

	return value, nil
}
