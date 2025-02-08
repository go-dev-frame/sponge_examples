package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// inventoryOperation business-level http error codes.
// the inventoryOperationNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	inventoryOperationNO       = 68
	inventoryOperationName     = "inventoryOperation"
	inventoryOperationBaseCode = errcode.HCode(inventoryOperationNO)

	ErrCreateInventoryOperation     = errcode.NewError(inventoryOperationBaseCode+1, "failed to create "+inventoryOperationName)
	ErrDeleteByIDInventoryOperation = errcode.NewError(inventoryOperationBaseCode+2, "failed to delete "+inventoryOperationName)
	ErrUpdateByIDInventoryOperation = errcode.NewError(inventoryOperationBaseCode+3, "failed to update "+inventoryOperationName)
	ErrGetByIDInventoryOperation    = errcode.NewError(inventoryOperationBaseCode+4, "failed to get "+inventoryOperationName+" details")
	ErrListInventoryOperation       = errcode.NewError(inventoryOperationBaseCode+5, "failed to list of "+inventoryOperationName)

	// error codes are globally unique, adding 1 to the previous error code
)
