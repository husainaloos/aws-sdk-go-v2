// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon ECS cluster. By default, your account receives a default
// cluster when you launch your first container instance. However, you can create
// your own cluster with a unique name with the CreateCluster action. When you call
// the CreateCluster () API operation, Amazon ECS attempts to create the Amazon ECS
// service-linked role for your account so that required resources in other AWS
// services can be managed on your behalf. However, if the IAM user that makes the
// call does not have permissions to create the service-linked role, it is not
// created. For more information, see Using Service-Linked Roles for Amazon ECS
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	stack := middleware.NewStack("CreateCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateClusterMiddlewares(stack)
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
	addOpCreateClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before)
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
			OperationName: "CreateCluster",
			Err:           err,
		}
	}
	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {
	// The short name of one or more capacity providers to associate with the cluster.
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity
	// provider must already be created and not already associated with another
	// cluster. New capacity providers can be created with the CreateCapacityProvider
	// () API operation. To use a AWS Fargate capacity provider, specify either the
	// FARGATE or FARGATE_SPOT capacity providers. The AWS Fargate capacity providers
	// are available to all accounts and only need to be associated with a cluster to
	// be used. The PutClusterCapacityProviders () API operation is used to update the
	// list of available capacity providers for a cluster after the cluster is created.
	CapacityProviders []*string
	// The name of your cluster. If you do not specify a name for your cluster, you
	// create a cluster named default. Up to 255 letters (uppercase and lowercase),
	// numbers, and hyphens are allowed.
	ClusterName *string
	// The metadata that you apply to the cluster to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define. The following basic restrictions apply to tags:
	//
	//     * Maximum number of
	// tags per resource - 50
	//
	//     * For each resource, each tag key must be unique,
	// and each tag key can have only one value.
	//
	//     * Maximum key length - 128
	// Unicode characters in UTF-8
	//
	//     * Maximum value length - 256 Unicode characters
	// in UTF-8
	//
	//     * If your tagging schema is used across multiple services and
	// resources, remember that other services may have restrictions on allowed
	// characters. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	//     *
	// Tag keys and values are case-sensitive.
	//
	//     * Do not use aws:, AWS:, or any
	// upper or lowercase combination of such as a prefix for either keys or values as
	// it is reserved for AWS use. You cannot edit or delete tag keys or values with
	// this prefix. Tags with this prefix do not count against your tags per resource
	// limit.
	Tags []*types.Tag
	// The capacity provider strategy to use by default for the cluster. When creating
	// a service or running a task on a cluster, if no capacity provider or launch type
	// is specified then the default capacity provider strategy for the cluster is
	// used. A capacity provider strategy consists of one or more capacity providers
	// along with the base and weight to assign to them. A capacity provider must be
	// associated with the cluster to be used in a capacity provider strategy. The
	// PutClusterCapacityProviders () API is used to associate a capacity provider with
	// a cluster. Only capacity providers with an ACTIVE or UPDATING status can be
	// used. If specifying a capacity provider that uses an Auto Scaling group, the
	// capacity provider must already be created. New capacity providers can be created
	// with the CreateCapacityProvider () API operation. To use a AWS Fargate capacity
	// provider, specify either the FARGATE or FARGATE_SPOT capacity providers. The AWS
	// Fargate capacity providers are available to all accounts and only need to be
	// associated with a cluster to be used. If a default capacity provider strategy is
	// not defined for a cluster during creation, it can be defined later with the
	// PutClusterCapacityProviders () API operation.
	DefaultCapacityProviderStrategy []*types.CapacityProviderStrategyItem
	// The setting to use when creating a cluster. This parameter is used to enable
	// CloudWatch Container Insights for a cluster. If this value is specified, it will
	// override the containerInsights value set with PutAccountSetting () or
	// PutAccountSettingDefault ().
	Settings []*types.ClusterSetting
}

type CreateClusterOutput struct {
	// The full description of your new cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "CreateCluster",
	}
}
