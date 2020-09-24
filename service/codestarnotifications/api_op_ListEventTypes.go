// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the event types available for configuring
// notifications.
func (c *Client) ListEventTypes(ctx context.Context, params *ListEventTypesInput, optFns ...func(*Options)) (*ListEventTypesOutput, error) {
	stack := middleware.NewStack("ListEventTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListEventTypesMiddlewares(stack)
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
	addOpListEventTypesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEventTypes(options.Region), middleware.Before)
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
			OperationName: "ListEventTypes",
			Err:           err,
		}
	}
	out := result.(*ListEventTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventTypesInput struct {
	// The filters to use to return information by service or resource type.
	Filters []*types.ListEventTypesFilter
	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
	// A non-negative integer used to limit the number of returned results. The default
	// number is 50. The maximum number of results that can be returned is 100.
	MaxResults *int32
}

type ListEventTypesOutput struct {
	// Information about each event, including service name, resource type, event ID,
	// and event name.
	EventTypes []*types.EventTypeSummary
	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListEventTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListEventTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEventTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "ListEventTypes",
	}
}
