// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified manual snapshot. The snapshot must be in the available
// state, with no other users authorized to access the snapshot. Unlike automated
// snapshots, manual snapshots are retained even after you delete your cluster.
// Amazon Redshift does not delete your manual snapshots. You must delete manual
// snapshot explicitly to avoid getting charged. If other accounts are authorized
// to access the snapshot, you must revoke all of the authorizations before you can
// delete the snapshot.
func (c *Client) DeleteClusterSnapshot(ctx context.Context, params *DeleteClusterSnapshotInput, optFns ...func(*Options)) (*DeleteClusterSnapshotOutput, error) {
	stack := middleware.NewStack("DeleteClusterSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteClusterSnapshotMiddlewares(stack)
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
	addOpDeleteClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteClusterSnapshot(options.Region), middleware.Before)
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
			OperationName: "DeleteClusterSnapshot",
			Err:           err,
		}
	}
	out := result.(*DeleteClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteClusterSnapshotInput struct {
	// The unique identifier of the cluster the snapshot was created from. This
	// parameter is required if your IAM user has a policy containing a snapshot
	// resource element that specifies anything other than * for the cluster name.
	// Constraints: Must be the name of valid cluster.
	SnapshotClusterIdentifier *string
	// The unique identifier of the manual snapshot to be deleted. Constraints: Must be
	// the name of an existing snapshot that is in the available, failed, or cancelled
	// state.
	SnapshotIdentifier *string
}

type DeleteClusterSnapshotOutput struct {
	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteClusterSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteClusterSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteClusterSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeleteClusterSnapshot",
	}
}
