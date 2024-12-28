// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/shared.proto

package auth

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

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Email

	if m.Id != nil {
		// no validation rules for Id
	}

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
	var msgs []string
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

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReqMultiError, or nil
// if none found.
func (m *LoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if m.Otp != nil {
		// no validation rules for Otp
	}

	if len(errors) > 0 {
		return LoginReqMultiError(errors)
	}

	return nil
}

// LoginReqMultiError is an error wrapping multiple validation errors returned
// by LoginReq.ValidateAll() if the designated constraints aren't met.
type LoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReqMultiError) AllErrors() []error { return m }

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on LoginResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRespMultiError, or nil
// if none found.
func (m *LoginResp) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *LoginResp_Data_:
		if v == nil {
			err := LoginRespValidationError{
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
					errors = append(errors, LoginRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LoginRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LoginRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LoginResp_Error:
		if v == nil {
			err := LoginRespValidationError{
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
					errors = append(errors, LoginRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LoginRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LoginRespValidationError{
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
		return LoginRespMultiError(errors)
	}

	return nil
}

// LoginRespMultiError is an error wrapping multiple validation errors returned
// by LoginResp.ValidateAll() if the designated constraints aren't met.
type LoginRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRespMultiError) AllErrors() []error { return m }

// LoginRespValidationError is the validation error returned by
// LoginResp.Validate if the designated constraints aren't met.
type LoginRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRespValidationError) ErrorName() string { return "LoginRespValidationError" }

// Error satisfies the builtin error interface
func (e LoginRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRespValidationError{}

// Validate checks the field values on EnableOrDisable2FAReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EnableOrDisable2FAReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnableOrDisable2FAReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EnableOrDisable2FAReqMultiError, or nil if none found.
func (m *EnableOrDisable2FAReq) ValidateAll() error {
	return m.validate(true)
}

func (m *EnableOrDisable2FAReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enable

	if len(errors) > 0 {
		return EnableOrDisable2FAReqMultiError(errors)
	}

	return nil
}

// EnableOrDisable2FAReqMultiError is an error wrapping multiple validation
// errors returned by EnableOrDisable2FAReq.ValidateAll() if the designated
// constraints aren't met.
type EnableOrDisable2FAReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnableOrDisable2FAReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnableOrDisable2FAReqMultiError) AllErrors() []error { return m }

// EnableOrDisable2FAReqValidationError is the validation error returned by
// EnableOrDisable2FAReq.Validate if the designated constraints aren't met.
type EnableOrDisable2FAReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnableOrDisable2FAReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnableOrDisable2FAReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnableOrDisable2FAReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnableOrDisable2FAReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnableOrDisable2FAReqValidationError) ErrorName() string {
	return "EnableOrDisable2FAReqValidationError"
}

// Error satisfies the builtin error interface
func (e EnableOrDisable2FAReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnableOrDisable2FAReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnableOrDisable2FAReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnableOrDisable2FAReqValidationError{}

// Validate checks the field values on EnableOrDisable2FAResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EnableOrDisable2FAResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnableOrDisable2FAResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EnableOrDisable2FARespMultiError, or nil if none found.
func (m *EnableOrDisable2FAResp) ValidateAll() error {
	return m.validate(true)
}

func (m *EnableOrDisable2FAResp) validate(all bool) error {
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
					errors = append(errors, EnableOrDisable2FARespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnableOrDisable2FARespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnableOrDisable2FARespValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.QrCodeImageUrl != nil {
		// no validation rules for QrCodeImageUrl
	}

	if len(errors) > 0 {
		return EnableOrDisable2FARespMultiError(errors)
	}

	return nil
}

// EnableOrDisable2FARespMultiError is an error wrapping multiple validation
// errors returned by EnableOrDisable2FAResp.ValidateAll() if the designated
// constraints aren't met.
type EnableOrDisable2FARespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnableOrDisable2FARespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnableOrDisable2FARespMultiError) AllErrors() []error { return m }

// EnableOrDisable2FARespValidationError is the validation error returned by
// EnableOrDisable2FAResp.Validate if the designated constraints aren't met.
type EnableOrDisable2FARespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnableOrDisable2FARespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnableOrDisable2FARespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnableOrDisable2FARespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnableOrDisable2FARespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnableOrDisable2FARespValidationError) ErrorName() string {
	return "EnableOrDisable2FARespValidationError"
}

// Error satisfies the builtin error interface
func (e EnableOrDisable2FARespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnableOrDisable2FAResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnableOrDisable2FARespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnableOrDisable2FARespValidationError{}

// Validate checks the field values on Action with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Action) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Action with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ActionMultiError, or nil if none found.
func (m *Action) ValidateAll() error {
	return m.validate(true)
}

func (m *Action) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	if len(errors) > 0 {
		return ActionMultiError(errors)
	}

	return nil
}

// ActionMultiError is an error wrapping multiple validation errors returned by
// Action.ValidateAll() if the designated constraints aren't met.
type ActionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActionMultiError) AllErrors() []error { return m }

// ActionValidationError is the validation error returned by Action.Validate if
// the designated constraints aren't met.
type ActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActionValidationError) ErrorName() string { return "ActionValidationError" }

