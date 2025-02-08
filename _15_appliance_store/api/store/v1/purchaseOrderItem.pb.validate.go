// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/store/v1/purchaseOrderItem.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreatePurchaseOrderItemRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePurchaseOrderItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePurchaseOrderItemRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreatePurchaseOrderItemRequestMultiError, or nil if none found.
func (m *CreatePurchaseOrderItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePurchaseOrderItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderID

	// no validation rules for SkuID

	// no validation rules for Quantity

	// no validation rules for Price

	// no validation rules for UnitPrice

	if len(errors) > 0 {
		return CreatePurchaseOrderItemRequestMultiError(errors)
	}

	return nil
}

// CreatePurchaseOrderItemRequestMultiError is an error wrapping multiple
// validation errors returned by CreatePurchaseOrderItemRequest.ValidateAll()
// if the designated constraints aren't met.
type CreatePurchaseOrderItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePurchaseOrderItemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePurchaseOrderItemRequestMultiError) AllErrors() []error { return m }

// CreatePurchaseOrderItemRequestValidationError is the validation error
// returned by CreatePurchaseOrderItemRequest.Validate if the designated
// constraints aren't met.
type CreatePurchaseOrderItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePurchaseOrderItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePurchaseOrderItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePurchaseOrderItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePurchaseOrderItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePurchaseOrderItemRequestValidationError) ErrorName() string {
	return "CreatePurchaseOrderItemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePurchaseOrderItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePurchaseOrderItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePurchaseOrderItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePurchaseOrderItemRequestValidationError{}

// Validate checks the field values on CreatePurchaseOrderItemReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePurchaseOrderItemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePurchaseOrderItemReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePurchaseOrderItemReplyMultiError, or nil if none found.
func (m *CreatePurchaseOrderItemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePurchaseOrderItemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreatePurchaseOrderItemReplyMultiError(errors)
	}

	return nil
}

// CreatePurchaseOrderItemReplyMultiError is an error wrapping multiple
// validation errors returned by CreatePurchaseOrderItemReply.ValidateAll() if
// the designated constraints aren't met.
type CreatePurchaseOrderItemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePurchaseOrderItemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePurchaseOrderItemReplyMultiError) AllErrors() []error { return m }

// CreatePurchaseOrderItemReplyValidationError is the validation error returned
// by CreatePurchaseOrderItemReply.Validate if the designated constraints
// aren't met.
type CreatePurchaseOrderItemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePurchaseOrderItemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePurchaseOrderItemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePurchaseOrderItemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePurchaseOrderItemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePurchaseOrderItemReplyValidationError) ErrorName() string {
	return "CreatePurchaseOrderItemReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePurchaseOrderItemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePurchaseOrderItemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePurchaseOrderItemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePurchaseOrderItemReplyValidationError{}

// Validate checks the field values on DeletePurchaseOrderItemByIDRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *DeletePurchaseOrderItemByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePurchaseOrderItemByIDRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// DeletePurchaseOrderItemByIDRequestMultiError, or nil if none found.
func (m *DeletePurchaseOrderItemByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePurchaseOrderItemByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeletePurchaseOrderItemByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeletePurchaseOrderItemByIDRequestMultiError(errors)
	}

	return nil
}

// DeletePurchaseOrderItemByIDRequestMultiError is an error wrapping multiple
// validation errors returned by
// DeletePurchaseOrderItemByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type DeletePurchaseOrderItemByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePurchaseOrderItemByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePurchaseOrderItemByIDRequestMultiError) AllErrors() []error { return m }

