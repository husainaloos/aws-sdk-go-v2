// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This examples uses a @mediaType trait on the payload to force a custom
// content-type to be serialized.
func (c *Client) HttpPayloadTraitsWithMediaType(ctx context.Context, params *HttpPayloadTraitsWithMediaTypeInput, optFns ...func(*Options)) (*HttpPayloadTraitsWithMediaTypeOutput, error) {
	stack := middleware.NewStack("HttpPayloadTraitsWithMediaType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpHttpPayloadTraitsWithMediaTypeMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(options.Region), middleware.Before)

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
			OperationName: "HttpPayloadTraitsWithMediaType",
			Err:           err,
		}
	}
	out := result.(*HttpPayloadTraitsWithMediaTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadTraitsWithMediaTypeInput struct {
	Foo  *string
	Blob []byte
}

type HttpPayloadTraitsWithMediaTypeOutput struct {
	Foo  *string
	Blob []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpHttpPayloadTraitsWithMediaTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadTraitsWithMediaType",
	}
}
