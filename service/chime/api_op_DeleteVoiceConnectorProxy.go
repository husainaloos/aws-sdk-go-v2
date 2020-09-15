// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the proxy configuration from the specified Amazon Chime Voice Connector.
func (c *Client) DeleteVoiceConnectorProxy(ctx context.Context, params *DeleteVoiceConnectorProxyInput, optFns ...func(*Options)) (*DeleteVoiceConnectorProxyOutput, error) {
	stack := middleware.NewStack("DeleteVoiceConnectorProxy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteVoiceConnectorProxyMiddlewares(stack)
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
	addOpDeleteVoiceConnectorProxyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVoiceConnectorProxy(options.Region), middleware.Before)

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
			OperationName: "DeleteVoiceConnectorProxy",
			Err:           err,
		}
	}
	out := result.(*DeleteVoiceConnectorProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVoiceConnectorProxyInput struct {
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

type DeleteVoiceConnectorProxyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteVoiceConnectorProxyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVoiceConnectorProxy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVoiceConnectorProxy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVoiceConnectorProxy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteVoiceConnectorProxy",
	}
}