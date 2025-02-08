// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/store/v1/transferDetail.proto

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

// Validate checks the field values on CreateTransferDetailRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTransferDetailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTransferDetailRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTransferDetailRequestMultiError, or nil if none found.
func (m *CreateTransferDetailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTransferDetailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransferID

	// no validation rules for SkuID

	// no validation rules for Quantity

	if len(errors) > 0 {
		return CreateTransferDetailRequestMultiError(errors)
	}

	return nil
}

// CreateTransferDetailRequestMultiError is an error wrapping multiple
// validation errors returned by CreateTransferDetailRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateTransferDetailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTransferDetailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTransferDetailRequestMultiError) AllErrors() []error { return m }

// CreateTransferDetailRequestValidationError is the validation error returned
// by CreateTransferDetailRequest.Validate if the designated constraints
// aren't met.
type CreateTransferDetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTransferDetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTransferDetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTransferDetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTransferDetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTransferDetailRequestValidationError) ErrorName() string {
	return "CreateTransferDetailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTransferDetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTransferDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTransferDetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTransferDetailRequestValidationError{}

// Validate checks the field values on CreateTransferDetailReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTransferDetailReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTransferDetailReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTransferDetailReplyMultiError, or nil if none found.
func (m *CreateTransferDetailReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTransferDetailReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransferID

	if len(errors) > 0 {
		return CreateTransferDetailReplyMultiError(errors)
	}

	return nil
}

// CreateTransferDetailReplyMultiError is an error wrapping multiple validation
// errors returned by CreateTransferDetailReply.ValidateAll() if the
// designated constraints aren't met.
type CreateTransferDetailReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTransferDetailReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTransferDetailReplyMultiError) AllErrors() []error { return m }

// CreateTransferDetailReplyValidationError is the validation error returned by
// CreateTransferDetailReply.Validate if the designated constraints aren't met.
type CreateTransferDetailReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTransferDetailReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTransferDetailReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTransferDetailReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTransferDetailReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTransferDetailReplyValidationError) ErrorName() string {
	return "CreateTransferDetailReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTransferDetailReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTransferDetailReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTransferDetailReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTransferDetailReplyValidationError{}

// Validate checks the field values on DeleteTransferDetailByTransferIDRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DeleteTransferDetailByTransferIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// DeleteTransferDetailByTransferIDRequest with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// DeleteTransferDetailByTransferIDRequestMultiError, or nil if none found.
func (m *DeleteTransferDetailByTransferIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTransferDetailByTransferIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTransferID()) < 1 {
		err := DeleteTransferDetailByTransferIDRequestValidationError{
			field:  "TransferID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteTransferDetailByTransferIDRequestMultiError(errors)
	}

	return nil
}

// DeleteTransferDetailByTransferIDRequestMultiError is an error wrapping
// multiple validation errors returned by
// DeleteTransferDetailByTransferIDRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteTransferDetailByTransferIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTransferDetailByTransferIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTransferDetailByTransferIDRequestMultiError) AllErrors() []error { return m }

// DeleteTransferDetailByTransferIDRequestValidationError is the validation
// error returned by DeleteTransferDetailByTransferIDRequest.Validate if the
// designated constraints aren't met.
type DeleteTransferDetailByTransferIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTransferDetailByTransferIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTransferDetailByTransferIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTransferDetailByTransferIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTransferDetailByTransferIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTransferDetailByTransferIDRequestValidationError) ErrorName() string {
	return "DeleteTransferDetailByTransferIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTransferDetailByTransferIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTransferDetailByTransferIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTransferDetailByTransferIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTransferDetailByTransferIDRequestValidationError{}

// Validate checks the field values on DeleteTransferDetailByTransferIDReply
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DeleteTransferDetailByTransferIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTransferDetailByTransferIDReply
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// DeleteTransferDetailByTransferIDReplyMultiError, or nil if none found.
func (m *DeleteTransferDetailByTransferIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTransferDetailByTransferIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteTransferDetailByTransferIDReplyMultiError(errors)
	}

	return nil
}

// DeleteTransferDetailByTransferIDReplyMultiError is an error wrapping
// multiple validation errors returned by
// DeleteTransferDetailByTransferIDReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteTransferDetailByTransferIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTransferDetailByTransferIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTransferDetailByTransferIDReplyMultiError) AllErrors() []error { return m }

