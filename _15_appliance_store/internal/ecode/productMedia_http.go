package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// productMedia business-level http error codes.
// the productMediaNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	productMediaNO       = 77
	productMediaName     = "productMedia"
	productMediaBaseCode = errcode.HCode(productMediaNO)

	ErrCreateProductMedia     = errcode.NewError(productMediaBaseCode+1, "failed to create "+productMediaName)
	ErrDeleteByIDProductMedia = errcode.NewError(productMediaBaseCode+2, "failed to delete "+productMediaName)
	ErrUpdateByIDProductMedia = errcode.NewError(productMediaBaseCode+3, "failed to update "+productMediaName)
	ErrGetByIDProductMedia    = errcode.NewError(productMediaBaseCode+4, "failed to get "+productMediaName+" details")
	ErrListProductMedia       = errcode.NewError(productMediaBaseCode+5, "failed to list of "+productMediaName)

	// error codes are globally unique, adding 1 to the previous error code
)
