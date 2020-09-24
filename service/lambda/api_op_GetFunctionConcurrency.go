// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns details about the reserved concurrency configuration for a function. To
// set a concurrency limit for a function, use PutFunctionConcurrency ().
func (c *Client) GetFunctionConcurrency(ctx context.Context, params *GetFunctionConcurrencyInput, optFns ...func(*Options)) (*GetFunctionConcurrencyOutput, error) {
	stack := middleware.NewStack("GetFunctionConcurrency", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetFunctionConcurrencyMiddlewares(stack)
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
	addOpGetFunctionConcurrencyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionConcurrency(options.Region), middleware.Before)
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
			OperationName: "GetFunctionConcurrency",
			Err:           err,
		}
	}
	out := result.(*GetFunctionConcurrencyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionConcurrencyInput struct {
	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// my-function.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN -
	// 123456789012:function:my-function.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	FunctionName *string
}

type GetFunctionConcurrencyOutput struct {
	// The number of simultaneous executions that are reserved for the function.
	ReservedConcurrentExecutions *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetFunctionConcurrencyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionConcurrency{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionConcurrency{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFunctionConcurrency(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetFunctionConcurrency",
	}
}
