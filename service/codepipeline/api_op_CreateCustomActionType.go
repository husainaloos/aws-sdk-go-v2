// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new custom action that can be used in all pipelines associated with
// the AWS account. Only used for custom actions.
func (c *Client) CreateCustomActionType(ctx context.Context, params *CreateCustomActionTypeInput, optFns ...func(*Options)) (*CreateCustomActionTypeOutput, error) {
	stack := middleware.NewStack("CreateCustomActionType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateCustomActionTypeMiddlewares(stack)
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
	addOpCreateCustomActionTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomActionType(options.Region), middleware.Before)
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
			OperationName: "CreateCustomActionType",
			Err:           err,
		}
	}
	out := result.(*CreateCustomActionTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCustomActionType operation.
type CreateCustomActionTypeInput struct {
	// The details of the input artifact for the action, such as its commit ID.
	InputArtifactDetails *types.ArtifactDetails
	// The details of the output artifact of the action, such as its commit ID.
	OutputArtifactDetails *types.ArtifactDetails
	// URLs that provide users information about this custom action.
	Settings *types.ActionTypeSettings
	// The version identifier of the custom action.
	Version *string
	// The provider of the service used in the custom action, such as AWS CodeDeploy.
	Provider *string
	// The tags for the custom action.
	Tags []*types.Tag
	// The configuration properties for the custom action. You can refer to a name in
	// the configuration properties of the custom action within the URL templates by
	// following the format of {Config:name}, as long as the configuration property is
	// both required and not secret. For more information, see Create a Custom Action
	// for a Pipeline
	// (https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html).
	ConfigurationProperties []*types.ActionConfigurationProperty
	// The category of the custom action, such as a build action or a test action.
	// Although Source and Approval are listed as valid values, they are not currently
	// functional. These values are reserved for future use.
	Category types.ActionCategory
}

// Represents the output of a CreateCustomActionType operation.
type CreateCustomActionTypeOutput struct {
	// Returns information about the details of an action type.
	ActionType *types.ActionType
	// Specifies the tags applied to the custom action.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateCustomActionTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCustomActionType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCustomActionType{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCustomActionType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "CreateCustomActionType",
	}
}
