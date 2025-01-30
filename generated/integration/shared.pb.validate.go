// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: integration/shared.proto

package integration

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

// Validate checks the field values on StravaAccount with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StravaAccount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StravaAccount with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StravaAccountMultiError, or
// nil if none found.
func (m *StravaAccount) ValidateAll() error {
	return m.validate(true)
}

func (m *StravaAccount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	// no validation rules for ExpiresAt

	// no validation rules for TokenType

	// no validation rules for AthleteId

	// no validation rules for Username

	// no validation rules for FirstName

	// no validation rules for LastName

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return StravaAccountMultiError(errors)
	}

	return nil
}

// StravaAccountMultiError is an error wrapping multiple validation errors
// returned by StravaAccount.ValidateAll() if the designated constraints
// aren't met.
type StravaAccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StravaAccountMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StravaAccountMultiError) AllErrors() []error { return m }

// StravaAccountValidationError is the validation error returned by
// StravaAccount.Validate if the designated constraints aren't met.
type StravaAccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StravaAccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StravaAccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StravaAccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StravaAccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StravaAccountValidationError) ErrorName() string { return "StravaAccountValidationError" }

// Error satisfies the builtin error interface
func (e StravaAccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStravaAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StravaAccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StravaAccountValidationError{}
