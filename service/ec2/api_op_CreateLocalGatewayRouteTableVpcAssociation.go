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

// Associates the specified VPC with the specified local gateway route table.
func (c *Client) CreateLocalGatewayRouteTableVpcAssociation(ctx context.Context, params *CreateLocalGatewayRouteTableVpcAssociationInput, optFns ...func(*Options)) (*CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	stack := middleware.NewStack("CreateLocalGatewayRouteTableVpcAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateLocalGatewayRouteTableVpcAssociationMiddlewares(stack)
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
	addOpCreateLocalGatewayRouteTableVpcAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocalGatewayRouteTableVpcAssociation(options.Region), middleware.Before)
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
			OperationName: "CreateLocalGatewayRouteTableVpcAssociation",
			Err:           err,
		}
	}
	out := result.(*CreateLocalGatewayRouteTableVpcAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocalGatewayRouteTableVpcAssociationInput struct {
	// The ID of the VPC.
	VpcId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId *string
	// The tags to assign to the local gateway route table VPC association.
	TagSpecifications []*types.TagSpecification
}

type CreateLocalGatewayRouteTableVpcAssociationOutput struct {
	// Information about the association.
	LocalGatewayRouteTableVpcAssociation *types.LocalGatewayRouteTableVpcAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateLocalGatewayRouteTableVpcAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateLocalGatewayRouteTableVpcAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLocalGatewayRouteTableVpcAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLocalGatewayRouteTableVpcAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLocalGatewayRouteTableVpcAssociation",
	}
}
