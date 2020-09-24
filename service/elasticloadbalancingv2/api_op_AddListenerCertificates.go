// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds the specified SSL server certificate to the certificate list for the
// specified HTTPS or TLS listener. If the certificate in already in the
// certificate list, the call is successful but the certificate is not added again.
// To get the certificate list for a listener, use DescribeListenerCertificates ().
// To remove certificates from the certificate list for a listener, use
// RemoveListenerCertificates (). To replace the default certificate for a
// listener, use ModifyListener (). For more information, see SSL Certificates
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#https-listener-certificates)
// in the Application Load Balancers Guide.
func (c *Client) AddListenerCertificates(ctx context.Context, params *AddListenerCertificatesInput, optFns ...func(*Options)) (*AddListenerCertificatesOutput, error) {
	stack := middleware.NewStack("AddListenerCertificates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAddListenerCertificatesMiddlewares(stack)
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
	addOpAddListenerCertificatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddListenerCertificates(options.Region), middleware.Before)
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
			OperationName: "AddListenerCertificates",
			Err:           err,
		}
	}
	out := result.(*AddListenerCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddListenerCertificatesInput struct {
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string
	// The certificate to add. You can specify one certificate per call. Set
	// CertificateArn to the certificate ARN but do not set IsDefault.
	Certificates []*types.Certificate
}

type AddListenerCertificatesOutput struct {
	// Information about the certificates in the certificate list.
	Certificates []*types.Certificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAddListenerCertificatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAddListenerCertificates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAddListenerCertificates{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddListenerCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "AddListenerCertificates",
	}
}
