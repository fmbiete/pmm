// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/node.proto

package managementpb

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

	inventorypb "github.com/percona/pmm/api/inventorypb"
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

	_ = inventorypb.NodeType(0)
)

// Validate checks the field values on RegisterNodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterNodeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterNodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterNodeRequestMultiError, or nil if none found.
func (m *RegisterNodeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterNodeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeType

	if utf8.RuneCountInString(m.GetNodeName()) < 1 {
		err := RegisterNodeRequestValidationError{
			field:  "NodeName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Address

	// no validation rules for MachineId

	// no validation rules for Distro

	// no validation rules for ContainerId

	// no validation rules for ContainerName

	// no validation rules for NodeModel

	// no validation rules for Region

	// no validation rules for Az

	// no validation rules for CustomLabels

	// no validation rules for Reregister

	// no validation rules for MetricsMode

	// no validation rules for AgentPassword

	if len(errors) > 0 {
		return RegisterNodeRequestMultiError(errors)
	}

	return nil
}

// RegisterNodeRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterNodeRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterNodeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterNodeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterNodeRequestMultiError) AllErrors() []error { return m }

// RegisterNodeRequestValidationError is the validation error returned by
// RegisterNodeRequest.Validate if the designated constraints aren't met.
type RegisterNodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterNodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterNodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterNodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterNodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterNodeRequestValidationError) ErrorName() string {
	return "RegisterNodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterNodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterNodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterNodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterNodeRequestValidationError{}

// Validate checks the field values on RegisterNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterNodeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterNodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterNodeResponseMultiError, or nil if none found.
func (m *RegisterNodeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterNodeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetGenericNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterNodeResponseValidationError{
					field:  "GenericNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterNodeResponseValidationError{
					field:  "GenericNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGenericNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterNodeResponseValidationError{
				field:  "GenericNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetContainerNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterNodeResponseValidationError{
					field:  "ContainerNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterNodeResponseValidationError{
					field:  "ContainerNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContainerNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterNodeResponseValidationError{
				field:  "ContainerNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPmmAgent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterNodeResponseValidationError{
					field:  "PmmAgent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterNodeResponseValidationError{
					field:  "PmmAgent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPmmAgent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterNodeResponseValidationError{
				field:  "PmmAgent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Token

	// no validation rules for Warning

	if len(errors) > 0 {
		return RegisterNodeResponseMultiError(errors)
	}

	return nil
}

// RegisterNodeResponseMultiError is an error wrapping multiple validation
// errors returned by RegisterNodeResponse.ValidateAll() if the designated
// constraints aren't met.
type RegisterNodeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterNodeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterNodeResponseMultiError) AllErrors() []error { return m }

// RegisterNodeResponseValidationError is the validation error returned by
// RegisterNodeResponse.Validate if the designated constraints aren't met.
type RegisterNodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterNodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterNodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterNodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterNodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterNodeResponseValidationError) ErrorName() string {
	return "RegisterNodeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterNodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterNodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterNodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterNodeResponseValidationError{}
