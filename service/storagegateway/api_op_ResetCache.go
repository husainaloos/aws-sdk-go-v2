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

// Resets all cache disks that have encountered an error and makes the disks
// available for reconfiguration as cache storage. If your cache disk encounters an
// error, the gateway prevents read and write operations on virtual tapes in the
// gateway. For example, an error can occur when a disk is corrupted or removed
// from the gateway. When a cache is reset, the gateway loses its cache storage. At
// this point, you can reconfigure the disks as cache disks. This operation is only
// supported in the cached volume and tape types.  <important> <p>If the cache disk
// you are resetting contains data that has not been uploaded to Amazon S3 yet,
// that data can be lost. After you reset cache disks, there will be no configured
// cache disks left in the gateway, so you must configure at least one new cache
// disk for your gateway to function properly.</p> </important>
func (c *Client) ResetCache(ctx context.Context, params *ResetCacheInput, optFns ...func(*Options)) (*ResetCacheOutput, error) {
	stack := middleware.NewStack("ResetCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResetCacheMiddlewares(stack)
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
	addOpResetCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetCache(options.Region), middleware.Before)
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
			OperationName: "ResetCache",
			Err:           err,
		}
	}
	out := result.(*ResetCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetCacheInput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
}

type ResetCacheOutput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResetCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResetCache{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ResetCache",
	}
}
