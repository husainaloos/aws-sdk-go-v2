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

// Task runners call SetTaskStatus to notify AWS Data Pipeline that a task is
// completed and provide information about the final status. A task runner makes
// this call regardless of whether the task was sucessful. A task runner does not
// need to call SetTaskStatus for tasks that are canceled by the web service during
// a call to ReportTaskProgress ().
func (c *Client) SetTaskStatus(ctx context.Context, params *SetTaskStatusInput, optFns ...func(*Options)) (*SetTaskStatusOutput, error) {
	stack := middleware.NewStack("SetTaskStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSetTaskStatusMiddlewares(stack)
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
	addOpSetTaskStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetTaskStatus(options.Region), middleware.Before)
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
			OperationName: "SetTaskStatus",
			Err:           err,
		}
	}
	out := result.(*SetTaskStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetTaskStatus.
type SetTaskStatusInput struct {
	// If an error occurred during the task, this value specifies the stack trace
	// associated with the error. This value is set on the physical attempt object. It
	// is used to display error information to the user. The web service does not parse
	// this value.
	ErrorStackTrace *string
	// If FINISHED, the task successfully completed. If FAILED, the task ended
	// unsuccessfully. Preconditions use false.
	TaskStatus types.TaskStatus
	// If an error occurred during the task, this value specifies the error code. This
	// value is set on the physical attempt object. It is used to display error
	// information to the user. It should not start with string "Service_" which is
	// reserved by the system.
	ErrorId *string
	// The ID of the task assigned to the task runner. This value is provided in the
	// response for PollForTask ().
	TaskId *string
	// If an error occurred during the task, this value specifies a text description of
	// the error. This value is set on the physical attempt object. It is used to
	// display error information to the user. The web service does not parse this
	// value.
	ErrorMessage *string
}

// Contains the output of SetTaskStatus.
type SetTaskStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSetTaskStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSetTaskStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetTaskStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetTaskStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "SetTaskStatus",
	}
}
