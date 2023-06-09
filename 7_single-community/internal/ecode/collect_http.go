// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// collectService http service level error code
var (
	collectServiceNO       = 6 // number range 1~100, if there is the same number, trigger panic.
	collectServiceName     = "collectService"
	collectServiceBaseCode = errcode.HCode(collectServiceNO)

	ErrCreateCollectService = errcode.NewError(collectServiceBaseCode+1, "failed to Create "+collectServiceName)
	ErrDeleteCollectService = errcode.NewError(collectServiceBaseCode+2, "failed to Delete "+collectServiceName)
	ErrListCollectService   = errcode.NewError(collectServiceBaseCode+3, "failed to List "+collectServiceName)
	// add +1 to the previous error code
)
