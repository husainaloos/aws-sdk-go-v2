// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches a typed link from a specified source and target object. For more
// information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) DetachTypedLink(ctx context.Context, params *DetachTypedLinkInput, optFns ...func(*Options)) (*DetachTypedLinkOutput, error) {
	if params == nil {
		params = &DetachTypedLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachTypedLink", params, optFns, addOperationDetachTypedLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachTypedLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachTypedLinkInput struct {

	// The Amazon Resource Name (ARN) of the directory where you want to detach the
	// typed link.
	//
	// This member is required.
	DirectoryArn *string

	// Used to accept a typed link specifier as input.
	//
	// This member is required.
	TypedLinkSpecifier *types.TypedLinkSpecifier
}

type DetachTypedLinkOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachTypedLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDetachTypedLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachTypedLink{}, middleware.After)
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
	addOpDetachTypedLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachTypedLink(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetachTypedLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "DetachTypedLink",
	}
}
