// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description of the gateway volumes specified in the request. The
// list of gateway volumes in the request must be from one gateway. In the
// response, AWS Storage Gateway returns volume information sorted by volume ARNs.
// This operation is only supported in stored volume gateway type.
func (c *Client) DescribeStorediSCSIVolumes(ctx context.Context, params *DescribeStorediSCSIVolumesInput, optFns ...func(*Options)) (*DescribeStorediSCSIVolumesOutput, error) {
	stack := middleware.NewStack("DescribeStorediSCSIVolumes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeStorediSCSIVolumesMiddlewares(stack)
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
	addOpDescribeStorediSCSIVolumesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorediSCSIVolumes(options.Region), middleware.Before)
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
			OperationName: "DescribeStorediSCSIVolumes",
			Err:           err,
		}
	}
	out := result.(*DescribeStorediSCSIVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing a list of DescribeStorediSCSIVolumesInput$VolumeARNs
// ().
type DescribeStorediSCSIVolumesInput struct {
	// An array of strings where each string represents the Amazon Resource Name (ARN)
	// of a stored volume. All of the specified stored volumes must be from the same
	// gateway. Use ListVolumes () to get volume ARNs for a gateway.
	VolumeARNs []*string
}

type DescribeStorediSCSIVolumesOutput struct {
	// Describes a single unit of output from DescribeStorediSCSIVolumes (). The
	// following fields are returned:  <ul> <li> <p> <code>ChapEnabled</code>:
	// Indicates whether mutual CHAP is enabled for the iSCSI target.</p> </li> <li>
	// <p> <code>LunNumber</code>: The logical disk number.</p> </li> <li> <p>
	// <code>NetworkInterfaceId</code>: The network interface ID of the stored volume
	// that initiator use to map the stored volume as an iSCSI target.</p> </li> <li>
	// <p> <code>NetworkInterfacePort</code>: The port used to communicate with iSCSI
	// targets.</p> </li> <li> <p> <code>PreservedExistingData</code>: Indicates when
	// the stored volume was created, existing data on the underlying local disk was
	// preserved.</p> </li> <li> <p> <code>SourceSnapshotId</code>: If the stored
	// volume was created from a snapshot, this field contains the snapshot ID used,
	// e.g. <code>snap-1122aabb</code>. Otherwise, this field is not included.</p>
	// </li> <li> <p> <code>StorediSCSIVolumes</code>: An array of StorediSCSIVolume
	// objects where each object contains metadata about one stored volume.</p> </li>
	// <li> <p> <code>TargetARN</code>: The Amazon Resource Name (ARN) of the volume
	// target.</p> </li> <li> <p> <code>VolumeARN</code>: The Amazon Resource Name
	// (ARN) of the stored volume.</p> </li> <li> <p> <code>VolumeDiskId</code>: The
	// disk ID of the local disk that was specified in the
	// <a>CreateStorediSCSIVolume</a> operation.</p> </li> <li> <p>
	// <code>VolumeId</code>: The unique identifier of the storage volume, e.g.
	// <code>vol-1122AABB</code>.</p> </li> <li> <p>
	// <code>VolumeiSCSIAttributes</code>: An <a>VolumeiSCSIAttributes</a> object that
	// represents a collection of iSCSI attributes for one stored volume.</p> </li>
	// <li> <p> <code>VolumeProgress</code>: Represents the percentage complete if the
	// volume is restoring or bootstrapping that represents the percent of data
	// transferred. This field does not appear in the response if the stored volume is
	// not restoring or bootstrapping.</p> </li> <li> <p>
	// <code>VolumeSizeInBytes</code>: The size of the volume in bytes.</p> </li> <li>
	// <p> <code>VolumeStatus</code>: One of the <code>VolumeStatus</code> values that
	// indicates the state of the volume.</p> </li> <li> <p> <code>VolumeType</code>:
	// One of the enumeration values describing the type of the volume. Currently, only
	// <code>STORED</code> volumes are supported.</p> </li> </ul>
	StorediSCSIVolumes []*types.StorediSCSIVolume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeStorediSCSIVolumesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStorediSCSIVolumes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStorediSCSIVolumes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStorediSCSIVolumes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeStorediSCSIVolumes",
	}
}
