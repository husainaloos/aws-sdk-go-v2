// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a forecast created using the CreateForecast () operation. In addition
// to listing the properties provided in the CreateForecast request, this operation
// lists the following properties:
//
//     * DatasetGroupArn - The dataset group that
// provided the training data.
//
//     * CreationTime
//
//     * LastModificationTime
//
//
// * Status
//
//     * Message - If an error occurred, information about the error.
func (c *Client) DescribeForecast(ctx context.Context, params *DescribeForecastInput, optFns ...func(*Options)) (*DescribeForecastOutput, error) {
	stack := middleware.NewStack("DescribeForecast", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeForecastMiddlewares(stack)
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
	addOpDescribeForecastValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeForecast(options.Region), middleware.Before)

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
			OperationName: "DescribeForecast",
			Err:           err,
		}
	}
	out := result.(*DescribeForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeForecastInput struct {
	// The Amazon Resource Name (ARN) of the forecast.
	ForecastArn *string
}

type DescribeForecastOutput struct {
	// The ARN of the dataset group that provided the data used to train the predictor.
	DatasetGroupArn *string
	// The forecast ARN as specified in the request.
	ForecastArn *string
	// The ARN of the predictor used to generate the forecast.
	PredictorArn *string
	// If an error occurred, an informational message about the error.
	Message *string
	// The quantiles at which probabilistic forecasts were generated.
	ForecastTypes []*string
	// The name of the forecast.
	ForecastName *string
	// Initially, the same as CreationTime (status is CREATE_PENDING). Updated when
	// inference (creating the forecast) starts (status changed to CREATE_IN_PROGRESS),
	// and when inference is complete (status changed to ACTIVE) or fails (status
	// changed to CREATE_FAILED).
	LastModificationTime *time.Time
	// The status of the forecast. States include:
	//
	//     * ACTIVE
	//
	//     * CREATE_PENDING,
	// CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//     * DELETE_PENDING, DELETE_IN_PROGRESS,
	// DELETE_FAILED
	//
	// The Status of the forecast must be ACTIVE before you can query or
	// export the forecast.
	Status *string
	// When the forecast creation task was created.
	CreationTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeForecastMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeForecast{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeForecast{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeForecast(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeForecast",
	}
}