package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// coupon business-level http error codes.
// the couponNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	couponNO       = 3
	couponName     = "coupon"
	couponBaseCode = errcode.HCode(couponNO)

	ErrCreateCoupon       = errcode.NewError(couponBaseCode+1, "failed to create "+couponName)
	ErrDeleteByIDCoupon   = errcode.NewError(couponBaseCode+2, "failed to delete "+couponName)
	ErrUpdateByIDCoupon   = errcode.NewError(couponBaseCode+3, "failed to update "+couponName)
	ErrGetByIDCoupon      = errcode.NewError(couponBaseCode+4, "failed to get "+couponName+" details")
	ErrListCoupon         = errcode.NewError(couponBaseCode+5, "failed to list of "+couponName)
	ErrRedeemCouponCoupon = errcode.NewError(couponBaseCode+6, "failed to RedeemCoupon "+couponName)

	// error codes are globally unique, adding 1 to the previous error code
)
