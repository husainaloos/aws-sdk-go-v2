// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified link aggregation group (LAG). You cannot delete a LAG if
// it has active virtual interfaces or hosted connections.
func (c *Client) DeleteLag(ctx context.Context, params *DeleteLagInput, optFns ...func(*Options)) (*DeleteLagOutput, error) {
	stack := middleware.NewStack("DeleteLag", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteLagMiddlewares(stack)
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
	addOpDeleteLagValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLag(options.Region), middleware.Before)
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
			OperationName: "DeleteLag",
			Err:           err,
		}
	}
	out := result.(*DeleteLagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLagInput struct {
	// The ID of the LAG.
	LagId *string
}

// Information about a link aggregation group (LAG).
type DeleteLagOutput struct {
	// The name of the service provider associated with the LAG.
	ProviderName *string
	// Indicates whether the LAG supports a secondary BGP peer in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy
	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDeviceV2 *string
	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDevice *string
	// The ID of the LAG.
	LagId *string
	// The state of the LAG. The following are the possible values:
	//
	//     * requested:
	// The initial state of a LAG. The LAG stays in the requested state until the
	// Letter of Authorization (LOA) is available.
	//
	//     * pending: The LAG has been
	// approved and is being initialized.
	//
	//     * available: The network link is
	// established and the LAG is ready for use.
	//
	//     * down: The network link is
	// down.
	//
	//     * deleting: The LAG is being deleted.
	//
	//     * deleted: The LAG is
	// deleted.
	//
	//     * unknown: The state of the LAG is not available.
	LagState types.LagState
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool
	// Indicates whether the LAG can host other connections.
	AllowsHostedConnections *bool
	// The individual bandwidth of the physical connections bundled by the LAG. The
	// possible values are 1Gbps and 10Gbps.
	ConnectionsBandwidth *string
	// The tags associated with the LAG.
	Tags []*types.Tag
	// The location of the LAG.
	Location *string
	// The name of the LAG.
	LagName *string
	// The number of physical connections bundled by the LAG, up to a maximum of 10.
	NumberOfConnections *int32
	// The ID of the AWS account that owns the LAG.
	OwnerAccount *string
	// The connections bundled by the LAG.
	Connections []*types.Connection
	// The AWS Region where the connection is located.
	Region *string
	// The minimum number of physical connections that must be operational for the LAG
	// itself to be operational.
	MinimumLinks *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteLagMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteLag{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteLag{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLag(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteLag",
	}
}