// DeletePurchaseOrderItemByIDRequestValidationError is the validation error
// returned by DeletePurchaseOrderItemByIDRequest.Validate if the designated
// constraints aren't met.
type DeletePurchaseOrderItemByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePurchaseOrderItemByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePurchaseOrderItemByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePurchaseOrderItemByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePurchaseOrderItemByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePurchaseOrderItemByIDRequestValidationError) ErrorName() string {
	return "DeletePurchaseOrderItemByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePurchaseOrderItemByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePurchaseOrderItemByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePurchaseOrderItemByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePurchaseOrderItemByIDRequestValidationError{}

// Validate checks the field values on DeletePurchaseOrderItemByIDReply with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *DeletePurchaseOrderItemByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePurchaseOrderItemByIDReply with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeletePurchaseOrderItemByIDReplyMultiError, or nil if none found.
func (m *DeletePurchaseOrderItemByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePurchaseOrderItemByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeletePurchaseOrderItemByIDReplyMultiError(errors)
	}

	return nil
}

// DeletePurchaseOrderItemByIDReplyMultiError is an error wrapping multiple
// validation errors returned by
// DeletePurchaseOrderItemByIDReply.ValidateAll() if the designated
// constraints aren't met.
type DeletePurchaseOrderItemByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePurchaseOrderItemByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePurchaseOrderItemByIDReplyMultiError) AllErrors() []error { return m }

// DeletePurchaseOrderItemByIDReplyValidationError is the validation error
// returned by DeletePurchaseOrderItemByIDReply.Validate if the designated
// constraints aren't met.
type DeletePurchaseOrderItemByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePurchaseOrderItemByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePurchaseOrderItemByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePurchaseOrderItemByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePurchaseOrderItemByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePurchaseOrderItemByIDReplyValidationError) ErrorName() string {
	return "DeletePurchaseOrderItemByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePurchaseOrderItemByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePurchaseOrderItemByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePurchaseOrderItemByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePurchaseOrderItemByIDReplyValidationError{}

// Validate checks the field values on UpdatePurchaseOrderItemByIDRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *UpdatePurchaseOrderItemByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePurchaseOrderItemByIDRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdatePurchaseOrderItemByIDRequestMultiError, or nil if none found.
func (m *UpdatePurchaseOrderItemByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePurchaseOrderItemByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdatePurchaseOrderItemByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for OrderID

	// no validation rules for SkuID

	// no validation rules for Quantity

	// no validation rules for Price

	// no validation rules for UnitPrice

	if len(errors) > 0 {
		return UpdatePurchaseOrderItemByIDRequestMultiError(errors)
	}

	return nil
}

// UpdatePurchaseOrderItemByIDRequestMultiError is an error wrapping multiple
// validation errors returned by
// UpdatePurchaseOrderItemByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdatePurchaseOrderItemByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePurchaseOrderItemByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePurchaseOrderItemByIDRequestMultiError) AllErrors() []error { return m }

// UpdatePurchaseOrderItemByIDRequestValidationError is the validation error
// returned by UpdatePurchaseOrderItemByIDRequest.Validate if the designated
// constraints aren't met.
type UpdatePurchaseOrderItemByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePurchaseOrderItemByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePurchaseOrderItemByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePurchaseOrderItemByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePurchaseOrderItemByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePurchaseOrderItemByIDRequestValidationError) ErrorName() string {
	return "UpdatePurchaseOrderItemByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePurchaseOrderItemByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePurchaseOrderItemByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePurchaseOrderItemByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePurchaseOrderItemByIDRequestValidationError{}

// Validate checks the field values on UpdatePurchaseOrderItemByIDReply with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *UpdatePurchaseOrderItemByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePurchaseOrderItemByIDReply with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdatePurchaseOrderItemByIDReplyMultiError, or nil if none found.
func (m *UpdatePurchaseOrderItemByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePurchaseOrderItemByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdatePurchaseOrderItemByIDReplyMultiError(errors)
	}

	return nil
}

// UpdatePurchaseOrderItemByIDReplyMultiError is an error wrapping multiple
// validation errors returned by
// UpdatePurchaseOrderItemByIDReply.ValidateAll() if the designated
// constraints aren't met.
type UpdatePurchaseOrderItemByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePurchaseOrderItemByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePurchaseOrderItemByIDReplyMultiError) AllErrors() []error { return m }

