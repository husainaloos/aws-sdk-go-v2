// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches the specified principal to the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
func (c *Client) AttachThingPrincipal(ctx context.Context, params *AttachThingPrincipalInput, optFns ...func(*Options)) (*AttachThingPrincipalOutput, error) {
	if params == nil {
		params = &AttachThingPrincipalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachThingPrincipal", params, optFns, addOperationAttachThingPrincipalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachThingPrincipalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the AttachThingPrincipal operation.
type AttachThingPrincipalInput struct {

	// The principal, which can be a certificate ARN (as returned from the
	// CreateCertificate operation) or an Amazon Cognito ID.
	//
	// This member is required.
	Principal *string

	// The name of the thing.
	//
	// This member is required.
	ThingName *string
}

// The output from the AttachThingPrincipal operation.
type AttachThingPrincipalOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachThingPrincipalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAttachThingPrincipal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAttachThingPrincipal{}, middleware.After)
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
	addOpAttachThingPrincipalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachThingPrincipal(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAttachThingPrincipal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "AttachThingPrincipal",
	}
}
