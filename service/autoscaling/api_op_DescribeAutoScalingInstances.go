// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more Auto Scaling instances.
func (c *Client) DescribeAutoScalingInstances(ctx context.Context, params *DescribeAutoScalingInstancesInput, optFns ...func(*Options)) (*DescribeAutoScalingInstancesOutput, error) {
	if params == nil {
		params = &DescribeAutoScalingInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutoScalingInstances", params, optFns, addOperationDescribeAutoScalingInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutoScalingInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoScalingInstancesInput struct {

	// The IDs of the instances. You can specify up to MaxRecords IDs. If you omit this
	// parameter, all Auto Scaling instances are described. If you specify an ID that
	// does not exist, it is ignored with no error.
	InstanceIds []*string

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 50.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeAutoScalingInstancesOutput struct {

	// The instances.
	AutoScalingInstances []*types.AutoScalingInstanceDetails

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAutoScalingInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAutoScalingInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAutoScalingInstances{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoScalingInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAutoScalingInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeAutoScalingInstances",
	}
}
