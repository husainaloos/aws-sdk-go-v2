// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a task assigned to a maintenance window. You can't change the task
// type, but you can change the following values:
//
//     * TaskARN. For example, you
// can change a RUN_COMMAND task from AWS-RunPowerShellScript to
// AWS-RunShellScript.
//
//     * ServiceRoleArn
//
//     * TaskInvocationParameters
//
//     *
// Priority
//
//     * MaxConcurrency
//
//     * MaxErrors
//
// If a parameter is null, then
// the corresponding field is not modified. Also, if you set Replace to true, then
// all fields required by the RegisterTaskWithMaintenanceWindow () action are
// required for this request. Optional fields that aren't specified are set to
// null.
func (c *Client) UpdateMaintenanceWindowTask(ctx context.Context, params *UpdateMaintenanceWindowTaskInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowTaskOutput, error) {
	if params == nil {
		params = &UpdateMaintenanceWindowTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMaintenanceWindowTask", params, optFns, addOperationUpdateMaintenanceWindowTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMaintenanceWindowTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMaintenanceWindowTaskInput struct {

	// The maintenance window ID that contains the task to modify.
	//
	// This member is required.
	WindowId *string

	// The task ID to modify.
	//
	// This member is required.
	WindowTaskId *string

	// The new task description to specify.
	Description *string

	// The new logging location in Amazon S3 to specify. LoggingInfo has been
	// deprecated. To specify an S3 bucket to contain logs, instead use the
	// OutputS3BucketName and OutputS3KeyPrefix options in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options for
	// the supported maintenance window task types, see
	// MaintenanceWindowTaskInvocationParameters ().
	LoggingInfo *types.LoggingInfo

	// The new MaxConcurrency value you want to specify. MaxConcurrency is the number
	// of targets that are allowed to run this task in parallel.
	MaxConcurrency *string

	// The new MaxErrors value to specify. MaxErrors is the maximum number of errors
	// that are allowed before the task stops being scheduled.
	MaxErrors *string

	// The new task name to specify.
	Name *string

	// The new task priority to specify. The lower the number, the higher the priority.
	// Tasks that have the same priority are scheduled in parallel.
	Priority *int32

	// If True, then all fields that are required by the
	// RegisterTaskWithMaintenanceWndow action are also required for this API request.
	// Optional fields that are not specified are set to null.
	Replace *bool

	// The ARN of the IAM service role for Systems Manager to assume when running a
	// maintenance window task. If you do not specify a service role ARN, Systems
	// Manager uses your account's service-linked role. If no service-linked role for
	// Systems Manager exists in your account, it is created when you run
	// RegisterTaskWithMaintenanceWindow. For more information, see the following
	// topics in the in the AWS Systems Manager User Guide:
	//
	//     * Using service-linked
	// roles for Systems Manager
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/using-service-linked-roles.html#slr-permissions)
	//
	//
	// * Should I use a service-linked role or a custom service role to run maintenance
	// window tasks?
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-permissions.html#maintenance-window-tasks-service-role)
	ServiceRoleArn *string

	// The targets (either instances or tags) to modify. Instances are specified using
	// Key=instanceids,Values=instanceID_1,instanceID_2. Tags are specified using
	// Key=tag_name,Values=tag_value.
	Targets []*types.Target

	// The task ARN to modify.
	TaskArn *string

	// The parameters that the task should use during execution. Populate only the
	// fields that match the task type. All other fields should be empty.
	TaskInvocationParameters *types.MaintenanceWindowTaskInvocationParameters

	// The parameters to modify. TaskParameters has been deprecated. To specify
	// parameters to pass to a task when it runs, instead use the Parameters option in
	// the TaskInvocationParameters structure. For information about how Systems
	// Manager handles these options for the supported maintenance window task types,
	// see MaintenanceWindowTaskInvocationParameters (). The map has the following
	// format: Key: string, between 1 and 255 characters Value: an array of strings,
	// each string is between 1 and 255 characters
	TaskParameters map[string]*types.MaintenanceWindowTaskParameterValueExpression
}

type UpdateMaintenanceWindowTaskOutput struct {

	// The updated task description.
	Description *string

	// The updated logging information in Amazon S3. LoggingInfo has been deprecated.
	// To specify an S3 bucket to contain logs, instead use the OutputS3BucketName and
	// OutputS3KeyPrefix options in the TaskInvocationParameters structure. For
	// information about how Systems Manager handles these options for the supported
	// maintenance window task types, see MaintenanceWindowTaskInvocationParameters ().
	LoggingInfo *types.LoggingInfo

	// The updated MaxConcurrency value.
	MaxConcurrency *string

	// The updated MaxErrors value.
	MaxErrors *string

	// The updated task name.
	Name *string

	// The updated priority value.
	Priority *int32

	// The ARN of the IAM service role to use to publish Amazon Simple Notification
	// Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn *string

	// The updated target values.
	Targets []*types.Target

	// The updated task ARN value.
	TaskArn *string

	// The updated parameter values.
	TaskInvocationParameters *types.MaintenanceWindowTaskInvocationParameters

	// The updated parameter values. TaskParameters has been deprecated. To specify
	// parameters to pass to a task when it runs, instead use the Parameters option in
	// the TaskInvocationParameters structure. For information about how Systems
	// Manager handles these options for the supported maintenance window task types,
	// see MaintenanceWindowTaskInvocationParameters ().
	TaskParameters map[string]*types.MaintenanceWindowTaskParameterValueExpression

	// The ID of the maintenance window that was updated.
	WindowId *string

	// The task ID of the maintenance window that was updated.
	WindowTaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateMaintenanceWindowTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceWindowTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceWindowTask{}, middleware.After)
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
	addOpUpdateMaintenanceWindowTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceWindowTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateMaintenanceWindowTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateMaintenanceWindowTask",
	}
}
