// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about vocabulary filters.
func (c *Client) ListVocabularyFilters(ctx context.Context, params *ListVocabularyFiltersInput, optFns ...func(*Options)) (*ListVocabularyFiltersOutput, error) {
	stack := middleware.NewStack("ListVocabularyFilters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListVocabularyFiltersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVocabularyFilters(options.Region), middleware.Before)
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
			OperationName: "ListVocabularyFilters",
			Err:           err,
		}
	}
	out := result.(*ListVocabularyFiltersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVocabularyFiltersInput struct {
	// If the result of the previous request to ListVocabularyFilters was truncated,
	// include the NextToken to fetch the next set of collections.
	NextToken *string
	// Filters the response so that it only contains vocabulary filters whose name
	// contains the specified string.
	NameContains *string
	// The maximum number of filters to return in the response. If there are fewer
	// results in the list, this response contains only the actual results.
	MaxResults *int32
}

type ListVocabularyFiltersOutput struct {
	// The ListVocabularyFilters operation returns a page of collections at a time. The
	// maximum size of the page is set by the MaxResults parameter. If there are more
	// jobs in the list than the page size, Amazon Transcribe returns the NextPage
	// token. Include the token in the next request to the ListVocabularyFilters
	// operation to return in the next page of jobs.
	NextToken *string
	// The list of vocabulary filters. It contains at most MaxResults number of
	// filters. If there are more filters, call the ListVocabularyFilters operation
	// again with the NextToken parameter in the request set to the value of the
	// NextToken field in the response.
	VocabularyFilters []*types.VocabularyFilterInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListVocabularyFiltersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListVocabularyFilters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVocabularyFilters{}, middleware.After)
}

func newServiceMetadataMiddleware_opListVocabularyFilters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListVocabularyFilters",
	}
}
