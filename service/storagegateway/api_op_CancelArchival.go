// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels archiving of a virtual tape to the virtual tape shelf (VTS) after the
// archiving process is initiated. This operation is only supported in the tape
// gateway type.
func (c *Client) CancelArchival(ctx context.Context, params *CancelArchivalInput, optFns ...func(*Options)) (*CancelArchivalOutput, error) {
	stack := middleware.NewStack("CancelArchival", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCancelArchivalMiddlewares(stack)
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
	addOpCancelArchivalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelArchival(options.Region), middleware.Before)
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
			OperationName: "CancelArchival",
			Err:           err,
		}
	}
	out := result.(*CancelArchivalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CancelArchivalInput
type CancelArchivalInput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
	// The Amazon Resource Name (ARN) of the virtual tape you want to cancel archiving
	// for.
	TapeARN *string
}

// CancelArchivalOutput
type CancelArchivalOutput struct {
	// The Amazon Resource Name (ARN) of the virtual tape for which archiving was
	// canceled.
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCancelArchivalMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCancelArchival{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelArchival{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelArchival(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CancelArchival",
	}
}
