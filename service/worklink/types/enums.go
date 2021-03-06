// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthorizationProviderType string

// Enum values for AuthorizationProviderType
const (
	AuthorizationProviderTypeSaml AuthorizationProviderType = "SAML"
)

type DeviceStatus string

// Enum values for DeviceStatus
const (
	DeviceStatusActive     DeviceStatus = "ACTIVE"
	DeviceStatusSigned_out DeviceStatus = "SIGNED_OUT"
)

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusPending_validation     DomainStatus = "PENDING_VALIDATION"
	DomainStatusAssociating            DomainStatus = "ASSOCIATING"
	DomainStatusActive                 DomainStatus = "ACTIVE"
	DomainStatusInactive               DomainStatus = "INACTIVE"
	DomainStatusDisassociating         DomainStatus = "DISASSOCIATING"
	DomainStatusDisassociated          DomainStatus = "DISASSOCIATED"
	DomainStatusFailed_to_associate    DomainStatus = "FAILED_TO_ASSOCIATE"
	DomainStatusFailed_to_disassociate DomainStatus = "FAILED_TO_DISASSOCIATE"
)

type FleetStatus string

// Enum values for FleetStatus
const (
	FleetStatusCreating         FleetStatus = "CREATING"
	FleetStatusActive           FleetStatus = "ACTIVE"
	FleetStatusDeleting         FleetStatus = "DELETING"
	FleetStatusDeleted          FleetStatus = "DELETED"
	FleetStatusFailed_to_create FleetStatus = "FAILED_TO_CREATE"
	FleetStatusFailed_to_delete FleetStatus = "FAILED_TO_DELETE"
)

type IdentityProviderType string

// Enum values for IdentityProviderType
const (
	IdentityProviderTypeSaml IdentityProviderType = "SAML"
)
