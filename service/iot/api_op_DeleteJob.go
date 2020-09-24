// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a job and its related job executions. Deleting a job may take time,
// depending on the number of job executions created for the job and various other
// factors. While the job is being deleted, the status of the job will be shown as
// "DELETION_IN_PROGRESS". Attempting to delete or cancel a job whose status is
// already "DELETION_IN_PROGRESS" will result in an error. Only 10 jobs may have
// status "DELETION_IN_PROGRESS" at the same time, or a LimitExceededException will
// occur.
func (c *Client) DeleteJob(ctx context.Context, params *DeleteJobInput, optFns ...func(*Options)) (*DeleteJobOutput, error) {
	stack := middleware.NewStack("DeleteJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteJobMiddlewares(stack)
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
	addOpDeleteJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJob(options.Region), middleware.Before)
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
			OperationName: "DeleteJob",
			Err:           err,
		}
	}
	out := result.(*DeleteJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobInput struct {
	// (Optional) When true, you can delete a job which is "IN_PROGRESS". Otherwise,
	// you can only delete a job which is in a terminal state ("COMPLETED" or
	// "CANCELED") or an exception will occur. The default is false. Deleting a job
	// which is "IN_PROGRESS", will cause a device which is executing the job to be
	// unable to access job information or update the job execution status. Use caution
	// and ensure that each device executing a job which is deleted is able to recover
	// to a valid state.
	Force *bool
	// The ID of the job to be deleted. After a job deletion is completed, you may
	// reuse this jobId when you create a new job. However, this is not recommended,
	// and you must ensure that your devices are not using the jobId to refer to the
	// deleted job.
	JobId *string
}

type DeleteJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteJob",
	}
}
