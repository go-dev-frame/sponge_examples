package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// salesOrder business-level http error codes.
// the salesOrderNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	salesOrderNO       = 6
	salesOrderName     = "salesOrder"
	salesOrderBaseCode = errcode.HCode(salesOrderNO)

	ErrCreateSalesOrder                 = errcode.NewError(salesOrderBaseCode+1, "failed to create "+salesOrderName)
	ErrDeleteByIDSalesOrder             = errcode.NewError(salesOrderBaseCode+2, "failed to delete "+salesOrderName)
	ErrUpdateByIDSalesOrder             = errcode.NewError(salesOrderBaseCode+3, "failed to update "+salesOrderName)
	ErrGetByIDSalesOrder                = errcode.NewError(salesOrderBaseCode+4, "failed to get "+salesOrderName+" details")
	ErrListSalesOrder                   = errcode.NewError(salesOrderBaseCode+5, "failed to list of "+salesOrderName)
	ErrGenerateHotSalesReportSalesOrder = errcode.NewError(salesOrderBaseCode+6, "failed to GenerateHotSalesReport "+salesOrderName)

	// error codes are globally unique, adding 1 to the previous error code
)
