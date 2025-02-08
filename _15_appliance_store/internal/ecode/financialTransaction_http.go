package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// financialTransaction business-level http error codes.
// the financialTransactionNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	financialTransactionNO       = 75
	financialTransactionName     = "financialTransaction"
	financialTransactionBaseCode = errcode.HCode(financialTransactionNO)

	ErrCreateFinancialTransaction     = errcode.NewError(financialTransactionBaseCode+1, "failed to create "+financialTransactionName)
	ErrDeleteByIDFinancialTransaction = errcode.NewError(financialTransactionBaseCode+2, "failed to delete "+financialTransactionName)
	ErrUpdateByIDFinancialTransaction = errcode.NewError(financialTransactionBaseCode+3, "failed to update "+financialTransactionName)
	ErrGetByIDFinancialTransaction    = errcode.NewError(financialTransactionBaseCode+4, "failed to get "+financialTransactionName+" details")
	ErrListFinancialTransaction       = errcode.NewError(financialTransactionBaseCode+5, "failed to list of "+financialTransactionName)

	// error codes are globally unique, adding 1 to the previous error code
)
