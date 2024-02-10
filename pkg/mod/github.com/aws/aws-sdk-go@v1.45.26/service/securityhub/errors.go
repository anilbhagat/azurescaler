// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have permission to perform the action specified in the request.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeInternalException for service response error code
	// "InternalException".
	//
	// Internal server error.
	ErrCodeInternalException = "InternalException"

	// ErrCodeInvalidAccessException for service response error code
	// "InvalidAccessException".
	//
	// The account doesn't have permission to perform this action.
	ErrCodeInvalidAccessException = "InvalidAccessException"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// The request was rejected because you supplied an invalid or out-of-range
	// value for an input parameter.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request was rejected because it attempted to create resources beyond
	// the current Amazon Web Services account or throttling limits. The error code
	// describes the limit exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// The resource specified in the request conflicts with an existing resource.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request was rejected because we can't find the specified resource.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"InternalException":         newErrorInternalException,
	"InvalidAccessException":    newErrorInvalidAccessException,
	"InvalidInputException":     newErrorInvalidInputException,
	"LimitExceededException":    newErrorLimitExceededException,
	"ResourceConflictException": newErrorResourceConflictException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
}
