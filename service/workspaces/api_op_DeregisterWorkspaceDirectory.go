// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is deregistered. If any WorkSpaces are registered
// to this directory, you must remove them before you can deregister the directory.
func (c *Client) DeregisterWorkspaceDirectory(ctx context.Context, params *DeregisterWorkspaceDirectoryInput, optFns ...func(*Options)) (*DeregisterWorkspaceDirectoryOutput, error) {
	stack := middleware.NewStack("DeregisterWorkspaceDirectory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterWorkspaceDirectoryMiddlewares(stack)
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
	addOpDeregisterWorkspaceDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterWorkspaceDirectory(options.Region), middleware.Before)
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
			OperationName: "DeregisterWorkspaceDirectory",
			Err:           err,
		}
	}
	out := result.(*DeregisterWorkspaceDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterWorkspaceDirectoryInput struct {
	// The identifier of the directory. If any WorkSpaces are registered to this
	// directory, you must remove them before you deregister the directory, or you will
	// receive an OperationNotSupportedException error.
	DirectoryId *string
}

type DeregisterWorkspaceDirectoryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterWorkspaceDirectoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterWorkspaceDirectory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterWorkspaceDirectory{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterWorkspaceDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DeregisterWorkspaceDirectory",
	}
}
