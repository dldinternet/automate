// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/interservice/authz/common/subject_purge.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ErrorShouldUseV1 with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ErrorShouldUseV1) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ErrorShouldUseV1ValidationError is the validation error returned by
// ErrorShouldUseV1.Validate if the designated constraints aren't met.
type ErrorShouldUseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorShouldUseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorShouldUseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorShouldUseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorShouldUseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorShouldUseV1ValidationError) ErrorName() string { return "ErrorShouldUseV1ValidationError" }

// Error satisfies the builtin error interface
func (e ErrorShouldUseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorShouldUseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorShouldUseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorShouldUseV1ValidationError{}

// Validate checks the field values on ErrorShouldUseV2 with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ErrorShouldUseV2) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ErrorShouldUseV2ValidationError is the validation error returned by
// ErrorShouldUseV2.Validate if the designated constraints aren't met.
type ErrorShouldUseV2ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorShouldUseV2ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorShouldUseV2ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorShouldUseV2ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorShouldUseV2ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorShouldUseV2ValidationError) ErrorName() string { return "ErrorShouldUseV2ValidationError" }

// Error satisfies the builtin error interface
func (e ErrorShouldUseV2ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorShouldUseV2.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorShouldUseV2ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorShouldUseV2ValidationError{}

// Validate checks the field values on PurgeSubjectFromPoliciesReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PurgeSubjectFromPoliciesReq) Validate() error {
	if m == nil {
		return nil
	}

	if !_PurgeSubjectFromPoliciesReq_Subject_Pattern.MatchString(m.GetSubject()) {
		return PurgeSubjectFromPoliciesReqValidationError{
			field:  "Subject",
			reason: "value does not match regex pattern \"^(?:team|user):(?:local|ldap|saml):(?:[^:*]+|[*])$|^(?:(?:team|user|token|service):)?[*]$|^token:[^:*]+$|^tls:service:(?:[^:*]+:)?(?:[^:*]+|[*])$\"",
		}
	}

	return nil
}

// PurgeSubjectFromPoliciesReqValidationError is the validation error returned
// by PurgeSubjectFromPoliciesReq.Validate if the designated constraints
// aren't met.
type PurgeSubjectFromPoliciesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurgeSubjectFromPoliciesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurgeSubjectFromPoliciesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurgeSubjectFromPoliciesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurgeSubjectFromPoliciesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurgeSubjectFromPoliciesReqValidationError) ErrorName() string {
	return "PurgeSubjectFromPoliciesReqValidationError"
}

// Error satisfies the builtin error interface
func (e PurgeSubjectFromPoliciesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurgeSubjectFromPoliciesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurgeSubjectFromPoliciesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurgeSubjectFromPoliciesReqValidationError{}

var _PurgeSubjectFromPoliciesReq_Subject_Pattern = regexp.MustCompile("^(?:team|user):(?:local|ldap|saml):(?:[^:*]+|[*])$|^(?:(?:team|user|token|service):)?[*]$|^token:[^:*]+$|^tls:service:(?:[^:*]+:)?(?:[^:*]+|[*])$")

// Validate checks the field values on PurgeSubjectFromPoliciesResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PurgeSubjectFromPoliciesResp) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PurgeSubjectFromPoliciesRespValidationError is the validation error returned
// by PurgeSubjectFromPoliciesResp.Validate if the designated constraints
// aren't met.
type PurgeSubjectFromPoliciesRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurgeSubjectFromPoliciesRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurgeSubjectFromPoliciesRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurgeSubjectFromPoliciesRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurgeSubjectFromPoliciesRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurgeSubjectFromPoliciesRespValidationError) ErrorName() string {
	return "PurgeSubjectFromPoliciesRespValidationError"
}

// Error satisfies the builtin error interface
func (e PurgeSubjectFromPoliciesRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurgeSubjectFromPoliciesResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurgeSubjectFromPoliciesRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurgeSubjectFromPoliciesRespValidationError{}