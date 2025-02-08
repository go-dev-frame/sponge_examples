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

var _ storeV1.SupplierLogicer = (*supplierHandler)(nil)
var _ time.Time

type supplierHandler struct {
	supplierDao dao.SupplierDao
}

// NewSupplierHandler create a handler
func NewSupplierHandler() storeV1.SupplierLogicer {
	return &supplierHandler{
		supplierDao: dao.NewSupplierDao(
			database.GetDB(), // db driver is mysql
			cache.NewSupplierCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *supplierHandler) Create(ctx context.Context, req *storeV1.CreateSupplierRequest) (*storeV1.CreateSupplierReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	supplier := &model.Supplier{}
	err = copier.Copy(supplier, req)
	if err != nil {
		return nil, ecode.ErrCreateSupplier.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.supplierDao.Create(ctx, supplier)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("supplier", supplier), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateSupplierReply{Id: supplier.ID}, nil
}

// DeleteByID delete a record by id
func (h *supplierHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteSupplierByIDRequest) (*storeV1.DeleteSupplierByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.supplierDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteSupplierByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *supplierHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateSupplierByIDRequest) (*storeV1.UpdateSupplierByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	supplier := &model.Supplier{}
	err = copier.Copy(supplier, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDSupplier.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	supplier.ID = req.Id

	err = h.supplierDao.UpdateByID(ctx, supplier)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("supplier", supplier), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateSupplierByIDReply{}, nil
}

// GetByID get a record by id
func (h *supplierHandler) GetByID(ctx context.Context, req *storeV1.GetSupplierByIDRequest) (*storeV1.GetSupplierByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.supplierDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertSupplier(record)
	if err != nil {
		logger.Warn("convertSupplier error", logger.Err(err), logger.Any("supplier", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDSupplier.Err()
	}

	return &storeV1.GetSupplierByIDReply{
		Supplier: data,
	}, nil
}

// List of records by query parameters
func (h *supplierHandler) List(ctx context.Context, req *storeV1.ListSupplierRequest) (*storeV1.ListSupplierReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListSupplier.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.supplierDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	suppliers := []*storeV1.Supplier{}
	for _, record := range records {
		data, err := convertSupplier(record)
		if err != nil {
			logger.Warn("convertSupplier error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		suppliers = append(suppliers, data)
	}

	return &storeV1.ListSupplierReply{
		Total:     total,
		Suppliers: suppliers,
	}, nil
}

func convertSupplier(record *model.Supplier) (*storeV1.Supplier, error) {
	value := &storeV1.Supplier{}
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
