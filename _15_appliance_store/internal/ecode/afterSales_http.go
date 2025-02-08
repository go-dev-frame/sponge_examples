package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// afterSales business-level http error codes.
// the afterSalesNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	afterSalesNO       = 86
	afterSalesName     = "afterSales"
	afterSalesBaseCode = errcode.HCode(afterSalesNO)

	ErrCreateAfterSales                           = errcode.NewError(afterSalesBaseCode+1, "failed to create "+afterSalesName)
	ErrDeleteByIDAfterSales                       = errcode.NewError(afterSalesBaseCode+2, "failed to delete "+afterSalesName)
	ErrUpdateByIDAfterSales                       = errcode.NewError(afterSalesBaseCode+3, "failed to update "+afterSalesName)
	ErrGetByIDAfterSales                          = errcode.NewError(afterSalesBaseCode+4, "failed to get "+afterSalesName+" details")
	ErrListAfterSales                             = errcode.NewError(afterSalesBaseCode+5, "failed to list of "+afterSalesName)
	ErrCreateServiceOrderWithAssignmentAfterSales = errcode.NewError(afterSalesBaseCode+6, "failed to CreateServiceOrderWithAssignment "+afterSalesName)

	// error codes are globally unique, adding 1 to the previous error code
)
