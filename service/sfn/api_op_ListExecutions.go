// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the executions of a state machine that meet the filtering criteria.
// Results are sorted by time, with the most recent execution first. If nextToken
// is returned, there are more results available. The value of nextToken is a
// unique pagination token for each page. Make the call again using the returned
// token to retrieve the next page. Keep all other arguments unchanged. Each
// pagination token expires after 24 hours. Using an expired pagination token will
// return an HTTP 400 InvalidToken error. This operation is eventually consistent.
// The results are best effort and may not reflect very recent updates and changes.
// This API action is not supported by EXPRESS state machines.
func (c *Client) ListExecutions(ctx context.Context, params *ListExecutionsInput, optFns ...func(*Options)) (*ListExecutionsOutput, error) {
	stack := middleware.NewStack("ListExecutions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListExecutionsMiddlewares(stack)
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
	addOpListExecutionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListExecutions(options.Region), middleware.Before)
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
			OperationName: "ListExecutions",
			Err:           err,
		}
	}
	out := result.(*ListExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExecutionsInput struct {
	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string
	// The Amazon Resource Name (ARN) of the state machine whose executions is listed.
	StateMachineArn *string
	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	MaxResults *int32
	// If specified, only list the executions whose current execution status matches
	// the given filter.
	StatusFilter types.ExecutionStatus
}

type ListExecutionsOutput struct {
	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string
	// The list of matching executions.
	Executions []*types.ExecutionListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListExecutionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListExecutions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListExecutions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListExecutions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "ListExecutions",
	}
}
