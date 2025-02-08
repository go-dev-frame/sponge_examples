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

var _ storeV1.CouponLogicer = (*couponHandler)(nil)
var _ time.Time

type couponHandler struct {
	couponDao dao.CouponDao
}

// NewCouponHandler create a handler
func NewCouponHandler() storeV1.CouponLogicer {
	return &couponHandler{
		couponDao: dao.NewCouponDao(
			database.GetDB(), // db driver is mysql
			cache.NewCouponCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *couponHandler) Create(ctx context.Context, req *storeV1.CreateCouponRequest) (*storeV1.CreateCouponReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	coupon := &model.Coupon{}
	err = copier.Copy(coupon, req)
	if err != nil {
		return nil, ecode.ErrCreateCoupon.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.couponDao.Create(ctx, coupon)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("coupon", coupon), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateCouponReply{Id: coupon.ID}, nil
}

// DeleteByID delete a record by id
func (h *couponHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteCouponByIDRequest) (*storeV1.DeleteCouponByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.couponDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteCouponByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *couponHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateCouponByIDRequest) (*storeV1.UpdateCouponByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	coupon := &model.Coupon{}
	err = copier.Copy(coupon, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDCoupon.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	coupon.ID = req.Id

	err = h.couponDao.UpdateByID(ctx, coupon)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("coupon", coupon), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateCouponByIDReply{}, nil
}

// GetByID get a record by id
func (h *couponHandler) GetByID(ctx context.Context, req *storeV1.GetCouponByIDRequest) (*storeV1.GetCouponByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.couponDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertCouponPb(record)
	if err != nil {
		logger.Warn("convertCoupon error", logger.Err(err), logger.Any("coupon", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDCoupon.Err()
	}

	return &storeV1.GetCouponByIDReply{
		Coupon: data,
	}, nil
}

// List of records by query parameters
func (h *couponHandler) List(ctx context.Context, req *storeV1.ListCouponRequest) (*storeV1.ListCouponReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListCoupon.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.couponDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	coupons := []*storeV1.Coupon{}
	for _, record := range records {
		data, err := convertCouponPb(record)
		if err != nil {
			logger.Warn("convertCoupon error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		coupons = append(coupons, data)
	}

	return &storeV1.ListCouponReply{
		Total:   total,
		Coupons: coupons,
	}, nil
}

// RedeemCoupon 优惠券核销
func (h *couponHandler) RedeemCoupon(ctx context.Context, req *storeV1.RedeemCouponRequest) (*storeV1.RedeemCouponReply, error) {
	panic("prompt: 优惠券核销")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.couponDao.RedeemCoupon(ctx, &model.Coupon{
	//     	CouponID: req.CouponID,
	//     	StoreID: req.StoreID,
	//     	OrderID: req.OrderID,
	//     	OperatorID: req.OperatorID,
	//     })
	//	    if err != nil {
	//			logger.Warn("RedeemCoupon error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.RedeemCouponReply{
	//     	TransactionID: reply.TransactionID,
	//     	RedeemedTime: reply.RedeemedTime,
	//     }, nil
}

func convertCouponPb(record *model.Coupon) (*storeV1.Coupon, error) {
	value := &storeV1.Coupon{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID

	return value, nil
}
