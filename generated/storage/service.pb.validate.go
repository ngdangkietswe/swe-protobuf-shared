// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: storage/service.proto

package storage

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

// Validate checks the field values on PresignedURLReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PresignedURLReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PresignedURLReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PresignedURLReqMultiError, or nil if none found.
func (m *PresignedURLReq) ValidateAll() error {
	return m.validate(true)
}

func (m *PresignedURLReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ObjectName

	// no validation rules for Type

	// no validation rules for BucketName

	// no validation rules for Duration

	if len(errors) > 0 {
		return PresignedURLReqMultiError(errors)
	}

	return nil
}

// PresignedURLReqMultiError is an error wrapping multiple validation errors
// returned by PresignedURLReq.ValidateAll() if the designated constraints
// aren't met.
type PresignedURLReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PresignedURLReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PresignedURLReqMultiError) AllErrors() []error { return m }

// PresignedURLReqValidationError is the validation error returned by
// PresignedURLReq.Validate if the designated constraints aren't met.
type PresignedURLReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PresignedURLReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PresignedURLReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PresignedURLReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PresignedURLReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PresignedURLReqValidationError) ErrorName() string { return "PresignedURLReqValidationError" }

// Error satisfies the builtin error interface
func (e PresignedURLReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPresignedURLReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PresignedURLReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PresignedURLReqValidationError{}

// Validate checks the field values on PresignedURLResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PresignedURLResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PresignedURLResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PresignedURLRespMultiError, or nil if none found.
func (m *PresignedURLResp) ValidateAll() error {
	return m.validate(true)
}

func (m *PresignedURLResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *PresignedURLResp_Url:
		if v == nil {
			err := PresignedURLRespValidationError{
				field:  "Resp",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Url
	case *PresignedURLResp_Error:
		if v == nil {
			err := PresignedURLRespValidationError{
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
					errors = append(errors, PresignedURLRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PresignedURLRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PresignedURLRespValidationError{
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
		return PresignedURLRespMultiError(errors)
	}

	return nil
}

// PresignedURLRespMultiError is an error wrapping multiple validation errors
// returned by PresignedURLResp.ValidateAll() if the designated constraints
// aren't met.
type PresignedURLRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PresignedURLRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PresignedURLRespMultiError) AllErrors() []error { return m }

// PresignedURLRespValidationError is the validation error returned by
// PresignedURLResp.Validate if the designated constraints aren't met.
type PresignedURLRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PresignedURLRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PresignedURLRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PresignedURLRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PresignedURLRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PresignedURLRespValidationError) ErrorName() string { return "PresignedURLRespValidationError" }

// Error satisfies the builtin error interface
func (e PresignedURLRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPresignedURLResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PresignedURLRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PresignedURLRespValidationError{}

// Validate checks the field values on ListFileObjectReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListFileObjectReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListFileObjectReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListFileObjectReqMultiError, or nil if none found.
func (m *ListFileObjectReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListFileObjectReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BucketName

	// no validation rules for Prefix

	if len(errors) > 0 {
		return ListFileObjectReqMultiError(errors)
	}

	return nil
}

// ListFileObjectReqMultiError is an error wrapping multiple validation errors
// returned by ListFileObjectReq.ValidateAll() if the designated constraints
// aren't met.
type ListFileObjectReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListFileObjectReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListFileObjectReqMultiError) AllErrors() []error { return m }

// ListFileObjectReqValidationError is the validation error returned by
// ListFileObjectReq.Validate if the designated constraints aren't met.
type ListFileObjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFileObjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFileObjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFileObjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFileObjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFileObjectReqValidationError) ErrorName() string {
	return "ListFileObjectReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListFileObjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFileObjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFileObjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFileObjectReqValidationError{}

// Validate checks the field values on ListFileObjectResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListFileObjectResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListFileObjectResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListFileObjectRespMultiError, or nil if none found.
func (m *ListFileObjectResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ListFileObjectResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	switch v := m.Resp.(type) {
	case *ListFileObjectResp_Data_:
		if v == nil {
			err := ListFileObjectRespValidationError{
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
					errors = append(errors, ListFileObjectRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListFileObjectRespValidationError{
						field:  "Data",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListFileObjectRespValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ListFileObjectResp_Error:
		if v == nil {
			err := ListFileObjectRespValidationError{
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
					errors = append(errors, ListFileObjectRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListFileObjectRespValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListFileObjectRespValidationError{
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
		return ListFileObjectRespMultiError(errors)
	}

	return nil
}

// ListFileObjectRespMultiError is an error wrapping multiple validation errors
// returned by ListFileObjectResp.ValidateAll() if the designated constraints
// aren't met.
type ListFileObjectRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListFileObjectRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListFileObjectRespMultiError) AllErrors() []error { return m }

// ListFileObjectRespValidationError is the validation error returned by
// ListFileObjectResp.Validate if the designated constraints aren't met.
type ListFileObjectRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFileObjectRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFileObjectRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFileObjectRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFileObjectRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFileObjectRespValidationError) ErrorName() string {
	return "ListFileObjectRespValidationError"
}

// Error satisfies the builtin error interface
func (e ListFileObjectRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFileObjectResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFileObjectRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFileObjectRespValidationError{}

// Validate checks the field values on DeleteFileObjectReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteFileObjectReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteFileObjectReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteFileObjectReqMultiError, or nil if none found.
func (m *DeleteFileObjectReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteFileObjectReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ObjectName

	// no validation rules for BucketName

	if len(errors) > 0 {
		return DeleteFileObjectReqMultiError(errors)
	}

	return nil
}

// DeleteFileObjectReqMultiError is an error wrapping multiple validation
// errors returned by DeleteFileObjectReq.ValidateAll() if the designated
// constraints aren't met.
type DeleteFileObjectReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteFileObjectReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteFileObjectReqMultiError) AllErrors() []error { return m }

// DeleteFileObjectReqValidationError is the validation error returned by
// DeleteFileObjectReq.Validate if the designated constraints aren't met.
type DeleteFileObjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteFileObjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteFileObjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteFileObjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteFileObjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteFileObjectReqValidationError) ErrorName() string {
	return "DeleteFileObjectReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteFileObjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteFileObjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteFileObjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteFileObjectReqValidationError{}

// Validate checks the field values on BatchDeleteFileObjectReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BatchDeleteFileObjectReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BatchDeleteFileObjectReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BatchDeleteFileObjectReqMultiError, or nil if none found.
func (m *BatchDeleteFileObjectReq) ValidateAll() error {
	return m.validate(true)
}

func (m *BatchDeleteFileObjectReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDeleteFileObjects() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BatchDeleteFileObjectReqValidationError{
						field:  fmt.Sprintf("DeleteFileObjects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BatchDeleteFileObjectReqValidationError{
						field:  fmt.Sprintf("DeleteFileObjects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BatchDeleteFileObjectReqValidationError{
					field:  fmt.Sprintf("DeleteFileObjects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BatchDeleteFileObjectReqMultiError(errors)
	}

	return nil
}

// BatchDeleteFileObjectReqMultiError is an error wrapping multiple validation
// errors returned by BatchDeleteFileObjectReq.ValidateAll() if the designated
// constraints aren't met.
type BatchDeleteFileObjectReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BatchDeleteFileObjectReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BatchDeleteFileObjectReqMultiError) AllErrors() []error { return m }

// BatchDeleteFileObjectReqValidationError is the validation error returned by
// BatchDeleteFileObjectReq.Validate if the designated constraints aren't met.
type BatchDeleteFileObjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchDeleteFileObjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchDeleteFileObjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchDeleteFileObjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchDeleteFileObjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchDeleteFileObjectReqValidationError) ErrorName() string {
	return "BatchDeleteFileObjectReqValidationError"
}

// Error satisfies the builtin error interface
func (e BatchDeleteFileObjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchDeleteFileObjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchDeleteFileObjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchDeleteFileObjectReqValidationError{}

// Validate checks the field values on ListFileObjectResp_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListFileObjectResp_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListFileObjectResp_Data with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListFileObjectResp_DataMultiError, or nil if none found.
func (m *ListFileObjectResp_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *ListFileObjectResp_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetFileObjects() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListFileObjectResp_DataValidationError{
						field:  fmt.Sprintf("FileObjects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListFileObjectResp_DataValidationError{
						field:  fmt.Sprintf("FileObjects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListFileObjectResp_DataValidationError{
					field:  fmt.Sprintf("FileObjects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListFileObjectResp_DataMultiError(errors)
	}

	return nil
}

// ListFileObjectResp_DataMultiError is an error wrapping multiple validation
// errors returned by ListFileObjectResp_Data.ValidateAll() if the designated
// constraints aren't met.
type ListFileObjectResp_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListFileObjectResp_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListFileObjectResp_DataMultiError) AllErrors() []error { return m }

// ListFileObjectResp_DataValidationError is the validation error returned by
// ListFileObjectResp_Data.Validate if the designated constraints aren't met.
type ListFileObjectResp_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFileObjectResp_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFileObjectResp_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFileObjectResp_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFileObjectResp_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFileObjectResp_DataValidationError) ErrorName() string {
	return "ListFileObjectResp_DataValidationError"
}

// Error satisfies the builtin error interface
func (e ListFileObjectResp_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFileObjectResp_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFileObjectResp_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFileObjectResp_DataValidationError{}
