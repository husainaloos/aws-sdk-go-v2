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

// Creates a managed worker node group for an Amazon EKS cluster. You can only
// create a node group for your cluster that is equal to the current Kubernetes
// version for the cluster. All node groups are created with the latest AMI release
// version for the respective minor Kubernetes version of the cluster. An Amazon
// EKS managed node group is an Amazon EC2 Auto Scaling group and associated Amazon
// EC2 instances that are managed by AWS for an Amazon EKS cluster. Each node group
// uses a version of the Amazon EKS-optimized Amazon Linux 2 AMI. For more
// information, see Managed Node Groups
// (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in
// the Amazon EKS User Guide.
func (c *Client) CreateNodegroup(ctx context.Context, params *CreateNodegroupInput, optFns ...func(*Options)) (*CreateNodegroupOutput, error) {
	stack := middleware.NewStack("CreateNodegroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateNodegroupMiddlewares(stack)
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
	addIdempotencyToken_opCreateNodegroupMiddleware(stack, options)
	addOpCreateNodegroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNodegroup(options.Region), middleware.Before)
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
			OperationName: "CreateNodegroup",
			Err:           err,
		}
	}
	out := result.(*CreateNodegroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNodegroupInput struct {
	// The subnets to use for the Auto Scaling group that is created for your node
	// group. These subnets must have the tag key kubernetes.io/cluster/CLUSTER_NAME
	// with a value of shared, where CLUSTER_NAME is replaced with the name of your
	// cluster.
	Subnets []*string
	// The unique name to give your node group.
	NodegroupName *string
	// The AMI type for your node group. GPU instance types should use the
	// AL2_x86_64_GPU AMI type, which uses the Amazon EKS-optimized Linux AMI with GPU
	// support. Non-GPU instances should use the AL2_x86_64 AMI type, which uses the
	// Amazon EKS-optimized Linux AMI.
	AmiType types.AMITypes
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node
	// group. The Amazon EKS worker node kubelet daemon makes calls to AWS APIs on your
	// behalf. Worker nodes receive permissions for these API calls through an IAM
	// instance profile and associated policies. Before you can launch worker nodes and
	// register them into a cluster, you must create an IAM role for those worker nodes
	// to use when they are launched. For more information, see Amazon EKS Worker Node
	// IAM Role
	// (https://docs.aws.amazon.com/eks/latest/userguide/worker_node_IAM_role.html) in
	// the Amazon EKS User Guide .
	NodeRole *string
	// The instance type to use for your node group. Currently, you can specify a
	// single instance type for a node group. The default value for this parameter is
	// t3.medium. If you choose a GPU instance type, be sure to specify the
	// AL2_x86_64_GPU with the amiType parameter.
	InstanceTypes []*string
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group. By
	// default, the latest available AMI version for the node group's current
	// Kubernetes version is used. For more information, see Amazon EKS-Optimized Linux
	// AMI Versions
	// (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
	// in the Amazon EKS User Guide.
	ReleaseVersion *string
	// The remote access (SSH) configuration to use with your node group.
	RemoteAccess *types.RemoteAccessConfig
	// The scaling configuration details for the Auto Scaling group that is created for
	// your node group.
	ScalingConfig *types.NodegroupScalingConfig
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string
	// The name of the cluster to create the node group in.
	ClusterName *string
	// The root device disk size (in GiB) for your node group instances. The default
	// disk size is 20 GiB.
	DiskSize *int32
	// The Kubernetes version to use for your managed nodes. By default, the Kubernetes
	// version of the cluster is used, and this is the only accepted specified value.
	Version *string
	// The metadata to apply to the node group to assist with categorization and
	// organization. Each tag consists of a key and an optional value, both of which
	// you define. Node group tags do not propagate to any other resources associated
	// with the node group, such as the Amazon EC2 instances or subnets.
	Tags map[string]*string
	// The Kubernetes labels to be applied to the nodes in the node group when they are
	// created.
	Labels map[string]*string
}

type CreateNodegroupOutput struct {
	// The full description of your new node group.
	Nodegroup *types.Nodegroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateNodegroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateNodegroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNodegroup{}, middleware.After)
}

type idempotencyToken_initializeOpCreateNodegroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNodegroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNodegroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNodegroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNodegroupInput ")
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
func addIdempotencyToken_opCreateNodegroupMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateNodegroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNodegroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "CreateNodegroup",
	}
}
