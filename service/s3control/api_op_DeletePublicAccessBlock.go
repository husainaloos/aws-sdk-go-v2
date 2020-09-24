// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the PublicAccessBlock configuration for an Amazon Web Services account.
func (c *Client) DeletePublicAccessBlock(ctx context.Context, params *DeletePublicAccessBlockInput, optFns ...func(*Options)) (*DeletePublicAccessBlockOutput, error) {
	stack := middleware.NewStack("DeletePublicAccessBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeletePublicAccessBlockMiddlewares(stack)
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
	addOpDeletePublicAccessBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePublicAccessBlock(options.Region), middleware.Before)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)

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
			OperationName: "DeletePublicAccessBlock",
			Err:           err,
		}
	}
	out := result.(*DeletePublicAccessBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePublicAccessBlockInput struct {
	// The account ID for the Amazon Web Services account whose PublicAccessBlock
	// configuration you want to remove.
	AccountId *string
}

type DeletePublicAccessBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeletePublicAccessBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeletePublicAccessBlock{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeletePublicAccessBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePublicAccessBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeletePublicAccessBlock",
	}
}
