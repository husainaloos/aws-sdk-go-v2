// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends messages to a channel.
func (c *Client) BatchPutMessage(ctx context.Context, params *BatchPutMessageInput, optFns ...func(*Options)) (*BatchPutMessageOutput, error) {
	stack := middleware.NewStack("BatchPutMessage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchPutMessageMiddlewares(stack)
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
	addOpBatchPutMessageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutMessage(options.Region), middleware.Before)
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
			OperationName: "BatchPutMessage",
			Err:           err,
		}
	}
	out := result.(*BatchPutMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutMessageInput struct {
	// The list of messages to be sent. Each message has format: '{ "messageId":
	// "string", "payload": "string"}'. Note that the field names of message payloads
	// (data) that you send to AWS IoT Analytics:
	//
	//     * Must contain only alphanumeric
	// characters and undescores (_); no other special characters are allowed.
	//
	//     *
	// Must begin with an alphabetic character or single underscore (_).
	//
	//     * Cannot
	// contain hyphens (-).
	//
	//     * In regular expression terms:
	// "^[A-Za-z_]([A-Za-z0-9]*|[A-Za-z0-9][A-Za-z0-9_]*)$".
	//
	//     * Cannot be greater
	// than 255 characters.
	//
	//     * Are case-insensitive. (Fields named "foo" and "FOO"
	// in the same payload are considered duplicates.)
	//
	// For example, {"temp_01": 29} or
	// {"_temp_01": 29} are valid, but {"temp-01": 29}, {"01_temp": 29} or
	// {"__temp_01": 29} are invalid in message payloads.
	Messages []*types.Message
	// The name of the channel where the messages are sent.
	ChannelName *string
}

type BatchPutMessageOutput struct {
	// A list of any errors encountered when sending the messages to the channel.
	BatchPutMessageErrorEntries []*types.BatchPutMessageErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchPutMessageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchPutMessage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchPutMessage{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchPutMessage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "BatchPutMessage",
	}
}
