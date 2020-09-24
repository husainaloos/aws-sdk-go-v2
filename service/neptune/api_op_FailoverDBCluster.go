// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Forces a failover for a DB cluster. A failover for a DB cluster promotes one of
// the Read Replicas (read-only instances) in the DB cluster to be the primary
// instance (the cluster writer). Amazon Neptune will automatically fail over to a
// Read Replica, if one exists, when the primary instance fails. You can force a
// failover when you want to simulate a failure of a primary instance for testing.
// Because each instance in a DB cluster has its own endpoint address, you will
// need to clean up and re-establish any existing connections that use those
// endpoint addresses when the failover is complete.
func (c *Client) FailoverDBCluster(ctx context.Context, params *FailoverDBClusterInput, optFns ...func(*Options)) (*FailoverDBClusterOutput, error) {
	stack := middleware.NewStack("FailoverDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpFailoverDBClusterMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverDBCluster(options.Region), middleware.Before)
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
			OperationName: "FailoverDBCluster",
			Err:           err,
		}
	}
	out := result.(*FailoverDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FailoverDBClusterInput struct {
	// A DB cluster identifier to force a failover for. This parameter is not
	// case-sensitive. Constraints:
	//
	//     * Must match the identifier of an existing
	// DBCluster.
	DBClusterIdentifier *string
	// The name of the instance to promote to the primary instance. You must specify
	// the instance identifier for an Read Replica in the DB cluster. For example,
	// mydbcluster-replica1.
	TargetDBInstanceIdentifier *string
}

type FailoverDBClusterOutput struct {
	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters () action.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpFailoverDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpFailoverDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opFailoverDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "FailoverDBCluster",
	}
}
