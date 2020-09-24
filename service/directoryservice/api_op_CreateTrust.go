// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// AWS Directory Service for Microsoft Active Directory allows you to configure
// trust relationships. For example, you can establish a trust between your AWS
// Managed Microsoft AD directory, and your existing on-premises Microsoft Active
// Directory. This would allow you to provide users and groups access to resources
// in either domain, with a single set of credentials. This action initiates the
// creation of the AWS side of a trust relationship between an AWS Managed
// Microsoft AD directory and an external domain. You can create either a forest
// trust or an external trust.
func (c *Client) CreateTrust(ctx context.Context, params *CreateTrustInput, optFns ...func(*Options)) (*CreateTrustOutput, error) {
	stack := middleware.NewStack("CreateTrust", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTrustMiddlewares(stack)
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
	addOpCreateTrustValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrust(options.Region), middleware.Before)
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
			OperationName: "CreateTrust",
			Err:           err,
		}
	}
	out := result.(*CreateTrustOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// AWS Directory Service for Microsoft Active Directory allows you to configure
// trust relationships. For example, you can establish a trust between your AWS
// Managed Microsoft AD directory, and your existing on-premises Microsoft Active
// Directory. This would allow you to provide users and groups access to resources
// in either domain, with a single set of credentials. This action initiates the
// creation of the AWS side of a trust relationship between an AWS Managed
// Microsoft AD directory and an external domain.
type CreateTrustInput struct {
	// The Directory ID of the AWS Managed Microsoft AD directory for which to
	// establish the trust relationship.
	DirectoryId *string
	// The Fully Qualified Domain Name (FQDN) of the external domain for which to
	// create the trust relationship.
	RemoteDomainName *string
	// Optional parameter to enable selective authentication for the trust.
	SelectiveAuth types.SelectiveAuth
	// The trust relationship type. Forest is the default.
	TrustType types.TrustType
	// The direction of the trust relationship.
	TrustDirection types.TrustDirection
	// The IP addresses of the remote DNS server associated with RemoteDomainName.
	ConditionalForwarderIpAddrs []*string
	// The trust password. The must be the same password that was used when creating
	// the trust relationship on the external domain.
	TrustPassword *string
}

// The result of a CreateTrust request.
type CreateTrustOutput struct {
	// A unique identifier for the trust relationship that was created.
	TrustId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTrustMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrust{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrust{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTrust(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CreateTrust",
	}
}
