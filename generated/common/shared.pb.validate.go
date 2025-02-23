// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/shared.proto

package common

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

// Validate checks the field values on Pageable with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Pageable) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pageable with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageableMultiError, or nil
// if none found.
func (m *Pageable) ValidateAll() error {
	return m.validate(true)
}

func (m *Pageable) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Size

	// no validation rules for Sort

	// no validation rules for Direction

	// no validation rules for UnPaged

	if m.Search != nil {
		// no validation rules for Search
	}

	if len(errors) > 0 {
		return PageableMultiError(errors)
	}

	return nil
}

// PageableMultiError is an error wrapping multiple validation errors returned
// by Pageable.ValidateAll() if the designated constraints aren't met.
type PageableMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageableMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageableMultiError) AllErrors() []error { return m }

// PageableValidationError is the validation error returned by
// Pageable.Validate if the designated constraints aren't met.
type PageableValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageableValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageableValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageableValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageableValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageableValidationError) ErrorName() string { return "PageableValidationError" }

// Error satisfies the builtin error interface
func (e PageableValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageable.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageableValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageableValidationError{}

// Validate checks the field values on PageMetaData with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageMetaData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageMetaData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageMetaDataMultiError, or
// nil if none found.
func (m *PageMetaData) ValidateAll() error {
	return m.validate(true)
}

func (m *PageMetaData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Size

	// no validation rules for TotalPages

	// no validation rules for TotalElements

	// no validation rules for Next

	// no validation rules for Previous

	if len(errors) > 0 {
		return PageMetaDataMultiError(errors)
	}

	return nil
}

// PageMetaDataMultiError is an error wrapping multiple validation errors
// returned by PageMetaData.ValidateAll() if the designated constraints aren't met.
type PageMetaDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageMetaDataMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageMetaDataMultiError) AllErrors() []error { return m }

// PageMetaDataValidationError is the validation error returned by
// PageMetaData.Validate if the designated constraints aren't met.
type PageMetaDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageMetaDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageMetaDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageMetaDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageMetaDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageMetaDataValidationError) ErrorName() string { return "PageMetaDataValidationError" }

// Error satisfies the builtin error interface
func (e PageMetaDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageMetaData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageMetaDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageMetaDataValidationError{}

// Validate checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Error) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ErrorMultiError, or nil if none found.
func (m *Error) ValidateAll() error {
	return m.validate(true)
}

func (m *Error) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Details

	if len(errors) > 0 {
		return ErrorMultiError(errors)
	}

	return nil
}

// ErrorMultiError is an error wrapping multiple validation errors returned by
// Error.ValidateAll() if the designated constraints aren't met.
type ErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ErrorMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ErrorMultiError) AllErrors() []error { return m }

// ErrorValidationError is the validation error returned by Error.Validate if
// the designated constraints aren't met.
type ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorValidationError) ErrorName() string { return "ErrorValidationError" }

// Error satisfies the builtin error interface
func (e ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorValidationError{}

// Validate checks the field values on IdReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdReqMultiError, or nil if none found.
func (m *IdReq) ValidateAll() error {
	return m.validate(true)
}

func (m *IdReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IdReqMultiError(errors)
	}

	return nil
}

// IdReqMultiError is an error wrapping multiple validation errors returned by
// IdReq.ValidateAll() if the designated constraints aren't met.
type IdReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdReqMultiError) AllErrors() []error { return m }

// IdReqValidationError is the validation error returned by IdReq.Validate if
// the designated constraints aren't met.
type IdReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdReqValidationError) ErrorName() string { return "IdReqValidationError" }

// Error satisfies the builtin error interface
func (e IdReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdReqValidationError{}

// Validate checks the field values on IdsReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdsReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdsReqMultiError, or nil if none found.
func (m *IdsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *IdsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IdsReqMultiError(errors)
	}

	return nil
}

// IdsReqMultiError is an error wrapping multiple validation errors returned by
// IdsReq.ValidateAll() if the designated constraints aren't met.
type IdsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdsReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdsReqMultiError) AllErrors() []error { return m }

// IdsReqValidationError is the validation error returned by IdsReq.Validate if
// the designated constraints aren't met.
type IdsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdsReqValidationError) ErrorName() string { return "IdsReqValidationError" }

// Error satisfies the builtin error interface
func (e IdsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdsReqValidationError{}

// Validate checks the field values on EmptyReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyReqMultiError, or nil
// if none found.
func (m *EmptyReq) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyReqMultiError(errors)
	}

	return nil
}

// EmptyReqMultiError is an error wrapping multiple validation errors returned
// by EmptyReq.ValidateAll() if the designated constraints aren't met.
type EmptyReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyReqMultiError) AllErrors() []error { return m }

