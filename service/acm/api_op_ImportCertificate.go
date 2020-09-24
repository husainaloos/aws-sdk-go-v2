// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Imports a certificate into AWS Certificate Manager (ACM) to use with services
// that are integrated with ACM. Note that integrated services
// (https://docs.aws.amazon.com/acm/latest/userguide/acm-services.html) allow only
// certificate types and keys they support to be associated with their resources.
// Further, their support differs depending on whether the certificate is imported
// into IAM or into ACM. For more information, see the documentation for each
// service. For more information about importing certificates into ACM, see
// Importing Certificates
// (https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in
// the AWS Certificate Manager User Guide.  <note> <p>ACM does not provide <a
// href="https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html">managed
// renewal</a> for certificates that you import.</p> </note> <p>Note the following
// guidelines when importing third party certificates:</p> <ul> <li> <p>You must
// enter the private key that matches the certificate you are importing.</p> </li>
// <li> <p>The private key must be unencrypted. You cannot import a private key
// that is protected by a password or a passphrase.</p> </li> <li> <p>If the
// certificate you are importing is not self-signed, you must enter its certificate
// chain.</p> </li> <li> <p>If a certificate chain is included, the issuer must be
// the subject of one of the certificates in the chain.</p> </li> <li> <p>The
// certificate, private key, and certificate chain must be PEM-encoded.</p> </li>
// <li> <p>The current time must be between the <code>Not Before</code> and
// <code>Not After</code> certificate fields.</p> </li> <li> <p>The
// <code>Issuer</code> field must not be empty.</p> </li> <li> <p>The OCSP
// authority URL, if present, must not exceed 1000 characters.</p> </li> <li> <p>To
// import a new certificate, omit the <code>CertificateArn</code> argument. Include
// this argument only when you want to replace a previously imported certifica</p>
// </li> <li> <p>When you import a certificate by using the CLI, you must specify
// the certificate, the certificate chain, and the private key by their file names
// preceded by <code>file://</code>. For example, you can specify a certificate
// saved in the <code>C:\temp</code> folder as
// <code>file://C:\temp\certificate_to_import.pem</code>. If you are making an HTTP
// or HTTPS Query request, include these arguments as BLOBs. </p> </li> <li>
// <p>When you import a certificate by using an SDK, you must specify the
// certificate, the certificate chain, and the private key files in the manner
// required by the programming language you're using. </p> </li> <li> <p>The
// cryptographic algorithm of an imported certificate must match the algorithm of
// the signing CA. For example, if the signing CA key type is RSA, then the
// certificate key type must also be RSA.</p> </li> </ul> <p>This operation returns
// the <a
// href="https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html">Amazon
// Resource Name (ARN)</a> of the imported certificate.</p>
func (c *Client) ImportCertificate(ctx context.Context, params *ImportCertificateInput, optFns ...func(*Options)) (*ImportCertificateOutput, error) {
	stack := middleware.NewStack("ImportCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpImportCertificateMiddlewares(stack)
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
	addOpImportCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportCertificate(options.Region), middleware.Before)
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
			OperationName: "ImportCertificate",
			Err:           err,
		}
	}
	out := result.(*ImportCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCertificateInput struct {
	// The Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// an imported certificate to replace. To import a new certificate, omit this
	// field.
	CertificateArn *string
	// The private key that matches the public key in the certificate.
	PrivateKey []byte
	// One or more resource tags to associate with the imported certificate. Note: You
	// cannot apply tags when reimporting a certificate.
	Tags []*types.Tag
	// The certificate to import.
	Certificate []byte
	// The PEM encoded certificate chain.
	CertificateChain []byte
}

type ImportCertificateOutput struct {
	// The Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the imported certificate.
	CertificateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpImportCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpImportCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opImportCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "ImportCertificate",
	}
}
