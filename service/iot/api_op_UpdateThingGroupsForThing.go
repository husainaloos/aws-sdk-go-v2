// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the groups to which the thing belongs.
func (c *Client) UpdateThingGroupsForThing(ctx context.Context, params *UpdateThingGroupsForThingInput, optFns ...func(*Options)) (*UpdateThingGroupsForThingOutput, error) {
	stack := middleware.NewStack("UpdateThingGroupsForThing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateThingGroupsForThingMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThingGroupsForThing(options.Region), middleware.Before)
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
			OperationName: "UpdateThingGroupsForThing",
			Err:           err,
		}
	}
	out := result.(*UpdateThingGroupsForThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateThingGroupsForThingInput struct {
	// Override dynamic thing groups with static thing groups when 10-group limit is
	// reached. If a thing belongs to 10 thing groups, and one or more of those groups
	// are dynamic thing groups, adding a thing to a static group removes the thing
	// from the last dynamic group.
	OverrideDynamicGroups *bool
	// The groups from which the thing will be removed.
	ThingGroupsToRemove []*string
	// The thing whose group memberships will be updated.
	ThingName *string
	// The groups to which the thing will be added.
	ThingGroupsToAdd []*string
}

type UpdateThingGroupsForThingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateThingGroupsForThingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThingGroupsForThing{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThingGroupsForThing{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateThingGroupsForThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateThingGroupsForThing",
	}
}
