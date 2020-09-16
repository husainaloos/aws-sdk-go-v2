// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove a public key you previously added to CloudFront.
func (c *Client) DeletePublicKey(ctx context.Context, params *DeletePublicKeyInput, optFns ...func(*Options)) (*DeletePublicKeyOutput, error) {
	stack := middleware.NewStack("DeletePublicKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeletePublicKeyMiddlewares(stack)
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
	addOpDeletePublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePublicKey(options.Region), middleware.Before)

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
			OperationName: "DeletePublicKey",
			Err:           err,
		}
	}
	out := result.(*DeletePublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePublicKeyInput struct {
	// The ID of the public key you want to remove from CloudFront.
	Id *string
	// The value of the ETag header that you received when retrieving the public key
	// identity to delete. For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type DeletePublicKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeletePublicKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeletePublicKey{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeletePublicKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeletePublicKey",
	}
}