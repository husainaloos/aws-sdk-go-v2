// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes AWS Shield Advanced from an account. AWS Shield Advanced requires a
// 1-year subscription commitment. You cannot delete a subscription prior to the
// completion of that commitment.
func (c *Client) DeleteSubscription(ctx context.Context, params *DeleteSubscriptionInput, optFns ...func(*Options)) (*DeleteSubscriptionOutput, error) {
	stack := middleware.NewStack("DeleteSubscription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteSubscriptionMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSubscription(options.Region), middleware.Before)
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
			OperationName: "DeleteSubscription",
			Err:           err,
		}
	}
	out := result.(*DeleteSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSubscriptionInput struct {
}

type DeleteSubscriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteSubscriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSubscription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSubscription{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DeleteSubscription",
	}
}
