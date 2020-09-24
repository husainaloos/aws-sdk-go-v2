// Code generated by smithy-go-codegen DO NOT EDIT.

package forecastquery

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecastquery/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a forecast for a single item, filtered by the supplied criteria. The
// criteria is a key-value pair. The key is either item_id (or the equivalent
// non-timestamp, non-target field) from the TARGET_TIME_SERIES dataset, or one of
// the forecast dimensions specified as part of the FeaturizationConfig object. By
// default, QueryForecast returns the complete date range for the filtered
// forecast. You can request a specific date range. To get the full forecast, use
// the CreateForecastExportJob
// (https://docs.aws.amazon.com/en_us/forecast/latest/dg/API_CreateForecastExportJob.html)
// operation. The forecasts generated by Amazon Forecast are in the same timezone
// as the dataset that was used to create the predictor.
func (c *Client) QueryForecast(ctx context.Context, params *QueryForecastInput, optFns ...func(*Options)) (*QueryForecastOutput, error) {
	stack := middleware.NewStack("QueryForecast", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpQueryForecastMiddlewares(stack)
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
	addOpQueryForecastValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryForecast(options.Region), middleware.Before)
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
			OperationName: "QueryForecast",
			Err:           err,
		}
	}
	out := result.(*QueryForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryForecastInput struct {
	// The filtering criteria to apply when retrieving the forecast. For example, to
	// get the forecast for client_21 in the electricity usage dataset, specify the
	// following: {"item_id" : "client_21"}
	//     <p>To get the full forecast, use the <a
	// href="https://docs.aws.amazon.com/en_us/forecast/latest/dg/API_CreateForecastExportJob.html">CreateForecastExportJob</a>
	// operation.</p>
	Filters map[string]*string
	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string
	// The Amazon Resource Name (ARN) of the forecast to query.
	ForecastArn *string
	// The start date for the forecast. Specify the date using this format:
	// yyyy-MM-dd'T'HH:mm:ss (ISO 8601 format). For example, 2015-01-01T08:00:00.
	StartDate *string
	// The end date for the forecast. Specify the date using this format:
	// yyyy-MM-dd'T'HH:mm:ss (ISO 8601 format). For example, 2015-01-01T20:00:00.
	EndDate *string
}

type QueryForecastOutput struct {
	// The forecast.
	Forecast *types.Forecast

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpQueryForecastMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpQueryForecast{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpQueryForecast{}, middleware.After)
}

func newServiceMetadataMiddleware_opQueryForecast(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "QueryForecast",
	}
}
