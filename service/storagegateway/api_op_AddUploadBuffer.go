// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures one or more gateway local disks as upload buffer for a specified
// gateway. This operation is supported for the stored volume, cached volume and
// tape gateway types.  <p>In the request, you specify the gateway Amazon Resource
// Name (ARN) to which you want to add upload buffer, and one or more disk IDs that
// you want to configure as upload buffer.</p>
func (c *Client) AddUploadBuffer(ctx context.Context, params *AddUploadBufferInput, optFns ...func(*Options)) (*AddUploadBufferOutput, error) {
	if params == nil {
		params = &AddUploadBufferInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddUploadBuffer", params, optFns, addOperationAddUploadBufferMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddUploadBufferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddUploadBufferInput struct {

	// An array of strings that identify disks that are to be configured as working
	// storage. Each string has a minimum length of 1 and maximum length of 300. You
	// can get the disk IDs from the ListLocalDisks () API.
	//
	// This member is required.
	DiskIds []*string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

type AddUploadBufferOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddUploadBufferMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddUploadBuffer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddUploadBuffer{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpAddUploadBufferValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddUploadBuffer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAddUploadBuffer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "AddUploadBuffer",
	}
}
