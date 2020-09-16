// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current status of Easy DKIM signing for an entity. For domain name
// identities, this operation also returns the DKIM tokens that are required for
// Easy DKIM signing, and whether Amazon SES has successfully verified that these
// tokens have been published. This operation takes a list of identities as input
// and returns the following information for each:
//
//     * Whether Easy DKIM signing
// is enabled or disabled.
//
//     * A set of DKIM tokens that represent the identity.
// If the identity is an email address, the tokens represent the domain of that
// address.
//
//     * Whether Amazon SES has successfully verified the DKIM tokens
// published in the domain's DNS. This information is only returned for domain name
// identities, not for email addresses.
//
// This operation is throttled at one request
// per second and can only get DKIM attributes for up to 100 identities at a time.
// For more information about creating DNS records using DKIM tokens, go to the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
func (c *Client) GetIdentityDkimAttributes(ctx context.Context, params *GetIdentityDkimAttributesInput, optFns ...func(*Options)) (*GetIdentityDkimAttributesOutput, error) {
	stack := middleware.NewStack("GetIdentityDkimAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetIdentityDkimAttributesMiddlewares(stack)
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
	addOpGetIdentityDkimAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityDkimAttributes(options.Region), middleware.Before)

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
			OperationName: "GetIdentityDkimAttributes",
			Err:           err,
		}
	}
	out := result.(*GetIdentityDkimAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request for the status of Amazon SES Easy DKIM signing for an
// identity. For domain identities, this request also returns the DKIM tokens that
// are required for Easy DKIM signing, and whether Amazon SES successfully verified
// that these tokens were published. For more information about Easy DKIM, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
type GetIdentityDkimAttributesInput struct {
	// A list of one or more verified identities - email addresses, domains, or both.
	Identities []*string
}

// Represents the status of Amazon SES Easy DKIM signing for an identity. For
// domain identities, this response also contains the DKIM tokens that are required
// for Easy DKIM signing, and whether Amazon SES successfully verified that these
// tokens were published.
type GetIdentityDkimAttributesOutput struct {
	// The DKIM attributes for an email address or a domain.
	DkimAttributes map[string]*types.IdentityDkimAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetIdentityDkimAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetIdentityDkimAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetIdentityDkimAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIdentityDkimAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetIdentityDkimAttributes",
	}
}