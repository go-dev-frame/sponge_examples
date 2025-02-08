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

	storeV1 "store/api/store/v1"
	"store/internal/cache"
	"store/internal/dao"
	"store/internal/database"
	"store/internal/ecode"
	"store/internal/model"
)

var _ storeV1.InventoryCheckLogicer = (*inventoryCheckHandler)(nil)
var _ time.Time

type inventoryCheckHandler struct {
	inventoryCheckDao dao.InventoryCheckDao
}

// NewInventoryCheckHandler create a handler
func NewInventoryCheckHandler() storeV1.InventoryCheckLogicer {
	return &inventoryCheckHandler{
		inventoryCheckDao: dao.NewInventoryCheckDao(
			database.GetDB(), // db driver is mysql
			cache.NewInventoryCheckCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *inventoryCheckHandler) Create(ctx context.Context, req *storeV1.CreateInventoryCheckRequest) (*storeV1.CreateInventoryCheckReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	inventoryCheck := &model.InventoryCheck{}
	err = copier.Copy(inventoryCheck, req)
	if err != nil {
		return nil, ecode.ErrCreateInventoryCheck.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.inventoryCheckDao.Create(ctx, inventoryCheck)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("inventoryCheck", inventoryCheck), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateInventoryCheckReply{Id: inventoryCheck.ID}, nil
}

// DeleteByID delete a record by id
func (h *inventoryCheckHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteInventoryCheckByIDRequest) (*storeV1.DeleteInventoryCheckByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.inventoryCheckDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteInventoryCheckByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *inventoryCheckHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateInventoryCheckByIDRequest) (*storeV1.UpdateInventoryCheckByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	inventoryCheck := &model.InventoryCheck{}
	err = copier.Copy(inventoryCheck, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDInventoryCheck.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	inventoryCheck.ID = req.Id

	err = h.inventoryCheckDao.UpdateByID(ctx, inventoryCheck)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("inventoryCheck", inventoryCheck), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateInventoryCheckByIDReply{}, nil
}

// GetByID get a record by id
func (h *inventoryCheckHandler) GetByID(ctx context.Context, req *storeV1.GetInventoryCheckByIDRequest) (*storeV1.GetInventoryCheckByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.inventoryCheckDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertInventoryCheckPb(record)
	if err != nil {
		logger.Warn("convertInventoryCheck error", logger.Err(err), logger.Any("inventoryCheck", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDInventoryCheck.Err()
	}

	return &storeV1.GetInventoryCheckByIDReply{
		InventoryCheck: data,
	}, nil
}

// List of records by query parameters
func (h *inventoryCheckHandler) List(ctx context.Context, req *storeV1.ListInventoryCheckRequest) (*storeV1.ListInventoryCheckReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListInventoryCheck.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.inventoryCheckDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	inventoryChecks := []*storeV1.InventoryCheck{}
	for _, record := range records {
		data, err := convertInventoryCheckPb(record)
		if err != nil {
			logger.Warn("convertInventoryCheck error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		inventoryChecks = append(inventoryChecks, data)
	}

	return &storeV1.ListInventoryCheckReply{
		Total:           total,
		InventoryChecks: inventoryChecks,
	}, nil
}

func convertInventoryCheckPb(record *model.InventoryCheck) (*storeV1.InventoryCheck, error) {
	value := &storeV1.InventoryCheck{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID

	return value, nil
}
