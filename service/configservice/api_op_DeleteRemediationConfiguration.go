// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the remediation configuration.
func (c *Client) DeleteRemediationConfiguration(ctx context.Context, params *DeleteRemediationConfigurationInput, optFns ...func(*Options)) (*DeleteRemediationConfigurationOutput, error) {
	if params == nil {
		params = &DeleteRemediationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRemediationConfiguration", params, optFns, addOperationDeleteRemediationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRemediationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRemediationConfigurationInput struct {

	// The name of the AWS Config rule for which you want to delete remediation
	// configuration.
	//
	// This member is required.
	ConfigRuleName *string

	// The type of a resource.
	ResourceType *string
}

type DeleteRemediationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteRemediationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRemediationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRemediationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteRemediationConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRemediationConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteRemediationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteRemediationConfiguration",
	}
}
