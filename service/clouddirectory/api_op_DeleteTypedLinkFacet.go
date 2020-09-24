// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a TypedLinkFacet (). For more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) DeleteTypedLinkFacet(ctx context.Context, params *DeleteTypedLinkFacetInput, optFns ...func(*Options)) (*DeleteTypedLinkFacetOutput, error) {
	stack := middleware.NewStack("DeleteTypedLinkFacet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteTypedLinkFacetMiddlewares(stack)
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
	addOpDeleteTypedLinkFacetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTypedLinkFacet(options.Region), middleware.Before)
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
			OperationName: "DeleteTypedLinkFacet",
			Err:           err,
		}
	}
	out := result.(*DeleteTypedLinkFacetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTypedLinkFacetInput struct {
	// The unique name of the typed link facet.
	Name *string
	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns ().
	SchemaArn *string
}

type DeleteTypedLinkFacetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteTypedLinkFacetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteTypedLinkFacet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteTypedLinkFacet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTypedLinkFacet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "DeleteTypedLinkFacet",
	}
}
