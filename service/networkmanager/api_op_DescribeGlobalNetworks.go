// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more global networks. By default, all global networks are
// described. To describe the objects in your global network, you must use the
// appropriate Get* action. For example, to list the transit gateways in your
// global network, use GetTransitGatewayRegistrations ().
func (c *Client) DescribeGlobalNetworks(ctx context.Context, params *DescribeGlobalNetworksInput, optFns ...func(*Options)) (*DescribeGlobalNetworksOutput, error) {
	stack := middleware.NewStack("DescribeGlobalNetworks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeGlobalNetworksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalNetworks(options.Region), middleware.Before)
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
			OperationName: "DescribeGlobalNetworks",
			Err:           err,
		}
	}
	out := result.(*DescribeGlobalNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalNetworksInput struct {
	// The maximum number of results to return.
	MaxResults *int32
	// The IDs of one or more global networks. The maximum is 10.
	GlobalNetworkIds []*string
	// The token for the next page of results.
	NextToken *string
}

type DescribeGlobalNetworksOutput struct {
	// The token for the next page of results.
	NextToken *string
	// Information about the global networks.
	GlobalNetworks []*types.GlobalNetwork

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeGlobalNetworksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGlobalNetworks{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGlobalNetworks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGlobalNetworks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "DescribeGlobalNetworks",
	}
}
