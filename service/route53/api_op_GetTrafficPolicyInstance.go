// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a specified traffic policy instance. After you submit a
// CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance request, there's a
// brief delay while Amazon Route 53 creates the resource record sets that are
// specified in the traffic policy definition. For more information, see the State
// response element. In the Route 53 console, traffic policy instances are known as
// policy records.
func (c *Client) GetTrafficPolicyInstance(ctx context.Context, params *GetTrafficPolicyInstanceInput, optFns ...func(*Options)) (*GetTrafficPolicyInstanceOutput, error) {
	stack := middleware.NewStack("GetTrafficPolicyInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetTrafficPolicyInstanceMiddlewares(stack)
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
	addOpGetTrafficPolicyInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTrafficPolicyInstance(options.Region), middleware.Before)
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
			OperationName: "GetTrafficPolicyInstance",
			Err:           err,
		}
	}
	out := result.(*GetTrafficPolicyInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets information about a specified traffic policy instance.
type GetTrafficPolicyInstanceInput struct {
	// The ID of the traffic policy instance that you want to get information about.
	Id *string
}

// A complex type that contains information about the resource record sets that
// Amazon Route 53 created based on a specified traffic policy.
type GetTrafficPolicyInstanceOutput struct {
	// A complex type that contains settings for the traffic policy instance.
	TrafficPolicyInstance *types.TrafficPolicyInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetTrafficPolicyInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetTrafficPolicyInstance{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetTrafficPolicyInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTrafficPolicyInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetTrafficPolicyInstance",
	}
}
