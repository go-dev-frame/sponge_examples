package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// purchaseOrder business-level http error codes.
// the purchaseOrderNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	purchaseOrderNO       = 5
	purchaseOrderName     = "purchaseOrder"
	purchaseOrderBaseCode = errcode.HCode(purchaseOrderNO)

	ErrCreatePurchaseOrder     = errcode.NewError(purchaseOrderBaseCode+1, "failed to create "+purchaseOrderName)
	ErrDeleteByIDPurchaseOrder = errcode.NewError(purchaseOrderBaseCode+2, "failed to delete "+purchaseOrderName)
	ErrUpdateByIDPurchaseOrder = errcode.NewError(purchaseOrderBaseCode+3, "failed to update "+purchaseOrderName)
	ErrGetByIDPurchaseOrder    = errcode.NewError(purchaseOrderBaseCode+4, "failed to get "+purchaseOrderName+" details")
	ErrListPurchaseOrder       = errcode.NewError(purchaseOrderBaseCode+5, "failed to list of "+purchaseOrderName)

	// error codes are globally unique, adding 1 to the previous error code
)
