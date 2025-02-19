// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/store/v1/store.proto

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

// Validate checks the field values on CreateStoreRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateStoreRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateStoreRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateStoreRequestMultiError, or nil if none found.
func (m *CreateStoreRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateStoreRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Address

	// no validation rules for ContactPhone

	// no validation rules for ManagerID

	if len(errors) > 0 {
		return CreateStoreRequestMultiError(errors)
	}

	return nil
}

// CreateStoreRequestMultiError is an error wrapping multiple validation errors
// returned by CreateStoreRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateStoreRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateStoreRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateStoreRequestMultiError) AllErrors() []error { return m }

// CreateStoreRequestValidationError is the validation error returned by
// CreateStoreRequest.Validate if the designated constraints aren't met.
type CreateStoreRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateStoreRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateStoreRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateStoreRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateStoreRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateStoreRequestValidationError) ErrorName() string {
	return "CreateStoreRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateStoreRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateStoreRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateStoreRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateStoreRequestValidationError{}

// Validate checks the field values on CreateStoreReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateStoreReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateStoreReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateStoreReplyMultiError, or nil if none found.
func (m *CreateStoreReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateStoreReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateStoreReplyMultiError(errors)
	}

	return nil
}

// CreateStoreReplyMultiError is an error wrapping multiple validation errors
// returned by CreateStoreReply.ValidateAll() if the designated constraints
// aren't met.
type CreateStoreReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateStoreReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateStoreReplyMultiError) AllErrors() []error { return m }

// CreateStoreReplyValidationError is the validation error returned by
// CreateStoreReply.Validate if the designated constraints aren't met.
type CreateStoreReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateStoreReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateStoreReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateStoreReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateStoreReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateStoreReplyValidationError) ErrorName() string { return "CreateStoreReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateStoreReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateStoreReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateStoreReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateStoreReplyValidationError{}

// Validate checks the field values on DeleteStoreByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteStoreByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteStoreByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteStoreByIDRequestMultiError, or nil if none found.
func (m *DeleteStoreByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteStoreByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteStoreByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteStoreByIDRequestMultiError(errors)
	}

	return nil
}

// DeleteStoreByIDRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteStoreByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteStoreByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteStoreByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteStoreByIDRequestMultiError) AllErrors() []error { return m }

// DeleteStoreByIDRequestValidationError is the validation error returned by
// DeleteStoreByIDRequest.Validate if the designated constraints aren't met.
type DeleteStoreByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteStoreByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteStoreByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteStoreByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteStoreByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteStoreByIDRequestValidationError) ErrorName() string {
	return "DeleteStoreByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteStoreByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteStoreByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteStoreByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteStoreByIDRequestValidationError{}

// Validate checks the field values on DeleteStoreByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteStoreByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteStoreByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteStoreByIDReplyMultiError, or nil if none found.
func (m *DeleteStoreByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteStoreByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteStoreByIDReplyMultiError(errors)
	}

	return nil
}

// DeleteStoreByIDReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteStoreByIDReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteStoreByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteStoreByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteStoreByIDReplyMultiError) AllErrors() []error { return m }

// DeleteStoreByIDReplyValidationError is the validation error returned by
// DeleteStoreByIDReply.Validate if the designated constraints aren't met.
type DeleteStoreByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteStoreByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteStoreByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteStoreByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteStoreByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteStoreByIDReplyValidationError) ErrorName() string {
	return "DeleteStoreByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteStoreByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteStoreByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteStoreByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteStoreByIDReplyValidationError{}

// Validate checks the field values on UpdateStoreByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateStoreByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateStoreByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateStoreByIDRequestMultiError, or nil if none found.
func (m *UpdateStoreByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateStoreByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateStoreByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for Address

	// no validation rules for ContactPhone

	// no validation rules for ManagerID

	if len(errors) > 0 {
		return UpdateStoreByIDRequestMultiError(errors)
	}

	return nil
}

// UpdateStoreByIDRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateStoreByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateStoreByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateStoreByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateStoreByIDRequestMultiError) AllErrors() []error { return m }

// UpdateStoreByIDRequestValidationError is the validation error returned by
// UpdateStoreByIDRequest.Validate if the designated constraints aren't met.
type UpdateStoreByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStoreByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStoreByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStoreByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStoreByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStoreByIDRequestValidationError) ErrorName() string {
	return "UpdateStoreByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStoreByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStoreByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStoreByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStoreByIDRequestValidationError{}

// Validate checks the field values on UpdateStoreByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateStoreByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateStoreByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateStoreByIDReplyMultiError, or nil if none found.
func (m *UpdateStoreByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateStoreByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateStoreByIDReplyMultiError(errors)
	}

	return nil
}

// UpdateStoreByIDReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateStoreByIDReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateStoreByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateStoreByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateStoreByIDReplyMultiError) AllErrors() []error { return m }

