// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets the personal meeting PIN for the specified user on an Amazon Chime
// account. Returns the User () object with the updated personal meeting PIN.
func (c *Client) ResetPersonalPIN(ctx context.Context, params *ResetPersonalPINInput, optFns ...func(*Options)) (*ResetPersonalPINOutput, error) {
	if params == nil {
		params = &ResetPersonalPINInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetPersonalPIN", params, optFns, addOperationResetPersonalPINMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetPersonalPINOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetPersonalPINInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The user ID.
	//
	// This member is required.
	UserId *string
}

type ResetPersonalPINOutput struct {

	// The user details and new personal meeting PIN.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResetPersonalPINMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpResetPersonalPIN{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpResetPersonalPIN{}, middleware.After)
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
	addOpResetPersonalPINValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetPersonalPIN(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opResetPersonalPIN(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ResetPersonalPIN",
	}
}
