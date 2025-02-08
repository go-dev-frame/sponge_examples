package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// productSku business-level http error codes.
// the productSkuNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	productSkuNO       = 49
	productSkuName     = "productSku"
	productSkuBaseCode = errcode.HCode(productSkuNO)

	ErrCreateProductSku     = errcode.NewError(productSkuBaseCode+1, "failed to create "+productSkuName)
	ErrDeleteByIDProductSku = errcode.NewError(productSkuBaseCode+2, "failed to delete "+productSkuName)
	ErrUpdateByIDProductSku = errcode.NewError(productSkuBaseCode+3, "failed to update "+productSkuName)
	ErrGetByIDProductSku    = errcode.NewError(productSkuBaseCode+4, "failed to get "+productSkuName+" details")
	ErrListProductSku       = errcode.NewError(productSkuBaseCode+5, "failed to list of "+productSkuName)

	// error codes are globally unique, adding 1 to the previous error code
)
