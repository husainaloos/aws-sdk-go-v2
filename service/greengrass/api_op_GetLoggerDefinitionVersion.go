// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a logger definition version.
func (c *Client) GetLoggerDefinitionVersion(ctx context.Context, params *GetLoggerDefinitionVersionInput, optFns ...func(*Options)) (*GetLoggerDefinitionVersionOutput, error) {
	stack := middleware.NewStack("GetLoggerDefinitionVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetLoggerDefinitionVersionMiddlewares(stack)
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
	addOpGetLoggerDefinitionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(options.Region), middleware.Before)

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
			OperationName: "GetLoggerDefinitionVersion",
			Err:           err,
		}
	}
	out := result.(*GetLoggerDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoggerDefinitionVersionInput struct {
	// The ID of the logger definition.
	LoggerDefinitionId *string
	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
	// The ID of the logger definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListLoggerDefinitionVersions'' requests. If the version is the last one
	// that was associated with a logger definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	LoggerDefinitionVersionId *string
}

type GetLoggerDefinitionVersionOutput struct {
	// The version of the logger definition version.
	Version *string
	// The ID of the logger definition version.
	Id *string
	// The ARN of the logger definition version.
	Arn *string
	// Information about the logger definition version.
	Definition *types.LoggerDefinitionVersion
	// The time, in milliseconds since the epoch, when the logger definition version
	// was created.
	CreationTimestamp *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetLoggerDefinitionVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetLoggerDefinitionVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLoggerDefinitionVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetLoggerDefinitionVersion",
	}
}