// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/stock/v1/final.proto

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

// Validate checks the field values on UpdateFinalRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateFinalRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateFinalRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateFinalRequestMultiError, or nil if none found.
func (m *UpdateFinalRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateFinalRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateFinalRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStock() <= 0 {
		err := UpdateFinalRequestValidationError{
			field:  "Stock",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateFinalRequestMultiError(errors)
	}

	return nil
}

// UpdateFinalRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateFinalRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateFinalRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateFinalRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateFinalRequestMultiError) AllErrors() []error { return m }

// UpdateFinalRequestValidationError is the validation error returned by
// UpdateFinalRequest.Validate if the designated constraints aren't met.
type UpdateFinalRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateFinalRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateFinalRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateFinalRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateFinalRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateFinalRequestValidationError) ErrorName() string {
	return "UpdateFinalRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateFinalRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateFinalRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateFinalRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateFinalRequestValidationError{}

// Validate checks the field values on UpdateFinalRequestReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateFinalRequestReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateFinalRequestReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateFinalRequestReplyMultiError, or nil if none found.
func (m *UpdateFinalRequestReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateFinalRequestReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateFinalRequestReplyMultiError(errors)
	}

	return nil
}

// UpdateFinalRequestReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateFinalRequestReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateFinalRequestReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateFinalRequestReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateFinalRequestReplyMultiError) AllErrors() []error { return m }

// UpdateFinalRequestReplyValidationError is the validation error returned by
// UpdateFinalRequestReply.Validate if the designated constraints aren't met.
type UpdateFinalRequestReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateFinalRequestReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateFinalRequestReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateFinalRequestReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateFinalRequestReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateFinalRequestReplyValidationError) ErrorName() string {
	return "UpdateFinalRequestReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateFinalRequestReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateFinalRequestReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateFinalRequestReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateFinalRequestReplyValidationError{}

// Validate checks the field values on QueryFinalRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QueryFinalRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryFinalRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryFinalRequestMultiError, or nil if none found.
func (m *QueryFinalRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryFinalRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := QueryFinalRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QueryFinalRequestMultiError(errors)
	}

	return nil
}

// QueryFinalRequestMultiError is an error wrapping multiple validation errors
// returned by QueryFinalRequest.ValidateAll() if the designated constraints
// aren't met.
type QueryFinalRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryFinalRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryFinalRequestMultiError) AllErrors() []error { return m }

// QueryFinalRequestValidationError is the validation error returned by
// QueryFinalRequest.Validate if the designated constraints aren't met.
type QueryFinalRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryFinalRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryFinalRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryFinalRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryFinalRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryFinalRequestValidationError) ErrorName() string {
	return "QueryFinalRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QueryFinalRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryFinalRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryFinalRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryFinalRequestValidationError{}

// Validate checks the field values on QueryFinalReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QueryFinalReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryFinalReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryFinalReplyMultiError, or nil if none found.
func (m *QueryFinalReply) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryFinalReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Stock

	if len(errors) > 0 {
		return QueryFinalReplyMultiError(errors)
	}

	return nil
}

// QueryFinalReplyMultiError is an error wrapping multiple validation errors
// returned by QueryFinalReply.ValidateAll() if the designated constraints
// aren't met.
type QueryFinalReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryFinalReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryFinalReplyMultiError) AllErrors() []error { return m }

// QueryFinalReplyValidationError is the validation error returned by
// QueryFinalReply.Validate if the designated constraints aren't met.
type QueryFinalReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryFinalReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryFinalReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryFinalReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryFinalReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryFinalReplyValidationError) ErrorName() string { return "QueryFinalReplyValidationError" }

// Error satisfies the builtin error interface
func (e QueryFinalReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryFinalReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryFinalReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryFinalReplyValidationError{}
