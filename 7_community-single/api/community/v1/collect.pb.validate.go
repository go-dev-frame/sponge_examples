// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/community/v1/collect.proto

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

// Validate checks the field values on CollectInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CollectInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CollectInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CollectInfoMultiError, or
// nil if none found.
func (m *CollectInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *CollectInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for PostId

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return CollectInfoMultiError(errors)
	}

	return nil
}

// CollectInfoMultiError is an error wrapping multiple validation errors
// returned by CollectInfo.ValidateAll() if the designated constraints aren't met.
type CollectInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CollectInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CollectInfoMultiError) AllErrors() []error { return m }

// CollectInfoValidationError is the validation error returned by
// CollectInfo.Validate if the designated constraints aren't met.
type CollectInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CollectInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CollectInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CollectInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CollectInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CollectInfoValidationError) ErrorName() string { return "CollectInfoValidationError" }

// Error satisfies the builtin error interface
func (e CollectInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCollectInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CollectInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CollectInfoValidationError{}

// Validate checks the field values on CreateCollectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCollectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCollectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCollectRequestMultiError, or nil if none found.
func (m *CreateCollectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCollectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := CreateCollectRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPostId() < 1 {
		err := CreateCollectRequestValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateCollectRequestMultiError(errors)
	}

	return nil
}

// CreateCollectRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCollectRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCollectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCollectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCollectRequestMultiError) AllErrors() []error { return m }

// CreateCollectRequestValidationError is the validation error returned by
// CreateCollectRequest.Validate if the designated constraints aren't met.
type CreateCollectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCollectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCollectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCollectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCollectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCollectRequestValidationError) ErrorName() string {
	return "CreateCollectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCollectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCollectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCollectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCollectRequestValidationError{}

// Validate checks the field values on CreateCollectReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCollectReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCollectReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCollectReplyMultiError, or nil if none found.
func (m *CreateCollectReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCollectReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateCollectReplyMultiError(errors)
	}

	return nil
}

// CreateCollectReplyMultiError is an error wrapping multiple validation errors
// returned by CreateCollectReply.ValidateAll() if the designated constraints
// aren't met.
type CreateCollectReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCollectReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCollectReplyMultiError) AllErrors() []error { return m }

// CreateCollectReplyValidationError is the validation error returned by
// CreateCollectReply.Validate if the designated constraints aren't met.
type CreateCollectReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCollectReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCollectReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCollectReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCollectReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCollectReplyValidationError) ErrorName() string {
	return "CreateCollectReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCollectReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCollectReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCollectReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCollectReplyValidationError{}

// Validate checks the field values on DeleteCollectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCollectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCollectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCollectRequestMultiError, or nil if none found.
func (m *DeleteCollectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCollectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 1 {
		err := DeleteCollectRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPostId() < 1 {
		err := DeleteCollectRequestValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteCollectRequestMultiError(errors)
	}

	return nil
}

// DeleteCollectRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteCollectRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteCollectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCollectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCollectRequestMultiError) AllErrors() []error { return m }

// DeleteCollectRequestValidationError is the validation error returned by
// DeleteCollectRequest.Validate if the designated constraints aren't met.
type DeleteCollectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCollectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCollectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCollectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCollectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCollectRequestValidationError) ErrorName() string {
	return "DeleteCollectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCollectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCollectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCollectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCollectRequestValidationError{}

// Validate checks the field values on DeleteCollectReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCollectReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCollectReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCollectReplyMultiError, or nil if none found.
func (m *DeleteCollectReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCollectReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteCollectReplyMultiError(errors)
	}

	return nil
}

// DeleteCollectReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteCollectReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteCollectReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCollectReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCollectReplyMultiError) AllErrors() []error { return m }

// DeleteCollectReplyValidationError is the validation error returned by
// DeleteCollectReply.Validate if the designated constraints aren't met.
type DeleteCollectReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCollectReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCollectReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCollectReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCollectReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCollectReplyValidationError) ErrorName() string {
	return "DeleteCollectReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCollectReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCollectReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCollectReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCollectReplyValidationError{}

// Validate checks the field values on ListCollectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCollectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCollectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCollectRequestMultiError, or nil if none found.
func (m *ListCollectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCollectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := ListCollectRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() < 0 {
		err := ListCollectRequestValidationError{
			field:  "Page",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLimit(); val <= 0 || val > 100 {
		err := ListCollectRequestValidationError{
			field:  "Limit",
			reason: "value must be inside range (0, 100]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListCollectRequestMultiError(errors)
	}

	return nil
}

// ListCollectRequestMultiError is an error wrapping multiple validation errors
// returned by ListCollectRequest.ValidateAll() if the designated constraints
// aren't met.
type ListCollectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCollectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCollectRequestMultiError) AllErrors() []error { return m }

// ListCollectRequestValidationError is the validation error returned by
// ListCollectRequest.Validate if the designated constraints aren't met.
type ListCollectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCollectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCollectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCollectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCollectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCollectRequestValidationError) ErrorName() string {
	return "ListCollectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCollectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCollectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCollectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCollectRequestValidationError{}

// Validate checks the field values on ListCollectReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCollectReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCollectReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCollectReplyMultiError, or nil if none found.
func (m *ListCollectReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCollectReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCollects() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCollectReplyValidationError{
						field:  fmt.Sprintf("Collects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCollectReplyValidationError{
						field:  fmt.Sprintf("Collects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCollectReplyValidationError{
					field:  fmt.Sprintf("Collects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListCollectReplyMultiError(errors)
	}

	return nil
}

// ListCollectReplyMultiError is an error wrapping multiple validation errors
// returned by ListCollectReply.ValidateAll() if the designated constraints
// aren't met.
type ListCollectReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCollectReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCollectReplyMultiError) AllErrors() []error { return m }

// ListCollectReplyValidationError is the validation error returned by
// ListCollectReply.Validate if the designated constraints aren't met.
type ListCollectReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCollectReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCollectReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCollectReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCollectReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCollectReplyValidationError) ErrorName() string { return "ListCollectReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListCollectReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCollectReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCollectReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCollectReplyValidationError{}
