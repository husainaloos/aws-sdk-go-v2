// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an SSL/TLS certificate for your Amazon Lightsail content delivery
// network (CDN) distribution.  <p>Certificates that are currently attached to a
// distribution cannot be deleted. Use the
// <code>DetachCertificateFromDistribution</code> action to detach a certificate
// from a distribution.</p>
func (c *Client) DeleteCertificate(ctx context.Context, params *DeleteCertificateInput, optFns ...func(*Options)) (*DeleteCertificateOutput, error) {
	if params == nil {
		params = &DeleteCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCertificate", params, optFns, addOperationDeleteCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCertificateInput struct {

	// The name of the certificate to delete.  <p>Use the <code>GetCertificates</code>
	// action to get a list of certificate names that you can specify.</p>
	//
	// This member is required.
	CertificateName *string
}

type DeleteCertificateOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCertificate{}, middleware.After)
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
	addOpDeleteCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteCertificate",
	}
}
