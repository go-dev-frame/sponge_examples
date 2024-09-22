// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// final business-level rpc error codes.
// the _finalNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_finalNO       = 89
	_finalName     = "final"
	_finalBaseCode = errcode.RCode(_finalNO)

	StatusUpdateFinal = errcode.NewRPCStatus(_finalBaseCode+1, "failed to Update "+_finalName)
	StatusQueryFinal  = errcode.NewRPCStatus(_finalBaseCode+2, "failed to Query "+_finalName)

	// error codes are globally unique, adding 1 to the previous error code
)
