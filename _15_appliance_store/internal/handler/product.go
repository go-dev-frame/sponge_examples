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

var _ storeV1.ProductLogicer = (*productHandler)(nil)
var _ time.Time

type productHandler struct {
	productDao dao.ProductDao
}

// NewProductHandler create a handler
func NewProductHandler() storeV1.ProductLogicer {
	return &productHandler{
		productDao: dao.NewProductDao(
			database.GetDB(), // db driver is mysql
			cache.NewProductCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *productHandler) Create(ctx context.Context, req *storeV1.CreateProductRequest) (*storeV1.CreateProductReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	product := &model.Product{}
	err = copier.Copy(product, req)
	if err != nil {
		return nil, ecode.ErrCreateProduct.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.productDao.Create(ctx, product)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("product", product), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateProductReply{Id: product.ID}, nil
}

// DeleteByID delete a record by id
func (h *productHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteProductByIDRequest) (*storeV1.DeleteProductByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.productDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteProductByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *productHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateProductByIDRequest) (*storeV1.UpdateProductByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	product := &model.Product{}
	err = copier.Copy(product, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDProduct.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	product.ID = req.Id

	err = h.productDao.UpdateByID(ctx, product)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("product", product), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateProductByIDReply{}, nil
}

// GetByID get a record by id
func (h *productHandler) GetByID(ctx context.Context, req *storeV1.GetProductByIDRequest) (*storeV1.GetProductByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.productDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertProduct(record)
	if err != nil {
		logger.Warn("convertProduct error", logger.Err(err), logger.Any("product", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDProduct.Err()
	}

	return &storeV1.GetProductByIDReply{
		Product: data,
	}, nil
}

// List of records by query parameters
func (h *productHandler) List(ctx context.Context, req *storeV1.ListProductRequest) (*storeV1.ListProductReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListProduct.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.productDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	products := []*storeV1.Product{}
	for _, record := range records {
		data, err := convertProduct(record)
		if err != nil {
			logger.Warn("convertProduct error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		products = append(products, data)
	}

	return &storeV1.ListProductReply{
		Total:    total,
		Products: products,
	}, nil
}

func convertProduct(record *model.Product) (*storeV1.Product, error) {
	value := &storeV1.Product{}
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
