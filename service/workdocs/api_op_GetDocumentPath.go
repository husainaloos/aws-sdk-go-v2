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
// requested document. By default, Amazon WorkDocs returns a maximum of 100 levels
// upwards from the requested document and only includes the IDs of the parent
// folders in the path. You can limit the maximum number of levels. You can also
// request the names of the parent folders.
func (c *Client) GetDocumentPath(ctx context.Context, params *GetDocumentPathInput, optFns ...func(*Options)) (*GetDocumentPathOutput, error) {
	stack := middleware.NewStack("GetDocumentPath", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDocumentPathMiddlewares(stack)
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
	addOpGetDocumentPathValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentPath(options.Region), middleware.Before)
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
			OperationName: "GetDocumentPath",
			Err:           err,
		}
	}
	out := result.(*GetDocumentPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentPathInput struct {
	// The ID of the document.
	DocumentId *string
	// The maximum number of levels in the hierarchy to return.
	Limit *int32
	// This value is not supported.
	Marker *string
	// A comma-separated list of values. Specify NAME to include the names of the
	// parent folders.
	Fields *string
	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string
}

type GetDocumentPathOutput struct {
	// The path information.
	Path *types.ResourcePath

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDocumentPathMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentPath{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentPath{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDocumentPath(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "GetDocumentPath",
	}
}
