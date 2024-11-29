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