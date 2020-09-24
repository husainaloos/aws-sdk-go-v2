// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update an event destination in a configuration set. An event destination is a
// location that you publish information about your voice calls to. For example,
// you can log an event to an Amazon CloudWatch destination when a call fails.
func (c *Client) UpdateConfigurationSetEventDestination(ctx context.Context, params *UpdateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*UpdateConfigurationSetEventDestinationOutput, error) {
	stack := middleware.NewStack("UpdateConfigurationSetEventDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateConfigurationSetEventDestinationMiddlewares(stack)
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
	addOpUpdateConfigurationSetEventDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfigurationSetEventDestination(options.Region), middleware.Before)
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
			OperationName: "UpdateConfigurationSetEventDestination",
			Err:           err,
		}
	}
	out := result.(*UpdateConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateConfigurationSetEventDestinationRequest
type UpdateConfigurationSetEventDestinationInput struct {
	// An object that defines a single event destination.
	EventDestination *types.EventDestinationDefinition
	// EventDestinationName
	EventDestinationName *string
	// ConfigurationSetName
	ConfigurationSetName *string
}

// An empty object that indicates that the event destination was updated
// successfully.
type UpdateConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConfigurationSetEventDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConfigurationSetEventDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateConfigurationSetEventDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "UpdateConfigurationSetEventDestination",
	}
}
