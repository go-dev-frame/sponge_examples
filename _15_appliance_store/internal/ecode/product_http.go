package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// product business-level http error codes.
// the productNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	productNO       = 2
	productName     = "product"
	productBaseCode = errcode.HCode(productNO)

	ErrCreateProduct     = errcode.NewError(productBaseCode+1, "failed to create "+productName)
	ErrDeleteByIDProduct = errcode.NewError(productBaseCode+2, "failed to delete "+productName)
	ErrUpdateByIDProduct = errcode.NewError(productBaseCode+3, "failed to update "+productName)
	ErrGetByIDProduct    = errcode.NewError(productBaseCode+4, "failed to get "+productName+" details")
	ErrListProduct       = errcode.NewError(productBaseCode+5, "failed to list of "+productName)

	// error codes are globally unique, adding 1 to the previous error code
)
