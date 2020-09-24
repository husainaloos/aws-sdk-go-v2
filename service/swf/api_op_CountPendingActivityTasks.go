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

// Returns the estimated number of activity tasks in the specified task list. The
// count returned is an approximation and isn't guaranteed to be exact. If you
// specify a task list that no activity task was ever scheduled in then 0 is
// returned. Access Control You can use IAM policies to control this action's
// access to Amazon SWF resources as follows:
//
//     * Use a Resource element with
// the domain name to limit the action to only specified domains.
//
//     * Use an
// Action element to allow or deny permission to call this action.
//
//     * Constrain
// the taskList.name parameter by using a Condition element with the
// swf:taskList.name key to allow the action to access only certain task lists.
//
// If
// the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) CountPendingActivityTasks(ctx context.Context, params *CountPendingActivityTasksInput, optFns ...func(*Options)) (*CountPendingActivityTasksOutput, error) {
	stack := middleware.NewStack("CountPendingActivityTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpCountPendingActivityTasksMiddlewares(stack)
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
	addOpCountPendingActivityTasksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCountPendingActivityTasks(options.Region), middleware.Before)
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
			OperationName: "CountPendingActivityTasks",
			Err:           err,
		}
	}
	out := result.(*CountPendingActivityTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CountPendingActivityTasksInput struct {
	// The name of the task list.
	TaskList *types.TaskList
	// The name of the domain that contains the task list.
	Domain *string
}

// Contains the count of tasks in a task list.
type CountPendingActivityTasksOutput struct {
	// If set to true, indicates that the actual count was more than the maximum
	// supported by this API and the count returned is the truncated value.
	Truncated *bool
	// The number of tasks in the task list.
	Count *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpCountPendingActivityTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpCountPendingActivityTasks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpCountPendingActivityTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opCountPendingActivityTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "CountPendingActivityTasks",
	}
}
