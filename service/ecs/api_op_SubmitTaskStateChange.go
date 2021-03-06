// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This action is only used by the Amazon ECS agent, and it is not intended for use
// outside of the agent. Sent to acknowledge that a task changed states.
func (c *Client) SubmitTaskStateChange(ctx context.Context, params *SubmitTaskStateChangeInput, optFns ...func(*Options)) (*SubmitTaskStateChangeOutput, error) {
	if params == nil {
		params = &SubmitTaskStateChangeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitTaskStateChange", params, optFns, addOperationSubmitTaskStateChangeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitTaskStateChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitTaskStateChangeInput struct {

	// Any attachments associated with the state change request.
	Attachments []*types.AttachmentStateChange

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// task.
	Cluster *string

	// Any containers associated with the state change request.
	Containers []*types.ContainerStateChange

	// The Unix timestamp for when the task execution stopped.
	ExecutionStoppedAt *time.Time

	// The Unix timestamp for when the container image pull began.
	PullStartedAt *time.Time

	// The Unix timestamp for when the container image pull completed.
	PullStoppedAt *time.Time

	// The reason for the state change request.
	Reason *string

	// The status of the state change request.
	Status *string

	// The task ID or full ARN of the task in the state change request.
	Task *string
}

type SubmitTaskStateChangeOutput struct {

	// Acknowledgement of the state change.
	Acknowledgment *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSubmitTaskStateChangeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSubmitTaskStateChange{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubmitTaskStateChange{}, middleware.After)
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
	addOpSubmitTaskStateChangeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitTaskStateChange(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSubmitTaskStateChange(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "SubmitTaskStateChange",
	}
}
