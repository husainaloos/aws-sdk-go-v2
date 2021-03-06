// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing gateway route.
func (c *Client) DeleteGatewayRoute(ctx context.Context, params *DeleteGatewayRouteInput, optFns ...func(*Options)) (*DeleteGatewayRouteOutput, error) {
	if params == nil {
		params = &DeleteGatewayRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGatewayRoute", params, optFns, addOperationDeleteGatewayRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGatewayRouteInput struct {

	// The name of the gateway route to delete.
	//
	// This member is required.
	GatewayRouteName *string

	// The name of the service mesh to delete the gateway route from.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual gateway to delete the route from.
	//
	// This member is required.
	VirtualGatewayName *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

type DeleteGatewayRouteOutput struct {

	// The gateway route that was deleted.
	//
	// This member is required.
	GatewayRoute *types.GatewayRouteData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGatewayRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGatewayRoute{}, middleware.After)
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
	addOpDeleteGatewayRouteValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
