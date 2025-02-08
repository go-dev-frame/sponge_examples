package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// employee business-level http error codes.
// the employeeNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	employeeNO       = 56
	employeeName     = "employee"
	employeeBaseCode = errcode.HCode(employeeNO)

	ErrCreateEmployee               = errcode.NewError(employeeBaseCode+1, "failed to create "+employeeName)
	ErrDeleteByIDEmployee           = errcode.NewError(employeeBaseCode+2, "failed to delete "+employeeName)
	ErrUpdateByIDEmployee           = errcode.NewError(employeeBaseCode+3, "failed to update "+employeeName)
	ErrGetByIDEmployee              = errcode.NewError(employeeBaseCode+4, "failed to get "+employeeName+" details")
	ErrListEmployee                 = errcode.NewError(employeeBaseCode+5, "failed to list of "+employeeName)
	ErrLoginEmployee                = errcode.NewError(employeeBaseCode+6, "failed to Login "+employeeName)
	ErrLogoutEmployee               = errcode.NewError(employeeBaseCode+7, "failed to Logout "+employeeName)
	ErrChangePasswordEmployee       = errcode.NewError(employeeBaseCode+8, "failed to ChangePassword "+employeeName)
	ErrResetPasswordEmployee        = errcode.NewError(employeeBaseCode+9, "failed to ResetPassword "+employeeName)
	ErrSendVerificationCodeEmployee = errcode.NewError(employeeBaseCode+10, "failed to SendVerificationCode "+employeeName)

	// error codes are globally unique, adding 1 to the previous error code
)
