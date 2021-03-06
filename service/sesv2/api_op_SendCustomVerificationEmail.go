// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an email address to the list of identities for your Amazon SES account in
// the current AWS Region and attempts to verify it. As a result of executing this
// operation, a customized verification email is sent to the specified address. To
// use this operation, you must first create a custom verification email template.
// For more information about creating and using custom verification email
// templates, see Using Custom Verification Email Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-verify-address-custom.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) SendCustomVerificationEmail(ctx context.Context, params *SendCustomVerificationEmailInput, optFns ...func(*Options)) (*SendCustomVerificationEmailOutput, error) {
	if params == nil {
		params = &SendCustomVerificationEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendCustomVerificationEmail", params, optFns, addOperationSendCustomVerificationEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendCustomVerificationEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a custom verification email to a specified
// recipient.
type SendCustomVerificationEmailInput struct {

	// The email address to verify.
	//
	// This member is required.
	EmailAddress *string

	// The name of the custom verification email template to use when sending the
	// verification email.
	//
	// This member is required.
	TemplateName *string

	// Name of a configuration set to use when sending the verification email.
	ConfigurationSetName *string
}

// The following element is returned by the service.
type SendCustomVerificationEmailOutput struct {

	// The unique message identifier returned from the SendCustomVerificationEmail
	// operation.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSendCustomVerificationEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendCustomVerificationEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendCustomVerificationEmail{}, middleware.After)
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
	addOpSendCustomVerificationEmailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendCustomVerificationEmail(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSendCustomVerificationEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendCustomVerificationEmail",
	}
}
