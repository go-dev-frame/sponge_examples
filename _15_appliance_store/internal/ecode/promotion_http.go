package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// promotion business-level http error codes.
// the promotionNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	promotionNO       = 14
	promotionName     = "promotion"
	promotionBaseCode = errcode.HCode(promotionNO)

	ErrCreatePromotion             = errcode.NewError(promotionBaseCode+1, "failed to create "+promotionName)
	ErrDeleteByIDPromotion         = errcode.NewError(promotionBaseCode+2, "failed to delete "+promotionName)
	ErrUpdateByIDPromotion         = errcode.NewError(promotionBaseCode+3, "failed to update "+promotionName)
	ErrGetByIDPromotion            = errcode.NewError(promotionBaseCode+4, "failed to get "+promotionName+" details")
	ErrListPromotion               = errcode.NewError(promotionBaseCode+5, "failed to list of "+promotionName)
	ErrBindCouponTemplatePromotion = errcode.NewError(promotionBaseCode+6, "failed to BindCouponTemplate "+promotionName)

	// error codes are globally unique, adding 1 to the previous error code
)
