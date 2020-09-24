// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a ReqeustValidator () of a given RestApi ().
func (c *Client) CreateRequestValidator(ctx context.Context, params *CreateRequestValidatorInput, optFns ...func(*Options)) (*CreateRequestValidatorOutput, error) {
	stack := middleware.NewStack("CreateRequestValidator", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateRequestValidatorMiddlewares(stack)
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
	addOpCreateRequestValidatorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRequestValidator(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "CreateRequestValidator",
			Err:           err,
		}
	}
	out := result.(*CreateRequestValidatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a RequestValidator () of a given RestApi ().
type CreateRequestValidatorInput struct {
	Template *bool
	Title    *string
	// [Required] The string identifier of the associated RestApi ().
	RestApiId        *string
	TemplateSkipList []*string
	// The name of the to-be-created RequestValidator ().
	Name *string
	// A Boolean flag to indicate whether to validate request body according to the
	// configured model schema for the method (true) or not (false).
	ValidateRequestBody *bool
	// A Boolean flag to indicate whether to validate request parameters, true, or not
	// false.
	ValidateRequestParameters *bool
}

// A set of validation rules for incoming Method () requests. In OpenAPI, a
// RequestValidator () of an API is defined by the
// x-amazon-apigateway-request-validators.requestValidator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validators.requestValidator.html)
// object. It the referenced using the x-amazon-apigateway-request-validator
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions.html#api-gateway-swagger-extensions-request-validator)
// property. Enable Basic Request Validation in API Gateway
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html)
type CreateRequestValidatorOutput struct {
	// The identifier of this RequestValidator ().
	Id *string
	// A Boolean flag to indicate whether to validate a request body according to the
	// configured Model () schema.
	ValidateRequestBody *bool
	// A Boolean flag to indicate whether to validate request parameters (true) or not
	// (false).
	ValidateRequestParameters *bool
	// The name of this RequestValidator ()
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateRequestValidatorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateRequestValidator{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRequestValidator{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRequestValidator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateRequestValidator",
	}
}
