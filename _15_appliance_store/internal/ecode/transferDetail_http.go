package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// transferDetail business-level http error codes.
// the transferDetailNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	transferDetailNO       = 10
	transferDetailName     = "transferDetail"
	transferDetailBaseCode = errcode.HCode(transferDetailNO)

	ErrCreateTransferDetail             = errcode.NewError(transferDetailBaseCode+1, "failed to create "+transferDetailName)
	ErrDeleteByTransferIDTransferDetail = errcode.NewError(transferDetailBaseCode+2, "failed to delete "+transferDetailName)
	ErrUpdateByTransferIDTransferDetail = errcode.NewError(transferDetailBaseCode+3, "failed to update "+transferDetailName)
	ErrGetByTransferIDTransferDetail    = errcode.NewError(transferDetailBaseCode+4, "failed to get "+transferDetailName+" details")
	ErrListTransferDetail               = errcode.NewError(transferDetailBaseCode+5, "failed to list of "+transferDetailName)

	// error codes are globally unique, adding 1 to the previous error code
)
