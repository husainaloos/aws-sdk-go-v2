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

// Configures one or more gateway local disks as cache for a gateway. This
// operation is only supported in the cached volume, tape, and file gateway type
// (see How AWS Storage Gateway works (architecture)
// (https://docs.aws.amazon.com/storagegateway/latest/userguide/StorageGatewayConcepts.html).
// <p>In the request, you specify the gateway Amazon Resource Name (ARN) to which
// you want to add cache, and one or more disk IDs that you want to configure as
// cache.</p>
func (c *Client) AddCache(ctx context.Context, params *AddCacheInput, optFns ...func(*Options)) (*AddCacheOutput, error) {
	stack := middleware.NewStack("AddCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddCacheMiddlewares(stack)
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
	addOpAddCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddCache(options.Region), middleware.Before)
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
			OperationName: "AddCache",
			Err:           err,
		}
	}
	out := result.(*AddCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddCacheInput struct {
	// An array of strings that identify disks that are to be configured as working
	// storage. Each string has a minimum length of 1 and maximum length of 300. You
	// can get the disk IDs from the ListLocalDisks () API.
	DiskIds []*string
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
}

type AddCacheOutput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddCache{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "AddCache",
	}
}
