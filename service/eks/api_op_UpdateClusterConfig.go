// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an Amazon EKS cluster configuration. Your cluster continues to function
// during the update. The response output includes an update ID that you can use to
// track the status of your cluster update with the DescribeUpdate () API
// operation. You can use this API operation to enable or disable exporting the
// Kubernetes control plane logs for your cluster to CloudWatch Logs. By default,
// cluster control plane logs aren't exported to CloudWatch Logs. For more
// information, see Amazon EKS Cluster Control Plane Logs
// (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in
// the Amazon EKS User Guide . CloudWatch Logs ingestion, archive storage, and data
// scanning rates apply to exported control plane logs. For more information, see
// Amazon CloudWatch Pricing (http://aws.amazon.com/cloudwatch/pricing/). You can
// also use this API operation to enable or disable public and private access to
// your cluster's Kubernetes API server endpoint. By default, public access is
// enabled, and private access is disabled. For more information, see Amazon EKS
// Cluster Endpoint Access Control
// (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the
// Amazon EKS User Guide . At this time, you can not update the subnets or security
// group IDs for an existing cluster. Cluster updates are asynchronous, and they
// should finish within a few minutes. During an update, the cluster status moves
// to UPDATING (this status transition is eventually consistent). When the update
// is complete (either Failed or Successful), the cluster status moves to Active.
func (c *Client) UpdateClusterConfig(ctx context.Context, params *UpdateClusterConfigInput, optFns ...func(*Options)) (*UpdateClusterConfigOutput, error) {
	if params == nil {
		params = &UpdateClusterConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClusterConfig", params, optFns, addOperationUpdateClusterConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterConfigInput struct {

	// The name of the Amazon EKS cluster to update.
	//
	// This member is required.
	Name *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// Enable or disable exporting the Kubernetes control plane logs for your cluster
	// to CloudWatch Logs. By default, cluster control plane logs aren't exported to
	// CloudWatch Logs. For more information, see Amazon EKS Cluster Control Plane Logs
	// (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in
	// the Amazon EKS User Guide . CloudWatch Logs ingestion, archive storage, and data
	// scanning rates apply to exported control plane logs. For more information, see
	// Amazon CloudWatch Pricing (http://aws.amazon.com/cloudwatch/pricing/).
	Logging *types.Logging

	// An object representing the VPC configuration to use for an Amazon EKS cluster.
	ResourcesVpcConfig *types.VpcConfigRequest
}

type UpdateClusterConfigOutput struct {

	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateClusterConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClusterConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClusterConfig{}, middleware.After)
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
	addIdempotencyToken_opUpdateClusterConfigMiddleware(stack, options)
	addOpUpdateClusterConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClusterConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpUpdateClusterConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateClusterConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateClusterConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateClusterConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateClusterConfigInput ")
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
func addIdempotencyToken_opUpdateClusterConfigMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateClusterConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateClusterConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "UpdateClusterConfig",
	}
}
