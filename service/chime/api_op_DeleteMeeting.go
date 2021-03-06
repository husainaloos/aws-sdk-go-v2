// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified Amazon Chime SDK meeting. When a meeting is deleted, its
// attendees are also deleted and clients can no longer join it. For more
// information about the Amazon Chime SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) DeleteMeeting(ctx context.Context, params *DeleteMeetingInput, optFns ...func(*Options)) (*DeleteMeetingOutput, error) {
	if params == nil {
		params = &DeleteMeetingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMeeting", params, optFns, addOperationDeleteMeetingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMeetingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMeetingInput struct {

	// The Amazon Chime SDK meeting ID.
	//
	// This member is required.
	MeetingId *string
}

type DeleteMeetingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteMeetingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMeeting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMeeting{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteMeetingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMeeting(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteMeeting(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteMeeting",
	}
}
