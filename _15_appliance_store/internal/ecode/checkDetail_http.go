package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// checkDetail business-level http error codes.
// the checkDetailNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	checkDetailNO       = 1
	checkDetailName     = "checkDetail"
	checkDetailBaseCode = errcode.HCode(checkDetailNO)

	ErrCreateCheckDetail          = errcode.NewError(checkDetailBaseCode+1, "failed to create "+checkDetailName)
	ErrDeleteByCheckIDCheckDetail = errcode.NewError(checkDetailBaseCode+2, "failed to delete "+checkDetailName)
	ErrUpdateByCheckIDCheckDetail = errcode.NewError(checkDetailBaseCode+3, "failed to update "+checkDetailName)
	ErrGetByCheckIDCheckDetail    = errcode.NewError(checkDetailBaseCode+4, "failed to get "+checkDetailName+" details")
	ErrListCheckDetail            = errcode.NewError(checkDetailBaseCode+5, "failed to list of "+checkDetailName)

	// error codes are globally unique, adding 1 to the previous error code
)
