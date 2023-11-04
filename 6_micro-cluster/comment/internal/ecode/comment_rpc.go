// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// comment business-level rpc error codes.
// the _commentNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_commentNO       = 94
	_commentName     = "comment"
	_commentBaseCode = errcode.RCode(_commentNO)

	StatusListByProductIDComment   = errcode.NewRPCStatus(_commentBaseCode+1, "failed to ListByProductID "+_commentName)
	// error codes are globally unique, adding 1 to the previous error code
)

