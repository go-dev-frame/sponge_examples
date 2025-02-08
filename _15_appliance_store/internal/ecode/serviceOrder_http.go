package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// serviceOrder business-level http error codes.
// the serviceOrderNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	serviceOrderNO       = 7
	serviceOrderName     = "serviceOrder"
	serviceOrderBaseCode = errcode.HCode(serviceOrderNO)

	ErrCreateServiceOrder     = errcode.NewError(serviceOrderBaseCode+1, "failed to create "+serviceOrderName)
	ErrDeleteByIDServiceOrder = errcode.NewError(serviceOrderBaseCode+2, "failed to delete "+serviceOrderName)
	ErrUpdateByIDServiceOrder = errcode.NewError(serviceOrderBaseCode+3, "failed to update "+serviceOrderName)
	ErrGetByIDServiceOrder    = errcode.NewError(serviceOrderBaseCode+4, "failed to get "+serviceOrderName+" details")
	ErrListServiceOrder       = errcode.NewError(serviceOrderBaseCode+5, "failed to list of "+serviceOrderName)

	// error codes are globally unique, adding 1 to the previous error code
)
