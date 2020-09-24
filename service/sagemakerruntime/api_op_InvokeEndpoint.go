// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// After you deploy a model into production using Amazon SageMaker hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint. For an overview of Amazon SageMaker, see How
// It Works (https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html)
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
func (c *Client) InvokeEndpoint(ctx context.Context, params *InvokeEndpointInput, optFns ...func(*Options)) (*InvokeEndpointOutput, error) {
	stack := middleware.NewStack("InvokeEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInvokeEndpointMiddlewares(stack)
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
	addOpInvokeEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeEndpoint(options.Region), middleware.Before)
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
			OperationName: "InvokeEndpoint",
			Err:           err,
		}
	}
	out := result.(*InvokeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeEndpointInput struct {
	// The MIME type of the input data in the request body.
	ContentType *string
	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model.
	Body []byte
	// The name of the endpoint that you specified when you created the endpoint using
	// the CreateEndpoint
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html) API.
	EndpointName *string
	// The desired MIME type of the inference in the response.
	Accept *string
}

type InvokeEndpointOutput struct {
	// Includes the inference provided by the model.
	Body []byte
	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string
	// The MIME type of the inference returned in the response body.
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInvokeEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInvokeEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opInvokeEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "InvokeEndpoint",
	}
}
