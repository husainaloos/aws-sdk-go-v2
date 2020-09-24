// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a core definition version.
func (c *Client) GetCoreDefinition(ctx context.Context, params *GetCoreDefinitionInput, optFns ...func(*Options)) (*GetCoreDefinitionOutput, error) {
	stack := middleware.NewStack("GetCoreDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCoreDefinitionMiddlewares(stack)
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
	addOpGetCoreDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCoreDefinition(options.Region), middleware.Before)
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
			OperationName: "GetCoreDefinition",
			Err:           err,
		}
	}
	out := result.(*GetCoreDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCoreDefinitionInput struct {
	// The ID of the core definition.
	CoreDefinitionId *string
}

type GetCoreDefinitionOutput struct {
	// The ARN of the definition.
	Arn *string
	// The ID of the definition.
	Id *string
	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string
	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string
	// The ID of the latest version associated with the definition.
	LatestVersion *string
	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string
	// The name of the definition.
	Name *string
	// Tag(s) attached to the resource arn.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCoreDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCoreDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCoreDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCoreDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetCoreDefinition",
	}
}
