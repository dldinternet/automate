// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/interservice/authz/v2/project.proto

package v2

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

// Validate checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Project) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Id

	// no validation rules for Type

	return nil
}

// ProjectValidationError is the validation error returned by Project.Validate
// if the designated constraints aren't met.
type ProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectValidationError) ErrorName() string { return "ProjectValidationError" }

// Error satisfies the builtin error interface
func (e ProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectValidationError{}

// Validate checks the field values on CreateProjectReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateProjectReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if !_CreateProjectReq_Id_Pattern.MatchString(m.GetId()) {
		return CreateProjectReqValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[a-z0-9-]{1,64}$\"",
		}
	}

	return nil
}

// CreateProjectReqValidationError is the validation error returned by
// CreateProjectReq.Validate if the designated constraints aren't met.
type CreateProjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectReqValidationError) ErrorName() string { return "CreateProjectReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateProjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectReqValidationError{}

var _CreateProjectReq_Id_Pattern = regexp.MustCompile("^[a-z0-9-]{1,64}$")

// Validate checks the field values on CreateProjectResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateProjectResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProjectRespValidationError{
				field:  "Project",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateProjectRespValidationError is the validation error returned by
// CreateProjectResp.Validate if the designated constraints aren't met.
type CreateProjectRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectRespValidationError) ErrorName() string {
	return "CreateProjectRespValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProjectRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectRespValidationError{}

// Validate checks the field values on GetProjectReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetProjectReq) Validate() error {
	if m == nil {
		return nil
	}

	if !_GetProjectReq_Id_Pattern.MatchString(m.GetId()) {
		return GetProjectReqValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[a-z0-9-]{1,64}$\"",
		}
	}

	return nil
}

// GetProjectReqValidationError is the validation error returned by
// GetProjectReq.Validate if the designated constraints aren't met.
type GetProjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectReqValidationError) ErrorName() string { return "GetProjectReqValidationError" }

// Error satisfies the builtin error interface
func (e GetProjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectReqValidationError{}

var _GetProjectReq_Id_Pattern = regexp.MustCompile("^[a-z0-9-]{1,64}$")

// Validate checks the field values on GetProjectResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetProjectResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProjectRespValidationError{
				field:  "Project",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetProjectRespValidationError is the validation error returned by
// GetProjectResp.Validate if the designated constraints aren't met.
type GetProjectRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectRespValidationError) ErrorName() string { return "GetProjectRespValidationError" }

// Error satisfies the builtin error interface
func (e GetProjectRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectRespValidationError{}

// Validate checks the field values on ListProjectsReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListProjectsReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListProjectsReqValidationError is the validation error returned by
// ListProjectsReq.Validate if the designated constraints aren't met.
type ListProjectsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsReqValidationError) ErrorName() string { return "ListProjectsReqValidationError" }

// Error satisfies the builtin error interface
func (e ListProjectsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsReqValidationError{}

// Validate checks the field values on ListProjectsResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListProjectsResp) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProjectsRespValidationError{
					field:  fmt.Sprintf("Projects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListProjectsRespValidationError is the validation error returned by
// ListProjectsResp.Validate if the designated constraints aren't met.
type ListProjectsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsRespValidationError) ErrorName() string { return "ListProjectsRespValidationError" }

// Error satisfies the builtin error interface
func (e ListProjectsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsRespValidationError{}

// Validate checks the field values on UpdateProjectReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateProjectReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if !_UpdateProjectReq_Id_Pattern.MatchString(m.GetId()) {
		return UpdateProjectReqValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[a-z0-9-]{1,64}$\"",
		}
	}

	return nil
}

// UpdateProjectReqValidationError is the validation error returned by
// UpdateProjectReq.Validate if the designated constraints aren't met.
type UpdateProjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProjectReqValidationError) ErrorName() string { return "UpdateProjectReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateProjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProjectReqValidationError{}

var _UpdateProjectReq_Id_Pattern = regexp.MustCompile("^[a-z0-9-]{1,64}$")

// Validate checks the field values on UpdateProjectResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateProjectResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProjectRespValidationError{
				field:  "Project",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateProjectRespValidationError is the validation error returned by
// UpdateProjectResp.Validate if the designated constraints aren't met.
type UpdateProjectRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProjectRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProjectRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProjectRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProjectRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProjectRespValidationError) ErrorName() string {
	return "UpdateProjectRespValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProjectRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProjectResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProjectRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProjectRespValidationError{}

// Validate checks the field values on DeleteProjectReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteProjectReq) Validate() error {
	if m == nil {
		return nil
	}

	if !_DeleteProjectReq_Id_Pattern.MatchString(m.GetId()) {
		return DeleteProjectReqValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[a-z0-9-]{1,64}$\"",
		}
	}

	return nil
}

