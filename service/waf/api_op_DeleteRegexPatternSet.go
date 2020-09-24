// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
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
// global use. Permanently deletes a RegexPatternSet (). You can't delete a
// RegexPatternSet if it's still used in any RegexMatchSet or if the
// RegexPatternSet is not empty.
func (c *Client) DeleteRegexPatternSet(ctx context.Context, params *DeleteRegexPatternSetInput, optFns ...func(*Options)) (*DeleteRegexPatternSetOutput, error) {
	stack := middleware.NewStack("DeleteRegexPatternSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteRegexPatternSetMiddlewares(stack)
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
	addOpDeleteRegexPatternSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRegexPatternSet(options.Region), middleware.Before)
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
			OperationName: "DeleteRegexPatternSet",
			Err:           err,
		}
	}
	out := result.(*DeleteRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRegexPatternSetInput struct {
	// The value returned by the most recent call to GetChangeToken ().
	ChangeToken *string
	// The RegexPatternSetId of the RegexPatternSet () that you want to delete.
	// RegexPatternSetId is returned by CreateRegexPatternSet () and by
	// ListRegexPatternSets ().
	RegexPatternSetId *string
}

type DeleteRegexPatternSetOutput struct {
	// The ChangeToken that you used to submit the DeleteRegexPatternSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteRegexPatternSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRegexPatternSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRegexPatternSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRegexPatternSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteRegexPatternSet",
	}
}
