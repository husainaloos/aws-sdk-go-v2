// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches one or more Lightsail instances to a load balancer. After some time,
// the instances are attached to the load balancer and the health check status is
// available. The attach instances to load balancer operation supports tag-based
// access control via resource tags applied to the resource identified by load
// balancer name. For more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) AttachInstancesToLoadBalancer(ctx context.Context, params *AttachInstancesToLoadBalancerInput, optFns ...func(*Options)) (*AttachInstancesToLoadBalancerOutput, error) {
	stack := middleware.NewStack("AttachInstancesToLoadBalancer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAttachInstancesToLoadBalancerMiddlewares(stack)
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
	addOpAttachInstancesToLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachInstancesToLoadBalancer(options.Region), middleware.Before)
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
			OperationName: "AttachInstancesToLoadBalancer",
			Err:           err,
		}
	}
	out := result.(*AttachInstancesToLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachInstancesToLoadBalancerInput struct {
	// An array of strings representing the instance name(s) you want to attach to your
	// load balancer. An instance must be running before you can attach it to your load
	// balancer. There are no additional limits on the number of instances you can
	// attach to your load balancer, aside from the limit of Lightsail instances you
	// can create in your account (20).
	InstanceNames []*string
	// The name of the load balancer.
	LoadBalancerName *string
}

type AttachInstancesToLoadBalancerOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAttachInstancesToLoadBalancerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAttachInstancesToLoadBalancer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachInstancesToLoadBalancer{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachInstancesToLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "AttachInstancesToLoadBalancer",
	}
}
