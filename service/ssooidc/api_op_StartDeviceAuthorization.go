// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates device authorization by requesting a pair of verification codes from
// the authorization service.
func (c *Client) StartDeviceAuthorization(ctx context.Context, params *StartDeviceAuthorizationInput, optFns ...func(*Options)) (*StartDeviceAuthorizationOutput, error) {
	stack := middleware.NewStack("StartDeviceAuthorization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartDeviceAuthorizationMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpStartDeviceAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeviceAuthorization(options.Region), middleware.Before)
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
			OperationName: "StartDeviceAuthorization",
			Err:           err,
		}
	}
	out := result.(*StartDeviceAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeviceAuthorizationInput struct {
	// A secret string that is generated for the client. This value should come from
	// the persisted result of the RegisterClient () API operation.
	ClientSecret *string
	// The URL for the AWS SSO user portal. For more information, see Using the User
	// Portal
	// (https://docs.aws.amazon.com/singlesignon/latest/userguide/using-the-portal.html)
	// in the AWS Single Sign-On User Guide.
	StartUrl *string
	// The unique identifier string for the client that is registered with AWS SSO.
	// This value should come from the persisted result of the RegisterClient () API
	// operation.
	ClientId *string
}

type StartDeviceAuthorizationOutput struct {
	// The URI of the verification page that takes the userCode to authorize the
	// device.
	VerificationUri *string
	// A one-time user verification code. This is needed to authorize an in-use device.
	UserCode *string
	// Indicates the number of seconds the client must wait between attempts when
	// polling for a session.
	Interval *int32
	// Indicates the number of seconds in which the verification code will become
	// invalid.
	ExpiresIn *int32
	// An alternate URL that the client can use to automatically launch a browser. This
	// process skips the manual step in which the user visits the verification page and
	// enters their code.
	VerificationUriComplete *string
	// The short-lived code that is used by the device when polling for a session
	// token.
	DeviceCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartDeviceAuthorizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartDeviceAuthorization{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDeviceAuthorization{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartDeviceAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDeviceAuthorization",
	}
}
