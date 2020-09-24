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

// Adds or removes the specified license configurations for the specified AWS
// resource. You can update the license specifications of AMIs, instances, and
// hosts. You cannot update the license specifications for launch templates and AWS
// CloudFormation templates, as they send license configurations to the operation
// that creates the resource.
func (c *Client) UpdateLicenseSpecificationsForResource(ctx context.Context, params *UpdateLicenseSpecificationsForResourceInput, optFns ...func(*Options)) (*UpdateLicenseSpecificationsForResourceOutput, error) {
	stack := middleware.NewStack("UpdateLicenseSpecificationsForResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateLicenseSpecificationsForResourceMiddlewares(stack)
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
	addOpUpdateLicenseSpecificationsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLicenseSpecificationsForResource(options.Region), middleware.Before)
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
			OperationName: "UpdateLicenseSpecificationsForResource",
			Err:           err,
		}
	}
	out := result.(*UpdateLicenseSpecificationsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLicenseSpecificationsForResourceInput struct {
	// ARNs of the license configurations to add.
	AddLicenseSpecifications []*types.LicenseSpecification
	// Amazon Resource Name (ARN) of the AWS resource.
	ResourceArn *string
	// ARNs of the license configurations to remove.
	RemoveLicenseSpecifications []*types.LicenseSpecification
}

type UpdateLicenseSpecificationsForResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateLicenseSpecificationsForResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLicenseSpecificationsForResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLicenseSpecificationsForResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateLicenseSpecificationsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "UpdateLicenseSpecificationsForResource",
	}
}
