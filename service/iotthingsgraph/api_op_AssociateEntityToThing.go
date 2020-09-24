// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a device with a concrete thing that is in the user's registry. A
// thing can be associated with only one device at a time. If you associate a thing
// with a new device id, its previous association will be removed.
func (c *Client) AssociateEntityToThing(ctx context.Context, params *AssociateEntityToThingInput, optFns ...func(*Options)) (*AssociateEntityToThingOutput, error) {
	stack := middleware.NewStack("AssociateEntityToThing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateEntityToThingMiddlewares(stack)
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
	addOpAssociateEntityToThingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateEntityToThing(options.Region), middleware.Before)
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
			OperationName: "AssociateEntityToThing",
			Err:           err,
		}
	}
	out := result.(*AssociateEntityToThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateEntityToThingInput struct {
	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64
	// The name of the thing to which the entity is to be associated.
	ThingName *string
	// The ID of the device to be associated with the thing. The ID should be in the
	// following format. urn:tdm:REGION/ACCOUNT ID/default:device:DEVICENAME
	EntityId *string
}

type AssociateEntityToThingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateEntityToThingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateEntityToThing{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateEntityToThing{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateEntityToThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "AssociateEntityToThing",
	}
}
