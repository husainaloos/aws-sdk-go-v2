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

// Adds an email address to the list of identities for your Amazon SES account in
// the current AWS Region and attempts to verify it. As a result of executing this
// operation, a customized verification email is sent to the specified address. To
// use this operation, you must first create a custom verification email template.
// For more information about creating and using custom verification email
// templates, see Using Custom Verification Email Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) SendCustomVerificationEmail(ctx context.Context, params *SendCustomVerificationEmailInput, optFns ...func(*Options)) (*SendCustomVerificationEmailOutput, error) {
	stack := middleware.NewStack("SendCustomVerificationEmail", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSendCustomVerificationEmailMiddlewares(stack)
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
	addOpSendCustomVerificationEmailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendCustomVerificationEmail(options.Region), middleware.Before)
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
			OperationName: "SendCustomVerificationEmail",
			Err:           err,
		}
	}
	out := result.(*SendCustomVerificationEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a custom verification email to a specified
// recipient.
type SendCustomVerificationEmailInput struct {
	// The email address to verify.
	EmailAddress *string
	// Name of a configuration set to use when sending the verification email.
	ConfigurationSetName *string
	// The name of the custom verification email template to use when sending the
	// verification email.
	TemplateName *string
}

// The response received when attempting to send the custom verification email.
type SendCustomVerificationEmailOutput struct {
	// The unique message identifier returned from the SendCustomVerificationEmail
	// operation.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSendCustomVerificationEmailMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSendCustomVerificationEmail{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSendCustomVerificationEmail{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendCustomVerificationEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendCustomVerificationEmail",
	}
}
