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

var _ storeV1.ProductCategoryLogicer = (*productCategoryHandler)(nil)
var _ time.Time

type productCategoryHandler struct {
	productCategoryDao dao.ProductCategoryDao
}

// NewProductCategoryHandler create a handler
func NewProductCategoryHandler() storeV1.ProductCategoryLogicer {
	return &productCategoryHandler{
		productCategoryDao: dao.NewProductCategoryDao(
			database.GetDB(), // db driver is mysql
			cache.NewProductCategoryCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *productCategoryHandler) Create(ctx context.Context, req *storeV1.CreateProductCategoryRequest) (*storeV1.CreateProductCategoryReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	productCategory := &model.ProductCategory{}
	err = copier.Copy(productCategory, req)
	if err != nil {
		return nil, ecode.ErrCreateProductCategory.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.productCategoryDao.Create(ctx, productCategory)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("productCategory", productCategory), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateProductCategoryReply{Id: productCategory.ID}, nil
}

// DeleteByID delete a record by id
func (h *productCategoryHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteProductCategoryByIDRequest) (*storeV1.DeleteProductCategoryByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.productCategoryDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteProductCategoryByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *productCategoryHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateProductCategoryByIDRequest) (*storeV1.UpdateProductCategoryByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	productCategory := &model.ProductCategory{}
	err = copier.Copy(productCategory, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDProductCategory.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	productCategory.ID = req.Id

	err = h.productCategoryDao.UpdateByID(ctx, productCategory)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("productCategory", productCategory), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateProductCategoryByIDReply{}, nil
}

// GetByID get a record by id
func (h *productCategoryHandler) GetByID(ctx context.Context, req *storeV1.GetProductCategoryByIDRequest) (*storeV1.GetProductCategoryByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.productCategoryDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertProductCategory(record)
	if err != nil {
		logger.Warn("convertProductCategory error", logger.Err(err), logger.Any("productCategory", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDProductCategory.Err()
	}

	return &storeV1.GetProductCategoryByIDReply{
		ProductCategory: data,
	}, nil
}

// List of records by query parameters
func (h *productCategoryHandler) List(ctx context.Context, req *storeV1.ListProductCategoryRequest) (*storeV1.ListProductCategoryReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListProductCategory.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.productCategoryDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	productCategorys := []*storeV1.ProductCategory{}
	for _, record := range records {
		data, err := convertProductCategory(record)
		if err != nil {
			logger.Warn("convertProductCategory error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		productCategorys = append(productCategorys, data)
	}

	return &storeV1.ListProductCategoryReply{
		Total:            total,
		ProductCategorys: productCategorys,
	}, nil
}

func convertProductCategory(record *model.ProductCategory) (*storeV1.ProductCategory, error) {
	value := &storeV1.ProductCategory{}
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
