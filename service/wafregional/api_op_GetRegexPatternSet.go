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
// global use. Returns the RegexPatternSet () specified by RegexPatternSetId.
func (c *Client) GetRegexPatternSet(ctx context.Context, params *GetRegexPatternSetInput, optFns ...func(*Options)) (*GetRegexPatternSetOutput, error) {
	stack := middleware.NewStack("GetRegexPatternSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRegexPatternSetMiddlewares(stack)
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
	addOpGetRegexPatternSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRegexPatternSet(options.Region), middleware.Before)
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
			OperationName: "GetRegexPatternSet",
			Err:           err,
		}
	}
	out := result.(*GetRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRegexPatternSetInput struct {
	// The RegexPatternSetId of the RegexPatternSet () that you want to get.
	// RegexPatternSetId is returned by CreateRegexPatternSet () and by
	// ListRegexPatternSets ().
	RegexPatternSetId *string
}

type GetRegexPatternSetOutput struct {
	// Information about the RegexPatternSet () that you specified in the
	// GetRegexPatternSet request, including the identifier of the pattern set and the
	// regular expression patterns you want AWS WAF to search for.
	RegexPatternSet *types.RegexPatternSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRegexPatternSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRegexPatternSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRegexPatternSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRegexPatternSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetRegexPatternSet",
	}
}
