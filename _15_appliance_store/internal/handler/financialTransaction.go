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

var _ storeV1.FinancialTransactionLogicer = (*financialTransactionHandler)(nil)
var _ time.Time

type financialTransactionHandler struct {
	financialTransactionDao dao.FinancialTransactionDao
}

// NewFinancialTransactionHandler create a handler
func NewFinancialTransactionHandler() storeV1.FinancialTransactionLogicer {
	return &financialTransactionHandler{
		financialTransactionDao: dao.NewFinancialTransactionDao(
			database.GetDB(), // db driver is mysql
			cache.NewFinancialTransactionCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *financialTransactionHandler) Create(ctx context.Context, req *storeV1.CreateFinancialTransactionRequest) (*storeV1.CreateFinancialTransactionReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	financialTransaction := &model.FinancialTransaction{}
	err = copier.Copy(financialTransaction, req)
	if err != nil {
		return nil, ecode.ErrCreateFinancialTransaction.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.financialTransactionDao.Create(ctx, financialTransaction)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("financialTransaction", financialTransaction), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateFinancialTransactionReply{Id: financialTransaction.ID}, nil
}

// DeleteByID delete a record by id
func (h *financialTransactionHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteFinancialTransactionByIDRequest) (*storeV1.DeleteFinancialTransactionByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.financialTransactionDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteFinancialTransactionByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *financialTransactionHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateFinancialTransactionByIDRequest) (*storeV1.UpdateFinancialTransactionByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	financialTransaction := &model.FinancialTransaction{}
	err = copier.Copy(financialTransaction, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDFinancialTransaction.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	financialTransaction.ID = req.Id

	err = h.financialTransactionDao.UpdateByID(ctx, financialTransaction)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("financialTransaction", financialTransaction), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateFinancialTransactionByIDReply{}, nil
}

// GetByID get a record by id
func (h *financialTransactionHandler) GetByID(ctx context.Context, req *storeV1.GetFinancialTransactionByIDRequest) (*storeV1.GetFinancialTransactionByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.financialTransactionDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertFinancialTransaction(record)
	if err != nil {
		logger.Warn("convertFinancialTransaction error", logger.Err(err), logger.Any("financialTransaction", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDFinancialTransaction.Err()
	}

	return &storeV1.GetFinancialTransactionByIDReply{
		FinancialTransaction: data,
	}, nil
}

// List of records by query parameters
func (h *financialTransactionHandler) List(ctx context.Context, req *storeV1.ListFinancialTransactionRequest) (*storeV1.ListFinancialTransactionReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListFinancialTransaction.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.financialTransactionDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	financialTransactions := []*storeV1.FinancialTransaction{}
	for _, record := range records {
		data, err := convertFinancialTransaction(record)
		if err != nil {
			logger.Warn("convertFinancialTransaction error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		financialTransactions = append(financialTransactions, data)
	}

	return &storeV1.ListFinancialTransactionReply{
		Total:                 total,
		FinancialTransactions: financialTransactions,
	}, nil
}

func convertFinancialTransaction(record *model.FinancialTransaction) (*storeV1.FinancialTransaction, error) {
	value := &storeV1.FinancialTransaction{}
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
