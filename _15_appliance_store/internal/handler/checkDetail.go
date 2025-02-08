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

var _ storeV1.CheckDetailLogicer = (*checkDetailHandler)(nil)
var _ time.Time

type checkDetailHandler struct {
	checkDetailDao dao.CheckDetailDao
}

// NewCheckDetailHandler create a handler
func NewCheckDetailHandler() storeV1.CheckDetailLogicer {
	return &checkDetailHandler{
		checkDetailDao: dao.NewCheckDetailDao(
			database.GetDB(), // db driver is mysql
			cache.NewCheckDetailCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *checkDetailHandler) Create(ctx context.Context, req *storeV1.CreateCheckDetailRequest) (*storeV1.CreateCheckDetailReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	checkDetail := &model.CheckDetail{}
	err = copier.Copy(checkDetail, req)
	if err != nil {
		return nil, ecode.ErrCreateCheckDetail.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.checkDetailDao.Create(ctx, checkDetail)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("checkDetail", checkDetail), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateCheckDetailReply{CheckID: checkDetail.CheckID}, nil
}

// DeleteByCheckID delete a record by checkID
func (h *checkDetailHandler) DeleteByCheckID(ctx context.Context, req *storeV1.DeleteCheckDetailByCheckIDRequest) (*storeV1.DeleteCheckDetailByCheckIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.checkDetailDao.DeleteByCheckID(ctx, req.CheckID)
	if err != nil {
		logger.Warn("DeleteByCheckID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteCheckDetailByCheckIDReply{}, nil
}

// UpdateByCheckID update a record by checkID
func (h *checkDetailHandler) UpdateByCheckID(ctx context.Context, req *storeV1.UpdateCheckDetailByCheckIDRequest) (*storeV1.UpdateCheckDetailByCheckIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	checkDetail := &model.CheckDetail{}
	err = copier.Copy(checkDetail, req)
	if err != nil {
		return nil, ecode.ErrUpdateByCheckIDCheckDetail.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	checkDetail.CheckID = req.CheckID

	err = h.checkDetailDao.UpdateByCheckID(ctx, checkDetail)
	if err != nil {
		logger.Error("UpdateByCheckID error", logger.Err(err), logger.Any("checkDetail", checkDetail), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateCheckDetailByCheckIDReply{}, nil
}

// GetByCheckID get a record by checkID
func (h *checkDetailHandler) GetByCheckID(ctx context.Context, req *storeV1.GetCheckDetailByCheckIDRequest) (*storeV1.GetCheckDetailByCheckIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.checkDetailDao.GetByCheckID(ctx, req.CheckID)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByCheckID error", logger.Err(err), logger.Any("checkID", req.CheckID), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByCheckID error", logger.Err(err), logger.Any("checkID", req.CheckID), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertCheckDetailPb(record)
	if err != nil {
		logger.Warn("convertCheckDetail error", logger.Err(err), logger.Any("checkDetail", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByCheckIDCheckDetail.Err()
	}

	return &storeV1.GetCheckDetailByCheckIDReply{
		CheckDetail: data,
	}, nil
}

// List of records by query parameters
func (h *checkDetailHandler) List(ctx context.Context, req *storeV1.ListCheckDetailRequest) (*storeV1.ListCheckDetailReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListCheckDetail.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.checkDetailDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	checkDetails := []*storeV1.CheckDetail{}
	for _, record := range records {
		data, err := convertCheckDetailPb(record)
		if err != nil {
			logger.Warn("convertCheckDetail error", logger.Err(err), logger.Any("checkID", record.CheckID), middleware.CtxRequestIDField(ctx))
			continue
		}
		checkDetails = append(checkDetails, data)
	}

	return &storeV1.ListCheckDetailReply{
		Total:        total,
		CheckDetails: checkDetails,
	}, nil
}

func convertCheckDetailPb(record *model.CheckDetail) (*storeV1.CheckDetail, error) {
	value := &storeV1.CheckDetail{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.CheckID = record.CheckID

	return value, nil
}