// UpdateStoreByIDReplyValidationError is the validation error returned by
// UpdateStoreByIDReply.Validate if the designated constraints aren't met.
type UpdateStoreByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStoreByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStoreByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStoreByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStoreByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStoreByIDReplyValidationError) ErrorName() string {
	return "UpdateStoreByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStoreByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStoreByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStoreByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStoreByIDReplyValidationError{}

// Validate checks the field values on Store with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Store) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Store with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StoreMultiError, or nil if none found.
func (m *Store) ValidateAll() error {
	return m.validate(true)
}

func (m *Store) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Address

	// no validation rules for ContactPhone

	// no validation rules for ManagerID

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return StoreMultiError(errors)
	}

	return nil
}

// StoreMultiError is an error wrapping multiple validation errors returned by
// Store.ValidateAll() if the designated constraints aren't met.
type StoreMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StoreMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StoreMultiError) AllErrors() []error { return m }

// StoreValidationError is the validation error returned by Store.Validate if
// the designated constraints aren't met.
type StoreValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StoreValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StoreValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StoreValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StoreValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StoreValidationError) ErrorName() string { return "StoreValidationError" }

// Error satisfies the builtin error interface
func (e StoreValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStore.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StoreValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StoreValidationError{}

// Validate checks the field values on GetStoreByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStoreByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStoreByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStoreByIDRequestMultiError, or nil if none found.
func (m *GetStoreByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStoreByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetStoreByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetStoreByIDRequestMultiError(errors)
	}

	return nil
}

// GetStoreByIDRequestMultiError is an error wrapping multiple validation
// errors returned by GetStoreByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type GetStoreByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStoreByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStoreByIDRequestMultiError) AllErrors() []error { return m }

// GetStoreByIDRequestValidationError is the validation error returned by
// GetStoreByIDRequest.Validate if the designated constraints aren't met.
type GetStoreByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStoreByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStoreByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStoreByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStoreByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStoreByIDRequestValidationError) ErrorName() string {
	return "GetStoreByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetStoreByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStoreByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStoreByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStoreByIDRequestValidationError{}

// Validate checks the field values on GetStoreByIDReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStoreByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStoreByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStoreByIDReplyMultiError, or nil if none found.
func (m *GetStoreByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStoreByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStore()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetStoreByIDReplyValidationError{
					field:  "Store",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetStoreByIDReplyValidationError{
					field:  "Store",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStore()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStoreByIDReplyValidationError{
				field:  "Store",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetStoreByIDReplyMultiError(errors)
	}

	return nil
}

// GetStoreByIDReplyMultiError is an error wrapping multiple validation errors
// returned by GetStoreByIDReply.ValidateAll() if the designated constraints
// aren't met.
type GetStoreByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStoreByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStoreByIDReplyMultiError) AllErrors() []error { return m }

// GetStoreByIDReplyValidationError is the validation error returned by
// GetStoreByIDReply.Validate if the designated constraints aren't met.
type GetStoreByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStoreByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStoreByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStoreByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStoreByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStoreByIDReplyValidationError) ErrorName() string {
	return "GetStoreByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetStoreByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStoreByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStoreByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStoreByIDReplyValidationError{}

// Validate checks the field values on ListStoreRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListStoreRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListStoreRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListStoreRequestMultiError, or nil if none found.
func (m *ListStoreRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListStoreRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetParams()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListStoreRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListStoreRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListStoreRequestValidationError{
				field:  "Params",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListStoreRequestMultiError(errors)
	}

	return nil
}

// ListStoreRequestMultiError is an error wrapping multiple validation errors
// returned by ListStoreRequest.ValidateAll() if the designated constraints
// aren't met.
type ListStoreRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListStoreRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListStoreRequestMultiError) AllErrors() []error { return m }

// ListStoreRequestValidationError is the validation error returned by
// ListStoreRequest.Validate if the designated constraints aren't met.
type ListStoreRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStoreRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStoreRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStoreRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStoreRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStoreRequestValidationError) ErrorName() string { return "ListStoreRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListStoreRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStoreRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStoreRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStoreRequestValidationError{}

// Validate checks the field values on ListStoreReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListStoreReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListStoreReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListStoreReplyMultiError,
// or nil if none found.
func (m *ListStoreReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListStoreReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetStores() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListStoreReplyValidationError{
						field:  fmt.Sprintf("Stores[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListStoreReplyValidationError{
						field:  fmt.Sprintf("Stores[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListStoreReplyValidationError{
					field:  fmt.Sprintf("Stores[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListStoreReplyMultiError(errors)
	}

	return nil
}

// ListStoreReplyMultiError is an error wrapping multiple validation errors
// returned by ListStoreReply.ValidateAll() if the designated constraints
// aren't met.
type ListStoreReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListStoreReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListStoreReplyMultiError) AllErrors() []error { return m }

// ListStoreReplyValidationError is the validation error returned by
// ListStoreReply.Validate if the designated constraints aren't met.
type ListStoreReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStoreReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStoreReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStoreReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStoreReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStoreReplyValidationError) ErrorName() string { return "ListStoreReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListStoreReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStoreReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStoreReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStoreReplyValidationError{}
