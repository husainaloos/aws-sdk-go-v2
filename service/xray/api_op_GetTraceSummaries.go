// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves IDs and annotations for traces available for a specified time frame
// using an optional filter. To get the full traces, pass the trace IDs to
// BatchGetTraces. A filter expression can target traced requests that hit specific
// service nodes or edges, have errors, or come from a known user. For example, the
// following filter expression targets traces that pass through api.example.com:
// service("api.example.com") This filter expression finds traces that have an
// annotation named account with the value 12345: annotation.account = "12345" For
// a full list of indexed fields and keywords that you can use in filter
// expressions, see Using Filter Expressions
// (https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html) in
// the AWS X-Ray Developer Guide.
func (c *Client) GetTraceSummaries(ctx context.Context, params *GetTraceSummariesInput, optFns ...func(*Options)) (*GetTraceSummariesOutput, error) {
	stack := middleware.NewStack("GetTraceSummaries", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetTraceSummariesMiddlewares(stack)
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
	addOpGetTraceSummariesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTraceSummaries(options.Region), middleware.Before)
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
			OperationName: "GetTraceSummaries",
			Err:           err,
		}
	}
	out := result.(*GetTraceSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTraceSummariesInput struct {
	// The end of the time frame for which to retrieve traces.
	EndTime *time.Time
	// Set to true to get summaries for only a subset of available traces.
	Sampling *bool
	// A paramater to indicate whether to enable sampling on trace summaries. Input
	// parameters are Name and Value.
	SamplingStrategy *types.SamplingStrategy
	// Specify the pagination token returned by a previous request to retrieve the next
	// page of results.
	NextToken *string
	// Specify a filter expression to retrieve trace summaries for services or requests
	// that meet certain requirements.
	FilterExpression *string
	// A parameter to indicate whether to query trace summaries by TraceId or Event
	// time.
	TimeRangeType types.TimeRangeType
	// The start of the time frame for which to retrieve traces.
	StartTime *time.Time
}

type GetTraceSummariesOutput struct {
	// Trace IDs and annotations for traces that were found in the specified time
	// frame.
	TraceSummaries []*types.TraceSummary
	// The start time of this page of results.
	ApproximateTime *time.Time
	// If the requested time frame contained more than one page of results, you can use
	// this token to retrieve the next page. The first page contains the most most
	// recent results, closest to the end of the time frame.
	NextToken *string
	// The total number of traces processed, including traces that did not match the
	// specified filter expression.
	TracesProcessedCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetTraceSummariesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetTraceSummaries{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTraceSummaries{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTraceSummaries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetTraceSummaries",
	}
}
