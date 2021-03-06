// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CodeGenerationStatus string

// Enum values for CodeGenerationStatus
const (
	CodeGenerationStatusCreate_in_progress CodeGenerationStatus = "CREATE_IN_PROGRESS"
	CodeGenerationStatusCreate_complete    CodeGenerationStatus = "CREATE_COMPLETE"
	CodeGenerationStatusCreate_failed      CodeGenerationStatus = "CREATE_FAILED"
)

type DiscovererState string

// Enum values for DiscovererState
const (
	DiscovererStateStarted DiscovererState = "STARTED"
	DiscovererStateStopped DiscovererState = "STOPPED"
)

type Type string

// Enum values for Type
const (
	TypeOpenapi3 Type = "OpenApi3"
)
