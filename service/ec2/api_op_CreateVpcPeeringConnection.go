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
)

// Requests a VPC peering connection between two VPCs: a requester VPC that you own
// and an accepter VPC with which to create the connection. The accepter VPC can
// belong to another AWS account and can be in a different Region to the requester
// VPC. The requester VPC and accepter VPC cannot have overlapping CIDR blocks.
// Limitations and rules apply to a VPC peering connection. For more information,
// see the limitations
// (https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html#vpc-peering-limitations)
// section in the VPC Peering Guide. The owner of the accepter VPC must accept the
// peering request to activate the peering connection. The VPC peering connection
// request expires after 7 days, after which it cannot be accepted or rejected. If
// you create a VPC peering connection request between VPCs with overlapping CIDR
// blocks, the VPC peering connection has a status of failed.
func (c *Client) CreateVpcPeeringConnection(ctx context.Context, params *CreateVpcPeeringConnectionInput, optFns ...func(*Options)) (*CreateVpcPeeringConnectionOutput, error) {
	stack := middleware.NewStack("CreateVpcPeeringConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateVpcPeeringConnectionMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcPeeringConnection(options.Region), middleware.Before)
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
			OperationName: "CreateVpcPeeringConnection",
			Err:           err,
		}
	}
	out := result.(*CreateVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVpcPeeringConnectionInput struct {
	// The ID of the requester VPC. You must specify this parameter in the request.
	VpcId *string
	// The Region code for the accepter VPC, if the accepter VPC is located in a Region
	// other than the Region in which you make the request. Default: The Region in
	// which you make the request.
	PeerRegion *string
	// The AWS account ID of the owner of the accepter VPC. Default: Your AWS account
	// ID
	PeerOwnerId *string
	// The tags to assign to the peering connection.
	TagSpecifications []*types.TagSpecification
	// The ID of the VPC with which you are creating the VPC peering connection. You
	// must specify this parameter in the request.
	PeerVpcId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type CreateVpcPeeringConnectionOutput struct {
	// Information about the VPC peering connection.
	VpcPeeringConnection *types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateVpcPeeringConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateVpcPeeringConnection{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpcPeeringConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateVpcPeeringConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVpcPeeringConnection",
	}
}
