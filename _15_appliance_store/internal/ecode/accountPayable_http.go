package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// accountPayable business-level http error codes.
// the accountPayableNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	accountPayableNO       = 43
	accountPayableName     = "accountPayable"
	accountPayableBaseCode = errcode.HCode(accountPayableNO)

	ErrCreateAccountPayable     = errcode.NewError(accountPayableBaseCode+1, "failed to create "+accountPayableName)
	ErrDeleteByIDAccountPayable = errcode.NewError(accountPayableBaseCode+2, "failed to delete "+accountPayableName)
	ErrUpdateByIDAccountPayable = errcode.NewError(accountPayableBaseCode+3, "failed to update "+accountPayableName)
	ErrGetByIDAccountPayable    = errcode.NewError(accountPayableBaseCode+4, "failed to get "+accountPayableName+" details")
	ErrListAccountPayable       = errcode.NewError(accountPayableBaseCode+5, "failed to list of "+accountPayableName)

	// error codes are globally unique, adding 1 to the previous error code
)