// UpdatePurchaseOrderItemByIDReplyValidationError is the validation error
// returned by UpdatePurchaseOrderItemByIDReply.Validate if the designated
// constraints aren't met.
type UpdatePurchaseOrderItemByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePurchaseOrderItemByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePurchaseOrderItemByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePurchaseOrderItemByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePurchaseOrderItemByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePurchaseOrderItemByIDReplyValidationError) ErrorName() string {
	return "UpdatePurchaseOrderItemByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePurchaseOrderItemByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePurchaseOrderItemByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePurchaseOrderItemByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePurchaseOrderItemByIDReplyValidationError{}

// Validate checks the field values on PurchaseOrderItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PurchaseOrderItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PurchaseOrderItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PurchaseOrderItemMultiError, or nil if none found.
func (m *PurchaseOrderItem) ValidateAll() error {
	return m.validate(true)
}

func (m *PurchaseOrderItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OrderID

	// no validation rules for SkuID

	// no validation rules for Quantity

	// no validation rules for Price

	// no validation rules for UnitPrice

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return PurchaseOrderItemMultiError(errors)
	}

	return nil
}

// PurchaseOrderItemMultiError is an error wrapping multiple validation errors
// returned by PurchaseOrderItem.ValidateAll() if the designated constraints
// aren't met.
type PurchaseOrderItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PurchaseOrderItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PurchaseOrderItemMultiError) AllErrors() []error { return m }

// PurchaseOrderItemValidationError is the validation error returned by
// PurchaseOrderItem.Validate if the designated constraints aren't met.
type PurchaseOrderItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurchaseOrderItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurchaseOrderItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurchaseOrderItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurchaseOrderItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurchaseOrderItemValidationError) ErrorName() string {
	return "PurchaseOrderItemValidationError"
}

// Error satisfies the builtin error interface
func (e PurchaseOrderItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurchaseOrderItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurchaseOrderItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurchaseOrderItemValidationError{}

// Validate checks the field values on GetPurchaseOrderItemByIDRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPurchaseOrderItemByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPurchaseOrderItemByIDRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPurchaseOrderItemByIDRequestMultiError, or nil if none found.
func (m *GetPurchaseOrderItemByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPurchaseOrderItemByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetPurchaseOrderItemByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetPurchaseOrderItemByIDRequestMultiError(errors)
	}

	return nil
}

// GetPurchaseOrderItemByIDRequestMultiError is an error wrapping multiple
// validation errors returned by GetPurchaseOrderItemByIDRequest.ValidateAll()
// if the designated constraints aren't met.
type GetPurchaseOrderItemByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPurchaseOrderItemByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPurchaseOrderItemByIDRequestMultiError) AllErrors() []error { return m }

