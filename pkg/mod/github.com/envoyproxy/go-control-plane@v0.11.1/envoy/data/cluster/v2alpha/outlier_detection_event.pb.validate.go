// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/cluster/v2alpha/outlier_detection_event.proto

package v2alpha

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

// Validate checks the field values on OutlierDetectionEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OutlierDetectionEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OutlierDetectionEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OutlierDetectionEventMultiError, or nil if none found.
func (m *OutlierDetectionEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *OutlierDetectionEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := OutlierEjectionType_name[int32(m.GetType())]; !ok {
		err := OutlierDetectionEventValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OutlierDetectionEventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OutlierDetectionEventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionEventValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSecsSinceLastAction()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OutlierDetectionEventValidationError{
					field:  "SecsSinceLastAction",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OutlierDetectionEventValidationError{
					field:  "SecsSinceLastAction",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSecsSinceLastAction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OutlierDetectionEventValidationError{
				field:  "SecsSinceLastAction",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetClusterName()) < 1 {
		err := OutlierDetectionEventValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetUpstreamUrl()) < 1 {
		err := OutlierDetectionEventValidationError{
			field:  "UpstreamUrl",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Action_name[int32(m.GetAction())]; !ok {
		err := OutlierDetectionEventValidationError{
			field:  "Action",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NumEjections

	// no validation rules for Enforced

	oneofEventPresent := false
	switch v := m.Event.(type) {
	case *OutlierDetectionEvent_EjectSuccessRateEvent:
		if v == nil {
			err := OutlierDetectionEventValidationError{
				field:  "Event",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofEventPresent = true

		if all {
			switch v := interface{}(m.GetEjectSuccessRateEvent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OutlierDetectionEventValidationError{
						field:  "EjectSuccessRateEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OutlierDetectionEventValidationError{
						field:  "EjectSuccessRateEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEjectSuccessRateEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OutlierDetectionEventValidationError{
					field:  "EjectSuccessRateEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *OutlierDetectionEvent_EjectConsecutiveEvent:
		if v == nil {
			err := OutlierDetectionEventValidationError{
				field:  "Event",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofEventPresent = true

		if all {
			switch v := interface{}(m.GetEjectConsecutiveEvent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OutlierDetectionEventValidationError{
						field:  "EjectConsecutiveEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OutlierDetectionEventValidationError{
						field:  "EjectConsecutiveEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEjectConsecutiveEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OutlierDetectionEventValidationError{
					field:  "EjectConsecutiveEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *OutlierDetectionEvent_EjectFailurePercentageEvent:
		if v == nil {
			err := OutlierDetectionEventValidationError{
				field:  "Event",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofEventPresent = true

		if all {
			switch v := interface{}(m.GetEjectFailurePercentageEvent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OutlierDetectionEventValidationError{
						field:  "EjectFailurePercentageEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OutlierDetectionEventValidationError{
						field:  "EjectFailurePercentageEvent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEjectFailurePercentageEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OutlierDetectionEventValidationError{
					field:  "EjectFailurePercentageEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofEventPresent {
		err := OutlierDetectionEventValidationError{
			field:  "Event",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OutlierDetectionEventMultiError(errors)
	}

	return nil
}

// OutlierDetectionEventMultiError is an error wrapping multiple validation
// errors returned by OutlierDetectionEvent.ValidateAll() if the designated
// constraints aren't met.
type OutlierDetectionEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OutlierDetectionEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OutlierDetectionEventMultiError) AllErrors() []error { return m }

// OutlierDetectionEventValidationError is the validation error returned by
// OutlierDetectionEvent.Validate if the designated constraints aren't met.
type OutlierDetectionEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutlierDetectionEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutlierDetectionEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutlierDetectionEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutlierDetectionEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutlierDetectionEventValidationError) ErrorName() string {
	return "OutlierDetectionEventValidationError"
}

// Error satisfies the builtin error interface
func (e OutlierDetectionEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutlierDetectionEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutlierDetectionEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutlierDetectionEventValidationError{}

// Validate checks the field values on OutlierEjectSuccessRate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OutlierEjectSuccessRate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OutlierEjectSuccessRate with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OutlierEjectSuccessRateMultiError, or nil if none found.
func (m *OutlierEjectSuccessRate) ValidateAll() error {
	return m.validate(true)
}

func (m *OutlierEjectSuccessRate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetHostSuccessRate() > 100 {
		err := OutlierEjectSuccessRateValidationError{
			field:  "HostSuccessRate",
			reason: "value must be less than or equal to 100",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetClusterAverageSuccessRate() > 100 {
		err := OutlierEjectSuccessRateValidationError{
			field:  "ClusterAverageSuccessRate",
			reason: "value must be less than or equal to 100",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetClusterSuccessRateEjectionThreshold() > 100 {
		err := OutlierEjectSuccessRateValidationError{
			field:  "ClusterSuccessRateEjectionThreshold",
			reason: "value must be less than or equal to 100",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OutlierEjectSuccessRateMultiError(errors)
	}

	return nil
}

// OutlierEjectSuccessRateMultiError is an error wrapping multiple validation
// errors returned by OutlierEjectSuccessRate.ValidateAll() if the designated
// constraints aren't met.
type OutlierEjectSuccessRateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OutlierEjectSuccessRateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OutlierEjectSuccessRateMultiError) AllErrors() []error { return m }

// OutlierEjectSuccessRateValidationError is the validation error returned by
// OutlierEjectSuccessRate.Validate if the designated constraints aren't met.
type OutlierEjectSuccessRateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutlierEjectSuccessRateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutlierEjectSuccessRateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutlierEjectSuccessRateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutlierEjectSuccessRateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutlierEjectSuccessRateValidationError) ErrorName() string {
	return "OutlierEjectSuccessRateValidationError"
}

// Error satisfies the builtin error interface
func (e OutlierEjectSuccessRateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutlierEjectSuccessRate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutlierEjectSuccessRateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutlierEjectSuccessRateValidationError{}

// Validate checks the field values on OutlierEjectConsecutive with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OutlierEjectConsecutive) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OutlierEjectConsecutive with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OutlierEjectConsecutiveMultiError, or nil if none found.
func (m *OutlierEjectConsecutive) ValidateAll() error {
	return m.validate(true)
}

func (m *OutlierEjectConsecutive) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OutlierEjectConsecutiveMultiError(errors)
	}

	return nil
}

// OutlierEjectConsecutiveMultiError is an error wrapping multiple validation
// errors returned by OutlierEjectConsecutive.ValidateAll() if the designated
// constraints aren't met.
type OutlierEjectConsecutiveMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OutlierEjectConsecutiveMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OutlierEjectConsecutiveMultiError) AllErrors() []error { return m }

// OutlierEjectConsecutiveValidationError is the validation error returned by
// OutlierEjectConsecutive.Validate if the designated constraints aren't met.
type OutlierEjectConsecutiveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutlierEjectConsecutiveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutlierEjectConsecutiveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutlierEjectConsecutiveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutlierEjectConsecutiveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutlierEjectConsecutiveValidationError) ErrorName() string {
	return "OutlierEjectConsecutiveValidationError"
}

// Error satisfies the builtin error interface
func (e OutlierEjectConsecutiveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutlierEjectConsecutive.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutlierEjectConsecutiveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutlierEjectConsecutiveValidationError{}

// Validate checks the field values on OutlierEjectFailurePercentage with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OutlierEjectFailurePercentage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OutlierEjectFailurePercentage with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// OutlierEjectFailurePercentageMultiError, or nil if none found.
func (m *OutlierEjectFailurePercentage) ValidateAll() error {
	return m.validate(true)
}

func (m *OutlierEjectFailurePercentage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetHostSuccessRate() > 100 {
		err := OutlierEjectFailurePercentageValidationError{
			field:  "HostSuccessRate",
			reason: "value must be less than or equal to 100",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OutlierEjectFailurePercentageMultiError(errors)
	}

	return nil
}

// OutlierEjectFailurePercentageMultiError is an error wrapping multiple
// validation errors returned by OutlierEjectFailurePercentage.ValidateAll()
// if the designated constraints aren't met.
type OutlierEjectFailurePercentageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OutlierEjectFailurePercentageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OutlierEjectFailurePercentageMultiError) AllErrors() []error { return m }

// OutlierEjectFailurePercentageValidationError is the validation error
// returned by OutlierEjectFailurePercentage.Validate if the designated
// constraints aren't met.
type OutlierEjectFailurePercentageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutlierEjectFailurePercentageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutlierEjectFailurePercentageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutlierEjectFailurePercentageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutlierEjectFailurePercentageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutlierEjectFailurePercentageValidationError) ErrorName() string {
	return "OutlierEjectFailurePercentageValidationError"
}

// Error satisfies the builtin error interface
func (e OutlierEjectFailurePercentageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutlierEjectFailurePercentage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutlierEjectFailurePercentageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutlierEjectFailurePercentageValidationError{}
