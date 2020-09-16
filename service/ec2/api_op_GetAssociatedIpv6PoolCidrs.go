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

// Gets information about the IPv6 CIDR block associations for a specified IPv6
// address pool.
func (c *Client) GetAssociatedIpv6PoolCidrs(ctx context.Context, params *GetAssociatedIpv6PoolCidrsInput, optFns ...func(*Options)) (*GetAssociatedIpv6PoolCidrsOutput, error) {
	stack := middleware.NewStack("GetAssociatedIpv6PoolCidrs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpGetAssociatedIpv6PoolCidrsMiddlewares(stack)
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
	addOpGetAssociatedIpv6PoolCidrsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssociatedIpv6PoolCidrs(options.Region), middleware.Before)

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
			OperationName: "GetAssociatedIpv6PoolCidrs",
			Err:           err,
		}
	}
	out := result.(*GetAssociatedIpv6PoolCidrsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssociatedIpv6PoolCidrsInput struct {
	// The ID of the IPv6 address pool.
	PoolId *string
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	// The token for the next page of results.
	NextToken *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type GetAssociatedIpv6PoolCidrsOutput struct {
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// Information about the IPv6 CIDR block associations.
	Ipv6CidrAssociations []*types.Ipv6CidrAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpGetAssociatedIpv6PoolCidrsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpGetAssociatedIpv6PoolCidrs{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpGetAssociatedIpv6PoolCidrs{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAssociatedIpv6PoolCidrs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetAssociatedIpv6PoolCidrs",
	}
}