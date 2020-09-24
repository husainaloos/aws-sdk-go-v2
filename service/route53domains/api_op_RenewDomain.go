// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation renews a domain for the specified number of years. The cost of
// renewing your domain is billed to your AWS account. We recommend that you renew
// your domain several weeks before the expiration date. Some TLD registries delete
// domains before the expiration date if you haven't renewed far enough in advance.
// For more information about renewing domain registration, see Renewing
// Registration for a Domain
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-renew.html) in
// the Amazon Route 53 Developer Guide.
func (c *Client) RenewDomain(ctx context.Context, params *RenewDomainInput, optFns ...func(*Options)) (*RenewDomainOutput, error) {
	stack := middleware.NewStack("RenewDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRenewDomainMiddlewares(stack)
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
	addOpRenewDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRenewDomain(options.Region), middleware.Before)
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
			OperationName: "RenewDomain",
			Err:           err,
		}
	}
	out := result.(*RenewDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A RenewDomain request includes the number of years that you want to renew for
// and the current expiration year.
type RenewDomainInput struct {
	// The name of the domain that you want to renew.
	DomainName *string
	// The year when the registration for the domain is set to expire. This value must
	// match the current expiration date for the domain.
	CurrentExpiryYear *int32
	// The number of years that you want to renew the domain for. The maximum number of
	// years depends on the top-level domain. For the range of valid values for your
	// domain, see Domains that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. Default: 1
	DurationInYears *int32
}

type RenewDomainOutput struct {
	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRenewDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRenewDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRenewDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opRenewDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "RenewDomain",
	}
}
