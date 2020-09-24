// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Purchases the Scheduled Instances with the specified schedule. Scheduled
// Instances enable you to purchase Amazon EC2 compute capacity by the hour for a
// one-year term. Before you can purchase a Scheduled Instance, you must call
// DescribeScheduledInstanceAvailability () to check for available schedules and
// obtain a purchase token. After you purchase a Scheduled Instance, you must call
// RunScheduledInstances () during each scheduled time period. After you purchase a
// Scheduled Instance, you can't cancel, modify, or resell your purchase.
func (c *Client) PurchaseScheduledInstances(ctx context.Context, params *PurchaseScheduledInstancesInput, optFns ...func(*Options)) (*PurchaseScheduledInstancesOutput, error) {
	stack := middleware.NewStack("PurchaseScheduledInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpPurchaseScheduledInstancesMiddlewares(stack)
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
	addIdempotencyToken_opPurchaseScheduledInstancesMiddleware(stack, options)
	addOpPurchaseScheduledInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseScheduledInstances(options.Region), middleware.Before)
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
			OperationName: "PurchaseScheduledInstances",
			Err:           err,
		}
	}
	out := result.(*PurchaseScheduledInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for PurchaseScheduledInstances.
type PurchaseScheduledInstancesInput struct {
	// Unique, case-sensitive identifier that ensures the idempotency of the request.
	// For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
	// The purchase requests.
	PurchaseRequests []*types.PurchaseRequest
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Contains the output of PurchaseScheduledInstances.
type PurchaseScheduledInstancesOutput struct {
	// Information about the Scheduled Instances.
	ScheduledInstanceSet []*types.ScheduledInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpPurchaseScheduledInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpPurchaseScheduledInstances{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpPurchaseScheduledInstances{}, middleware.After)
}

type idempotencyToken_initializeOpPurchaseScheduledInstances struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPurchaseScheduledInstances) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPurchaseScheduledInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PurchaseScheduledInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PurchaseScheduledInstancesInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opPurchaseScheduledInstancesMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpPurchaseScheduledInstances{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPurchaseScheduledInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "PurchaseScheduledInstances",
	}
}
