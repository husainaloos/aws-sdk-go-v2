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

// Describes one or more of your VPN customer gateways. For more information, see
// AWS Site-to-Site VPN
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the AWS
// Site-to-Site VPN User Guide.
func (c *Client) DescribeCustomerGateways(ctx context.Context, params *DescribeCustomerGatewaysInput, optFns ...func(*Options)) (*DescribeCustomerGatewaysOutput, error) {
	stack := middleware.NewStack("DescribeCustomerGateways", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeCustomerGatewaysMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomerGateways(options.Region), middleware.Before)
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
			OperationName: "DescribeCustomerGateways",
			Err:           err,
		}
	}
	out := result.(*DescribeCustomerGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeCustomerGateways.
type DescribeCustomerGatewaysInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// One or more customer gateway IDs. Default: Describes all your customer gateways.
	CustomerGatewayIds []*string
	// One or more filters.
	//
	//     * bgp-asn - The customer gateway's Border Gateway
	// Protocol (BGP) Autonomous System Number (ASN).
	//
	//     * customer-gateway-id - The
	// ID of the customer gateway.
	//
	//     * ip-address - The IP address of the customer
	// gateway's Internet-routable external interface.
	//
	//     * state - The state of the
	// customer gateway (pending | available | deleting | deleted).
	//
	//     * type - The
	// type of customer gateway. Currently, the only supported type is ipsec.1.
	//
	//     *
	// tag: - The key/value combination of a tag assigned to the resource. Use the tag
	// key in the filter name and the tag value as the filter value. For example, to
	// find all resources that have a tag with the key Owner and the value TeamA,
	// specify tag:Owner for the filter name and TeamA for the filter value.
	//
	//     *
	// tag-key - The key of a tag assigned to the resource. Use this filter to find all
	// resources assigned a tag with a specific key, regardless of the tag value.
	Filters []*types.Filter
}

// Contains the output of DescribeCustomerGateways.
type DescribeCustomerGatewaysOutput struct {
	// Information about one or more customer gateways.
	CustomerGateways []*types.CustomerGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeCustomerGatewaysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeCustomerGateways{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeCustomerGateways{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCustomerGateways(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeCustomerGateways",
	}
}
