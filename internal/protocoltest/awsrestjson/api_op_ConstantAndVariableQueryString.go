// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example uses fixed query string params and variable query string params.
// The fixed query string parameters and variable parameters must both be
// serialized (implementations may need to merge them together).
func (c *Client) ConstantAndVariableQueryString(ctx context.Context, params *ConstantAndVariableQueryStringInput, optFns ...func(*Options)) (*ConstantAndVariableQueryStringOutput, error) {
	stack := middleware.NewStack("ConstantAndVariableQueryString", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpConstantAndVariableQueryStringMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConstantAndVariableQueryString(options.Region), middleware.Before)

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
			OperationName: "ConstantAndVariableQueryString",
			Err:           err,
		}
	}
	out := result.(*ConstantAndVariableQueryStringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConstantAndVariableQueryStringInput struct {
	Baz      *string
	MaybeSet *string
}

type ConstantAndVariableQueryStringOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpConstantAndVariableQueryStringMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpConstantAndVariableQueryString{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpConstantAndVariableQueryString{}, middleware.After)
}

func newServiceMetadataMiddleware_opConstantAndVariableQueryString(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConstantAndVariableQueryString",
	}
}
