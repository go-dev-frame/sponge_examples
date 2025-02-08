package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// productCategory business-level http error codes.
// the productCategoryNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	productCategoryNO       = 9
	productCategoryName     = "productCategory"
	productCategoryBaseCode = errcode.HCode(productCategoryNO)

	ErrCreateProductCategory     = errcode.NewError(productCategoryBaseCode+1, "failed to create "+productCategoryName)
	ErrDeleteByIDProductCategory = errcode.NewError(productCategoryBaseCode+2, "failed to delete "+productCategoryName)
	ErrUpdateByIDProductCategory = errcode.NewError(productCategoryBaseCode+3, "failed to update "+productCategoryName)
	ErrGetByIDProductCategory    = errcode.NewError(productCategoryBaseCode+4, "failed to get "+productCategoryName+" details")
	ErrListProductCategory       = errcode.NewError(productCategoryBaseCode+5, "failed to list of "+productCategoryName)

	// error codes are globally unique, adding 1 to the previous error code
)
