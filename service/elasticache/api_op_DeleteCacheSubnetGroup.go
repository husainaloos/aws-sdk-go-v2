// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a cache subnet group. You cannot delete a cache subnet group if it is
// associated with any clusters.
func (c *Client) DeleteCacheSubnetGroup(ctx context.Context, params *DeleteCacheSubnetGroupInput, optFns ...func(*Options)) (*DeleteCacheSubnetGroupOutput, error) {
	if params == nil {
		params = &DeleteCacheSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCacheSubnetGroup", params, optFns, addOperationDeleteCacheSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCacheSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteCacheSubnetGroup operation.
type DeleteCacheSubnetGroupInput struct {

	// The name of the cache subnet group to delete. Constraints: Must contain no more
	// than 255 alphanumeric characters or hyphens.
	//
	// This member is required.
	CacheSubnetGroupName *string
}

type DeleteCacheSubnetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCacheSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCacheSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCacheSubnetGroup{}, middleware.After)
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
	addOpDeleteCacheSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCacheSubnetGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCacheSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteCacheSubnetGroup",
	}
}
