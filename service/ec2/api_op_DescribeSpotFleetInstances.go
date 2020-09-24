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

// Describes the running instances for the specified Spot Fleet.
func (c *Client) DescribeSpotFleetInstances(ctx context.Context, params *DescribeSpotFleetInstancesInput, optFns ...func(*Options)) (*DescribeSpotFleetInstancesOutput, error) {
	stack := middleware.NewStack("DescribeSpotFleetInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeSpotFleetInstancesMiddlewares(stack)
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
	addOpDescribeSpotFleetInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotFleetInstances(options.Region), middleware.Before)
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
			OperationName: "DescribeSpotFleetInstances",
			Err:           err,
		}
	}
	out := result.(*DescribeSpotFleetInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotFleetInstances.
type DescribeSpotFleetInstancesInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int32
	// The token for the next set of results.
	NextToken *string
	// The ID of the Spot Fleet request.
	SpotFleetRequestId *string
}

// Contains the output of DescribeSpotFleetInstances.
type DescribeSpotFleetInstancesOutput struct {
	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// The ID of the Spot Fleet request.
	SpotFleetRequestId *string
	// The running instances. This list is refreshed periodically and might be out of
	// date.
	ActiveInstances []*types.ActiveInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeSpotFleetInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotFleetInstances{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotFleetInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSpotFleetInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotFleetInstances",
	}
}
