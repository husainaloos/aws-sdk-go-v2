// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the roles for an identity pool. You must use AWS Developer credentials to
// call this API.
func (c *Client) GetIdentityPoolRoles(ctx context.Context, params *GetIdentityPoolRolesInput, optFns ...func(*Options)) (*GetIdentityPoolRolesOutput, error) {
	stack := middleware.NewStack("GetIdentityPoolRoles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetIdentityPoolRolesMiddlewares(stack)
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
	addOpGetIdentityPoolRolesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityPoolRoles(options.Region), middleware.Before)
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
			OperationName: "GetIdentityPoolRoles",
			Err:           err,
		}
	}
	out := result.(*GetIdentityPoolRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the GetIdentityPoolRoles action.
type GetIdentityPoolRolesInput struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string
}

// Returned in response to a successful GetIdentityPoolRoles operation.
type GetIdentityPoolRolesOutput struct {
	// How users for a specific identity provider are to mapped to roles. This is a
	// String-to-RoleMapping () object map. The string identifies the identity
	// provider, for example, "graph.facebook.com" or
	// "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
	RoleMappings map[string]*types.RoleMapping
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string
	// The map of roles associated with this pool. Currently only authenticated and
	// unauthenticated roles are supported.
	Roles map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetIdentityPoolRolesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetIdentityPoolRoles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetIdentityPoolRoles{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIdentityPoolRoles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "GetIdentityPoolRoles",
	}
}
