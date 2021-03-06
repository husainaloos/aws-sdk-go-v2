// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the specified rule. If the rule does not exist, the operation fails.
// <p>When you enable a rule, incoming events might not immediately start matching
// to a newly enabled rule. Allow a short period of time for changes to take
// effect.</p>
func (c *Client) EnableRule(ctx context.Context, params *EnableRuleInput, optFns ...func(*Options)) (*EnableRuleOutput, error) {
	if params == nil {
		params = &EnableRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableRule", params, optFns, addOperationEnableRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableRuleInput struct {

	// The name of the rule.
	//
	// This member is required.
	Name *string

	// The event bus associated with the rule. If you omit this, the default event bus
	// is used.
	EventBusName *string
}

type EnableRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableRule{}, middleware.After)
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
	addOpEnableRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "EnableRule",
	}
}