// Error satisfies the builtin error interface
func (e ActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActionValidationError{}

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Resource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Resource with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResourceMultiError, or nil
// if none found.
func (m *Resource) ValidateAll() error {
	return m.validate(true)
}

func (m *Resource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	if len(errors) > 0 {
		return ResourceMultiError(errors)
	}

	return nil
}

// ResourceMultiError is an error wrapping multiple validation errors returned
// by Resource.ValidateAll() if the designated constraints aren't met.
type ResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResourceMultiError) AllErrors() []error { return m }

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on Permission with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Permission) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Permission with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PermissionMultiError, or
// nil if none found.
func (m *Permission) ValidateAll() error {
	return m.validate(true)
}

func (m *Permission) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetAction()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PermissionValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PermissionValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PermissionValidationError{
				field:  "Action",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PermissionValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PermissionValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PermissionValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PermissionMultiError(errors)
	}

	return nil
}

// PermissionMultiError is an error wrapping multiple validation errors
// returned by Permission.ValidateAll() if the designated constraints aren't met.
type PermissionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PermissionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PermissionMultiError) AllErrors() []error { return m }

// PermissionValidationError is the validation error returned by
// Permission.Validate if the designated constraints aren't met.
type PermissionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PermissionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PermissionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PermissionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PermissionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PermissionValidationError) ErrorName() string { return "PermissionValidationError" }

// Error satisfies the builtin error interface
func (e PermissionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPermission.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PermissionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PermissionValidationError{}

// Validate checks the field values on LoginResp_Data with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResp_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResp_Data with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResp_DataMultiError,
// or nil if none found.
func (m *LoginResp_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResp_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TokenType

	// no validation rules for AccessToken

	// no validation rules for AccessTokenExpiresIn

	// no validation rules for RefreshToken

	// no validation rules for RefreshTokenExpiresIn

	// no validation rules for TwoFactorAuth

	if len(errors) > 0 {
		return LoginResp_DataMultiError(errors)
	}

	return nil
}

// LoginResp_DataMultiError is an error wrapping multiple validation errors
// returned by LoginResp_Data.ValidateAll() if the designated constraints
// aren't met.
type LoginResp_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResp_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResp_DataMultiError) AllErrors() []error { return m }

// LoginResp_DataValidationError is the validation error returned by
// LoginResp_Data.Validate if the designated constraints aren't met.
type LoginResp_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResp_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResp_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResp_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResp_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResp_DataValidationError) ErrorName() string { return "LoginResp_DataValidationError" }

// Error satisfies the builtin error interface
func (e LoginResp_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResp_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResp_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResp_DataValidationError{}
