// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves (queries) pre-aggregated data for a standard execution metric that
// applies to a journey.
func (c *Client) GetJourneyExecutionMetrics(ctx context.Context, params *GetJourneyExecutionMetricsInput, optFns ...func(*Options)) (*GetJourneyExecutionMetricsOutput, error) {
	stack := middleware.NewStack("GetJourneyExecutionMetrics", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetJourneyExecutionMetricsMiddlewares(stack)
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
	addOpGetJourneyExecutionMetricsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetJourneyExecutionMetrics(options.Region), middleware.Before)
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
			OperationName: "GetJourneyExecutionMetrics",
			Err:           err,
		}
	}
	out := result.(*GetJourneyExecutionMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJourneyExecutionMetricsInput struct {
	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string
	// The unique identifier for the journey.
	JourneyId *string
	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	ApplicationId *string
	// The string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string
}

type GetJourneyExecutionMetricsOutput struct {
	// Provides the results of a query that retrieved the data for a standard execution
	// metric that applies to a journey, and provides information about that query.
	JourneyExecutionMetricsResponse *types.JourneyExecutionMetricsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetJourneyExecutionMetricsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetJourneyExecutionMetrics{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetJourneyExecutionMetrics{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetJourneyExecutionMetrics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetJourneyExecutionMetrics",
	}
}
