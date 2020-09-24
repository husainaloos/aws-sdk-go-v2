// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the state of the specified instances with respect to the specified
// load balancer. If no instances are specified, the call describes the state of
// all instances that are currently registered with the load balancer. If instances
// are specified, their state is returned even if they are no longer registered
// with the load balancer. The state of terminated instances is not returned.
func (c *Client) DescribeInstanceHealth(ctx context.Context, params *DescribeInstanceHealthInput, optFns ...func(*Options)) (*DescribeInstanceHealthOutput, error) {
	stack := middleware.NewStack("DescribeInstanceHealth", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeInstanceHealthMiddlewares(stack)
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
	addOpDescribeInstanceHealthValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceHealth(options.Region), middleware.Before)
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
			OperationName: "DescribeInstanceHealth",
			Err:           err,
		}
	}
	out := result.(*DescribeInstanceHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeInstanceHealth.
type DescribeInstanceHealthInput struct {
	// The name of the load balancer.
	LoadBalancerName *string
	// The IDs of the instances.
	Instances []*types.Instance
}

// Contains the output for DescribeInstanceHealth.
type DescribeInstanceHealthOutput struct {
	// Information about the health of the instances.
	InstanceStates []*types.InstanceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeInstanceHealthMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstanceHealth{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstanceHealth{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstanceHealth(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeInstanceHealth",
	}
}
