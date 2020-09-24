// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Creates a RuleGroup () per the specifications provided. A rule group defines a
// collection of rules to inspect and control web requests that you can use in a
// WebACL (). When you create a rule group, you define an immutable capacity limit.
// If you update a rule group, you must stay within the capacity. This allows
// others to reuse the rule group with confidence in its capacity requirements.
func (c *Client) CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) {
	stack := middleware.NewStack("CreateRuleGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateRuleGroupMiddlewares(stack)
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
	addOpCreateRuleGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRuleGroup(options.Region), middleware.Before)
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
			OperationName: "CreateRuleGroup",
			Err:           err,
		}
	}
	out := result.(*CreateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleGroupInput struct {
	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB) or
	// an API Gateway stage. To work with CloudFront, you must also specify the Region
	// US East (N. Virginia) as follows:
	//
	//     * CLI - Specify the Region when you use
	// the CloudFront scope: --scope=CLOUDFRONT --region=us-east-1.
	//
	//     * API and SDKs
	// - For all calls, use the Region endpoint us-east-1.
	Scope types.Scope
	// An array of key:value pairs to associate with the resource.
	Tags []*types.Tag
	// A description of the rule group that helps with identification. You cannot
	// change the description of a rule group after you create it.
	Description *string
	// The web ACL capacity units (WCUs) required for this rule group. When you create
	// your own rule group, you define this, and you cannot change it after creation.
	// When you add or modify the rules in a rule group, AWS WAF enforces this limit.
	// You can check the capacity for a set of rules using CheckCapacity (). AWS WAF
	// uses WCUs to calculate and control the operating resources that are used to run
	// your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently
	// for each rule type, to reflect the relative cost of each rule. Simple rules that
	// cost little to run use fewer WCUs than more complex rules that use more
	// processing power. Rule group capacity is fixed at creation, which helps users
	// plan their web ACL WCU usage when they use a rule group. The WCU limit for web
	// ACLs is 1,500.
	Capacity *int64
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig *types.VisibilityConfig
	// The name of the rule group. You cannot change the name of a rule group after you
	// create it.
	Name *string
	// The Rule () statements used to identify the web requests that you want to allow,
	// block, or count. Each rule includes one top-level statement that AWS WAF uses to
	// identify matching web requests, and parameters that govern how AWS WAF handles
	// them.
	Rules []*types.Rule
}

type CreateRuleGroupOutput struct {
	// High-level information about a RuleGroup (), returned by operations like create
	// and list. This provides information like the ID, that you can use to retrieve
	// and manage a RuleGroup, and the ARN, that you provide to the
	// RuleGroupReferenceStatement () to use the rule group in a Rule ().
	Summary *types.RuleGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateRuleGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRuleGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRuleGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRuleGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "CreateRuleGroup",
	}
}
