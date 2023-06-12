// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// likeService http service level error code
var (
	likeServiceNO       = 4 // number range 1~100, if there is the same number, trigger panic.
	likeServiceName     = "likeService"
	likeServiceBaseCode = errcode.HCode(likeServiceNO)

	ErrCreateLikeService      = errcode.NewError(likeServiceBaseCode+1, "failed to Create "+likeServiceName)
	ErrDeleteLikeService      = errcode.NewError(likeServiceBaseCode+2, "failed to Delete "+likeServiceName)
	ErrListPostLikeService    = errcode.NewError(likeServiceBaseCode+3, "failed to ListPost "+likeServiceName)
	ErrListCommentLikeService = errcode.NewError(likeServiceBaseCode+4, "failed to ListComment "+likeServiceName)
	// add +1 to the previous error code
)
