// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a task. A task is a set of two locations (source and destination) and a
// set of Options that you use to control the behavior of a task. If you don't
// specify Options when you create a task, AWS DataSync populates them with service
// defaults. When you create a task, it first enters the CREATING state. During
// CREATING AWS DataSync attempts to mount the on-premises Network File System
// (NFS) location. The task transitions to the AVAILABLE state without waiting for
// the AWS location to become mounted. If required, AWS DataSync mounts the AWS
// location before each task execution. If an agent that is associated with a
// source (NFS) location goes offline, the task transitions to the UNAVAILABLE
// status. If the status of the task remains in the CREATING status for more than a
// few minutes, it means that your agent might be having trouble mounting the
// source NFS file system. Check the task's ErrorCode and ErrorDetail. Mount issues
// are often caused by either a misconfigured firewall or a mistyped NFS server
// host name.
func (c *Client) CreateTask(ctx context.Context, params *CreateTaskInput, optFns ...func(*Options)) (*CreateTaskOutput, error) {
	if params == nil {
		params = &CreateTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTask", params, optFns, addOperationCreateTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateTaskRequest
type CreateTaskInput struct {

	// The Amazon Resource Name (ARN) of an AWS storage resource's location.
	//
	// This member is required.
	DestinationLocationArn *string

	// The Amazon Resource Name (ARN) of the source location for the task.
	//
	// This member is required.
	SourceLocationArn *string

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is used
	// to monitor and log events in the task.
	CloudWatchLogGroupArn *string

	// A list of filter rules that determines which files to exclude from a task. The
	// list should contain a single filter string that consists of the patterns to
	// exclude. The patterns are delimited by "|" (that is, a pipe), for example,
	// "/folder1|/folder2"
	Excludes []*types.FilterRule

	// The name of a task. This value is a text reference that is used to identify the
	// task in the console.
	Name *string

	// The set of configuration options that control the behavior of a single execution
	// of the task that occurs when you call StartTaskExecution. You can configure
	// these options to preserve metadata such as user ID (UID) and group ID (GID),
	// file permissions, data integrity verification, and so on. For each individual
	// task execution, you can override these options by specifying the OverrideOptions
	// before starting the task execution. For more information, see the operation.
	Options *types.Options

	// Specifies a schedule used to periodically transfer files from a source to a
	// destination location. The schedule should be specified in UTC time. For more
	// information, see task-scheduling ().
	Schedule *types.TaskSchedule

	// The key-value pair that represents the tag that you want to add to the resource.
	// The value can be an empty string.
	Tags []*types.TagListEntry
}

// CreateTaskResponse
type CreateTaskOutput struct {

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTask{}, middleware.After)
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
	addOpCreateTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateTask",
	}
}
