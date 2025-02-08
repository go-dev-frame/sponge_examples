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

var _ storeV1.TransferDetailLogicer = (*transferDetailHandler)(nil)
var _ time.Time

type transferDetailHandler struct {
	transferDetailDao dao.TransferDetailDao
}

// NewTransferDetailHandler create a handler
func NewTransferDetailHandler() storeV1.TransferDetailLogicer {
	return &transferDetailHandler{
		transferDetailDao: dao.NewTransferDetailDao(
			database.GetDB(), // db driver is mysql
			cache.NewTransferDetailCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *transferDetailHandler) Create(ctx context.Context, req *storeV1.CreateTransferDetailRequest) (*storeV1.CreateTransferDetailReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	transferDetail := &model.TransferDetail{}
	err = copier.Copy(transferDetail, req)
	if err != nil {
		return nil, ecode.ErrCreateTransferDetail.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.transferDetailDao.Create(ctx, transferDetail)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("transferDetail", transferDetail), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateTransferDetailReply{TransferID: transferDetail.TransferID}, nil
}

// DeleteByTransferID delete a record by transferID
func (h *transferDetailHandler) DeleteByTransferID(ctx context.Context, req *storeV1.DeleteTransferDetailByTransferIDRequest) (*storeV1.DeleteTransferDetailByTransferIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.transferDetailDao.DeleteByTransferID(ctx, req.TransferID)
	if err != nil {
		logger.Warn("DeleteByTransferID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteTransferDetailByTransferIDReply{}, nil
}

// UpdateByTransferID update a record by transferID
func (h *transferDetailHandler) UpdateByTransferID(ctx context.Context, req *storeV1.UpdateTransferDetailByTransferIDRequest) (*storeV1.UpdateTransferDetailByTransferIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	transferDetail := &model.TransferDetail{}
	err = copier.Copy(transferDetail, req)
	if err != nil {
		return nil, ecode.ErrUpdateByTransferIDTransferDetail.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	transferDetail.TransferID = req.TransferID

	err = h.transferDetailDao.UpdateByTransferID(ctx, transferDetail)
	if err != nil {
		logger.Error("UpdateByTransferID error", logger.Err(err), logger.Any("transferDetail", transferDetail), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateTransferDetailByTransferIDReply{}, nil
}

// GetByTransferID get a record by transferID
func (h *transferDetailHandler) GetByTransferID(ctx context.Context, req *storeV1.GetTransferDetailByTransferIDRequest) (*storeV1.GetTransferDetailByTransferIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.transferDetailDao.GetByTransferID(ctx, req.TransferID)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByTransferID error", logger.Err(err), logger.Any("transferID", req.TransferID), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByTransferID error", logger.Err(err), logger.Any("transferID", req.TransferID), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertTransferDetailPb(record)
	if err != nil {
		logger.Warn("convertTransferDetail error", logger.Err(err), logger.Any("transferDetail", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByTransferIDTransferDetail.Err()
	}

	return &storeV1.GetTransferDetailByTransferIDReply{
		TransferDetail: data,
	}, nil
}

// List of records by query parameters
func (h *transferDetailHandler) List(ctx context.Context, req *storeV1.ListTransferDetailRequest) (*storeV1.ListTransferDetailReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListTransferDetail.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.transferDetailDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	transferDetails := []*storeV1.TransferDetail{}
	for _, record := range records {
		data, err := convertTransferDetailPb(record)
		if err != nil {
			logger.Warn("convertTransferDetail error", logger.Err(err), logger.Any("transferID", record.TransferID), middleware.CtxRequestIDField(ctx))
			continue
		}
		transferDetails = append(transferDetails, data)
	}

	return &storeV1.ListTransferDetailReply{
		Total:           total,
		TransferDetails: transferDetails,
	}, nil
}

func convertTransferDetailPb(record *model.TransferDetail) (*storeV1.TransferDetail, error) {
	value := &storeV1.TransferDetail{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.TransferID = record.TransferID
	value.CreatedAt = utils.FormatDateTimeRFC3339(*record.CreatedAt)
	value.UpdatedAt = utils.FormatDateTimeRFC3339(*record.UpdatedAt)

	return value, nil
}