// EmptyReqValidationError is the validation error returned by
// EmptyReq.Validate if the designated constraints aren't met.
type EmptyReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyReqValidationError) ErrorName() string { return "EmptyReqValidationError" }

// Error satisfies the builtin error interface
func (e EmptyReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyReqValidationError{}

// Validate checks the field values on UpsertResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpsertResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpsertResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpsertRespMultiError, or
// nil if none found.
func (m *UpsertResp) ValidateAll() error {
	return m.validate(true)
}

func (m *UpsertResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *UpsertResp_Data_:
		if v == nil {
			err := UpsertRespValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetData()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpsertRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpsertRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpsertRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *UpsertResp_Error:
		if v == nil {
			err := UpsertRespValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetError()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpsertRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpsertRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpsertRespValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return UpsertRespMultiError(errors)
	}

	return nil
}

// UpsertRespMultiError is an error wrapping multiple validation errors
// returned by UpsertResp.ValidateAll() if the designated constraints aren't met.
type UpsertRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpsertRespMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpsertRespMultiError) AllErrors() []error { return m }

// UpsertRespValidationError is the validation error returned by
// UpsertResp.Validate if the designated constraints aren't met.
type UpsertRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertRespValidationError) ErrorName() string { return "UpsertRespValidationError" }

// Error satisfies the builtin error interface
func (e UpsertRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertRespValidationError{}

// Validate checks the field values on EmptyResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EmptyResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmptyResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EmptyRespMultiError, or nil
// if none found.
func (m *EmptyResp) ValidateAll() error {
	return m.validate(true)
}

func (m *EmptyResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if m.Error != nil {

		if all {
			switch v := interface{}(m.GetError()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EmptyRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EmptyRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmptyRespValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EmptyRespMultiError(errors)
	}

	return nil
}

// EmptyRespMultiError is an error wrapping multiple validation errors returned
// by EmptyResp.ValidateAll() if the designated constraints aren't met.
type EmptyRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyRespMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyRespMultiError) AllErrors() []error { return m }

// EmptyRespValidationError is the validation error returned by
// EmptyResp.Validate if the designated constraints aren't met.
type EmptyRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyRespValidationError) ErrorName() string { return "EmptyRespValidationError" }

// Error satisfies the builtin error interface
func (e EmptyRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyRespValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Email

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on Audit with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Audit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Audit with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AuditMultiError, or nil if none found.
func (m *Audit) ValidateAll() error {
	return m.validate(true)
}

func (m *Audit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for IsDeleted

	if all {
		switch v := interface{}(m.GetCreatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuditValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuditValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuditValidationError{
				field:  "CreatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuditValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuditValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuditValidationError{
				field:  "UpdatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.DeletedAt != nil {
		// no validation rules for DeletedAt
	}

	if m.DeletedBy != nil {

		if all {
			switch v := interface{}(m.GetDeletedBy()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditValidationError{
						field:  "DeletedBy",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditValidationError{
						field:  "DeletedBy",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeletedBy()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditValidationError{
					field:  "DeletedBy",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AuditMultiError(errors)
	}

	return nil
}

// AuditMultiError is an error wrapping multiple validation errors returned by
// Audit.ValidateAll() if the designated constraints aren't met.
type AuditMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuditMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuditMultiError) AllErrors() []error { return m }

// AuditValidationError is the validation error returned by Audit.Validate if
// the designated constraints aren't met.
type AuditValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuditValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuditValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuditValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuditValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuditValidationError) ErrorName() string { return "AuditValidationError" }

// Error satisfies the builtin error interface
func (e AuditValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAudit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuditValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuditValidationError{}

// Validate checks the field values on UpsertResp_Data with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpsertResp_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpsertResp_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpsertResp_DataMultiError, or nil if none found.
func (m *UpsertResp_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *UpsertResp_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return UpsertResp_DataMultiError(errors)
	}

	return nil
}

// UpsertResp_DataMultiError is an error wrapping multiple validation errors
// returned by UpsertResp_Data.ValidateAll() if the designated constraints
// aren't met.
type UpsertResp_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpsertResp_DataMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpsertResp_DataMultiError) AllErrors() []error { return m }

// UpsertResp_DataValidationError is the validation error returned by
// UpsertResp_Data.Validate if the designated constraints aren't met.
type UpsertResp_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertResp_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertResp_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertResp_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertResp_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertResp_DataValidationError) ErrorName() string { return "UpsertResp_DataValidationError" }

// Error satisfies the builtin error interface
func (e UpsertResp_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertResp_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertResp_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertResp_DataValidationError{}
