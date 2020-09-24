// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the schedule state of the specified crawler to SCHEDULED, unless the
// crawler is already running or the schedule state is already SCHEDULED.
func (c *Client) StartCrawlerSchedule(ctx context.Context, params *StartCrawlerScheduleInput, optFns ...func(*Options)) (*StartCrawlerScheduleOutput, error) {
	stack := middleware.NewStack("StartCrawlerSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartCrawlerScheduleMiddlewares(stack)
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
	addOpStartCrawlerScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartCrawlerSchedule(options.Region), middleware.Before)
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
			OperationName: "StartCrawlerSchedule",
			Err:           err,
		}
	}
	out := result.(*StartCrawlerScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCrawlerScheduleInput struct {
	// Name of the crawler to schedule.
	CrawlerName *string
}

type StartCrawlerScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartCrawlerScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartCrawlerSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartCrawlerSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartCrawlerSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartCrawlerSchedule",
	}
}
