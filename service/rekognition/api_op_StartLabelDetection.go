// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts asynchronous detection of labels in a stored video. Amazon Rekognition
// Video can detect labels in a video. Labels are instances of real-world entities.
// This includes objects like flower, tree, and table; events like wedding,
// graduation, and birthday party; concepts like landscape, evening, and nature;
// and activities like a person getting out of a car or a person skiing.  <p>The
// video must be stored in an Amazon S3 bucket. Use <a>Video</a> to specify the
// bucket name and the filename of the video. <code>StartLabelDetection</code>
// returns a job identifier (<code>JobId</code>) which you use to get the results
// of the operation. When label detection is finished, Amazon Rekognition Video
// publishes a completion status to the Amazon Simple Notification Service topic
// that you specify in <code>NotificationChannel</code>.</p> <p>To get the results
// of the label detection operation, first check that the status value published to
// the Amazon SNS topic is <code>SUCCEEDED</code>. If so, call
// <a>GetLabelDetection</a> and pass the job identifier (<code>JobId</code>) from
// the initial call to <code>StartLabelDetection</code>.</p> <p></p>
func (c *Client) StartLabelDetection(ctx context.Context, params *StartLabelDetectionInput, optFns ...func(*Options)) (*StartLabelDetectionOutput, error) {
	stack := middleware.NewStack("StartLabelDetection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartLabelDetectionMiddlewares(stack)
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
	addOpStartLabelDetectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartLabelDetection(options.Region), middleware.Before)
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
			OperationName: "StartLabelDetection",
			Err:           err,
		}
	}
	out := result.(*StartLabelDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartLabelDetectionInput struct {
	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartLabelDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string
	// Specifies the minimum confidence that Amazon Rekognition Video must have in
	// order to return a detected label. Confidence represents how certain Amazon
	// Rekognition is that a label is correctly identified.0 is the lowest confidence.
	// 100 is the highest confidence. Amazon Rekognition Video doesn't return any
	// labels with a confidence level lower than this specified value. If you don't
	// specify MinConfidence, the operation returns labels with confidence values
	// greater than or equal to 50 percent.
	MinConfidence *float32
	// The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the
	// completion status of the label detection operation to.
	NotificationChannel *types.NotificationChannel
	// The video in which you want to detect labels. The video must be stored in an
	// Amazon S3 bucket.
	Video *types.Video
	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string
}

type StartLabelDetectionOutput struct {
	// The identifier for the label detection job. Use JobId to identify the job in a
	// subsequent call to GetLabelDetection.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartLabelDetectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartLabelDetection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartLabelDetection{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartLabelDetection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartLabelDetection",
	}
}
