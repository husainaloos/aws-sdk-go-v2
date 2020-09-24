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

// Task runners call ReportTaskProgress when assigned a task to acknowledge that it
// has the task. If the web service does not receive this acknowledgement within 2
// minutes, it assigns the task in a subsequent PollForTask () call. After this
// initial acknowledgement, the task runner only needs to report progress every 15
// minutes to maintain its ownership of the task. You can change this reporting
// time from 15 minutes by specifying a reportProgressTimeout field in your
// pipeline.
func (c *Client) ReportTaskProgress(ctx context.Context, params *ReportTaskProgressInput, optFns ...func(*Options)) (*ReportTaskProgressOutput, error) {
	stack := middleware.NewStack("ReportTaskProgress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpReportTaskProgressMiddlewares(stack)
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
	addOpReportTaskProgressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReportTaskProgress(options.Region), middleware.Before)
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
			OperationName: "ReportTaskProgress",
			Err:           err,
		}
	}
	out := result.(*ReportTaskProgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ReportTaskProgress.
type ReportTaskProgressInput struct {
	// Key-value pairs that define the properties of the ReportTaskProgressInput
	// object.
	Fields []*types.Field
	// The ID of the task assigned to the task runner. This value is provided in the
	// response for PollForTask ().
	TaskId *string
}

// Contains the output of ReportTaskProgress.
type ReportTaskProgressOutput struct {
	// If true, the calling task runner should cancel processing of the task. The task
	// runner does not need to call SetTaskStatus () for canceled tasks.
	Canceled *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpReportTaskProgressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpReportTaskProgress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpReportTaskProgress{}, middleware.After)
}

func newServiceMetadataMiddleware_opReportTaskProgress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "ReportTaskProgress",
	}
}
