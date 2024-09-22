// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// downgrade business-level rpc error codes.
// the _downgradeNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_downgradeNO       = 74
	_downgradeName     = "downgrade"
	_downgradeBaseCode = errcode.RCode(_downgradeNO)

	StatusUpdateDowngrade          = errcode.NewRPCStatus(_downgradeBaseCode+1, "failed to Update "+_downgradeName)
	StatusQueryDowngrade           = errcode.NewRPCStatus(_downgradeBaseCode+2, "failed to Query "+_downgradeName)
	StatusDowngradeBranchDowngrade = errcode.NewRPCStatus(_downgradeBaseCode+3, "failed to DowngradeBranch "+_downgradeName)

	// error codes are globally unique, adding 1 to the previous error code
)
