package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// customer business-level http error codes.
// the customerNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	customerNO       = 36
	customerName     = "customer"
	customerBaseCode = errcode.HCode(customerNO)

	ErrCreateCustomer     = errcode.NewError(customerBaseCode+1, "failed to create "+customerName)
	ErrDeleteByIDCustomer = errcode.NewError(customerBaseCode+2, "failed to delete "+customerName)
	ErrUpdateByIDCustomer = errcode.NewError(customerBaseCode+3, "failed to update "+customerName)
	ErrGetByIDCustomer    = errcode.NewError(customerBaseCode+4, "failed to get "+customerName+" details")
	ErrListCustomer       = errcode.NewError(customerBaseCode+5, "failed to list of "+customerName)

	// error codes are globally unique, adding 1 to the previous error code
)
