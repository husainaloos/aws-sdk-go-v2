// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an AWS Directory Service directory. Before you call DeleteDirectory,
// ensure that all of the required permissions have been explicitly granted through
// a policy. For details about what permissions are required to run the
// DeleteDirectory operation, see AWS Directory Service API Permissions: Actions,
// Resources, and Conditions Reference
// (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
func (c *Client) DeleteDirectory(ctx context.Context, params *DeleteDirectoryInput, optFns ...func(*Options)) (*DeleteDirectoryOutput, error) {
	if params == nil {
		params = &DeleteDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDirectory", params, optFns, addOperationDeleteDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DeleteDirectory () operation.
type DeleteDirectoryInput struct {

	// The identifier of the directory to delete.
	//
	// This member is required.
	DirectoryId *string
}

// Contains the results of the DeleteDirectory () operation.
type DeleteDirectoryOutput struct {

	// The directory identifier.
	DirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDirectory{}, middleware.After)
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
	addOpDeleteDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DeleteDirectory",
	}
}
