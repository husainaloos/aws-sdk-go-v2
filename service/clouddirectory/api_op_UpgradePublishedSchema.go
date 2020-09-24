// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Upgrades a published schema under a new minor version revision using the current
// contents of DevelopmentSchemaArn.
func (c *Client) UpgradePublishedSchema(ctx context.Context, params *UpgradePublishedSchemaInput, optFns ...func(*Options)) (*UpgradePublishedSchemaOutput, error) {
	stack := middleware.NewStack("UpgradePublishedSchema", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpgradePublishedSchemaMiddlewares(stack)
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
	addOpUpgradePublishedSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpgradePublishedSchema(options.Region), middleware.Before)
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
			OperationName: "UpgradePublishedSchema",
			Err:           err,
		}
	}
	out := result.(*UpgradePublishedSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpgradePublishedSchemaInput struct {
	// Identifies the minor version of the published schema that will be created. This
	// parameter is NOT optional.
	MinorVersion *string
	// The ARN of the published schema to be upgraded.
	PublishedSchemaArn *string
	// Used for testing whether the Development schema provided is backwards
	// compatible, or not, with the publish schema provided by the user to be upgraded.
	// If schema compatibility fails, an exception would be thrown else the call would
	// succeed. This parameter is optional and defaults to false.
	DryRun *bool
	// The ARN of the development schema with the changes used for the upgrade.
	DevelopmentSchemaArn *string
}

type UpgradePublishedSchemaOutput struct {
	// The ARN of the upgraded schema that is returned as part of the response.
	UpgradedSchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpgradePublishedSchemaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpgradePublishedSchema{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpgradePublishedSchema{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpgradePublishedSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpgradePublishedSchema",
	}
}
