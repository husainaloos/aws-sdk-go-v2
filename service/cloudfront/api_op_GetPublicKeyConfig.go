// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Return public key configuration informaation
func (c *Client) GetPublicKeyConfig(ctx context.Context, params *GetPublicKeyConfigInput, optFns ...func(*Options)) (*GetPublicKeyConfigOutput, error) {
	stack := middleware.NewStack("GetPublicKeyConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetPublicKeyConfigMiddlewares(stack)
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
	addOpGetPublicKeyConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPublicKeyConfig(options.Region), middleware.Before)
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
			OperationName: "GetPublicKeyConfig",
			Err:           err,
		}
	}
	out := result.(*GetPublicKeyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPublicKeyConfigInput struct {
	// Request the ID for the public key configuration.
	Id *string
}

type GetPublicKeyConfigOutput struct {
	// The current version of the public key configuration. For example:
	// E2QWRUHAPOMQZL.
	ETag *string
	// Return the result for the public key configuration.
	PublicKeyConfig *types.PublicKeyConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetPublicKeyConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetPublicKeyConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetPublicKeyConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPublicKeyConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetPublicKeyConfig",
	}
}
