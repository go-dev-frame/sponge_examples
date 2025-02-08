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

var _ storeV1.AccountPayableLogicer = (*accountPayableHandler)(nil)
var _ time.Time

type accountPayableHandler struct {
	accountPayableDao dao.AccountPayableDao
}

// NewAccountPayableHandler create a handler
func NewAccountPayableHandler() storeV1.AccountPayableLogicer {
	return &accountPayableHandler{
		accountPayableDao: dao.NewAccountPayableDao(
			database.GetDB(), // db driver is mysql
			cache.NewAccountPayableCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *accountPayableHandler) Create(ctx context.Context, req *storeV1.CreateAccountPayableRequest) (*storeV1.CreateAccountPayableReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	accountPayable := &model.AccountPayable{}
	err = copier.Copy(accountPayable, req)
	if err != nil {
		return nil, ecode.ErrCreateAccountPayable.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.accountPayableDao.Create(ctx, accountPayable)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("accountPayable", accountPayable), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateAccountPayableReply{Id: accountPayable.ID}, nil
}

// DeleteByID delete a record by id
func (h *accountPayableHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteAccountPayableByIDRequest) (*storeV1.DeleteAccountPayableByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.accountPayableDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteAccountPayableByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *accountPayableHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateAccountPayableByIDRequest) (*storeV1.UpdateAccountPayableByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	accountPayable := &model.AccountPayable{}
	err = copier.Copy(accountPayable, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDAccountPayable.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	accountPayable.ID = req.Id

	err = h.accountPayableDao.UpdateByID(ctx, accountPayable)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("accountPayable", accountPayable), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateAccountPayableByIDReply{}, nil
}

// GetByID get a record by id
func (h *accountPayableHandler) GetByID(ctx context.Context, req *storeV1.GetAccountPayableByIDRequest) (*storeV1.GetAccountPayableByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.accountPayableDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertAccountPayable(record)
	if err != nil {
		logger.Warn("convertAccountPayable error", logger.Err(err), logger.Any("accountPayable", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDAccountPayable.Err()
	}

	return &storeV1.GetAccountPayableByIDReply{
		AccountPayable: data,
	}, nil
}

// List of records by query parameters
func (h *accountPayableHandler) List(ctx context.Context, req *storeV1.ListAccountPayableRequest) (*storeV1.ListAccountPayableReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListAccountPayable.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.accountPayableDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	accountPayables := []*storeV1.AccountPayable{}
	for _, record := range records {
		data, err := convertAccountPayable(record)
		if err != nil {
			logger.Warn("convertAccountPayable error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		accountPayables = append(accountPayables, data)
	}

	return &storeV1.ListAccountPayableReply{
		Total:           total,
		AccountPayables: accountPayables,
	}, nil
}

func convertAccountPayable(record *model.AccountPayable) (*storeV1.AccountPayable, error) {
	value := &storeV1.AccountPayable{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID

	return value, nil
}
