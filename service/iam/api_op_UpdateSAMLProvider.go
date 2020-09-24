// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the metadata document for an existing SAML provider resource object.
// This operation requires Signature Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) UpdateSAMLProvider(ctx context.Context, params *UpdateSAMLProviderInput, optFns ...func(*Options)) (*UpdateSAMLProviderOutput, error) {
	stack := middleware.NewStack("UpdateSAMLProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateSAMLProviderMiddlewares(stack)
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
	addOpUpdateSAMLProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSAMLProvider(options.Region), middleware.Before)
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
			OperationName: "UpdateSAMLProvider",
			Err:           err,
		}
	}
	out := result.(*UpdateSAMLProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSAMLProviderInput struct {
	// An XML document generated by an identity provider (IdP) that supports SAML 2.0.
	// The document includes the issuer's name, expiration information, and keys that
	// can be used to validate the SAML authentication response (assertions) that are
	// received from the IdP. You must generate the metadata document using the
	// identity management software that is used as your organization's IdP.
	SAMLMetadataDocument *string
	// The Amazon Resource Name (ARN) of the SAML provider to update. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	SAMLProviderArn *string
}

// Contains the response to a successful UpdateSAMLProvider () request.
type UpdateSAMLProviderOutput struct {
	// The Amazon Resource Name (ARN) of the SAML provider that was updated.
	SAMLProviderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateSAMLProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateSAMLProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateSAMLProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSAMLProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateSAMLProvider",
	}
}
