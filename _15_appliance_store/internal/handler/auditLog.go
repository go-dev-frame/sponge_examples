package handler

import (
	"context"
	"errors"
	"github.com/go-dev-frame/sponge/pkg/utils"
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

var _ storeV1.AuditLogLogicer = (*auditLogHandler)(nil)
var _ time.Time

type auditLogHandler struct {
	auditLogDao dao.AuditLogDao
}

// NewAuditLogHandler create a handler
func NewAuditLogHandler() storeV1.AuditLogLogicer {
	return &auditLogHandler{
		auditLogDao: dao.NewAuditLogDao(
			database.GetDB(), // db driver is mysql
			cache.NewAuditLogCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *auditLogHandler) Create(ctx context.Context, req *storeV1.CreateAuditLogRequest) (*storeV1.CreateAuditLogReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	auditLog := &model.AuditLog{}
	err = copier.Copy(auditLog, req)
	if err != nil {
		return nil, ecode.ErrCreateAuditLog.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.auditLogDao.Create(ctx, auditLog)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("auditLog", auditLog), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateAuditLogReply{Id: auditLog.ID}, nil
}

// DeleteByID delete a record by id
func (h *auditLogHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteAuditLogByIDRequest) (*storeV1.DeleteAuditLogByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.auditLogDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteAuditLogByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *auditLogHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateAuditLogByIDRequest) (*storeV1.UpdateAuditLogByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	auditLog := &model.AuditLog{}
	err = copier.Copy(auditLog, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDAuditLog.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	auditLog.ID = req.Id

	err = h.auditLogDao.UpdateByID(ctx, auditLog)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("auditLog", auditLog), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateAuditLogByIDReply{}, nil
}

// GetByID get a record by id
func (h *auditLogHandler) GetByID(ctx context.Context, req *storeV1.GetAuditLogByIDRequest) (*storeV1.GetAuditLogByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.auditLogDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertAuditLog(record)
	if err != nil {
		logger.Warn("convertAuditLog error", logger.Err(err), logger.Any("auditLog", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDAuditLog.Err()
	}

	return &storeV1.GetAuditLogByIDReply{
		AuditLog: data,
	}, nil
}

// List of records by query parameters
func (h *auditLogHandler) List(ctx context.Context, req *storeV1.ListAuditLogRequest) (*storeV1.ListAuditLogReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListAuditLog.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.auditLogDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	auditLogs := []*storeV1.AuditLog{}
	for _, record := range records {
		data, err := convertAuditLog(record)
		if err != nil {
			logger.Warn("convertAuditLog error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		auditLogs = append(auditLogs, data)
	}

	return &storeV1.ListAuditLogReply{
		Total:     total,
		AuditLogs: auditLogs,
	}, nil
}

func convertAuditLog(record *model.AuditLog) (*storeV1.AuditLog, error) {
	value := &storeV1.AuditLog{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID
	value.CreatedAt = utils.FormatDateTimeRFC3339(*record.CreatedAt)

	return value, nil
}
