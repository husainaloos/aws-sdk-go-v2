// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates metadata in a build resource, including the build name and version. To
// update the metadata, specify the build ID to update and provide the new values.
// If successful, a build object containing the updated metadata is returned. Learn
// more  Upload a Custom Server Build
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// Related operations
//
//     * CreateBuild ()
//
//     * ListBuilds ()
//
//     *
// DescribeBuild ()
//
//     * UpdateBuild ()
//
//     * DeleteBuild ()
func (c *Client) UpdateBuild(ctx context.Context, params *UpdateBuildInput, optFns ...func(*Options)) (*UpdateBuildOutput, error) {
	stack := middleware.NewStack("UpdateBuild", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateBuildMiddlewares(stack)
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
	addOpUpdateBuildValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBuild(options.Region), middleware.Before)
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
			OperationName: "UpdateBuild",
			Err:           err,
		}
	}
	out := result.(*UpdateBuildOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type UpdateBuildInput struct {
	// A unique identifier for a build to update. You can use either the build ID or
	// ARN value.
	BuildId *string
	// Version information that is associated with a build or script. Version strings
	// do not need to be unique.
	Version *string
	// A descriptive label that is associated with a build. Build names do not need to
	// be unique.
	Name *string
}

// Represents the returned data in response to a request action.
type UpdateBuildOutput struct {
	// The updated build resource.
	Build *types.Build

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateBuildMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBuild{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBuild{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateBuild(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateBuild",
	}
}
