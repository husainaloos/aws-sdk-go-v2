// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the specified workflow execution including its type
// and some statistics. This operation is eventually consistent. The results are
// best effort and may not exactly reflect recent updates and changes. Access
// Control You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
//     * Use a Resource element with the domain name to
// limit the action to only specified domains.
//
//     * Use an Action element to
// allow or deny permission to call this action.
//
//     * You cannot use an IAM
// policy to constrain this action's parameters.
//
// If the caller doesn't have
// sufficient permissions to invoke the action, or the parameter values fall
// outside the specified constraints, the action fails. The associated event
// attribute's cause parameter is set to OPERATION_NOT_PERMITTED. For details and
// example IAM policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) DescribeWorkflowExecution(ctx context.Context, params *DescribeWorkflowExecutionInput, optFns ...func(*Options)) (*DescribeWorkflowExecutionOutput, error) {
	stack := middleware.NewStack("DescribeWorkflowExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeWorkflowExecutionMiddlewares(stack)
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
	addOpDescribeWorkflowExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkflowExecution(options.Region), middleware.Before)
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
			OperationName: "DescribeWorkflowExecution",
			Err:           err,
		}
	}
	out := result.(*DescribeWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkflowExecutionInput struct {
	// The workflow execution to describe.
	Execution *types.WorkflowExecution
	// The name of the domain containing the workflow execution.
	Domain *string
}

// Contains details about a workflow execution.
type DescribeWorkflowExecutionOutput struct {
	// The configuration settings for this workflow execution including timeout values,
	// tasklist etc.
	ExecutionConfiguration *types.WorkflowExecutionConfiguration
	// The time when the last activity task was scheduled for this workflow execution.
	// You can use this information to determine if the workflow has not made progress
	// for an unusually long period of time and might require a corrective action.
	LatestActivityTaskTimestamp *time.Time
	// Information about the workflow execution.
	ExecutionInfo *types.WorkflowExecutionInfo
	// The number of tasks for this workflow execution. This includes open and closed
	// tasks of all types.
	OpenCounts *types.WorkflowExecutionOpenCounts
	// The latest executionContext provided by the decider for this workflow execution.
	// A decider can provide an executionContext (a free-form string) when closing a
	// decision task using RespondDecisionTaskCompleted ().
	LatestExecutionContext *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeWorkflowExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeWorkflowExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeWorkflowExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeWorkflowExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "DescribeWorkflowExecution",
	}
}
