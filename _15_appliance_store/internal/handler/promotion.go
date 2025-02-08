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

var _ storeV1.PromotionLogicer = (*promotionHandler)(nil)
var _ time.Time

type promotionHandler struct {
	promotionDao dao.PromotionDao
}

// NewPromotionHandler create a handler
func NewPromotionHandler() storeV1.PromotionLogicer {
	return &promotionHandler{
		promotionDao: dao.NewPromotionDao(
			database.GetDB(), // db driver is mysql
			cache.NewPromotionCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *promotionHandler) Create(ctx context.Context, req *storeV1.CreatePromotionRequest) (*storeV1.CreatePromotionReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	promotion := &model.Promotion{}
	err = copier.Copy(promotion, req)
	if err != nil {
		return nil, ecode.ErrCreatePromotion.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.promotionDao.Create(ctx, promotion)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("promotion", promotion), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreatePromotionReply{Id: promotion.ID}, nil
}

// DeleteByID delete a record by id
func (h *promotionHandler) DeleteByID(ctx context.Context, req *storeV1.DeletePromotionByIDRequest) (*storeV1.DeletePromotionByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.promotionDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeletePromotionByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *promotionHandler) UpdateByID(ctx context.Context, req *storeV1.UpdatePromotionByIDRequest) (*storeV1.UpdatePromotionByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	promotion := &model.Promotion{}
	err = copier.Copy(promotion, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDPromotion.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	promotion.ID = req.Id

	err = h.promotionDao.UpdateByID(ctx, promotion)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("promotion", promotion), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdatePromotionByIDReply{}, nil
}

// GetByID get a record by id
func (h *promotionHandler) GetByID(ctx context.Context, req *storeV1.GetPromotionByIDRequest) (*storeV1.GetPromotionByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.promotionDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertPromotion(record)
	if err != nil {
		logger.Warn("convertPromotion error", logger.Err(err), logger.Any("promotion", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDPromotion.Err()
	}

	return &storeV1.GetPromotionByIDReply{
		Promotion: data,
	}, nil
}

// List of records by query parameters
func (h *promotionHandler) List(ctx context.Context, req *storeV1.ListPromotionRequest) (*storeV1.ListPromotionReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListPromotion.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.promotionDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	promotions := []*storeV1.Promotion{}
	for _, record := range records {
		data, err := convertPromotion(record)
		if err != nil {
			logger.Warn("convertPromotion error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		promotions = append(promotions, data)
	}

	return &storeV1.ListPromotionReply{
		Total:      total,
		Promotions: promotions,
	}, nil
}

// BindCouponTemplate 创建促销活动与优惠券绑定关系
// 实现逻辑：
// 1. 验证促销活动是否存在且处于可编辑状态
// 2. 检查优惠券模板是否有效（类型匹配、有效期覆盖）
// 3. 建立促销活动与优惠券的绑定关系
func (h *promotionHandler) BindCouponTemplate(ctx context.Context, req *storeV1.BindCouponTemplateRequest) (*storeV1.BindCouponTemplateReply, error) {
	panic("prompt: 创建促销活动与优惠券绑定关系  实现逻辑：  1. 验证促销活动是否存在且处于可编辑状态  2. 检查优惠券模板是否有效（类型匹配、有效期覆盖）  3. 建立促销活动与优惠券的绑定关系")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.promotionDao.BindCouponTemplate(ctx, &model.Promotion{
	//     	PromotionID: req.PromotionID,
	//     	CouponIDs: req.CouponIDs,
	//     })
	//	    if err != nil {
	//			logger.Warn("BindCouponTemplate error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.BindCouponTemplateReply{
	//     	BoundCount: reply.BoundCount,
	//     }, nil
}

func convertPromotion(record *model.Promotion) (*storeV1.Promotion, error) {
	value := &storeV1.Promotion{}
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
