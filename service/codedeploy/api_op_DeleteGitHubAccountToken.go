// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a GitHub account connection.
func (c *Client) DeleteGitHubAccountToken(ctx context.Context, params *DeleteGitHubAccountTokenInput, optFns ...func(*Options)) (*DeleteGitHubAccountTokenOutput, error) {
	stack := middleware.NewStack("DeleteGitHubAccountToken", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteGitHubAccountTokenMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGitHubAccountToken(options.Region), middleware.Before)
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
			OperationName: "DeleteGitHubAccountToken",
			Err:           err,
		}
	}
	out := result.(*DeleteGitHubAccountTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteGitHubAccount operation.
type DeleteGitHubAccountTokenInput struct {
	// The name of the GitHub account connection to delete.
	TokenName *string
}

// Represents the output of a DeleteGitHubAccountToken operation.
type DeleteGitHubAccountTokenOutput struct {
	// The name of the GitHub account connection that was deleted.
	TokenName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteGitHubAccountTokenMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGitHubAccountToken{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGitHubAccountToken{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteGitHubAccountToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "DeleteGitHubAccountToken",
	}
}
