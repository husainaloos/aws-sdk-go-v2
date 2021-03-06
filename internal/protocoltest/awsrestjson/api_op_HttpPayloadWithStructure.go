// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This examples serializes a structure in the payload. Note that serializing a
// structure changes the wrapper element name to match the targeted structure.
func (c *Client) HttpPayloadWithStructure(ctx context.Context, params *HttpPayloadWithStructureInput, optFns ...func(*Options)) (*HttpPayloadWithStructureOutput, error) {
	if params == nil {
		params = &HttpPayloadWithStructureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPayloadWithStructure", params, optFns, addOperationHttpPayloadWithStructureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPayloadWithStructureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithStructureInput struct {
	Nested *types.NestedPayload
}

type HttpPayloadWithStructureOutput struct {
	Nested *types.NestedPayload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHttpPayloadWithStructureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpHttpPayloadWithStructure{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpPayloadWithStructure{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithStructure(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opHttpPayloadWithStructure(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadWithStructure",
	}
}
