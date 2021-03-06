// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers an existing and disabled user, group, or resource for Amazon WorkMail
// use by associating a mailbox and calendaring capabilities. It performs no change
// if the user, group, or resource is enabled and fails if the user, group, or
// resource is deleted. This operation results in the accumulation of costs. For
// more information, see Pricing (https://aws.amazon.com//workmail/pricing). The
// equivalent console functionality for this operation is Enable. Users can either
// be created by calling the CreateUser () API operation or they can be
// synchronized from your directory. For more information, see
// DeregisterFromWorkMail ().
func (c *Client) RegisterToWorkMail(ctx context.Context, params *RegisterToWorkMailInput, optFns ...func(*Options)) (*RegisterToWorkMailOutput, error) {
	if params == nil {
		params = &RegisterToWorkMailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterToWorkMail", params, optFns, addOperationRegisterToWorkMailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterToWorkMailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterToWorkMailInput struct {

	// The email for the user, group, or resource to be updated.
	//
	// This member is required.
	Email *string

	// The identifier for the user, group, or resource to be updated.
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the user, group, or resource
	// exists.
	//
	// This member is required.
	OrganizationId *string
}

type RegisterToWorkMailOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterToWorkMailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterToWorkMail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterToWorkMail{}, middleware.After)
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
	addOpRegisterToWorkMailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterToWorkMail(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterToWorkMail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "RegisterToWorkMail",
	}
}
