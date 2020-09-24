// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecated. Use the VerifyEmailIdentity operation to verify a new email address.
func (c *Client) VerifyEmailAddress(ctx context.Context, params *VerifyEmailAddressInput, optFns ...func(*Options)) (*VerifyEmailAddressOutput, error) {
	stack := middleware.NewStack("VerifyEmailAddress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpVerifyEmailAddressMiddlewares(stack)
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
	addOpVerifyEmailAddressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyEmailAddress(options.Region), middleware.Before)
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
			OperationName: "VerifyEmailAddress",
			Err:           err,
		}
	}
	out := result.(*VerifyEmailAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to begin email address verification with Amazon SES. For
// information about email address verification, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html).
type VerifyEmailAddressInput struct {
	// The email address to be verified.
	EmailAddress *string
}

type VerifyEmailAddressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpVerifyEmailAddressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpVerifyEmailAddress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpVerifyEmailAddress{}, middleware.After)
}

func newServiceMetadataMiddleware_opVerifyEmailAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "VerifyEmailAddress",
	}
}
