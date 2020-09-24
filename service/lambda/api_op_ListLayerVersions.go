// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the versions of an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// Versions that have been deleted aren't listed. Specify a runtime identifier
// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) to list only
// versions that indicate that they're compatible with that runtime.
func (c *Client) ListLayerVersions(ctx context.Context, params *ListLayerVersionsInput, optFns ...func(*Options)) (*ListLayerVersionsOutput, error) {
	stack := middleware.NewStack("ListLayerVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListLayerVersionsMiddlewares(stack)
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
	addOpListLayerVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLayerVersions(options.Region), middleware.Before)
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
			OperationName: "ListLayerVersions",
			Err:           err,
		}
	}
	out := result.(*ListLayerVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLayerVersionsInput struct {
	// A runtime identifier. For example, go1.x.
	CompatibleRuntime types.Runtime
	// A pagination token returned by a previous call.
	Marker *string
	// The name or Amazon Resource Name (ARN) of the layer.
	LayerName *string
	// The maximum number of versions to return.
	MaxItems *int32
}

type ListLayerVersionsOutput struct {
	// A pagination token returned when the response doesn't contain all versions.
	NextMarker *string
	// A list of versions.
	LayerVersions []*types.LayerVersionsListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListLayerVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListLayerVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListLayerVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListLayerVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListLayerVersions",
	}
}