// DeleteTransferDetailByTransferIDReplyValidationError is the validation error
// returned by DeleteTransferDetailByTransferIDReply.Validate if the
// designated constraints aren't met.
type DeleteTransferDetailByTransferIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTransferDetailByTransferIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTransferDetailByTransferIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTransferDetailByTransferIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTransferDetailByTransferIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTransferDetailByTransferIDReplyValidationError) ErrorName() string {
	return "DeleteTransferDetailByTransferIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTransferDetailByTransferIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTransferDetailByTransferIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTransferDetailByTransferIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTransferDetailByTransferIDReplyValidationError{}

// Validate checks the field values on UpdateTransferDetailByTransferIDRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateTransferDetailByTransferIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// UpdateTransferDetailByTransferIDRequest with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// UpdateTransferDetailByTransferIDRequestMultiError, or nil if none found.
func (m *UpdateTransferDetailByTransferIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTransferDetailByTransferIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTransferID()) < 1 {
		err := UpdateTransferDetailByTransferIDRequestValidationError{
			field:  "TransferID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for SkuID

	// no validation rules for Quantity

	if len(errors) > 0 {
		return UpdateTransferDetailByTransferIDRequestMultiError(errors)
	}

	return nil
}

// UpdateTransferDetailByTransferIDRequestMultiError is an error wrapping
// multiple validation errors returned by
// UpdateTransferDetailByTransferIDRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateTransferDetailByTransferIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTransferDetailByTransferIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTransferDetailByTransferIDRequestMultiError) AllErrors() []error { return m }

// UpdateTransferDetailByTransferIDRequestValidationError is the validation
// error returned by UpdateTransferDetailByTransferIDRequest.Validate if the
// designated constraints aren't met.
type UpdateTransferDetailByTransferIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTransferDetailByTransferIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTransferDetailByTransferIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTransferDetailByTransferIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTransferDetailByTransferIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTransferDetailByTransferIDRequestValidationError) ErrorName() string {
	return "UpdateTransferDetailByTransferIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTransferDetailByTransferIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTransferDetailByTransferIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTransferDetailByTransferIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTransferDetailByTransferIDRequestValidationError{}

// Validate checks the field values on UpdateTransferDetailByTransferIDReply
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateTransferDetailByTransferIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTransferDetailByTransferIDReply
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdateTransferDetailByTransferIDReplyMultiError, or nil if none found.
func (m *UpdateTransferDetailByTransferIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTransferDetailByTransferIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateTransferDetailByTransferIDReplyMultiError(errors)
	}

	return nil
}

// UpdateTransferDetailByTransferIDReplyMultiError is an error wrapping
// multiple validation errors returned by
// UpdateTransferDetailByTransferIDReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateTransferDetailByTransferIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTransferDetailByTransferIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTransferDetailByTransferIDReplyMultiError) AllErrors() []error { return m }

// UpdateTransferDetailByTransferIDReplyValidationError is the validation error
// returned by UpdateTransferDetailByTransferIDReply.Validate if the
// designated constraints aren't met.
type UpdateTransferDetailByTransferIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTransferDetailByTransferIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTransferDetailByTransferIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTransferDetailByTransferIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTransferDetailByTransferIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTransferDetailByTransferIDReplyValidationError) ErrorName() string {
	return "UpdateTransferDetailByTransferIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTransferDetailByTransferIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTransferDetailByTransferIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTransferDetailByTransferIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTransferDetailByTransferIDReplyValidationError{}

// Validate checks the field values on TransferDetail with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TransferDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransferDetail with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TransferDetailMultiError,
// or nil if none found.
func (m *TransferDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *TransferDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransferID

	// no validation rules for SkuID

	// no validation rules for Quantity

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return TransferDetailMultiError(errors)
	}

	return nil
}

// TransferDetailMultiError is an error wrapping multiple validation errors
// returned by TransferDetail.ValidateAll() if the designated constraints
// aren't met.
type TransferDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransferDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransferDetailMultiError) AllErrors() []error { return m }

// TransferDetailValidationError is the validation error returned by
// TransferDetail.Validate if the designated constraints aren't met.
type TransferDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransferDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransferDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransferDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransferDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransferDetailValidationError) ErrorName() string { return "TransferDetailValidationError" }

