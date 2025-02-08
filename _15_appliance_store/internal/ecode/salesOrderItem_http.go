package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// salesOrderItem business-level http error codes.
// the salesOrderItemNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	salesOrderItemNO       = 73
	salesOrderItemName     = "salesOrderItem"
	salesOrderItemBaseCode = errcode.HCode(salesOrderItemNO)

	ErrCreateSalesOrderItem     = errcode.NewError(salesOrderItemBaseCode+1, "failed to create "+salesOrderItemName)
	ErrDeleteByIDSalesOrderItem = errcode.NewError(salesOrderItemBaseCode+2, "failed to delete "+salesOrderItemName)
	ErrUpdateByIDSalesOrderItem = errcode.NewError(salesOrderItemBaseCode+3, "failed to update "+salesOrderItemName)
	ErrGetByIDSalesOrderItem    = errcode.NewError(salesOrderItemBaseCode+4, "failed to get "+salesOrderItemName+" details")
	ErrListSalesOrderItem       = errcode.NewError(salesOrderItemBaseCode+5, "failed to list of "+salesOrderItemName)

	// error codes are globally unique, adding 1 to the previous error code
)
