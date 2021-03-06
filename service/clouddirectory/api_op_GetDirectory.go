// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves metadata about a directory.
func (c *Client) GetDirectory(ctx context.Context, params *GetDirectoryInput, optFns ...func(*Options)) (*GetDirectoryOutput, error) {
	if params == nil {
		params = &GetDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDirectory", params, optFns, addOperationGetDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDirectoryInput struct {

	// The ARN of the directory.
	//
	// This member is required.
	DirectoryArn *string
}

type GetDirectoryOutput struct {

	// Metadata about the directory.
	//
	// This member is required.
	Directory *types.Directory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDirectory{}, middleware.After)
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
	addOpGetDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDirectory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "GetDirectory",
	}
}
