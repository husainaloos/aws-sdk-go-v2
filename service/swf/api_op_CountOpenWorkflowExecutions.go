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
)

// Returns the number of open workflow executions within the given domain that meet
// the specified filtering criteria. This operation is eventually consistent. The
// results are best effort and may not exactly reflect recent updates and changes.
// Access Control You can use IAM policies to control this action's access to
// Amazon SWF resources as follows:
//
//     * Use a Resource element with the domain
// name to limit the action to only specified domains.
//
//     * Use an Action element
// to allow or deny permission to call this action.
//
//     * Constrain the following
// parameters by using a Condition element with the appropriate keys.
//
//         *
// tagFilter.tag: String constraint. The key is swf:tagFilter.tag.
//
//         *
// typeFilter.name: String constraint. The key is swf:typeFilter.name.
//
//         *
// typeFilter.version: String constraint. The key is swf:typeFilter.version.
//
// If
// the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) CountOpenWorkflowExecutions(ctx context.Context, params *CountOpenWorkflowExecutionsInput, optFns ...func(*Options)) (*CountOpenWorkflowExecutionsOutput, error) {
	stack := middleware.NewStack("CountOpenWorkflowExecutions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpCountOpenWorkflowExecutionsMiddlewares(stack)
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
	addOpCountOpenWorkflowExecutionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCountOpenWorkflowExecutions(options.Region), middleware.Before)
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
			OperationName: "CountOpenWorkflowExecutions",
			Err:           err,
		}
	}
	out := result.(*CountOpenWorkflowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CountOpenWorkflowExecutionsInput struct {
	// If specified, only workflow executions matching the WorkflowId in the filter are
	// counted. executionFilter, typeFilter and tagFilter are mutually exclusive. You
	// can specify at most one of these in a request.
	ExecutionFilter *types.WorkflowExecutionFilter
	// Specifies the start time criteria that workflow executions must meet in order to
	// be counted.
	StartTimeFilter *types.ExecutionTimeFilter
	// The name of the domain containing the workflow executions to count.
	Domain *string
	// If specified, only executions that have a tag that matches the filter are
	// counted. executionFilter, typeFilter and tagFilter are mutually exclusive. You
	// can specify at most one of these in a request.
	TagFilter *types.TagFilter
	// Specifies the type of the workflow executions to be counted. executionFilter,
	// typeFilter and tagFilter are mutually exclusive. You can specify at most one of
	// these in a request.
	TypeFilter *types.WorkflowTypeFilter
}

// Contains the count of workflow executions returned from
// CountOpenWorkflowExecutions () or CountClosedWorkflowExecutions ()
type CountOpenWorkflowExecutionsOutput struct {
	// If set to true, indicates that the actual count was more than the maximum
	// supported by this API and the count returned is the truncated value.
	Truncated *bool
	// The number of workflow executions.
	Count *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpCountOpenWorkflowExecutionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpCountOpenWorkflowExecutions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpCountOpenWorkflowExecutions{}, middleware.After)
}

func newServiceMetadataMiddleware_opCountOpenWorkflowExecutions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "CountOpenWorkflowExecutions",
	}
}
