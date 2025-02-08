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

var _ storeV1.ServiceOrderLogicer = (*serviceOrderHandler)(nil)
var _ time.Time

type serviceOrderHandler struct {
	serviceOrderDao dao.ServiceOrderDao
}

// NewServiceOrderHandler create a handler
func NewServiceOrderHandler() storeV1.ServiceOrderLogicer {
	return &serviceOrderHandler{
		serviceOrderDao: dao.NewServiceOrderDao(
			database.GetDB(), // db driver is mysql
			cache.NewServiceOrderCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *serviceOrderHandler) Create(ctx context.Context, req *storeV1.CreateServiceOrderRequest) (*storeV1.CreateServiceOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	serviceOrder := &model.ServiceOrder{}
	err = copier.Copy(serviceOrder, req)
	if err != nil {
		return nil, ecode.ErrCreateServiceOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.serviceOrderDao.Create(ctx, serviceOrder)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("serviceOrder", serviceOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateServiceOrderReply{Id: serviceOrder.ID}, nil
}

// DeleteByID delete a record by id
func (h *serviceOrderHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteServiceOrderByIDRequest) (*storeV1.DeleteServiceOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.serviceOrderDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteServiceOrderByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *serviceOrderHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateServiceOrderByIDRequest) (*storeV1.UpdateServiceOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	serviceOrder := &model.ServiceOrder{}
	err = copier.Copy(serviceOrder, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDServiceOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	serviceOrder.ID = req.Id

	err = h.serviceOrderDao.UpdateByID(ctx, serviceOrder)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("serviceOrder", serviceOrder), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateServiceOrderByIDReply{}, nil
}

// GetByID get a record by id
func (h *serviceOrderHandler) GetByID(ctx context.Context, req *storeV1.GetServiceOrderByIDRequest) (*storeV1.GetServiceOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.serviceOrderDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertServiceOrderPb(record)
	if err != nil {
		logger.Warn("convertServiceOrder error", logger.Err(err), logger.Any("serviceOrder", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDServiceOrder.Err()
	}

	return &storeV1.GetServiceOrderByIDReply{
		ServiceOrder: data,
	}, nil
}

// List of records by query parameters
func (h *serviceOrderHandler) List(ctx context.Context, req *storeV1.ListServiceOrderRequest) (*storeV1.ListServiceOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListServiceOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.serviceOrderDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	serviceOrders := []*storeV1.ServiceOrder{}
	for _, record := range records {
		data, err := convertServiceOrderPb(record)
		if err != nil {
			logger.Warn("convertServiceOrder error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		serviceOrders = append(serviceOrders, data)
	}

	return &storeV1.ListServiceOrderReply{
		Total:         total,
		ServiceOrders: serviceOrders,
	}, nil
}

func convertServiceOrderPb(record *model.ServiceOrder) (*storeV1.ServiceOrder, error) {
	value := &storeV1.ServiceOrder{}
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
