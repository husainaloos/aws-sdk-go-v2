// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the existing signaling channel. This is an asynchronous operation and
// takes time to complete. If the MessageTtlSeconds value is updated (either
// increased or reduced), it only applies to new messages sent via this channel
// after it's been updated. Existing messages are still expired as per the previous
// MessageTtlSeconds value.
func (c *Client) UpdateSignalingChannel(ctx context.Context, params *UpdateSignalingChannelInput, optFns ...func(*Options)) (*UpdateSignalingChannelOutput, error) {
	stack := middleware.NewStack("UpdateSignalingChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateSignalingChannelMiddlewares(stack)
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
	addOpUpdateSignalingChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSignalingChannel(options.Region), middleware.Before)

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
			OperationName: "UpdateSignalingChannel",
			Err:           err,
		}
	}
	out := result.(*UpdateSignalingChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSignalingChannelInput struct {
	// The structure containing the configuration for the SINGLE_MASTER type of the
	// signaling channel that you want to update.
	SingleMasterConfiguration *types.SingleMasterConfiguration
	// The Amazon Resource Name (ARN) of the signaling channel that you want to update.
	ChannelARN *string
	// The current version of the signaling channel that you want to update.
	CurrentVersion *string
}

type UpdateSignalingChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateSignalingChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSignalingChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSignalingChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSignalingChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "UpdateSignalingChannel",
	}
}