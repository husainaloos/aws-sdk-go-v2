// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified target groups or all of your target groups. By default,
// all target groups are described. Alternatively, you can specify one of the
// following to filter the results: the ARN of the load balancer, the names of one
// or more target groups, or the ARNs of one or more target groups. To describe the
// targets for a target group, use DescribeTargetHealth (). To describe the
// attributes of a target group, use DescribeTargetGroupAttributes ().
func (c *Client) DescribeTargetGroups(ctx context.Context, params *DescribeTargetGroupsInput, optFns ...func(*Options)) (*DescribeTargetGroupsOutput, error) {
	if params == nil {
		params = &DescribeTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTargetGroups", params, optFns, addOperationDescribeTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTargetGroupsInput struct {

	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The names of the target groups.
	Names []*string

	// The maximum number of results to return with this call.
	PageSize *int32

	// The Amazon Resource Names (ARN) of the target groups.
	TargetGroupArns []*string
}

type DescribeTargetGroupsOutput struct {

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string

	// Information about the target groups.
	TargetGroups []*types.TargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTargetGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTargetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeTargetGroups",
	}
}