// DeleteProjectReqValidationError is the validation error returned by
// DeleteProjectReq.Validate if the designated constraints aren't met.
type DeleteProjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProjectReqValidationError) ErrorName() string { return "DeleteProjectReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteProjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProjectReqValidationError{}

var _DeleteProjectReq_Id_Pattern = regexp.MustCompile("^[a-z0-9-]{1,64}$")

// Validate checks the field values on DeleteProjectResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteProjectResp) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteProjectRespValidationError is the validation error returned by
// DeleteProjectResp.Validate if the designated constraints aren't met.
type DeleteProjectRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProjectRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProjectRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProjectRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProjectRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProjectRespValidationError) ErrorName() string {
	return "DeleteProjectRespValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProjectRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProjectResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProjectRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProjectRespValidationError{}

// Validate checks the field values on ListProjectRulesReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProjectRulesReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListProjectRulesReqValidationError is the validation error returned by
// ListProjectRulesReq.Validate if the designated constraints aren't met.
type ListProjectRulesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectRulesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectRulesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectRulesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectRulesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectRulesReqValidationError) ErrorName() string {
	return "ListProjectRulesReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectRulesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectRulesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectRulesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectRulesReqValidationError{}

// Validate checks the field values on ProjectCollectionRulesResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectCollectionRulesResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProjectRules

	return nil
}

// ProjectCollectionRulesRespValidationError is the validation error returned
// by ProjectCollectionRulesResp.Validate if the designated constraints aren't met.
type ProjectCollectionRulesRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectCollectionRulesRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectCollectionRulesRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectCollectionRulesRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectCollectionRulesRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectCollectionRulesRespValidationError) ErrorName() string {
	return "ProjectCollectionRulesRespValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectCollectionRulesRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectCollectionRulesResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectCollectionRulesRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectCollectionRulesRespValidationError{}

// Validate checks the field values on GetProjectRulesReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProjectRulesReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProjectId

	return nil
}

// GetProjectRulesReqValidationError is the validation error returned by
// GetProjectRulesReq.Validate if the designated constraints aren't met.
type GetProjectRulesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectRulesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectRulesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectRulesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectRulesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectRulesReqValidationError) ErrorName() string {
	return "GetProjectRulesReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectRulesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectRulesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectRulesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectRulesReqValidationError{}

// Validate checks the field values on GetProjectRulesResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProjectRulesResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRulesForProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProjectRulesRespValidationError{
				field:  "RulesForProject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetProjectRulesRespValidationError is the validation error returned by
// GetProjectRulesResp.Validate if the designated constraints aren't met.
type GetProjectRulesRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectRulesRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectRulesRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectRulesRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectRulesRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectRulesRespValidationError) ErrorName() string {
	return "GetProjectRulesRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectRulesRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectRulesResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectRulesRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectRulesRespValidationError{}

// Validate checks the field values on ProjectRules with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProjectRules) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectRulesValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProjectRulesValidationError is the validation error returned by
// ProjectRules.Validate if the designated constraints aren't met.
type ProjectRulesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectRulesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectRulesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectRulesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectRulesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectRulesValidationError) ErrorName() string { return "ProjectRulesValidationError" }

// Error satisfies the builtin error interface
func (e ProjectRulesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectRules.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectRulesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectRulesValidationError{}

// Validate checks the field values on ProjectRule with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProjectRule) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetConditions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectRuleValidationError{
					field:  fmt.Sprintf("Conditions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProjectRuleValidationError is the validation error returned by
// ProjectRule.Validate if the designated constraints aren't met.
type ProjectRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectRuleValidationError) ErrorName() string { return "ProjectRuleValidationError" }

// Error satisfies the builtin error interface
func (e ProjectRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectRuleValidationError{}

// Validate checks the field values on Condition with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Condition) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	return nil
}

// ConditionValidationError is the validation error returned by
// Condition.Validate if the designated constraints aren't met.
type ConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConditionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConditionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConditionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConditionValidationError) ErrorName() string { return "ConditionValidationError" }

// Error satisfies the builtin error interface
func (e ConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConditionValidationError{}