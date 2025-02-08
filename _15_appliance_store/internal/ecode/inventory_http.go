package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// inventory business-level http error codes.
// the inventoryNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	inventoryNO       = 79
	inventoryName     = "inventory"
	inventoryBaseCode = errcode.HCode(inventoryNO)

	ErrCreateInventory          = errcode.NewError(inventoryBaseCode+1, "failed to create "+inventoryName)
	ErrDeleteByIDInventory      = errcode.NewError(inventoryBaseCode+2, "failed to delete "+inventoryName)
	ErrUpdateByIDInventory      = errcode.NewError(inventoryBaseCode+3, "failed to update "+inventoryName)
	ErrGetByIDInventory         = errcode.NewError(inventoryBaseCode+4, "failed to get "+inventoryName+" details")
	ErrListInventory            = errcode.NewError(inventoryBaseCode+5, "failed to list of "+inventoryName)
	ErrExecuteTransferInventory = errcode.NewError(inventoryBaseCode+6, "failed to ExecuteTransfer "+inventoryName)

	// error codes are globally unique, adding 1 to the previous error code
)
