// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an Amazon EKS managed node group configuration. Your node group
// continues to function during the update. The response output includes an update
// ID that you can use to track the status of your node group update with the
// DescribeUpdate () API operation. Currently you can update the Kubernetes labels
// for a node group or the scaling configuration.
func (c *Client) UpdateNodegroupConfig(ctx context.Context, params *UpdateNodegroupConfigInput, optFns ...func(*Options)) (*UpdateNodegroupConfigOutput, error) {
	stack := middleware.NewStack("UpdateNodegroupConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateNodegroupConfigMiddlewares(stack)
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
	addIdempotencyToken_opUpdateNodegroupConfigMiddleware(stack, options)
	addOpUpdateNodegroupConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNodegroupConfig(options.Region), middleware.Before)
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
			OperationName: "UpdateNodegroupConfig",
			Err:           err,
		}
	}
	out := result.(*UpdateNodegroupConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNodegroupConfigInput struct {
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string
	// The scaling configuration details for the Auto Scaling group after the update.
	ScalingConfig *types.NodegroupScalingConfig
	// The name of the managed node group to update.
	NodegroupName *string
	// The Kubernetes labels to be applied to the nodes in the node group after the
	// update.
	Labels *types.UpdateLabelsPayload
	// The name of the Amazon EKS cluster that the managed node group resides in.
	ClusterName *string
}

type UpdateNodegroupConfigOutput struct {
	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateNodegroupConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNodegroupConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNodegroupConfig{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateNodegroupConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateNodegroupConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateNodegroupConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateNodegroupConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateNodegroupConfigInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateNodegroupConfigMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateNodegroupConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateNodegroupConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "UpdateNodegroupConfig",
	}
}
