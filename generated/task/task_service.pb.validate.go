// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: task/task_service.proto

package task

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

// Validate checks the field values on GetTaskResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTaskResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTaskRespMultiError, or
// nil if none found.
func (m *GetTaskResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *GetTaskResp_Task:
		if v == nil {
			err := GetTaskRespValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetTask()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTaskRespValidationError{
						field:  "Task",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTaskRespValidationError{
						field:  "Task",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTaskRespValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *GetTaskResp_Error:
		if v == nil {
			err := GetTaskRespValidationError{
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
					errors = append(errors, GetTaskRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTaskRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTaskRespValidationError{
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
		return GetTaskRespMultiError(errors)
	}

	return nil
}

// GetTaskRespMultiError is an error wrapping multiple validation errors
// returned by GetTaskResp.ValidateAll() if the designated constraints aren't met.
type GetTaskRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskRespMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskRespMultiError) AllErrors() []error { return m }

// GetTaskRespValidationError is the validation error returned by
// GetTaskResp.Validate if the designated constraints aren't met.
type GetTaskRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskRespValidationError) ErrorName() string { return "GetTaskRespValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskRespValidationError{}

// Validate checks the field values on UpsertTaskReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpsertTaskReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpsertTaskReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpsertTaskReqMultiError, or
// nil if none found.
func (m *UpsertTaskReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UpsertTaskReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for Status

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.ReporterId != nil {
		// no validation rules for ReporterId
	}

	if m.AssigneeId != nil {
		// no validation rules for AssigneeId
	}

	if len(errors) > 0 {
		return UpsertTaskReqMultiError(errors)
	}

	return nil
}

// UpsertTaskReqMultiError is an error wrapping multiple validation errors
// returned by UpsertTaskReq.ValidateAll() if the designated constraints
// aren't met.
type UpsertTaskReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpsertTaskReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpsertTaskReqMultiError) AllErrors() []error { return m }

// UpsertTaskReqValidationError is the validation error returned by
// UpsertTaskReq.Validate if the designated constraints aren't met.
type UpsertTaskReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertTaskReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertTaskReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertTaskReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertTaskReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertTaskReqValidationError) ErrorName() string { return "UpsertTaskReqValidationError" }

// Error satisfies the builtin error interface
func (e UpsertTaskReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertTaskReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertTaskReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertTaskReqValidationError{}

// Validate checks the field values on ListTaskReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListTaskReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTaskReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListTaskReqMultiError, or
// nil if none found.
func (m *ListTaskReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTaskReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPageable()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListTaskReqValidationError{
					field:  "Pageable",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTaskReqValidationError{
					field:  "Pageable",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPageable()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTaskReqValidationError{
				field:  "Pageable",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.CreatedById != nil {
		// no validation rules for CreatedById
	}

	if m.ReporterId != nil {
		// no validation rules for ReporterId
	}

	if m.AssigneeId != nil {
		// no validation rules for AssigneeId
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Search != nil {
		// no validation rules for Search
	}

	if len(errors) > 0 {
		return ListTaskReqMultiError(errors)
	}

	return nil
}

// ListTaskReqMultiError is an error wrapping multiple validation errors
// returned by ListTaskReq.ValidateAll() if the designated constraints aren't met.
type ListTaskReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTaskReqMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTaskReqMultiError) AllErrors() []error { return m }

// ListTaskReqValidationError is the validation error returned by
// ListTaskReq.Validate if the designated constraints aren't met.
type ListTaskReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTaskReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTaskReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTaskReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTaskReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTaskReqValidationError) ErrorName() string { return "ListTaskReqValidationError" }

// Error satisfies the builtin error interface
func (e ListTaskReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTaskReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTaskReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTaskReqValidationError{}

// Validate checks the field values on ListTaskResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListTaskResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTaskResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListTaskRespMultiError, or
// nil if none found.
func (m *ListTaskResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTaskResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *ListTaskResp_Data_:
		if v == nil {
			err := ListTaskRespValidationError{
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
					errors = append(errors, ListTaskRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTaskRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTaskRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ListTaskResp_Error:
		if v == nil {
			err := ListTaskRespValidationError{
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
					errors = append(errors, ListTaskRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTaskRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTaskRespValidationError{
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
		return ListTaskRespMultiError(errors)
	}

	return nil
}

// ListTaskRespMultiError is an error wrapping multiple validation errors
// returned by ListTaskResp.ValidateAll() if the designated constraints aren't met.
type ListTaskRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTaskRespMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTaskRespMultiError) AllErrors() []error { return m }

// ListTaskRespValidationError is the validation error returned by
// ListTaskResp.Validate if the designated constraints aren't met.
type ListTaskRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTaskRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTaskRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTaskRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTaskRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTaskRespValidationError) ErrorName() string { return "ListTaskRespValidationError" }

// Error satisfies the builtin error interface
func (e ListTaskRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTaskResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTaskRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTaskRespValidationError{}

// Validate checks the field values on ListTaskResp_Data with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTaskResp_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTaskResp_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTaskResp_DataMultiError, or nil if none found.
func (m *ListTaskResp_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTaskResp_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTasks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListTaskResp_DataValidationError{
						field:  fmt.Sprintf("Tasks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTaskResp_DataValidationError{
						field:  fmt.Sprintf("Tasks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTaskResp_DataValidationError{
					field:  fmt.Sprintf("Tasks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPageMetaData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListTaskResp_DataValidationError{
					field:  "PageMetaData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTaskResp_DataValidationError{
					field:  "PageMetaData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPageMetaData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTaskResp_DataValidationError{
				field:  "PageMetaData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListTaskResp_DataMultiError(errors)
	}

	return nil
}

// ListTaskResp_DataMultiError is an error wrapping multiple validation errors
// returned by ListTaskResp_Data.ValidateAll() if the designated constraints
// aren't met.
type ListTaskResp_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTaskResp_DataMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTaskResp_DataMultiError) AllErrors() []error { return m }

// ListTaskResp_DataValidationError is the validation error returned by
// ListTaskResp_Data.Validate if the designated constraints aren't met.
type ListTaskResp_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTaskResp_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTaskResp_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTaskResp_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTaskResp_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTaskResp_DataValidationError) ErrorName() string {
	return "ListTaskResp_DataValidationError"
}

// Error satisfies the builtin error interface
func (e ListTaskResp_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTaskResp_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTaskResp_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTaskResp_DataValidationError{}
