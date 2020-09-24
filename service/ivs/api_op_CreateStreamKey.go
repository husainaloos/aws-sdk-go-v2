// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a stream key, used to initiate a stream, for the specified channel ARN.
// Note that CreateChannel () creates a stream key. If you subsequently use
// CreateStreamKey on the same channel, it will fail because a stream key already
// exists and there is a limit of 1 stream key per channel. To reset the stream key
// on a channel, use DeleteStreamKey () and then CreateStreamKey.
func (c *Client) CreateStreamKey(ctx context.Context, params *CreateStreamKeyInput, optFns ...func(*Options)) (*CreateStreamKeyOutput, error) {
	stack := middleware.NewStack("CreateStreamKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateStreamKeyMiddlewares(stack)
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
	addOpCreateStreamKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamKey(options.Region), middleware.Before)
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
			OperationName: "CreateStreamKey",
			Err:           err,
		}
	}
	out := result.(*CreateStreamKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamKeyInput struct {
	// ARN of the channel for which to create the stream key.
	ChannelArn *string
	// See Channel$tags ().
	Tags map[string]*string
}

type CreateStreamKeyOutput struct {
	// Stream key used to authenticate an RTMPS stream for ingestion.
	StreamKey *types.StreamKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateStreamKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateStreamKey{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStreamKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateStreamKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "CreateStreamKey",
	}
}
