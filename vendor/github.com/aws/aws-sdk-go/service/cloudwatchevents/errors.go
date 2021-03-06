// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// There is concurrent modification on a resource.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeInternalException for service response error code
	// "InternalException".
	//
	// This exception occurs due to unexpected causes.
	ErrCodeInternalException = "InternalException"

	// ErrCodeInvalidEventPatternException for service response error code
	// "InvalidEventPatternException".
	//
	// The event pattern isn't valid.
	ErrCodeInvalidEventPatternException = "InvalidEventPatternException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// The specified state isn't a valid state for an event source.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// You tried to create more resources than is allowed.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeManagedRuleException for service response error code
	// "ManagedRuleException".
	//
	// An AWS service created this rule on behalf of your account. That service
	// manages it. If you see this error in response to DeleteRule or RemoveTargets,
	// you can use the Force parameter in those calls to delete the rule or remove
	// targets from the rule. You can't modify these managed rules by using DisableRule,
	// EnableRule, PutTargets, PutRule, TagResource, or UntagResource.
	ErrCodeManagedRuleException = "ManagedRuleException"

	// ErrCodePolicyLengthExceededException for service response error code
	// "PolicyLengthExceededException".
	//
	// The event bus policy is too long. For more information, see the limits.
	ErrCodePolicyLengthExceededException = "PolicyLengthExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The resource that you're trying to create already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// An entity that you specified doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConcurrentModificationException": newErrorConcurrentModificationException,
	"InternalException":               newErrorInternalException,
	"InvalidEventPatternException":    newErrorInvalidEventPatternException,
	"InvalidStateException":           newErrorInvalidStateException,
	"LimitExceededException":          newErrorLimitExceededException,
	"ManagedRuleException":            newErrorManagedRuleException,
	"PolicyLengthExceededException":   newErrorPolicyLengthExceededException,
	"ResourceAlreadyExistsException":  newErrorResourceAlreadyExistsException,
	"ResourceNotFoundException":       newErrorResourceNotFoundException,
}
