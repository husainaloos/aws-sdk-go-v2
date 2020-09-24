// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a snapshot of an EBS volume and stores it in Amazon S3. You can use
// snapshots for backups, to make copies of EBS volumes, and to save data before
// shutting down an instance. When a snapshot is created, any AWS Marketplace
// product codes that are associated with the source volume are propagated to the
// snapshot. You can take a snapshot of an attached volume that is in use. However,
// snapshots only capture data that has been written to your EBS volume at the time
// the snapshot command is issued; this may exclude any data that has been cached
// by any applications or the operating system. If you can pause any file systems
// on the volume long enough to take a snapshot, your snapshot should be complete.
// However, if you cannot pause all file writes to the volume, you should unmount
// the volume from within the instance, issue the snapshot command, and then
// remount the volume to ensure a consistent and complete snapshot. You may remount
// and use your volume while the snapshot status is pending. To create a snapshot
// for EBS volumes that serve as root devices, you should stop the instance before
// taking the snapshot. Snapshots that are taken from encrypted volumes are
// automatically encrypted. Volumes that are created from encrypted snapshots are
// also automatically encrypted. Your encrypted volumes and any associated
// snapshots always remain protected. You can tag your snapshots during creation.
// For more information, see Tagging Your Amazon EC2 Resources
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
// Amazon Elastic Compute Cloud User Guide. For more information, see Amazon
// Elastic Block Store
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AmazonEBS.html) and Amazon
// EBS Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateSnapshot(ctx context.Context, params *CreateSnapshotInput, optFns ...func(*Options)) (*CreateSnapshotOutput, error) {
	stack := middleware.NewStack("CreateSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateSnapshotMiddlewares(stack)
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
	addOpCreateSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshot(options.Region), middleware.Before)
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
			OperationName: "CreateSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSnapshotInput struct {
	// A description for the snapshot.
	Description *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The tags to apply to the snapshot during creation.
	TagSpecifications []*types.TagSpecification
	// The ID of the EBS volume.
	VolumeId *string
}

// Describes a snapshot.
type CreateSnapshotOutput struct {
	// The AWS account ID of the EBS snapshot owner.
	OwnerId *string
	// The size of the volume, in GiB.
	VolumeSize *int32
	// The AWS owner alias, as maintained by Amazon. The possible values are: amazon |
	// self | all | aws-marketplace | microsoft. This AWS owner alias is not to be
	// confused with the user-configured AWS account alias, which is set from the IAM
	// console.
	OwnerAlias *string
	// The ID of the snapshot. Each snapshot receives a unique identifier when it is
	// created.
	SnapshotId *string
	// The ID of the volume that was used to create the snapshot. Snapshots created by
	// the CopySnapshot () action have an arbitrary volume ID that should not be used
	// for any purpose.
	VolumeId *string
	// Encrypted Amazon EBS snapshots are copied asynchronously. If a snapshot copy
	// operation fails (for example, if the proper AWS Key Management Service (AWS KMS)
	// permissions are not obtained) this field displays error state details to help
	// you diagnose why the error occurred. This parameter is only returned by
	// DescribeSnapshots ().
	StateMessage *string
	// Indicates whether the snapshot is encrypted.
	Encrypted *bool
	// Any tags assigned to the snapshot.
	Tags []*types.Tag
	// The data encryption key identifier for the snapshot. This value is a unique
	// identifier that corresponds to the data encryption key that was used to encrypt
	// the original volume or snapshot copy. Because data encryption keys are inherited
	// by volumes created from snapshots, and vice versa, if snapshots share the same
	// data encryption key identifier, then they belong to the same volume/snapshot
	// lineage. This parameter is only returned by DescribeSnapshots ().
	DataEncryptionKeyId *string
	// The progress of the snapshot, as a percentage.
	Progress *string
	// The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS)
	// customer master key (CMK) that was used to protect the volume encryption key for
	// the parent volume.
	KmsKeyId *string
	// The description for the snapshot.
	Description *string
	// The snapshot state.
	State types.SnapshotState
	// The time stamp when the snapshot was initiated.
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateSnapshot",
	}
}
