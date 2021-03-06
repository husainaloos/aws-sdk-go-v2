// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified collection. Note that this operation removes all faces in
// the collection. For an example, see delete-collection-procedure ().  <p>This
// operation requires permissions to perform the
// <code>rekognition:DeleteCollection</code> action.</p>
func (c *Client) DeleteCollection(ctx context.Context, params *DeleteCollectionInput, optFns ...func(*Options)) (*DeleteCollectionOutput, error) {
	if params == nil {
		params = &DeleteCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCollection", params, optFns, addOperationDeleteCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCollectionInput struct {

	// ID of the collection to delete.
	//
	// This member is required.
	CollectionId *string
}

type DeleteCollectionOutput struct {

	// HTTP status code that indicates the result of the operation.
	StatusCode *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCollection{}, middleware.After)
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
	addOpDeleteCollectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCollection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCollection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DeleteCollection",
	}
}
