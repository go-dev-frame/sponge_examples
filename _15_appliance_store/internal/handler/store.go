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

var _ storeV1.StoreLogicer = (*storeHandler)(nil)
var _ time.Time

type storeHandler struct {
	storeDao dao.StoreDao
}

// NewStoreHandler create a handler
func NewStoreHandler() storeV1.StoreLogicer {
	return &storeHandler{
		storeDao: dao.NewStoreDao(
			database.GetDB(), // db driver is mysql
			cache.NewStoreCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *storeHandler) Create(ctx context.Context, req *storeV1.CreateStoreRequest) (*storeV1.CreateStoreReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	store := &model.Store{}
	err = copier.Copy(store, req)
	if err != nil {
		return nil, ecode.ErrCreateStore.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.storeDao.Create(ctx, store)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("store", store), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateStoreReply{Id: store.ID}, nil
}

// DeleteByID delete a record by id
func (h *storeHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteStoreByIDRequest) (*storeV1.DeleteStoreByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.storeDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteStoreByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *storeHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateStoreByIDRequest) (*storeV1.UpdateStoreByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	store := &model.Store{}
	err = copier.Copy(store, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDStore.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	store.ID = req.Id

	err = h.storeDao.UpdateByID(ctx, store)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("store", store), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateStoreByIDReply{}, nil
}

// GetByID get a record by id
func (h *storeHandler) GetByID(ctx context.Context, req *storeV1.GetStoreByIDRequest) (*storeV1.GetStoreByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.storeDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertStore(record)
	if err != nil {
		logger.Warn("convertStore error", logger.Err(err), logger.Any("store", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDStore.Err()
	}

	return &storeV1.GetStoreByIDReply{
		Store: data,
	}, nil
}

// List of records by query parameters
func (h *storeHandler) List(ctx context.Context, req *storeV1.ListStoreRequest) (*storeV1.ListStoreReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListStore.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.storeDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	stores := []*storeV1.Store{}
	for _, record := range records {
		data, err := convertStore(record)
		if err != nil {
			logger.Warn("convertStore error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		stores = append(stores, data)
	}

	return &storeV1.ListStoreReply{
		Total:  total,
		Stores: stores,
	}, nil
}

func convertStore(record *model.Store) (*storeV1.Store, error) {
	value := &storeV1.Store{}
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
