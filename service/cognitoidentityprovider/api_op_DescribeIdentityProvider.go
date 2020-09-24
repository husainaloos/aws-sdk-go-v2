// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a specific identity provider.
func (c *Client) DescribeIdentityProvider(ctx context.Context, params *DescribeIdentityProviderInput, optFns ...func(*Options)) (*DescribeIdentityProviderOutput, error) {
	stack := middleware.NewStack("DescribeIdentityProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeIdentityProviderMiddlewares(stack)
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
	addOpDescribeIdentityProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIdentityProvider(options.Region), middleware.Before)
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
			OperationName: "DescribeIdentityProvider",
			Err:           err,
		}
	}
	out := result.(*DescribeIdentityProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIdentityProviderInput struct {
	// The identity provider name.
	ProviderName *string
	// The user pool ID.
	UserPoolId *string
}

type DescribeIdentityProviderOutput struct {
	// The identity provider that was deleted.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeIdentityProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeIdentityProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeIdentityProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeIdentityProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DescribeIdentityProvider",
	}
}