// Error satisfies the builtin error interface
func (e TransferDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransferDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransferDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransferDetailValidationError{}

// Validate checks the field values on GetTransferDetailByTransferIDRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetTransferDetailByTransferIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTransferDetailByTransferIDRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetTransferDetailByTransferIDRequestMultiError, or nil if none found.
func (m *GetTransferDetailByTransferIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTransferDetailByTransferIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTransferID()) < 1 {
		err := GetTransferDetailByTransferIDRequestValidationError{
			field:  "TransferID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetTransferDetailByTransferIDRequestMultiError(errors)
	}

	return nil
}

// GetTransferDetailByTransferIDRequestMultiError is an error wrapping multiple
// validation errors returned by
// GetTransferDetailByTransferIDRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTransferDetailByTransferIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTransferDetailByTransferIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTransferDetailByTransferIDRequestMultiError) AllErrors() []error { return m }

// GetTransferDetailByTransferIDRequestValidationError is the validation error
// returned by GetTransferDetailByTransferIDRequest.Validate if the designated
// constraints aren't met.
type GetTransferDetailByTransferIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransferDetailByTransferIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransferDetailByTransferIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransferDetailByTransferIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransferDetailByTransferIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransferDetailByTransferIDRequestValidationError) ErrorName() string {
	return "GetTransferDetailByTransferIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransferDetailByTransferIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransferDetailByTransferIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransferDetailByTransferIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransferDetailByTransferIDRequestValidationError{}

// Validate checks the field values on GetTransferDetailByTransferIDReply with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetTransferDetailByTransferIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTransferDetailByTransferIDReply
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetTransferDetailByTransferIDReplyMultiError, or nil if none found.
func (m *GetTransferDetailByTransferIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTransferDetailByTransferIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTransferDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTransferDetailByTransferIDReplyValidationError{
					field:  "TransferDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTransferDetailByTransferIDReplyValidationError{
					field:  "TransferDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTransferDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTransferDetailByTransferIDReplyValidationError{
				field:  "TransferDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTransferDetailByTransferIDReplyMultiError(errors)
	}

	return nil
}

// GetTransferDetailByTransferIDReplyMultiError is an error wrapping multiple
// validation errors returned by
// GetTransferDetailByTransferIDReply.ValidateAll() if the designated
// constraints aren't met.
type GetTransferDetailByTransferIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTransferDetailByTransferIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTransferDetailByTransferIDReplyMultiError) AllErrors() []error { return m }

// GetTransferDetailByTransferIDReplyValidationError is the validation error
// returned by GetTransferDetailByTransferIDReply.Validate if the designated
// constraints aren't met.
type GetTransferDetailByTransferIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransferDetailByTransferIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransferDetailByTransferIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransferDetailByTransferIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransferDetailByTransferIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransferDetailByTransferIDReplyValidationError) ErrorName() string {
	return "GetTransferDetailByTransferIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransferDetailByTransferIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransferDetailByTransferIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransferDetailByTransferIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransferDetailByTransferIDReplyValidationError{}

// Validate checks the field values on ListTransferDetailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTransferDetailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTransferDetailRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTransferDetailRequestMultiError, or nil if none found.
func (m *ListTransferDetailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTransferDetailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetParams()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListTransferDetailRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTransferDetailRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTransferDetailRequestValidationError{
				field:  "Params",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListTransferDetailRequestMultiError(errors)
	}

	return nil
}

// ListTransferDetailRequestMultiError is an error wrapping multiple validation
// errors returned by ListTransferDetailRequest.ValidateAll() if the
// designated constraints aren't met.
type ListTransferDetailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTransferDetailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTransferDetailRequestMultiError) AllErrors() []error { return m }

// ListTransferDetailRequestValidationError is the validation error returned by
// ListTransferDetailRequest.Validate if the designated constraints aren't met.
type ListTransferDetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTransferDetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTransferDetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTransferDetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTransferDetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTransferDetailRequestValidationError) ErrorName() string {
	return "ListTransferDetailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListTransferDetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTransferDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTransferDetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTransferDetailRequestValidationError{}

// Validate checks the field values on ListTransferDetailReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTransferDetailReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTransferDetailReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTransferDetailReplyMultiError, or nil if none found.
func (m *ListTransferDetailReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTransferDetailReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetTransferDetails() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListTransferDetailReplyValidationError{
						field:  fmt.Sprintf("TransferDetails[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTransferDetailReplyValidationError{
						field:  fmt.Sprintf("TransferDetails[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTransferDetailReplyValidationError{
					field:  fmt.Sprintf("TransferDetails[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListTransferDetailReplyMultiError(errors)
	}

	return nil
}

// ListTransferDetailReplyMultiError is an error wrapping multiple validation
// errors returned by ListTransferDetailReply.ValidateAll() if the designated
// constraints aren't met.
type ListTransferDetailReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTransferDetailReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTransferDetailReplyMultiError) AllErrors() []error { return m }

// ListTransferDetailReplyValidationError is the validation error returned by
// ListTransferDetailReply.Validate if the designated constraints aren't met.
type ListTransferDetailReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTransferDetailReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTransferDetailReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTransferDetailReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTransferDetailReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTransferDetailReplyValidationError) ErrorName() string {
	return "ListTransferDetailReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListTransferDetailReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTransferDetailReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTransferDetailReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTransferDetailReplyValidationError{}
