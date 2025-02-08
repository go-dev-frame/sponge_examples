package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// transferOrder business-level http error codes.
// the transferOrderNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	transferOrderNO       = 11
	transferOrderName     = "transferOrder"
	transferOrderBaseCode = errcode.HCode(transferOrderNO)

	ErrCreateTransferOrder           = errcode.NewError(transferOrderBaseCode+1, "failed to create "+transferOrderName)
	ErrDeleteByIDTransferOrder       = errcode.NewError(transferOrderBaseCode+2, "failed to delete "+transferOrderName)
	ErrUpdateByIDTransferOrder       = errcode.NewError(transferOrderBaseCode+3, "failed to update "+transferOrderName)
	ErrGetByIDTransferOrder          = errcode.NewError(transferOrderBaseCode+4, "failed to get "+transferOrderName+" details")
	ErrListTransferOrder             = errcode.NewError(transferOrderBaseCode+5, "failed to list of "+transferOrderName)
	ErrPrecheckTransferTransferOrder = errcode.NewError(transferOrderBaseCode+6, "failed to PrecheckTransfer "+transferOrderName)

	// error codes are globally unique, adding 1 to the previous error code
)
