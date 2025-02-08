package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// notification business-level http error codes.
// the notificationNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	notificationNO       = 72
	notificationName     = "notification"
	notificationBaseCode = errcode.HCode(notificationNO)

	ErrCreateNotification                   = errcode.NewError(notificationBaseCode+1, "failed to create "+notificationName)
	ErrDeleteByIDNotification               = errcode.NewError(notificationBaseCode+2, "failed to delete "+notificationName)
	ErrUpdateByIDNotification               = errcode.NewError(notificationBaseCode+3, "failed to update "+notificationName)
	ErrGetByIDNotification                  = errcode.NewError(notificationBaseCode+4, "failed to get "+notificationName+" details")
	ErrListNotification                     = errcode.NewError(notificationBaseCode+5, "failed to list of "+notificationName)
	ErrSendRealTimeNotificationNotification = errcode.NewError(notificationBaseCode+6, "failed to SendRealTimeNotification "+notificationName)

	// error codes are globally unique, adding 1 to the previous error code
)
