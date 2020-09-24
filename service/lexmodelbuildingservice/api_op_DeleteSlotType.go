// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes all versions of the slot type, including the $LATEST version. To delete
// a specific version of the slot type, use the DeleteSlotTypeVersion () operation.
// You can delete a version of a slot type only if it is not referenced. To delete
// a slot type that is referred to in one or more intents, you must remove those
// references first. If you get the ResourceInUseException exception, the exception
// provides an example reference that shows the intent where the slot type is
// referenced. To remove the reference to the slot type, either update the intent
// or delete it. If you get the same exception when you attempt to delete the slot
// type again, repeat until the slot type has no references and the DeleteSlotType
// call is successful. This operation requires permission for the
// lex:DeleteSlotType action.
func (c *Client) DeleteSlotType(ctx context.Context, params *DeleteSlotTypeInput, optFns ...func(*Options)) (*DeleteSlotTypeOutput, error) {
	stack := middleware.NewStack("DeleteSlotType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteSlotTypeMiddlewares(stack)
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
	addOpDeleteSlotTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSlotType(options.Region), middleware.Before)
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
			OperationName: "DeleteSlotType",
			Err:           err,
		}
	}
	out := result.(*DeleteSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSlotTypeInput struct {
	// The name of the slot type. The name is case sensitive.
	Name *string
}

type DeleteSlotTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteSlotTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSlotType{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSlotType{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSlotType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteSlotType",
	}
}
