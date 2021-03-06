// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a job execution.
func (c *Client) DescribeJobExecution(ctx context.Context, params *DescribeJobExecutionInput, optFns ...func(*Options)) (*DescribeJobExecutionOutput, error) {
	if params == nil {
		params = &DescribeJobExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobExecution", params, optFns, addOperationDescribeJobExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobExecutionInput struct {

	// The unique identifier you assigned to this job when it was created.
	//
	// This member is required.
	JobId *string

	// The name of the thing on which the job execution is running.
	//
	// This member is required.
	ThingName *string

	// A string (consisting of the digits "0" through "9" which is used to specify a
	// particular job execution on a particular device.
	ExecutionNumber *int64
}

type DescribeJobExecutionOutput struct {

	// Information about the job execution.
	Execution *types.JobExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeJobExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobExecution{}, middleware.After)
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
	addOpDescribeJobExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeJobExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeJobExecution",
	}
}
