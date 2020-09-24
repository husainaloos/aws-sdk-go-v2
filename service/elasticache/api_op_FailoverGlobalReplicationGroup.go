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

// Used to failover the primary region to a selected secondary region. The selected
// secondary region will become primary, and all other clusters will become
// secondary.
func (c *Client) FailoverGlobalReplicationGroup(ctx context.Context, params *FailoverGlobalReplicationGroupInput, optFns ...func(*Options)) (*FailoverGlobalReplicationGroupOutput, error) {
	stack := middleware.NewStack("FailoverGlobalReplicationGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpFailoverGlobalReplicationGroupMiddlewares(stack)
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
	addOpFailoverGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverGlobalReplicationGroup(options.Region), middleware.Before)
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
			OperationName: "FailoverGlobalReplicationGroup",
			Err:           err,
		}
	}
	out := result.(*FailoverGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FailoverGlobalReplicationGroupInput struct {
	// The name of the primary replication group
	PrimaryReplicationGroupId *string
	// The name of the Global Datastore
	GlobalReplicationGroupId *string
	// The AWS region of the primary cluster of the Global Datastore
	PrimaryRegion *string
}

type FailoverGlobalReplicationGroupOutput struct {
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

func addawsAwsquery_serdeOpFailoverGlobalReplicationGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpFailoverGlobalReplicationGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverGlobalReplicationGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opFailoverGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "FailoverGlobalReplicationGroup",
	}
}
