package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// employeePerformance business-level http error codes.
// the employeePerformanceNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	employeePerformanceNO       = 21
	employeePerformanceName     = "employeePerformance"
	employeePerformanceBaseCode = errcode.HCode(employeePerformanceNO)

	ErrCreateEmployeePerformance                   = errcode.NewError(employeePerformanceBaseCode+1, "failed to create "+employeePerformanceName)
	ErrDeleteByIDEmployeePerformance               = errcode.NewError(employeePerformanceBaseCode+2, "failed to delete "+employeePerformanceName)
	ErrUpdateByIDEmployeePerformance               = errcode.NewError(employeePerformanceBaseCode+3, "failed to update "+employeePerformanceName)
	ErrGetByIDEmployeePerformance                  = errcode.NewError(employeePerformanceBaseCode+4, "failed to get "+employeePerformanceName+" details")
	ErrListEmployeePerformance                     = errcode.NewError(employeePerformanceBaseCode+5, "failed to list of "+employeePerformanceName)
	ErrCalculateSalesCommissionEmployeePerformance = errcode.NewError(employeePerformanceBaseCode+6, "failed to CalculateSalesCommission "+employeePerformanceName)

	// error codes are globally unique, adding 1 to the previous error code
)
