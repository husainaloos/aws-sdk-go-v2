// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns the RateBasedRule () that is specified by the RuleId that
// you included in the GetRateBasedRule request.
func (c *Client) GetRateBasedRule(ctx context.Context, params *GetRateBasedRuleInput, optFns ...func(*Options)) (*GetRateBasedRuleOutput, error) {
	if params == nil {
		params = &GetRateBasedRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRateBasedRule", params, optFns, addOperationGetRateBasedRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRateBasedRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRateBasedRuleInput struct {

	// The RuleId of the RateBasedRule () that you want to get. RuleId is returned by
	// CreateRateBasedRule () and by ListRateBasedRules ().
	//
	// This member is required.
	RuleId *string
}

type GetRateBasedRuleOutput struct {

	// Information about the RateBasedRule () that you specified in the
	// GetRateBasedRule request.
	Rule *types.RateBasedRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRateBasedRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRateBasedRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRateBasedRule{}, middleware.After)
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
	addOpGetRateBasedRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRateBasedRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRateBasedRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetRateBasedRule",
	}
}
