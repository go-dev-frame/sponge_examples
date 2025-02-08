package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// supplier business-level http error codes.
// the supplierNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	supplierNO       = 8
	supplierName     = "supplier"
	supplierBaseCode = errcode.HCode(supplierNO)

	ErrCreateSupplier     = errcode.NewError(supplierBaseCode+1, "failed to create "+supplierName)
	ErrDeleteByIDSupplier = errcode.NewError(supplierBaseCode+2, "failed to delete "+supplierName)
	ErrUpdateByIDSupplier = errcode.NewError(supplierBaseCode+3, "failed to update "+supplierName)
	ErrGetByIDSupplier    = errcode.NewError(supplierBaseCode+4, "failed to get "+supplierName+" details")
	ErrListSupplier       = errcode.NewError(supplierBaseCode+5, "failed to list of "+supplierName)

	// error codes are globally unique, adding 1 to the previous error code
)
