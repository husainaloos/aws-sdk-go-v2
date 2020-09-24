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

// Describes the specified load balancer policy types or all load balancer policy
// types. The description of each type indicates how it can be used. For example,
// some policies can be used only with layer 7 listeners, some policies can be used
// only with layer 4 listeners, and some policies can be used only with your EC2
// instances. You can use CreateLoadBalancerPolicy () to create a policy
// configuration for any of these policy types. Then, depending on the policy type,
// use either SetLoadBalancerPoliciesOfListener () or
// SetLoadBalancerPoliciesForBackendServer () to set the policy.
func (c *Client) DescribeLoadBalancerPolicyTypes(ctx context.Context, params *DescribeLoadBalancerPolicyTypesInput, optFns ...func(*Options)) (*DescribeLoadBalancerPolicyTypesOutput, error) {
	stack := middleware.NewStack("DescribeLoadBalancerPolicyTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeLoadBalancerPolicyTypesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancerPolicyTypes(options.Region), middleware.Before)
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
			OperationName: "DescribeLoadBalancerPolicyTypes",
			Err:           err,
		}
	}
	out := result.(*DescribeLoadBalancerPolicyTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeLoadBalancerPolicyTypes.
type DescribeLoadBalancerPolicyTypesInput struct {
	// The names of the policy types. If no names are specified, describes all policy
	// types defined by Elastic Load Balancing.
	PolicyTypeNames []*string
}

// Contains the output of DescribeLoadBalancerPolicyTypes.
type DescribeLoadBalancerPolicyTypesOutput struct {
	// Information about the policy types.
	PolicyTypeDescriptions []*types.PolicyTypeDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeLoadBalancerPolicyTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancerPolicyTypes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancerPolicyTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLoadBalancerPolicyTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeLoadBalancerPolicyTypes",
	}
}
