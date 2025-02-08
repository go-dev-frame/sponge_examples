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

var _ storeV1.ProductMediaLogicer = (*productMediaHandler)(nil)
var _ time.Time

type productMediaHandler struct {
	productMediaDao dao.ProductMediaDao
}

// NewProductMediaHandler create a handler
func NewProductMediaHandler() storeV1.ProductMediaLogicer {
	return &productMediaHandler{
		productMediaDao: dao.NewProductMediaDao(
			database.GetDB(), // db driver is mysql
			cache.NewProductMediaCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *productMediaHandler) Create(ctx context.Context, req *storeV1.CreateProductMediaRequest) (*storeV1.CreateProductMediaReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	productMedia := &model.ProductMedia{}
	err = copier.Copy(productMedia, req)
	if err != nil {
		return nil, ecode.ErrCreateProductMedia.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.productMediaDao.Create(ctx, productMedia)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("productMedia", productMedia), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateProductMediaReply{Id: productMedia.ID}, nil
}

// DeleteByID delete a record by id
func (h *productMediaHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteProductMediaByIDRequest) (*storeV1.DeleteProductMediaByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.productMediaDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteProductMediaByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *productMediaHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateProductMediaByIDRequest) (*storeV1.UpdateProductMediaByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	productMedia := &model.ProductMedia{}
	err = copier.Copy(productMedia, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDProductMedia.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	productMedia.ID = req.Id

	err = h.productMediaDao.UpdateByID(ctx, productMedia)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("productMedia", productMedia), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateProductMediaByIDReply{}, nil
}

// GetByID get a record by id
func (h *productMediaHandler) GetByID(ctx context.Context, req *storeV1.GetProductMediaByIDRequest) (*storeV1.GetProductMediaByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.productMediaDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertProductMedia(record)
	if err != nil {
		logger.Warn("convertProductMedia error", logger.Err(err), logger.Any("productMedia", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDProductMedia.Err()
	}

	return &storeV1.GetProductMediaByIDReply{
		ProductMedia: data,
	}, nil
}

// List of records by query parameters
func (h *productMediaHandler) List(ctx context.Context, req *storeV1.ListProductMediaRequest) (*storeV1.ListProductMediaReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListProductMedia.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.productMediaDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	productMedias := []*storeV1.ProductMedia{}
	for _, record := range records {
		data, err := convertProductMedia(record)
		if err != nil {
			logger.Warn("convertProductMedia error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		productMedias = append(productMedias, data)
	}

	return &storeV1.ListProductMediaReply{
		Total:         total,
		ProductMedias: productMedias,
	}, nil
}

func convertProductMedia(record *model.ProductMedia) (*storeV1.ProductMedia, error) {
	value := &storeV1.ProductMedia{}
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
