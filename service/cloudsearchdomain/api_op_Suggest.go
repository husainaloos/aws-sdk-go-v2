// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearchdomain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves autocomplete suggestions for a partial query string. You can use
// suggestions enable you to display likely matches before users finish typing. In
// Amazon CloudSearch, suggestions are based on the contents of a particular text
// field. When you request suggestions, Amazon CloudSearch finds all of the
// documents whose values in the suggester field start with the specified query
// string. The beginning of the field must match the query string to be considered
// a match. For more information about configuring suggesters and retrieving
// suggestions, see Getting Suggestions
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.  <p>The endpoint for submitting
// <code>Suggest</code> requests is domain-specific and requires the
// <code>--endpoint-url</code> option. You submit suggest requests to a domain's
// search endpoint. To get the search endpoint for your domain, use the Amazon
// CloudSearch configuration service <code>DescribeDomains</code> action. The
// endpoints are also available on the domain dashboard in the Amazon CloudSearch
// console.</p>
func (c *Client) Suggest(ctx context.Context, params *SuggestInput, optFns ...func(*Options)) (*SuggestOutput, error) {
	stack := middleware.NewStack("Suggest", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSuggestMiddlewares(stack)
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
	addOpSuggestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSuggest(options.Region), middleware.Before)
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
			OperationName: "Suggest",
			Err:           err,
		}
	}
	out := result.(*SuggestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the Suggest request.
type SuggestInput struct {
	// Specifies the string for which you want to get suggestions.
	Query *string
	// Specifies the maximum number of suggestions to return.
	Size *int64
	// Specifies the name of the suggester to use to find suggested matches.
	Suggester *string
}

// Contains the response to a Suggest request.
type SuggestOutput struct {
	// Container for the matching search suggestion information.
	Suggest *types.SuggestModel
	// The status of a SuggestRequest. Contains the resource ID (rid) and how long it
	// took to process the request (timems).
	Status *types.SuggestStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSuggestMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSuggest{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSuggest{}, middleware.After)
}

func newServiceMetadataMiddleware_opSuggest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "Suggest",
	}
}
