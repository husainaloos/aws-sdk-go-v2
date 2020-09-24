// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of FAQ lists associated with an index.
func (c *Client) ListFaqs(ctx context.Context, params *ListFaqsInput, optFns ...func(*Options)) (*ListFaqsOutput, error) {
	stack := middleware.NewStack("ListFaqs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListFaqsMiddlewares(stack)
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
	addOpListFaqsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFaqs(options.Region), middleware.Before)
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
			OperationName: "ListFaqs",
			Err:           err,
		}
	}
	out := result.(*ListFaqsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFaqsInput struct {
	// The index that contains the FAQ lists.
	IndexId *string
	// The maximum number of FAQs to return in the response. If there are fewer results
	// in the list, this response contains only the actual results.
	MaxResults *int32
	// If the result of the previous request to ListFaqs was truncated, include the
	// NextToken to fetch the next set of FAQs.
	NextToken *string
}

type ListFaqsOutput struct {
	// information about the FAQs associated with the specified index.
	FaqSummaryItems []*types.FaqSummary
	// The ListFaqs operation returns a page of FAQs at a time. The maximum size of the
	// page is set by the MaxResults parameter. If there are more jobs in the list than
	// the page size, Amazon Kendra returns the NextPage token. Include the token in
	// the next request to the ListFaqs operation to return the next page of FAQs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListFaqsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListFaqs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFaqs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFaqs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "ListFaqs",
	}
}
