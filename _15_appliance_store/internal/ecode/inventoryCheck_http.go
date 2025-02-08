package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// inventoryCheck business-level http error codes.
// the inventoryCheckNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	inventoryCheckNO       = 4
	inventoryCheckName     = "inventoryCheck"
	inventoryCheckBaseCode = errcode.HCode(inventoryCheckNO)

	ErrCreateInventoryCheck     = errcode.NewError(inventoryCheckBaseCode+1, "failed to create "+inventoryCheckName)
	ErrDeleteByIDInventoryCheck = errcode.NewError(inventoryCheckBaseCode+2, "failed to delete "+inventoryCheckName)
	ErrUpdateByIDInventoryCheck = errcode.NewError(inventoryCheckBaseCode+3, "failed to update "+inventoryCheckName)
	ErrGetByIDInventoryCheck    = errcode.NewError(inventoryCheckBaseCode+4, "failed to get "+inventoryCheckName+" details")
	ErrListInventoryCheck       = errcode.NewError(inventoryCheckBaseCode+5, "failed to list of "+inventoryCheckName)

	// error codes are globally unique, adding 1 to the previous error code
)
