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

var _ storeV1.InventoryOperationLogicer = (*inventoryOperationHandler)(nil)
var _ time.Time

type inventoryOperationHandler struct {
	inventoryOperationDao dao.InventoryOperationDao
}

// NewInventoryOperationHandler create a handler
func NewInventoryOperationHandler() storeV1.InventoryOperationLogicer {
	return &inventoryOperationHandler{
		inventoryOperationDao: dao.NewInventoryOperationDao(
			database.GetDB(), // db driver is mysql
			cache.NewInventoryOperationCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *inventoryOperationHandler) Create(ctx context.Context, req *storeV1.CreateInventoryOperationRequest) (*storeV1.CreateInventoryOperationReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	inventoryOperation := &model.InventoryOperation{}
	err = copier.Copy(inventoryOperation, req)
	if err != nil {
		return nil, ecode.ErrCreateInventoryOperation.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.inventoryOperationDao.Create(ctx, inventoryOperation)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("inventoryOperation", inventoryOperation), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateInventoryOperationReply{Id: inventoryOperation.ID}, nil
}

// DeleteByID delete a record by id
func (h *inventoryOperationHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteInventoryOperationByIDRequest) (*storeV1.DeleteInventoryOperationByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.inventoryOperationDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteInventoryOperationByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *inventoryOperationHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateInventoryOperationByIDRequest) (*storeV1.UpdateInventoryOperationByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	inventoryOperation := &model.InventoryOperation{}
	err = copier.Copy(inventoryOperation, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDInventoryOperation.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	inventoryOperation.ID = req.Id

	err = h.inventoryOperationDao.UpdateByID(ctx, inventoryOperation)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("inventoryOperation", inventoryOperation), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateInventoryOperationByIDReply{}, nil
}

// GetByID get a record by id
func (h *inventoryOperationHandler) GetByID(ctx context.Context, req *storeV1.GetInventoryOperationByIDRequest) (*storeV1.GetInventoryOperationByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.inventoryOperationDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertInventoryOperation(record)
	if err != nil {
		logger.Warn("convertInventoryOperation error", logger.Err(err), logger.Any("inventoryOperation", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDInventoryOperation.Err()
	}

	return &storeV1.GetInventoryOperationByIDReply{
		InventoryOperation: data,
	}, nil
}

// List of records by query parameters
func (h *inventoryOperationHandler) List(ctx context.Context, req *storeV1.ListInventoryOperationRequest) (*storeV1.ListInventoryOperationReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListInventoryOperation.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.inventoryOperationDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	inventoryOperations := []*storeV1.InventoryOperation{}
	for _, record := range records {
		data, err := convertInventoryOperation(record)
		if err != nil {
			logger.Warn("convertInventoryOperation error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		inventoryOperations = append(inventoryOperations, data)
	}

	return &storeV1.ListInventoryOperationReply{
		Total:               total,
		InventoryOperations: inventoryOperations,
	}, nil
}

func convertInventoryOperation(record *model.InventoryOperation) (*storeV1.InventoryOperation, error) {
	value := &storeV1.InventoryOperation{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID

	return value, nil
}
