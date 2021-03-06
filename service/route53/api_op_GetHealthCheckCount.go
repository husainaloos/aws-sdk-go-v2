// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the number of health checks that are associated with the current AWS
// account.
func (c *Client) GetHealthCheckCount(ctx context.Context, params *GetHealthCheckCountInput, optFns ...func(*Options)) (*GetHealthCheckCountOutput, error) {
	if params == nil {
		params = &GetHealthCheckCountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHealthCheckCount", params, optFns, addOperationGetHealthCheckCountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHealthCheckCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for the number of health checks that are associated with the current
// AWS account.
type GetHealthCheckCountInput struct {
}

// A complex type that contains the response to a GetHealthCheckCount request.
type GetHealthCheckCountOutput struct {

	// The number of health checks associated with the current AWS account.
	//
	// This member is required.
	HealthCheckCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetHealthCheckCountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetHealthCheckCount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetHealthCheckCount{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetHealthCheckCount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetHealthCheckCount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHealthCheckCount",
	}
}
