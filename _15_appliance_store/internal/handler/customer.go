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

var _ storeV1.CustomerLogicer = (*customerHandler)(nil)
var _ time.Time

type customerHandler struct {
	customerDao dao.CustomerDao
}

// NewCustomerHandler create a handler
func NewCustomerHandler() storeV1.CustomerLogicer {
	return &customerHandler{
		customerDao: dao.NewCustomerDao(
			database.GetDB(), // db driver is mysql
			cache.NewCustomerCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *customerHandler) Create(ctx context.Context, req *storeV1.CreateCustomerRequest) (*storeV1.CreateCustomerReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	customer := &model.Customer{}
	err = copier.Copy(customer, req)
	if err != nil {
		return nil, ecode.ErrCreateCustomer.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.customerDao.Create(ctx, customer)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("customer", customer), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateCustomerReply{Id: customer.ID}, nil
}

// DeleteByID delete a record by id
func (h *customerHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteCustomerByIDRequest) (*storeV1.DeleteCustomerByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.customerDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteCustomerByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *customerHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateCustomerByIDRequest) (*storeV1.UpdateCustomerByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	customer := &model.Customer{}
	err = copier.Copy(customer, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDCustomer.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	customer.ID = req.Id

	err = h.customerDao.UpdateByID(ctx, customer)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("customer", customer), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateCustomerByIDReply{}, nil
}

// GetByID get a record by id
func (h *customerHandler) GetByID(ctx context.Context, req *storeV1.GetCustomerByIDRequest) (*storeV1.GetCustomerByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.customerDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertCustomer(record)
	if err != nil {
		logger.Warn("convertCustomer error", logger.Err(err), logger.Any("customer", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDCustomer.Err()
	}

	return &storeV1.GetCustomerByIDReply{
		Customer: data,
	}, nil
}

// List of records by query parameters
func (h *customerHandler) List(ctx context.Context, req *storeV1.ListCustomerRequest) (*storeV1.ListCustomerReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListCustomer.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.customerDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	customers := []*storeV1.Customer{}
	for _, record := range records {
		data, err := convertCustomer(record)
		if err != nil {
			logger.Warn("convertCustomer error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		customers = append(customers, data)
	}

	return &storeV1.ListCustomerReply{
		Total:     total,
		Customers: customers,
	}, nil
}

func convertCustomer(record *model.Customer) (*storeV1.Customer, error) {
	value := &storeV1.Customer{}
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
