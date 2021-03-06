// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a preview of the MIME content of an email when provided with a template
// and a set of replacement data. You can execute this operation no more than once
// per second.
func (c *Client) TestRenderTemplate(ctx context.Context, params *TestRenderTemplateInput, optFns ...func(*Options)) (*TestRenderTemplateOutput, error) {
	if params == nil {
		params = &TestRenderTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestRenderTemplate", params, optFns, addOperationTestRenderTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestRenderTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestRenderTemplateInput struct {

	// A list of replacement values to apply to the template. This parameter is a JSON
	// object, typically consisting of key-value pairs in which the keys correspond to
	// replacement tags in the email template.
	//
	// This member is required.
	TemplateData *string

	// The name of the template that you want to render.
	//
	// This member is required.
	TemplateName *string
}

type TestRenderTemplateOutput struct {

	// The complete MIME message rendered by applying the data in the TemplateData
	// parameter to the template specified in the TemplateName parameter.
	RenderedTemplate *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestRenderTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTestRenderTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTestRenderTemplate{}, middleware.After)
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
	addOpTestRenderTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestRenderTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTestRenderTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "TestRenderTemplate",
	}
}
