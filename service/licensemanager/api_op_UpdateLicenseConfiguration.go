// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the attributes of an existing license configuration. A license
// configuration is an abstraction of a customer license agreement that can be
// consumed and enforced by License Manager. Components include specifications for
// the license type (licensing by instance, socket, CPU, or vCPU), allowed tenancy
// (shared tenancy, Dedicated Instance, Dedicated Host, or all of these), host
// affinity (how long a VM must be associated with a host), and the number of
// licenses purchased and used.
func (c *Client) UpdateLicenseConfiguration(ctx context.Context, params *UpdateLicenseConfigurationInput, optFns ...func(*Options)) (*UpdateLicenseConfigurationOutput, error) {
	stack := middleware.NewStack("UpdateLicenseConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateLicenseConfigurationMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateLicenseConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLicenseConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "UpdateLicenseConfiguration",
			Err:           err,
		}
	}
	out := result.(*UpdateLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLicenseConfigurationInput struct {
	// New hard limit of the number of available licenses.
	LicenseCountHardLimit *bool
	// New product information.
	ProductInformationList []*types.ProductInformation
	// New description of the license configuration.
	Description *string
	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string
	// New license rules.
	LicenseRules []*string
	// New number of licenses managed by the license configuration.
	LicenseCount *int64
	// New status of the license configuration.
	LicenseConfigurationStatus types.LicenseConfigurationStatus
	// New name of the license configuration.
	Name *string
}

type UpdateLicenseConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateLicenseConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLicenseConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLicenseConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateLicenseConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "UpdateLicenseConfiguration",
	}
}
