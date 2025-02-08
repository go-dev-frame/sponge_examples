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

var _ storeV1.ProductSkuLogicer = (*productSkuHandler)(nil)
var _ time.Time

type productSkuHandler struct {
	productSkuDao dao.ProductSkuDao
}

// NewProductSkuHandler create a handler
func NewProductSkuHandler() storeV1.ProductSkuLogicer {
	return &productSkuHandler{
		productSkuDao: dao.NewProductSkuDao(
			database.GetDB(), // db driver is mysql
			cache.NewProductSkuCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *productSkuHandler) Create(ctx context.Context, req *storeV1.CreateProductSkuRequest) (*storeV1.CreateProductSkuReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	productSku := &model.ProductSku{}
	err = copier.Copy(productSku, req)
	if err != nil {
		return nil, ecode.ErrCreateProductSku.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.productSkuDao.Create(ctx, productSku)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("productSku", productSku), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateProductSkuReply{Id: productSku.ID}, nil
}

// DeleteByID delete a record by id
func (h *productSkuHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteProductSkuByIDRequest) (*storeV1.DeleteProductSkuByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.productSkuDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteProductSkuByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *productSkuHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateProductSkuByIDRequest) (*storeV1.UpdateProductSkuByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	productSku := &model.ProductSku{}
	err = copier.Copy(productSku, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDProductSku.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	productSku.ID = req.Id

	err = h.productSkuDao.UpdateByID(ctx, productSku)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("productSku", productSku), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateProductSkuByIDReply{}, nil
}

// GetByID get a record by id
func (h *productSkuHandler) GetByID(ctx context.Context, req *storeV1.GetProductSkuByIDRequest) (*storeV1.GetProductSkuByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.productSkuDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertProductSku(record)
	if err != nil {
		logger.Warn("convertProductSku error", logger.Err(err), logger.Any("productSku", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDProductSku.Err()
	}

	return &storeV1.GetProductSkuByIDReply{
		ProductSku: data,
	}, nil
}

// List of records by query parameters
func (h *productSkuHandler) List(ctx context.Context, req *storeV1.ListProductSkuRequest) (*storeV1.ListProductSkuReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListProductSku.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.productSkuDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	productSkus := []*storeV1.ProductSku{}
	for _, record := range records {
		data, err := convertProductSku(record)
		if err != nil {
			logger.Warn("convertProductSku error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		productSkus = append(productSkus, data)
	}

	return &storeV1.ListProductSkuReply{
		Total:       total,
		ProductSkus: productSkus,
	}, nil
}

func convertProductSku(record *model.ProductSku) (*storeV1.ProductSku, error) {
	value := &storeV1.ProductSku{}
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
