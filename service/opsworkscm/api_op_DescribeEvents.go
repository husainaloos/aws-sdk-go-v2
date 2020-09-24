// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes events for a specified server. Results are ordered by time, with
// newest events first. This operation is synchronous. A ResourceNotFoundException
// is thrown when the server does not exist. A ValidationException is raised when
// parameters of the request are not valid.
func (c *Client) DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
	stack := middleware.NewStack("DescribeEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEventsMiddlewares(stack)
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
	addOpDescribeEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEvents(options.Region), middleware.Before)
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
			OperationName: "DescribeEvents",
			Err:           err,
		}
	}
	out := result.(*DescribeEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventsInput struct {
	// The name of the server for which you want to view events.
	ServerName *string
	// NextToken is a string that is returned in some command responses. It indicates
	// that not all entries have been returned, and that you must run at least one more
	// request to get remaining items. To get remaining results, call DescribeEvents
	// again, and assign the token from the previous results as the value of the
	// nextToken parameter. If there are no more results, the response object's
	// nextToken parameter value is null. Setting a nextToken value that was not
	// returned in your previous results causes an InvalidNextTokenException to occur.
	NextToken *string
	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that you
	// can assign to the NextToken request parameter to get the next set of results.
	MaxResults *int32
}

type DescribeEventsOutput struct {
	// Contains the response to a DescribeEvents request.
	ServerEvents []*types.ServerEvent
	// NextToken is a string that is returned in some command responses. It indicates
	// that not all entries have been returned, and that you must run at least one more
	// request to get remaining items. To get remaining results, call DescribeEvents
	// again, and assign the token from the previous results as the value of the
	// nextToken parameter. If there are no more results, the response object's
	// nextToken parameter value is null. Setting a nextToken value that was not
	// returned in your previous results causes an InvalidNextTokenException to occur.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEvents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "DescribeEvents",
	}
}
