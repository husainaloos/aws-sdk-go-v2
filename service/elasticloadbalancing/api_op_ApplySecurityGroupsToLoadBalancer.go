// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates one or more security groups with your load balancer in a virtual
// private cloud (VPC). The specified security groups override the previously
// associated security groups. For more information, see Security Groups for Load
// Balancers in a VPC
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-security-groups.html#elb-vpc-security-groups)
// in the Classic Load Balancers Guide.
func (c *Client) ApplySecurityGroupsToLoadBalancer(ctx context.Context, params *ApplySecurityGroupsToLoadBalancerInput, optFns ...func(*Options)) (*ApplySecurityGroupsToLoadBalancerOutput, error) {
	stack := middleware.NewStack("ApplySecurityGroupsToLoadBalancer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpApplySecurityGroupsToLoadBalancerMiddlewares(stack)
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
	addOpApplySecurityGroupsToLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opApplySecurityGroupsToLoadBalancer(options.Region), middleware.Before)
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
			OperationName: "ApplySecurityGroupsToLoadBalancer",
			Err:           err,
		}
	}
	out := result.(*ApplySecurityGroupsToLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ApplySecurityGroupsToLoadBalancer.
type ApplySecurityGroupsToLoadBalancerInput struct {
	// The IDs of the security groups to associate with the load balancer. Note that
	// you cannot specify the name of the security group.
	SecurityGroups []*string
	// The name of the load balancer.
	LoadBalancerName *string
}

// Contains the output of ApplySecurityGroupsToLoadBalancer.
type ApplySecurityGroupsToLoadBalancerOutput struct {
	// The IDs of the security groups associated with the load balancer.
	SecurityGroups []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpApplySecurityGroupsToLoadBalancerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpApplySecurityGroupsToLoadBalancer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpApplySecurityGroupsToLoadBalancer{}, middleware.After)
}

func newServiceMetadataMiddleware_opApplySecurityGroupsToLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "ApplySecurityGroupsToLoadBalancer",
	}
}
