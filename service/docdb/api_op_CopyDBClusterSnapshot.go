// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies a snapshot of a cluster. To copy a cluster snapshot from a shared manual
// cluster snapshot, SourceDBClusterSnapshotIdentifier must be the Amazon Resource
// Name (ARN) of the shared cluster snapshot. To cancel the copy operation after it
// is in progress, delete the target cluster snapshot identified by
// TargetDBClusterSnapshotIdentifier while that DB cluster snapshot is in the
// copying status.
func (c *Client) CopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*Options)) (*CopyDBClusterSnapshotOutput, error) {
	stack := middleware.NewStack("CopyDBClusterSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCopyDBClusterSnapshotMiddlewares(stack)
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
	addOpCopyDBClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBClusterSnapshot(options.Region), middleware.Before)
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
			OperationName: "CopyDBClusterSnapshot",
			Err:           err,
		}
	}
	out := result.(*CopyDBClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CopyDBClusterSnapshot ().
type CopyDBClusterSnapshotInput struct {
	// The tags to be assigned to the cluster snapshot.
	Tags []*types.Tag
	// Set to true to copy all tags from the source cluster snapshot to the target
	// cluster snapshot, and otherwise false. The default is false.
	CopyTags *bool
	// The AWS KMS key ID for an encrypted cluster snapshot. The AWS KMS key ID is the
	// Amazon Resource Name (ARN), AWS KMS key identifier, or the AWS KMS key alias for
	// the AWS KMS encryption key. If you copy an encrypted cluster snapshot from your
	// AWS account, you can specify a value for KmsKeyId to encrypt the copy with a new
	// AWS KMS encryption key. If you don't specify a value for KmsKeyId, then the copy
	// of the cluster snapshot is encrypted with the same AWS KMS key as the source
	// cluster snapshot. If you copy an encrypted cluster snapshot that is shared from
	// another AWS account, then you must specify a value for KmsKeyId. To copy an
	// encrypted cluster snapshot to another AWS Region, set KmsKeyId to the AWS KMS
	// key ID that you want to use to encrypt the copy of the cluster snapshot in the
	// destination Region. AWS KMS encryption keys are specific to the AWS Region that
	// they are created in, and you can't use encryption keys from one Region in
	// another Region. If you copy an unencrypted cluster snapshot and specify a value
	// for the KmsKeyId parameter, an error is returned.
	KmsKeyId *string
	// The identifier of the cluster snapshot to copy. This parameter is not case
	// sensitive. You can't copy an encrypted, shared cluster snapshot from one AWS
	// Region to another. Constraints:
	//
	//     * Must specify a valid system snapshot in
	// the "available" state.
	//
	//     * If the source snapshot is in the same AWS Region
	// as the copy, specify a valid snapshot identifier.
	//
	//     * If the source snapshot
	// is in a different AWS Region than the copy, specify a valid cluster snapshot
	// ARN.
	//
	// Example: my-cluster-snapshot1
	SourceDBClusterSnapshotIdentifier *string
	// The URL that contains a Signature Version 4 signed request for the
	// CopyDBClusterSnapshot API action in the AWS Region that contains the source
	// cluster snapshot to copy. You must use the PreSignedUrl parameter when copying
	// an encrypted cluster snapshot from another AWS Region. The presigned URL must be
	// a valid request for the CopyDBSClusterSnapshot API action that can be executed
	// in the source AWS Region that contains the encrypted DB cluster snapshot to be
	// copied. The presigned URL request must contain the following parameter values:
	//
	//
	// * KmsKeyId - The AWS KMS key identifier for the key to use to encrypt the copy
	// of the cluster snapshot in the destination AWS Region. This is the same
	// identifier for both the CopyDBClusterSnapshot action that is called in the
	// destination AWS Region, and the action contained in the presigned URL.
	//
	//     *
	// DestinationRegion - The name of the AWS Region that the DB cluster snapshot will
	// be created in.
	//
	//     * SourceDBClusterSnapshotIdentifier - The cluster snapshot
	// identifier for the encrypted cluster snapshot to be copied. This identifier must
	// be in the Amazon Resource Name (ARN) format for the source AWS Region. For
	// example, if you are copying an encrypted cluster snapshot from the us-west-2 AWS
	// Region, then your SourceDBClusterSnapshotIdentifier looks like the following
	// example:
	// arn:aws:rds:us-west-2:123456789012:cluster-snapshot:my-cluster-snapshot-20161115.
	PreSignedUrl *string
	// The identifier of the new cluster snapshot to create from the source cluster
	// snapshot. This parameter is not case sensitive. Constraints:
	//
	//     * Must contain
	// from 1 to 63 letters, numbers, or hyphens.
	//
	//     * The first character must be a
	// letter.
	//
	//     * Cannot end with a hyphen or contain two consecutive
	// hyphens.
	//
	// Example: my-cluster-snapshot2
	TargetDBClusterSnapshotIdentifier *string
}

type CopyDBClusterSnapshotOutput struct {
	// Detailed information about a cluster snapshot.
	DBClusterSnapshot *types.DBClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCopyDBClusterSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBClusterSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBClusterSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyDBClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBClusterSnapshot",
	}
}
