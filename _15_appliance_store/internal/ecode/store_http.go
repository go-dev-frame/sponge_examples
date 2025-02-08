package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// store business-level http error codes.
// the storeNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	storeNO       = 35
	storeName     = "store"
	storeBaseCode = errcode.HCode(storeNO)

	ErrCreateStore     = errcode.NewError(storeBaseCode+1, "failed to create "+storeName)
	ErrDeleteByIDStore = errcode.NewError(storeBaseCode+2, "failed to delete "+storeName)
	ErrUpdateByIDStore = errcode.NewError(storeBaseCode+3, "failed to update "+storeName)
	ErrGetByIDStore    = errcode.NewError(storeBaseCode+4, "failed to get "+storeName+" details")
	ErrListStore       = errcode.NewError(storeBaseCode+5, "failed to list of "+storeName)

	// error codes are globally unique, adding 1 to the previous error code
)
