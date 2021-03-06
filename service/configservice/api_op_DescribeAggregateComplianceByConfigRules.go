// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of compliant and noncompliant rules with the number of resources
// for compliant and noncompliant rules. The results can return an empty result
// page, but if you have a nextToken, the results are displayed on the next page.
func (c *Client) DescribeAggregateComplianceByConfigRules(ctx context.Context, params *DescribeAggregateComplianceByConfigRulesInput, optFns ...func(*Options)) (*DescribeAggregateComplianceByConfigRulesOutput, error) {
	if params == nil {
		params = &DescribeAggregateComplianceByConfigRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAggregateComplianceByConfigRules", params, optFns, addOperationDescribeAggregateComplianceByConfigRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAggregateComplianceByConfigRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAggregateComplianceByConfigRulesInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// Filters the results by ConfigRuleComplianceFilters object.
	Filters *types.ConfigRuleComplianceFilters

	// The maximum number of evaluation results returned on each page. The default is
	// maximum. If you specify 0, AWS Config uses the default.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type DescribeAggregateComplianceByConfigRulesOutput struct {

	// Returns a list of AggregateComplianceByConfigRule object.
	AggregateComplianceByConfigRules []*types.AggregateComplianceByConfigRule

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAggregateComplianceByConfigRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAggregateComplianceByConfigRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAggregateComplianceByConfigRules{}, middleware.After)
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
	addOpDescribeAggregateComplianceByConfigRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAggregateComplianceByConfigRules(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAggregateComplianceByConfigRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeAggregateComplianceByConfigRules",
	}
}
