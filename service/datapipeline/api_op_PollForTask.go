// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Task runners call PollForTask to receive a task to perform from AWS Data
// Pipeline. The task runner specifies which tasks it can perform by setting a
// value for the workerGroup parameter. The task returned can come from any of the
// pipelines that match the workerGroup value passed in by the task runner and that
// was launched using the IAM user credentials specified by the task runner. If
// tasks are ready in the work queue, PollForTask returns a response immediately.
// If no tasks are available in the queue, PollForTask uses long-polling and holds
// on to a poll connection for up to a 90 seconds, during which time the first
// newly scheduled task is handed to the task runner. To accomodate this, set the
// socket timeout in your task runner to 90 seconds. The task runner should not
// call PollForTask again on the same workerGroup until it receives a response, and
// this can take up to 90 seconds.
func (c *Client) PollForTask(ctx context.Context, params *PollForTaskInput, optFns ...func(*Options)) (*PollForTaskOutput, error) {
	stack := middleware.NewStack("PollForTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPollForTaskMiddlewares(stack)
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
	addOpPollForTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPollForTask(options.Region), middleware.Before)
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
			OperationName: "PollForTask",
			Err:           err,
		}
	}
	out := result.(*PollForTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for PollForTask.
type PollForTaskInput struct {
	// The type of task the task runner is configured to accept and process. The worker
	// group is set as a field on objects in the pipeline when they are created. You
	// can only specify a single value for workerGroup in the call to PollForTask.
	// There are no wildcard values permitted in workerGroup; the string must be an
	// exact, case-sensitive, match.
	WorkerGroup *string
	// Identity information for the EC2 instance that is hosting the task runner. You
	// can get this value from the instance using
	// http://169.254.169.254/latest/meta-data/instance-id. For more information, see
	// Instance Metadata
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AESDG-chapter-instancedata.html)
	// in the Amazon Elastic Compute Cloud User Guide. Passing in this value proves
	// that your task runner is running on an EC2 instance, and ensures the proper AWS
	// Data Pipeline service charges are applied to your pipeline.
	InstanceIdentity *types.InstanceIdentity
	// The public DNS name of the calling task runner.
	Hostname *string
}

// Contains the output of PollForTask.
type PollForTaskOutput struct {
	// The information needed to complete the task that is being assigned to the task
	// runner. One of the fields returned in this object is taskId, which contains an
	// identifier for the task being assigned. The calling task runner uses taskId in
	// subsequent calls to ReportTaskProgress () and SetTaskStatus ().
	TaskObject *types.TaskObject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPollForTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPollForTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPollForTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opPollForTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "PollForTask",
	}
}
