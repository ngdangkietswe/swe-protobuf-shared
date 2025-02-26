// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: timetracking/time_tracking_service.proto

package timetracking

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

// Validate checks the field values on CheckInOutReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckInOutReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckInOutReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckInOutReqMultiError, or
// nil if none found.
func (m *CheckInOutReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckInOutReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Latitude

	// no validation rules for Longitude

	if len(errors) > 0 {
		return CheckInOutReqMultiError(errors)
	}

	return nil
}

// CheckInOutReqMultiError is an error wrapping multiple validation errors
// returned by CheckInOutReq.ValidateAll() if the designated constraints
// aren't met.
type CheckInOutReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckInOutReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckInOutReqMultiError) AllErrors() []error { return m }

// CheckInOutReqValidationError is the validation error returned by
// CheckInOutReq.Validate if the designated constraints aren't met.
type CheckInOutReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckInOutReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckInOutReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckInOutReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckInOutReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckInOutReqValidationError) ErrorName() string { return "CheckInOutReqValidationError" }

// Error satisfies the builtin error interface
func (e CheckInOutReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckInOutReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckInOutReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckInOutReqValidationError{}

// Validate checks the field values on CheckInOutResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckInOutResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckInOutResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckInOutRespMultiError,
// or nil if none found.
func (m *CheckInOutResp) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckInOutResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *CheckInOutResp_Data_:
		if v == nil {
			err := CheckInOutRespValidationError{
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
					errors = append(errors, CheckInOutRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckInOutRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckInOutRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CheckInOutResp_Error:
		if v == nil {
			err := CheckInOutRespValidationError{
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
					errors = append(errors, CheckInOutRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CheckInOutRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckInOutRespValidationError{
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
		return CheckInOutRespMultiError(errors)
	}

	return nil
}

// CheckInOutRespMultiError is an error wrapping multiple validation errors
// returned by CheckInOutResp.ValidateAll() if the designated constraints
// aren't met.
type CheckInOutRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckInOutRespMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckInOutRespMultiError) AllErrors() []error { return m }

// CheckInOutRespValidationError is the validation error returned by
// CheckInOutResp.Validate if the designated constraints aren't met.
type CheckInOutRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckInOutRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckInOutRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckInOutRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckInOutRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckInOutRespValidationError) ErrorName() string { return "CheckInOutRespValidationError" }

// Error satisfies the builtin error interface
func (e CheckInOutRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckInOutResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckInOutRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckInOutRespValidationError{}

// Validate checks the field values on GetTimeTrackingReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTimeTrackingReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTimeTrackingReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTimeTrackingReqMultiError, or nil if none found.
func (m *GetTimeTrackingReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTimeTrackingReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if len(errors) > 0 {
		return GetTimeTrackingReqMultiError(errors)
	}

	return nil
}

// GetTimeTrackingReqMultiError is an error wrapping multiple validation errors
// returned by GetTimeTrackingReq.ValidateAll() if the designated constraints
// aren't met.
type GetTimeTrackingReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTimeTrackingReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTimeTrackingReqMultiError) AllErrors() []error { return m }

// GetTimeTrackingReqValidationError is the validation error returned by
// GetTimeTrackingReq.Validate if the designated constraints aren't met.
type GetTimeTrackingReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTimeTrackingReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTimeTrackingReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTimeTrackingReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTimeTrackingReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTimeTrackingReqValidationError) ErrorName() string {
	return "GetTimeTrackingReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetTimeTrackingReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTimeTrackingReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTimeTrackingReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTimeTrackingReqValidationError{}

// Validate checks the field values on GetTimeTrackingResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTimeTrackingResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTimeTrackingResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTimeTrackingRespMultiError, or nil if none found.
func (m *GetTimeTrackingResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTimeTrackingResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *GetTimeTrackingResp_TimeTracking:
		if v == nil {
			err := GetTimeTrackingRespValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetTimeTracking()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTimeTrackingRespValidationError{
						field:  "TimeTracking",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTimeTrackingRespValidationError{
						field:  "TimeTracking",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTimeTracking()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTimeTrackingRespValidationError{
					field:  "TimeTracking",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *GetTimeTrackingResp_Error:
		if v == nil {
			err := GetTimeTrackingRespValidationError{
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
					errors = append(errors, GetTimeTrackingRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTimeTrackingRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTimeTrackingRespValidationError{
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
		return GetTimeTrackingRespMultiError(errors)
	}

	return nil
}

// GetTimeTrackingRespMultiError is an error wrapping multiple validation
// errors returned by GetTimeTrackingResp.ValidateAll() if the designated
// constraints aren't met.
type GetTimeTrackingRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTimeTrackingRespMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTimeTrackingRespMultiError) AllErrors() []error { return m }

// GetTimeTrackingRespValidationError is the validation error returned by
// GetTimeTrackingResp.Validate if the designated constraints aren't met.
type GetTimeTrackingRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTimeTrackingRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTimeTrackingRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTimeTrackingRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTimeTrackingRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTimeTrackingRespValidationError) ErrorName() string {
	return "GetTimeTrackingRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetTimeTrackingRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTimeTrackingResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTimeTrackingRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTimeTrackingRespValidationError{}

// Validate checks the field values on GetListTimeTrackingReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListTimeTrackingReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListTimeTrackingReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListTimeTrackingReqMultiError, or nil if none found.
func (m *GetListTimeTrackingReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListTimeTrackingReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StartDate

	// no validation rules for EndDate

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if len(errors) > 0 {
		return GetListTimeTrackingReqMultiError(errors)
	}

	return nil
}

// GetListTimeTrackingReqMultiError is an error wrapping multiple validation
// errors returned by GetListTimeTrackingReq.ValidateAll() if the designated
// constraints aren't met.
type GetListTimeTrackingReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListTimeTrackingReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListTimeTrackingReqMultiError) AllErrors() []error { return m }

// GetListTimeTrackingReqValidationError is the validation error returned by
// GetListTimeTrackingReq.Validate if the designated constraints aren't met.
type GetListTimeTrackingReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListTimeTrackingReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListTimeTrackingReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListTimeTrackingReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListTimeTrackingReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListTimeTrackingReqValidationError) ErrorName() string {
	return "GetListTimeTrackingReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetListTimeTrackingReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListTimeTrackingReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListTimeTrackingReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListTimeTrackingReqValidationError{}

// Validate checks the field values on GetListTimeTrackingResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListTimeTrackingResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListTimeTrackingResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListTimeTrackingRespMultiError, or nil if none found.
func (m *GetListTimeTrackingResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListTimeTrackingResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *GetListTimeTrackingResp_Data_:
		if v == nil {
			err := GetListTimeTrackingRespValidationError{
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
					errors = append(errors, GetListTimeTrackingRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListTimeTrackingRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListTimeTrackingRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *GetListTimeTrackingResp_Error:
		if v == nil {
			err := GetListTimeTrackingRespValidationError{
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
					errors = append(errors, GetListTimeTrackingRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListTimeTrackingRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListTimeTrackingRespValidationError{
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
		return GetListTimeTrackingRespMultiError(errors)
	}

	return nil
}

// GetListTimeTrackingRespMultiError is an error wrapping multiple validation
// errors returned by GetListTimeTrackingResp.ValidateAll() if the designated
// constraints aren't met.
type GetListTimeTrackingRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListTimeTrackingRespMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListTimeTrackingRespMultiError) AllErrors() []error { return m }

// GetListTimeTrackingRespValidationError is the validation error returned by
// GetListTimeTrackingResp.Validate if the designated constraints aren't met.
type GetListTimeTrackingRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListTimeTrackingRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListTimeTrackingRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListTimeTrackingRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListTimeTrackingRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListTimeTrackingRespValidationError) ErrorName() string {
	return "GetListTimeTrackingRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetListTimeTrackingRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListTimeTrackingResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListTimeTrackingRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListTimeTrackingRespValidationError{}

// Validate checks the field values on OverTimeReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OverTimeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OverTimeReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OverTimeReqMultiError, or
// nil if none found.
func (m *OverTimeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *OverTimeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for OvertimeHours

	// no validation rules for Reason

	// no validation rules for ApproverId

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if len(errors) > 0 {
		return OverTimeReqMultiError(errors)
	}

	return nil
}

// OverTimeReqMultiError is an error wrapping multiple validation errors
// returned by OverTimeReq.ValidateAll() if the designated constraints aren't met.
type OverTimeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OverTimeReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OverTimeReqMultiError) AllErrors() []error { return m }

// OverTimeReqValidationError is the validation error returned by
// OverTimeReq.Validate if the designated constraints aren't met.
type OverTimeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OverTimeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OverTimeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OverTimeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OverTimeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OverTimeReqValidationError) ErrorName() string { return "OverTimeReqValidationError" }

// Error satisfies the builtin error interface
func (e OverTimeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverTimeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OverTimeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OverTimeReqValidationError{}

// Validate checks the field values on ApproveOvertimeReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ApproveOvertimeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApproveOvertimeReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApproveOvertimeReqMultiError, or nil if none found.
func (m *ApproveOvertimeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ApproveOvertimeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for IsApproved

	// no validation rules for OvertimeHours

	if m.Reason != nil {
		// no validation rules for Reason
	}

	if len(errors) > 0 {
		return ApproveOvertimeReqMultiError(errors)
	}

	return nil
}

// ApproveOvertimeReqMultiError is an error wrapping multiple validation errors
// returned by ApproveOvertimeReq.ValidateAll() if the designated constraints
// aren't met.
type ApproveOvertimeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApproveOvertimeReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApproveOvertimeReqMultiError) AllErrors() []error { return m }

// ApproveOvertimeReqValidationError is the validation error returned by
// ApproveOvertimeReq.Validate if the designated constraints aren't met.
type ApproveOvertimeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApproveOvertimeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApproveOvertimeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApproveOvertimeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApproveOvertimeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApproveOvertimeReqValidationError) ErrorName() string {
	return "ApproveOvertimeReqValidationError"
}

// Error satisfies the builtin error interface
func (e ApproveOvertimeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApproveOvertimeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApproveOvertimeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApproveOvertimeReqValidationError{}

// Validate checks the field values on CheckInOutResp_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckInOutResp_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckInOutResp_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckInOutResp_DataMultiError, or nil if none found.
func (m *CheckInOutResp_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckInOutResp_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Date

	// no validation rules for CheckInTime

	// no validation rules for CheckOutTime

	// no validation rules for Status

	if len(errors) > 0 {
		return CheckInOutResp_DataMultiError(errors)
	}

	return nil
}

// CheckInOutResp_DataMultiError is an error wrapping multiple validation
// errors returned by CheckInOutResp_Data.ValidateAll() if the designated
// constraints aren't met.
type CheckInOutResp_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckInOutResp_DataMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckInOutResp_DataMultiError) AllErrors() []error { return m }

// CheckInOutResp_DataValidationError is the validation error returned by
// CheckInOutResp_Data.Validate if the designated constraints aren't met.
type CheckInOutResp_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckInOutResp_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckInOutResp_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckInOutResp_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckInOutResp_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckInOutResp_DataValidationError) ErrorName() string {
	return "CheckInOutResp_DataValidationError"
}

// Error satisfies the builtin error interface
func (e CheckInOutResp_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckInOutResp_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckInOutResp_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckInOutResp_DataValidationError{}

// Validate checks the field values on GetListTimeTrackingResp_Data with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListTimeTrackingResp_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListTimeTrackingResp_Data with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListTimeTrackingResp_DataMultiError, or nil if none found.
func (m *GetListTimeTrackingResp_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListTimeTrackingResp_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTimeTrackings() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetListTimeTrackingResp_DataValidationError{
						field:  fmt.Sprintf("TimeTrackings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListTimeTrackingResp_DataValidationError{
						field:  fmt.Sprintf("TimeTrackings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListTimeTrackingResp_DataValidationError{
					field:  fmt.Sprintf("TimeTrackings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetListTimeTrackingResp_DataMultiError(errors)
	}

	return nil
}

// GetListTimeTrackingResp_DataMultiError is an error wrapping multiple
// validation errors returned by GetListTimeTrackingResp_Data.ValidateAll() if
// the designated constraints aren't met.
type GetListTimeTrackingResp_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListTimeTrackingResp_DataMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListTimeTrackingResp_DataMultiError) AllErrors() []error { return m }

// GetListTimeTrackingResp_DataValidationError is the validation error returned
// by GetListTimeTrackingResp_Data.Validate if the designated constraints
// aren't met.
type GetListTimeTrackingResp_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListTimeTrackingResp_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListTimeTrackingResp_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListTimeTrackingResp_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListTimeTrackingResp_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListTimeTrackingResp_DataValidationError) ErrorName() string {
	return "GetListTimeTrackingResp_DataValidationError"
}

// Error satisfies the builtin error interface
func (e GetListTimeTrackingResp_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListTimeTrackingResp_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListTimeTrackingResp_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListTimeTrackingResp_DataValidationError{}
