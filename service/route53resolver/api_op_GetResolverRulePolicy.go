// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a resolver rule policy. A resolver rule policy specifies
// the Resolver operations and resources that you want to allow another AWS account
// to be able to use.
func (c *Client) GetResolverRulePolicy(ctx context.Context, params *GetResolverRulePolicyInput, optFns ...func(*Options)) (*GetResolverRulePolicyOutput, error) {
	stack := middleware.NewStack("GetResolverRulePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetResolverRulePolicyMiddlewares(stack)
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
	addOpGetResolverRulePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResolverRulePolicy(options.Region), middleware.Before)
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
			OperationName: "GetResolverRulePolicy",
			Err:           err,
		}
	}
	out := result.(*GetResolverRulePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResolverRulePolicyInput struct {
	// The ID of the resolver rule policy that you want to get information about.
	Arn *string
}

type GetResolverRulePolicyOutput struct {
	// Information about the resolver rule policy that you specified in a
	// GetResolverRulePolicy request.
	ResolverRulePolicy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetResolverRulePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetResolverRulePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResolverRulePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetResolverRulePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetResolverRulePolicy",
	}
}
