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

// Detects unsafe content in a specified JPEG or PNG format image. Use
// DetectModerationLabels to moderate images depending on your requirements. For
// example, you might want to filter images that contain nudity, but not images
// containing suggestive content. To filter images, use the labels returned by
// DetectModerationLabels to determine which types of content are appropriate.
// <p>For information about moderation labels, see Detecting Unsafe Content in the
// Amazon Rekognition Developer Guide.</p> <p>You pass the input image either as
// base64-encoded image bytes or as a reference to an image in an Amazon S3 bucket.
// If you use the AWS CLI to call Amazon Rekognition operations, passing image
// bytes is not supported. The image must be either a PNG or JPEG formatted file.
// </p>
func (c *Client) DetectModerationLabels(ctx context.Context, params *DetectModerationLabelsInput, optFns ...func(*Options)) (*DetectModerationLabelsOutput, error) {
	stack := middleware.NewStack("DetectModerationLabels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetectModerationLabelsMiddlewares(stack)
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
	addOpDetectModerationLabelsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectModerationLabels(options.Region), middleware.Before)
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
			OperationName: "DetectModerationLabels",
			Err:           err,
		}
	}
	out := result.(*DetectModerationLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectModerationLabelsInput struct {
	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	Image *types.Image
	// Sets up the configuration for human evaluation, including the FlowDefinition the
	// image will be sent to.
	HumanLoopConfig *types.HumanLoopConfig
	// Specifies the minimum confidence level for the labels to return. Amazon
	// Rekognition doesn't return any labels with a confidence level lower than this
	// specified value. If you don't specify MinConfidence, the operation returns
	// labels with confidence values greater than or equal to 50 percent.
	MinConfidence *float32
}

type DetectModerationLabelsOutput struct {
	// Shows the results of the human in the loop evaluation.
	HumanLoopActivationOutput *types.HumanLoopActivationOutput
	// Version number of the moderation detection model that was used to detect unsafe
	// content.
	ModerationModelVersion *string
	// Array of detected Moderation labels and the time, in milliseconds from the start
	// of the video, they were detected.
	ModerationLabels []*types.ModerationLabel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetectModerationLabelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetectModerationLabels{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectModerationLabels{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetectModerationLabels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectModerationLabels",
	}
}
