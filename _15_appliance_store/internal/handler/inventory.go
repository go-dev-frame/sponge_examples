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

var _ storeV1.InventoryLogicer = (*inventoryHandler)(nil)
var _ time.Time

type inventoryHandler struct {
	inventoryDao dao.InventoryDao
}

// NewInventoryHandler create a handler
func NewInventoryHandler() storeV1.InventoryLogicer {
	return &inventoryHandler{
		inventoryDao: dao.NewInventoryDao(
			database.GetDB(), // db driver is mysql
			cache.NewInventoryCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *inventoryHandler) Create(ctx context.Context, req *storeV1.CreateInventoryRequest) (*storeV1.CreateInventoryReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	inventory := &model.Inventory{}
	err = copier.Copy(inventory, req)
	if err != nil {
		return nil, ecode.ErrCreateInventory.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.inventoryDao.Create(ctx, inventory)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("inventory", inventory), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateInventoryReply{Id: inventory.ID}, nil
}

// DeleteByID delete a record by id
func (h *inventoryHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteInventoryByIDRequest) (*storeV1.DeleteInventoryByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.inventoryDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteInventoryByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *inventoryHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateInventoryByIDRequest) (*storeV1.UpdateInventoryByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	inventory := &model.Inventory{}
	err = copier.Copy(inventory, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDInventory.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	inventory.ID = req.Id

	err = h.inventoryDao.UpdateByID(ctx, inventory)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("inventory", inventory), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateInventoryByIDReply{}, nil
}

// GetByID get a record by id
func (h *inventoryHandler) GetByID(ctx context.Context, req *storeV1.GetInventoryByIDRequest) (*storeV1.GetInventoryByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.inventoryDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertInventory(record)
	if err != nil {
		logger.Warn("convertInventory error", logger.Err(err), logger.Any("inventory", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDInventory.Err()
	}

	return &storeV1.GetInventoryByIDReply{
		Inventory: data,
	}, nil
}

// List of records by query parameters
func (h *inventoryHandler) List(ctx context.Context, req *storeV1.ListInventoryRequest) (*storeV1.ListInventoryReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListInventory.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.inventoryDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	inventorys := []*storeV1.Inventory{}
	for _, record := range records {
		data, err := convertInventory(record)
		if err != nil {
			logger.Warn("convertInventory error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		inventorys = append(inventorys, data)
	}

	return &storeV1.ListInventoryReply{
		Total:      total,
		Inventorys: inventorys,
	}, nil
}

// ExecuteTransfer 执行跨仓库调拨操作
// 实现逻辑：
// 1. 校验源仓库库存是否充足
// 2. 生成调拨单并锁定库存
// 3. 触发物流系统通知
// 4. 更新双方仓库库存记录
func (h *inventoryHandler) ExecuteTransfer(ctx context.Context, req *storeV1.ExecuteTransferRequest) (*storeV1.ExecuteTransferReply, error) {
	panic("prompt: 执行跨仓库调拨操作  实现逻辑：  1. 校验源仓库库存是否充足  2. 生成调拨单并锁定库存  3. 触发物流系统通知  4. 更新双方仓库库存记录")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.inventoryDao.ExecuteTransfer(ctx, &model.Inventory{
	//     	FromStoreID: req.FromStoreID,
	//     	ToStoreID: req.ToStoreID,
	//     	Items: req.Items,
	//     })
	//	    if err != nil {
	//			logger.Warn("ExecuteTransfer error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.ExecuteTransferReply{
	//     	TransferOrderID: reply.TransferOrderID,
	//     }, nil
}

func convertInventory(record *model.Inventory) (*storeV1.Inventory, error) {
	value := &storeV1.Inventory{}
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
