// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Dynamically increases the number of replics in a Redis (cluster mode disabled)
// replication group or the number of replica nodes in one or more node groups
// (shards) of a Redis (cluster mode enabled) replication group. This operation is
// performed with no cluster down time.
func (c *Client) IncreaseReplicaCount(ctx context.Context, params *IncreaseReplicaCountInput, optFns ...func(*Options)) (*IncreaseReplicaCountOutput, error) {
	stack := middleware.NewStack("IncreaseReplicaCount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpIncreaseReplicaCountMiddlewares(stack)
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
	addOpIncreaseReplicaCountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opIncreaseReplicaCount(options.Region), middleware.Before)
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
			OperationName: "IncreaseReplicaCount",
			Err:           err,
		}
	}
	out := result.(*IncreaseReplicaCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IncreaseReplicaCountInput struct {
	// A list of ConfigureShard objects that can be used to configure each shard in a
	// Redis (cluster mode enabled) replication group. The ConfigureShard has three
	// members: NewReplicaCount, NodeGroupId, and PreferredAvailabilityZones.
	ReplicaConfiguration []*types.ConfigureShard
	// The id of the replication group to which you want to add replica nodes.
	ReplicationGroupId *string
	// The number of read replica nodes you want at the completion of this operation.
	// For Redis (cluster mode disabled) replication groups, this is the number of
	// replica nodes in the replication group. For Redis (cluster mode enabled)
	// replication groups, this is the number of replica nodes in each of the
	// replication group's node groups.
	NewReplicaCount *int32
	// If True, the number of replica nodes is increased immediately.
	// <code>ApplyImmediately=False</code> is not currently supported.</p>
	ApplyImmediately *bool
}

type IncreaseReplicaCountOutput struct {
	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpIncreaseReplicaCountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpIncreaseReplicaCount{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpIncreaseReplicaCount{}, middleware.After)
}

func newServiceMetadataMiddleware_opIncreaseReplicaCount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "IncreaseReplicaCount",
	}
}