// GetPurchaseOrderItemByIDRequestValidationError is the validation error
// returned by GetPurchaseOrderItemByIDRequest.Validate if the designated
// constraints aren't met.
type GetPurchaseOrderItemByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPurchaseOrderItemByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPurchaseOrderItemByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPurchaseOrderItemByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPurchaseOrderItemByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPurchaseOrderItemByIDRequestValidationError) ErrorName() string {
	return "GetPurchaseOrderItemByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPurchaseOrderItemByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPurchaseOrderItemByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPurchaseOrderItemByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPurchaseOrderItemByIDRequestValidationError{}

// Validate checks the field values on GetPurchaseOrderItemByIDReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPurchaseOrderItemByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPurchaseOrderItemByIDReply with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPurchaseOrderItemByIDReplyMultiError, or nil if none found.
func (m *GetPurchaseOrderItemByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPurchaseOrderItemByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPurchaseOrderItem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetPurchaseOrderItemByIDReplyValidationError{
					field:  "PurchaseOrderItem",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetPurchaseOrderItemByIDReplyValidationError{
					field:  "PurchaseOrderItem",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPurchaseOrderItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPurchaseOrderItemByIDReplyValidationError{
				field:  "PurchaseOrderItem",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetPurchaseOrderItemByIDReplyMultiError(errors)
	}

	return nil
}

// GetPurchaseOrderItemByIDReplyMultiError is an error wrapping multiple
// validation errors returned by GetPurchaseOrderItemByIDReply.ValidateAll()
// if the designated constraints aren't met.
type GetPurchaseOrderItemByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPurchaseOrderItemByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPurchaseOrderItemByIDReplyMultiError) AllErrors() []error { return m }

// GetPurchaseOrderItemByIDReplyValidationError is the validation error
// returned by GetPurchaseOrderItemByIDReply.Validate if the designated
// constraints aren't met.
type GetPurchaseOrderItemByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPurchaseOrderItemByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPurchaseOrderItemByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPurchaseOrderItemByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPurchaseOrderItemByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPurchaseOrderItemByIDReplyValidationError) ErrorName() string {
	return "GetPurchaseOrderItemByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetPurchaseOrderItemByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPurchaseOrderItemByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPurchaseOrderItemByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPurchaseOrderItemByIDReplyValidationError{}

// Validate checks the field values on ListPurchaseOrderItemRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPurchaseOrderItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPurchaseOrderItemRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPurchaseOrderItemRequestMultiError, or nil if none found.
func (m *ListPurchaseOrderItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPurchaseOrderItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetParams()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListPurchaseOrderItemRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListPurchaseOrderItemRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPurchaseOrderItemRequestValidationError{
				field:  "Params",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListPurchaseOrderItemRequestMultiError(errors)
	}

	return nil
}

// ListPurchaseOrderItemRequestMultiError is an error wrapping multiple
// validation errors returned by ListPurchaseOrderItemRequest.ValidateAll() if
// the designated constraints aren't met.
type ListPurchaseOrderItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPurchaseOrderItemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPurchaseOrderItemRequestMultiError) AllErrors() []error { return m }

// ListPurchaseOrderItemRequestValidationError is the validation error returned
// by ListPurchaseOrderItemRequest.Validate if the designated constraints
// aren't met.
type ListPurchaseOrderItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPurchaseOrderItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPurchaseOrderItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPurchaseOrderItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPurchaseOrderItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPurchaseOrderItemRequestValidationError) ErrorName() string {
	return "ListPurchaseOrderItemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPurchaseOrderItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPurchaseOrderItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPurchaseOrderItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPurchaseOrderItemRequestValidationError{}

// Validate checks the field values on ListPurchaseOrderItemReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPurchaseOrderItemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPurchaseOrderItemReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPurchaseOrderItemReplyMultiError, or nil if none found.
func (m *ListPurchaseOrderItemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPurchaseOrderItemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetPurchaseOrderItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListPurchaseOrderItemReplyValidationError{
						field:  fmt.Sprintf("PurchaseOrderItems[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPurchaseOrderItemReplyValidationError{
						field:  fmt.Sprintf("PurchaseOrderItems[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPurchaseOrderItemReplyValidationError{
					field:  fmt.Sprintf("PurchaseOrderItems[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListPurchaseOrderItemReplyMultiError(errors)
	}

	return nil
}

// ListPurchaseOrderItemReplyMultiError is an error wrapping multiple
// validation errors returned by ListPurchaseOrderItemReply.ValidateAll() if
// the designated constraints aren't met.
type ListPurchaseOrderItemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPurchaseOrderItemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPurchaseOrderItemReplyMultiError) AllErrors() []error { return m }

// ListPurchaseOrderItemReplyValidationError is the validation error returned
// by ListPurchaseOrderItemReply.Validate if the designated constraints aren't met.
type ListPurchaseOrderItemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPurchaseOrderItemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPurchaseOrderItemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPurchaseOrderItemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPurchaseOrderItemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPurchaseOrderItemReplyValidationError) ErrorName() string {
	return "ListPurchaseOrderItemReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListPurchaseOrderItemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPurchaseOrderItemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPurchaseOrderItemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPurchaseOrderItemReplyValidationError{}
