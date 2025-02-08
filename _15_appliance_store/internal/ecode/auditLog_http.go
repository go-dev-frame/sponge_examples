package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// auditLog business-level http error codes.
// the auditLogNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	auditLogNO       = 82
	auditLogName     = "auditLog"
	auditLogBaseCode = errcode.HCode(auditLogNO)

	ErrCreateAuditLog     = errcode.NewError(auditLogBaseCode+1, "failed to create "+auditLogName)
	ErrDeleteByIDAuditLog = errcode.NewError(auditLogBaseCode+2, "failed to delete "+auditLogName)
	ErrUpdateByIDAuditLog = errcode.NewError(auditLogBaseCode+3, "failed to update "+auditLogName)
	ErrGetByIDAuditLog    = errcode.NewError(auditLogBaseCode+4, "failed to get "+auditLogName+" details")
	ErrListAuditLog       = errcode.NewError(auditLogBaseCode+5, "failed to list of "+auditLogName)

	// error codes are globally unique, adding 1 to the previous error code
)
