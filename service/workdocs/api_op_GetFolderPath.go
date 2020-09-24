// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the path information (the hierarchy from the root folder) for the
// specified folder. By default, Amazon WorkDocs returns a maximum of 100 levels
// upwards from the requested folder and only includes the IDs of the parent
// folders in the path. You can limit the maximum number of levels. You can also
// request the parent folder names.
func (c *Client) GetFolderPath(ctx context.Context, params *GetFolderPathInput, optFns ...func(*Options)) (*GetFolderPathOutput, error) {
	stack := middleware.NewStack("GetFolderPath", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetFolderPathMiddlewares(stack)
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
	addOpGetFolderPathValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFolderPath(options.Region), middleware.Before)
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
			OperationName: "GetFolderPath",
			Err:           err,
		}
	}
	out := result.(*GetFolderPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFolderPathInput struct {
	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string
	// A comma-separated list of values. Specify "NAME" to include the names of the
	// parent folders.
	Fields *string
	// This value is not supported.
	Marker *string
	// The ID of the folder.
	FolderId *string
	// The maximum number of levels in the hierarchy to return.
	Limit *int32
}

type GetFolderPathOutput struct {
	// The path information.
	Path *types.ResourcePath

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetFolderPathMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetFolderPath{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFolderPath{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFolderPath(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "GetFolderPath",
	}
}
