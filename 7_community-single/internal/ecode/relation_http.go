// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// relationService business-level http error codes.
// the relationServiceNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	relationServiceNO       = 5
	relationServiceName     = "relationService"
	relationServiceBaseCode = errcode.HCode(relationServiceNO)

	ErrFollowRelationService           = errcode.NewError(relationServiceBaseCode+1, "failed to Follow "+relationServiceName)
	ErrUnfollowRelationService         = errcode.NewError(relationServiceBaseCode+2, "failed to Unfollow "+relationServiceName)
	ErrListFollowingRelationService    = errcode.NewError(relationServiceBaseCode+3, "failed to ListFollowing "+relationServiceName)
	ErrListFollowerRelationService     = errcode.NewError(relationServiceBaseCode+4, "failed to ListFollower "+relationServiceName)
	ErrBatchGetRelationRelationService = errcode.NewError(relationServiceBaseCode+5, "failed to BatchGetRelation "+relationServiceName)
	ErrFollowSelfRelationService       = errcode.NewError(relationServiceBaseCode+6, "禁止关注自己")
	ErrAlreadyFollowRelationService    = errcode.NewError(relationServiceBaseCode+7, "已经关注过了")
	// error codes are globally unique, adding 1 to the previous error code
)
