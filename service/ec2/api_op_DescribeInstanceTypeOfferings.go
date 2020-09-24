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

// Returns a list of all instance types offered. The results can be filtered by
// location (Region or Availability Zone). If no location is specified, the
// instance types offered in the current Region are returned.
func (c *Client) DescribeInstanceTypeOfferings(ctx context.Context, params *DescribeInstanceTypeOfferingsInput, optFns ...func(*Options)) (*DescribeInstanceTypeOfferingsOutput, error) {
	stack := middleware.NewStack("DescribeInstanceTypeOfferings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeInstanceTypeOfferingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceTypeOfferings(options.Region), middleware.Before)
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
			OperationName: "DescribeInstanceTypeOfferings",
			Err:           err,
		}
	}
	out := result.(*DescribeInstanceTypeOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceTypeOfferingsInput struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the next token
	// value.
	MaxResults *int32
	// One or more filters. Filter names and values are case-sensitive.
	//
	//     * location
	// - This depends on the location type. For example, if the location type is region
	// (default), the location is the Region code (for example, us-east-2.)
	//
	//     *
	// instance-type - The instance type.
	Filters []*types.Filter
	// The location type.
	LocationType types.LocationType
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeInstanceTypeOfferingsOutput struct {
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// The instance types offered.
	InstanceTypeOfferings []*types.InstanceTypeOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeInstanceTypeOfferingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceTypeOfferings{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceTypeOfferings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstanceTypeOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceTypeOfferings",
	}
}
