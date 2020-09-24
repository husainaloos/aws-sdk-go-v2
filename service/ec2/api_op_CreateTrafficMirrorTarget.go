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

// Creates a target for your Traffic Mirror session. A Traffic Mirror target is the
// destination for mirrored traffic. The Traffic Mirror source and the Traffic
// Mirror target (monitoring appliances) can be in the same VPC, or in different
// VPCs connected via VPC peering or a transit gateway. A Traffic Mirror target can
// be a network interface, or a Network Load Balancer. To use the target in a
// Traffic Mirror session, use CreateTrafficMirrorSession
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorSession.htm).
func (c *Client) CreateTrafficMirrorTarget(ctx context.Context, params *CreateTrafficMirrorTargetInput, optFns ...func(*Options)) (*CreateTrafficMirrorTargetOutput, error) {
	stack := middleware.NewStack("CreateTrafficMirrorTarget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateTrafficMirrorTargetMiddlewares(stack)
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
	addIdempotencyToken_opCreateTrafficMirrorTargetMiddleware(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficMirrorTarget(options.Region), middleware.Before)
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
			OperationName: "CreateTrafficMirrorTarget",
			Err:           err,
		}
	}
	out := result.(*CreateTrafficMirrorTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrafficMirrorTargetInput struct {
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated
	// with the target.
	NetworkLoadBalancerArn *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The description of the Traffic Mirror target.
	Description *string
	// The tags to assign to the Traffic Mirror target.
	TagSpecifications []*types.TagSpecification
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string
}

type CreateTrafficMirrorTargetOutput struct {
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
	// Information about the Traffic Mirror target.
	TrafficMirrorTarget *types.TrafficMirrorTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateTrafficMirrorTargetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateTrafficMirrorTarget{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTrafficMirrorTarget{}, middleware.After)
}

type idempotencyToken_initializeOpCreateTrafficMirrorTarget struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTrafficMirrorTarget) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTrafficMirrorTarget) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTrafficMirrorTargetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTrafficMirrorTargetInput ")
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
func addIdempotencyToken_opCreateTrafficMirrorTargetMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateTrafficMirrorTarget{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTrafficMirrorTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTrafficMirrorTarget",
	}
}
