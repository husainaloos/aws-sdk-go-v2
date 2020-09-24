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

// Redistribute slots to ensure uniform distribution across existing shards in the
// cluster.
func (c *Client) RebalanceSlotsInGlobalReplicationGroup(ctx context.Context, params *RebalanceSlotsInGlobalReplicationGroupInput, optFns ...func(*Options)) (*RebalanceSlotsInGlobalReplicationGroupOutput, error) {
	stack := middleware.NewStack("RebalanceSlotsInGlobalReplicationGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRebalanceSlotsInGlobalReplicationGroupMiddlewares(stack)
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
	addOpRebalanceSlotsInGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebalanceSlotsInGlobalReplicationGroup(options.Region), middleware.Before)
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
			OperationName: "RebalanceSlotsInGlobalReplicationGroup",
			Err:           err,
		}
	}
	out := result.(*RebalanceSlotsInGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebalanceSlotsInGlobalReplicationGroupInput struct {
	// The name of the Global Datastore
	GlobalReplicationGroupId *string
	// If True, redistribution is applied immediately.
	ApplyImmediately *bool
}

type RebalanceSlotsInGlobalReplicationGroupOutput struct {
	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.  <ul> <li> <p>The <b>GlobalReplicationGroupIdSuffix</b>
	// represents the name of the Global Datastore, which is what you use to associate
	// a secondary cluster.</p> </li> </ul>
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRebalanceSlotsInGlobalReplicationGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRebalanceSlotsInGlobalReplicationGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRebalanceSlotsInGlobalReplicationGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opRebalanceSlotsInGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "RebalanceSlotsInGlobalReplicationGroup",
	}
}
