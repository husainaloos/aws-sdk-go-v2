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
// global use. Returns an array of IPSetSummary () objects in the response.
func (c *Client) ListIPSets(ctx context.Context, params *ListIPSetsInput, optFns ...func(*Options)) (*ListIPSetsOutput, error) {
	stack := middleware.NewStack("ListIPSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListIPSetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIPSets(options.Region), middleware.Before)
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
			OperationName: "ListIPSets",
			Err:           err,
		}
	}
	out := result.(*ListIPSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIPSetsInput struct {
	// AWS WAF returns a NextMarker value in the response that allows you to list
	// another group of IPSets. For the second and subsequent ListIPSets requests,
	// specify the value of NextMarker from the previous response to get information
	// about another batch of IPSets.
	NextMarker *string
	// Specifies the number of IPSet objects that you want AWS WAF to return for this
	// request. If you have more IPSet objects than the number you specify for Limit,
	// the response includes a NextMarker value that you can use to get another batch
	// of IPSet objects.
	Limit *int32
}

type ListIPSetsOutput struct {
	// An array of IPSetSummary () objects.
	IPSets []*types.IPSetSummary
	// To list more IPSet objects, submit another ListIPSets request, and in the next
	// request use the NextMarker response value as the NextMarker value.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListIPSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListIPSets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListIPSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListIPSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "ListIPSets",
	}
}
