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

// Enables the APNs VoIP sandbox channel for an application or updates the status
// and settings of the APNs VoIP sandbox channel for an application.
func (c *Client) UpdateApnsVoipSandboxChannel(ctx context.Context, params *UpdateApnsVoipSandboxChannelInput, optFns ...func(*Options)) (*UpdateApnsVoipSandboxChannelOutput, error) {
	stack := middleware.NewStack("UpdateApnsVoipSandboxChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateApnsVoipSandboxChannelMiddlewares(stack)
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
	addOpUpdateApnsVoipSandboxChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApnsVoipSandboxChannel(options.Region), middleware.Before)
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
			OperationName: "UpdateApnsVoipSandboxChannel",
			Err:           err,
		}
	}
	out := result.(*UpdateApnsVoipSandboxChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApnsVoipSandboxChannelInput struct {
	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	ApplicationId *string
	// Specifies the status and settings of the APNs (Apple Push Notification service)
	// VoIP sandbox channel for an application.
	APNSVoipSandboxChannelRequest *types.APNSVoipSandboxChannelRequest
}

type UpdateApnsVoipSandboxChannelOutput struct {
	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP sandbox channel for an application.
	APNSVoipSandboxChannelResponse *types.APNSVoipSandboxChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateApnsVoipSandboxChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApnsVoipSandboxChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApnsVoipSandboxChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApnsVoipSandboxChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateApnsVoipSandboxChannel",
	}
}
