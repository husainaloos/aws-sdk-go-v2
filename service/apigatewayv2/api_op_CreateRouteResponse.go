// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a RouteResponse for a Route.
func (c *Client) CreateRouteResponse(ctx context.Context, params *CreateRouteResponseInput, optFns ...func(*Options)) (*CreateRouteResponseOutput, error) {
	stack := middleware.NewStack("CreateRouteResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateRouteResponseMiddlewares(stack)
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
	addOpCreateRouteResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRouteResponse(options.Region), middleware.Before)
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
			OperationName: "CreateRouteResponse",
			Err:           err,
		}
	}
	out := result.(*CreateRouteResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new RouteResponse resource to represent a route response.
type CreateRouteResponseInput struct {
	// The route response parameters.
	ResponseParameters map[string]*types.ParameterConstraints
	// The route response key.
	RouteResponseKey *string
	// The response models for the route response.
	ResponseModels map[string]*string
	// The API identifier.
	ApiId *string
	// The route ID.
	RouteId *string
	// The model selection expression for the route response. Supported only for
	// WebSocket APIs.
	ModelSelectionExpression *string
}

type CreateRouteResponseOutput struct {
	// Represents the response models of a route response.
	ResponseModels map[string]*string
	// Represents the identifier of a route response.
	RouteResponseId *string
	// Represents the response parameters of a route response.
	ResponseParameters map[string]*types.ParameterConstraints
	// Represents the route response key of a route response.
	RouteResponseKey *string
	// Represents the model selection expression of a route response. Supported only
	// for WebSocket APIs.
	ModelSelectionExpression *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateRouteResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateRouteResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRouteResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRouteResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateRouteResponse",
	}
}
