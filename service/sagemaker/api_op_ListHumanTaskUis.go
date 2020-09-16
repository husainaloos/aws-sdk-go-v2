// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the human task user interfaces in your account.
func (c *Client) ListHumanTaskUis(ctx context.Context, params *ListHumanTaskUisInput, optFns ...func(*Options)) (*ListHumanTaskUisOutput, error) {
	stack := middleware.NewStack("ListHumanTaskUis", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListHumanTaskUisMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHumanTaskUis(options.Region), middleware.Before)

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
			OperationName: "ListHumanTaskUis",
			Err:           err,
		}
	}
	out := result.(*ListHumanTaskUisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHumanTaskUisInput struct {
	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults, then a NextToken will be provided
	// in the output that you can use to resume pagination.
	MaxResults *int32
	// A filter that returns only human task user interfaces that were created before
	// the specified timestamp.
	CreationTimeBefore *time.Time
	// An optional value that specifies whether you want the results sorted in
	// Ascending or Descending order.
	SortOrder types.SortOrder
	// A token to resume pagination.
	NextToken *string
	// A filter that returns only human task user interfaces with a creation time
	// greater than or equal to the specified timestamp.
	CreationTimeAfter *time.Time
}

type ListHumanTaskUisOutput struct {
	// An array of objects describing the human task user interfaces.
	HumanTaskUiSummaries []*types.HumanTaskUiSummary
	// A token to resume pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListHumanTaskUisMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListHumanTaskUis{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHumanTaskUis{}, middleware.After)
}

func newServiceMetadataMiddleware_opListHumanTaskUis(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListHumanTaskUis",
	}
}