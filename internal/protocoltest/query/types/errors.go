// Code generated by smithy-go-codegen DO NOT EDIT.
package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// This error is thrown when a request is invalid.
type ComplexError struct {
	Message *string

	TopLevel *string
	Nested   *ComplexNestedErrorData
}

func (e *ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ComplexError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ComplexError) ErrorCode() string             { return "ComplexError" }
func (e *ComplexError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ComplexError) GetTopLevel() string {
	return ptr.ToString(e.TopLevel)
}
func (e *ComplexError) HasTopLevel() bool {
	return e.TopLevel != nil
}
func (e *ComplexError) GetNested() *ComplexNestedErrorData {
	return e.Nested
}
func (e *ComplexError) HasNested() bool {
	return e.Nested != nil
}

// This error is thrown when an invalid greeting value is provided.
type InvalidGreeting struct {
	Message *string
}

func (e *InvalidGreeting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGreeting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGreeting) ErrorCode() string             { return "InvalidGreeting" }
func (e *InvalidGreeting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidGreeting) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidGreeting) HasMessage() bool {
	return e.Message != nil
}
