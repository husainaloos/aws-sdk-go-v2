// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns an array of RegexPatternSetSummary () objects.
func (c *Client) ListRegexPatternSets(ctx context.Context, params *ListRegexPatternSetsInput, optFns ...func(*Options)) (*ListRegexPatternSetsOutput, error) {
	stack := middleware.NewStack("ListRegexPatternSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRegexPatternSetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRegexPatternSets(options.Region), middleware.Before)

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
			OperationName: "ListRegexPatternSets",
			Err:           err,
		}
	}
	out := result.(*ListRegexPatternSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegexPatternSetsInput struct {
	// If you specify a value for Limit and you have more RegexPatternSet objects than
	// the value of Limit, AWS WAF returns a NextMarker value in the response that
	// allows you to list another group of RegexPatternSet objects. For the second and
	// subsequent ListRegexPatternSets requests, specify the value of NextMarker from
	// the previous response to get information about another batch of RegexPatternSet
	// objects.
	NextMarker *string
	// Specifies the number of RegexPatternSet objects that you want AWS WAF to return
	// for this request. If you have more RegexPatternSet objects than the number you
	// specify for Limit, the response includes a NextMarker value that you can use to
	// get another batch of RegexPatternSet objects.
	Limit *int32
}

type ListRegexPatternSetsOutput struct {
	// If you have more RegexPatternSet objects than the number that you specified for
	// Limit in the request, the response includes a NextMarker value. To list more
	// RegexPatternSet objects, submit another ListRegexPatternSets request, and
	// specify the NextMarker value from the response in the NextMarker value in the
	// next request.
	NextMarker *string
	// An array of RegexPatternSetSummary () objects.
	RegexPatternSets []*types.RegexPatternSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRegexPatternSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRegexPatternSets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRegexPatternSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRegexPatternSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "ListRegexPatternSets",
	}
}