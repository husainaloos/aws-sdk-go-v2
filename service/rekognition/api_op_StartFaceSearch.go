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

// Starts the asynchronous search for faces in a collection that match the faces of
// persons detected in a stored video. The video must be stored in an Amazon S3
// bucket. Use Video () to specify the bucket name and the filename of the video.
// StartFaceSearch returns a job identifier (JobId) which you use to get the search
// results once the search has completed. When searching is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel. To get the
// search results, first check that the status value published to the Amazon SNS
// topic is SUCCEEDED. If so, call GetFaceSearch () and pass the job identifier
// (JobId) from the initial call to StartFaceSearch. For more information, see
// procedure-person-search-videos ().
func (c *Client) StartFaceSearch(ctx context.Context, params *StartFaceSearchInput, optFns ...func(*Options)) (*StartFaceSearchOutput, error) {
	stack := middleware.NewStack("StartFaceSearch", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartFaceSearchMiddlewares(stack)
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
	addOpStartFaceSearchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartFaceSearch(options.Region), middleware.Before)
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
			OperationName: "StartFaceSearch",
			Err:           err,
		}
	}
	out := result.(*StartFaceSearchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFaceSearchInput struct {
	// The minimum confidence in the person match to return. For example, don't return
	// any matches where confidence in matches is less than 70%. The default value is
	// 80%.
	FaceMatchThreshold *float32
	// ID of the collection that contains the faces you want to search for.
	CollectionId *string
	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartFaceSearch requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string
	// The video you want to search. The video must be stored in an Amazon S3 bucket.
	Video *types.Video
	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string
	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to
	// publish the completion status of the search.
	NotificationChannel *types.NotificationChannel
}

type StartFaceSearchOutput struct {
	// The identifier for the search job. Use JobId to identify the job in a subsequent
	// call to GetFaceSearch.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartFaceSearchMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartFaceSearch{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartFaceSearch{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartFaceSearch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartFaceSearch",
	}
}
