// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns details about the IAM user or role whose credentials are used to call
// the operation. No permissions are required to perform this operation. If an
// administrator adds a policy to your IAM user or role that explicitly denies
// access to the sts:GetCallerIdentity action, you can still perform this
// operation. Permissions are not required because the same information is returned
// when an IAM user or role is denied access. To view an example response, see I Am
// Not Authorized to Perform: iam:DeleteVirtualMFADevice
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_access-denied-delete-mfa)
// in the IAM User Guide.
func (c *Client) GetCallerIdentity(ctx context.Context, params *GetCallerIdentityInput, optFns ...func(*Options)) (*GetCallerIdentityOutput, error) {
	if params == nil {
		params = &GetCallerIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCallerIdentity", params, optFns, addOperationGetCallerIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCallerIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCallerIdentityInput struct {
}

// Contains the response to a successful GetCallerIdentity () request, including
// information about the entity making the request.
type GetCallerIdentityOutput struct {

	// The AWS account ID number of the account that owns or contains the calling
	// entity.
	Account *string

	// The AWS ARN associated with the calling entity.
	Arn *string

	// The unique identifier of the calling entity. The exact value depends on the type
	// of entity that is making the call. The values returned are those listed in the
	// aws:userid column in the Principal table
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_variables.html#principaltable)
	// found on the Policy Variables reference page in the IAM User Guide.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCallerIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetCallerIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetCallerIdentity{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCallerIdentity(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCallerIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sts",
		OperationName: "GetCallerIdentity",
	}
}
