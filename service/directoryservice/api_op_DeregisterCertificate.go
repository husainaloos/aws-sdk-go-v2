// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes from the system the certificate that was registered for a secured LDAP
// connection.
func (c *Client) DeregisterCertificate(ctx context.Context, params *DeregisterCertificateInput, optFns ...func(*Options)) (*DeregisterCertificateOutput, error) {
	if params == nil {
		params = &DeregisterCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterCertificate", params, optFns, addOperationDeregisterCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterCertificateInput struct {

	// The identifier of the certificate.
	//
	// This member is required.
	CertificateId *string

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string
}

type DeregisterCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterCertificate{}, middleware.After)
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
	addOpDeregisterCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeregisterCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DeregisterCertificate",
	}
}
