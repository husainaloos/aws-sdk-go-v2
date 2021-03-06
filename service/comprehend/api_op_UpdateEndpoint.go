// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates information about the specified endpoint.
func (c *Client) UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) {
	if params == nil {
		params = &UpdateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEndpoint", params, optFns, addOperationUpdateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointInput struct {

	// The desired number of inference units to be used by the model using this
	// endpoint. Each inference unit represents of a throughput of 100 characters per
	// second.
	//
	// This member is required.
	DesiredInferenceUnits *int32

	// The Amazon Resource Number (ARN) of the endpoint being updated.
	//
	// This member is required.
	EndpointArn *string
}

type UpdateEndpointOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpoint{}, middleware.After)
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
	addOpUpdateEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpoint(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "UpdateEndpoint",
	}
}
