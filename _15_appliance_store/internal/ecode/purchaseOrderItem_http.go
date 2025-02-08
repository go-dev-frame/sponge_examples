package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// purchaseOrderItem business-level http error codes.
// the purchaseOrderItemNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	purchaseOrderItemNO       = 59
	purchaseOrderItemName     = "purchaseOrderItem"
	purchaseOrderItemBaseCode = errcode.HCode(purchaseOrderItemNO)

	ErrCreatePurchaseOrderItem     = errcode.NewError(purchaseOrderItemBaseCode+1, "failed to create "+purchaseOrderItemName)
	ErrDeleteByIDPurchaseOrderItem = errcode.NewError(purchaseOrderItemBaseCode+2, "failed to delete "+purchaseOrderItemName)
	ErrUpdateByIDPurchaseOrderItem = errcode.NewError(purchaseOrderItemBaseCode+3, "failed to update "+purchaseOrderItemName)
	ErrGetByIDPurchaseOrderItem    = errcode.NewError(purchaseOrderItemBaseCode+4, "failed to get "+purchaseOrderItemName+" details")
	ErrListPurchaseOrderItem       = errcode.NewError(purchaseOrderItemBaseCode+5, "failed to list of "+purchaseOrderItemName)

	// error codes are globally unique, adding 1 to the previous error code
)
