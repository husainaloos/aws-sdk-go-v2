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

// Describes the attributes for the specified load balancer.
func (c *Client) DescribeLoadBalancerAttributes(ctx context.Context, params *DescribeLoadBalancerAttributesInput, optFns ...func(*Options)) (*DescribeLoadBalancerAttributesOutput, error) {
	stack := middleware.NewStack("DescribeLoadBalancerAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeLoadBalancerAttributesMiddlewares(stack)
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
	addOpDescribeLoadBalancerAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancerAttributes(options.Region), middleware.Before)

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
			OperationName: "DescribeLoadBalancerAttributes",
			Err:           err,
		}
	}
	out := result.(*DescribeLoadBalancerAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeLoadBalancerAttributes.
type DescribeLoadBalancerAttributesInput struct {
	// The name of the load balancer.
	LoadBalancerName *string
}

// Contains the output of DescribeLoadBalancerAttributes.
type DescribeLoadBalancerAttributesOutput struct {
	// Information about the load balancer attributes.
	LoadBalancerAttributes *types.LoadBalancerAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeLoadBalancerAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancerAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancerAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLoadBalancerAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeLoadBalancerAttributes",
	}
}