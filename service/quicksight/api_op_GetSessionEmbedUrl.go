// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a session URL and authorization code that you can embed in your web
// server code.
func (c *Client) GetSessionEmbedUrl(ctx context.Context, params *GetSessionEmbedUrlInput, optFns ...func(*Options)) (*GetSessionEmbedUrlOutput, error) {
	stack := middleware.NewStack("GetSessionEmbedUrl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSessionEmbedUrlMiddlewares(stack)
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
	addOpGetSessionEmbedUrlValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSessionEmbedUrl(options.Region), middleware.Before)
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
			OperationName: "GetSessionEmbedUrl",
			Err:           err,
		}
	}
	out := result.(*GetSessionEmbedUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSessionEmbedUrlInput struct {
	// How many minutes the session is valid. The session lifetime must be 15-600
	// minutes.
	SessionLifetimeInMinutes *int64
	// The ID for the AWS account that contains the QuickSight session that you're
	// embedding.
	AwsAccountId *string
	// The Amazon QuickSight user's Amazon Resource Name (ARN), for use with QUICKSIGHT
	// identity type. You can use this for any Amazon QuickSight users in your account
	// (readers, authors, or admins) authenticated as one of the following:
	//
	//     *
	// Active Directory (AD) users or group members
	//
	//     * Invited nonfederated users
	//
	//
	// * IAM users and IAM role-based sessions authenticated through Federated Single
	// Sign-On using SAML, OpenID Connect, or IAM federation.
	UserArn *string
	// The entry point for the embedded session.
	EntryPoint *string
}

type GetSessionEmbedUrlOutput struct {
	// A single-use URL that you can put into your server-side web page to embed your
	// QuickSight session. This URL is valid for 5 minutes. The API provides the URL
	// with an auth_code value that enables one (and only one) sign-on to a user
	// session that is valid for 10 hours.
	EmbedUrl *string
	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSessionEmbedUrlMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSessionEmbedUrl{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSessionEmbedUrl{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSessionEmbedUrl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "GetSessionEmbedUrl",
	}
}
